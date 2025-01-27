package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestUnmarshalTopologySpreadConstraints(t *testing.T) {
	testCases := map[string]struct {
		input     string
		exp       TopologySpreadConstraints
		errString string
	}{
		"invalid config schema": {
			input: `- maxSkew: 1
  topologyKey: kubernetes.io/hostname
    whenUnsatisfiable: ScheduleAnyway`,
			errString: "yaml: line 3: mapping values are not allowed in this context",
		},
		"valid configs": {
			input: `- maxSkew: 1
  topologyKey: kubernetes.io/hostname
  whenUnsatisfiable: ScheduleAnyway
- maxSkew: 2
  topologyKey: kubernetes.io/hostname
  whenUnsatisfiable: DoNotSchedule
  labelSelector:
    matchLabels:
      app-label: spread
- maxSkew: 3
  topologyKey: kubernetes.io/hostname
  whenUnsatisfiable: DoNotSchedule
  labelSelector:
    matchLabels:
      app-label: spread
    matchExpressions:
      - key: app-expression
        operator: In
        values:
          - 1`,
			exp: []corev1.TopologySpreadConstraint{
				{
					MaxSkew:           1,
					TopologyKey:       "kubernetes.io/hostname",
					WhenUnsatisfiable: corev1.ScheduleAnyway,
				},
				{
					MaxSkew:           2,
					TopologyKey:       "kubernetes.io/hostname",
					WhenUnsatisfiable: corev1.DoNotSchedule,
					LabelSelector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app-label": "spread",
						},
					},
				},
				{
					MaxSkew:           3,
					TopologyKey:       "kubernetes.io/hostname",
					WhenUnsatisfiable: corev1.DoNotSchedule,
					LabelSelector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app-label": "spread",
						},
						MatchExpressions: []metav1.LabelSelectorRequirement{
							{
								Key:      "app-expression",
								Operator: metav1.LabelSelectorOpIn,
								Values:   []string{"1"},
							},
						},
					},
				},
			},
		},
	}

	for testName, tC := range testCases {
		t.Run(testName, func(t *testing.T) {
			var topologySpreadConstraints TopologySpreadConstraints
			inputByte := []byte(tC.input)
			err := yaml.Unmarshal(inputByte, &topologySpreadConstraints)

			if tC.errString == "" {
				assert.NoError(t, err)
				assert.Equal(t, tC.exp, topologySpreadConstraints)
			} else {
				assert.EqualError(t, err, tC.errString)
				assert.Nil(t, topologySpreadConstraints)
			}
		})
	}
}

func TestPDBConfig(t *testing.T) {
	defaultMinAvailablePercentage := 20

	testCases := []struct {
		desc              string
		envConfigPath     string
		expectedPDBConfig PodDisruptionBudgetConfig
		validateErr       error
	}{
		{
			desc:          "Success: valid pdb config",
			envConfigPath: "./testdata/valid-environment-1.yaml",
			expectedPDBConfig: PodDisruptionBudgetConfig{
				Enabled:                true,
				MinAvailablePercentage: &defaultMinAvailablePercentage,
			},
		},
		{
			desc:          "Invalid: both pdb maxUnavailable and minUnavailable is set",
			envConfigPath: "./testdata/invalid-pdb-environment-both-set.yaml",
			validateErr:   fmt.Errorf("invalid environment config: Key: 'EnvironmentConfig.PodDisruptionBudget.max_unavailable_percentage' Error:Field validation for 'max_unavailable_percentage' failed on the 'choose_one[max_unavailable_percentage,min_available_percentage]' tag\nKey: 'EnvironmentConfig.PodDisruptionBudget.min_available_percentage' Error:Field validation for 'min_available_percentage' failed on the 'choose_one[max_unavailable_percentage,min_available_percentage]' tag"),
		},
		{
			desc:          "Invalid: both pdb maxUnavailable and minUnavailable is set",
			envConfigPath: "./testdata/invalid-pdb-environment-both-nil.yaml",
			validateErr:   fmt.Errorf("invalid environment config: Key: 'EnvironmentConfig.PodDisruptionBudget.max_unavailable_percentage' Error:Field validation for 'max_unavailable_percentage' failed on the 'choose_one[max_unavailable_percentage,min_available_percentage]' tag\nKey: 'EnvironmentConfig.PodDisruptionBudget.min_available_percentage' Error:Field validation for 'min_available_percentage' failed on the 'choose_one[max_unavailable_percentage,min_available_percentage]' tag"),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			envConfigs, err := InitEnvironmentConfigs(tC.envConfigPath)
			if tC.validateErr == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tC.validateErr.Error(), err.Error())
				return
			}

			cfg := &Config{
				ClusterConfig: ClusterConfig{
					EnvironmentConfigs: envConfigs,
				},
			}

			for _, envCfg := range cfg.ClusterConfig.EnvironmentConfigs {
				deploymentCfg := ParseDeploymentConfig(envCfg, "")
				assert.Equal(t, tC.expectedPDBConfig, deploymentCfg.PodDisruptionBudget)
			}
		})
	}
}

func TestImageBuilderConfig(t *testing.T) {
	baseFilePath := "./testdata/config-1.yaml"
	imageBuilderCfgFilePath := "./testdata/valid-imagebuilder-nodeselectors.yaml"

	os.Clearenv()
	setRequiredEnvironmentVariables()

	filePaths := []string{baseFilePath}
	filePaths = append(filePaths, imageBuilderCfgFilePath)

	var emptyCfg Config
	cfg, err := Load(&emptyCfg, filePaths...)
	require.NoError(t, err)
	require.NotNil(t, cfg)

	expected := map[string]string{
		"cloud.google.com/gke-nodepool": "image-build-job-node-pool",
	}
	assert.Equal(t, expected, cfg.ImageBuilderConfig.NodeSelectors)
}

func TestGPUsConfig(t *testing.T) {
	testCases := []struct {
		desc               string
		envConfigPath      string
		expectedGPUsConfig []GPUConfig
		validateErr        error
	}{
		{
			desc:          "Success: valid pdb config",
			envConfigPath: "./testdata/valid-environment-1.yaml",
			expectedGPUsConfig: []GPUConfig{
				{
					Name:         "NVIDIA T4",
					Values:       []string{"None", "1"},
					ResourceType: "nvidia.com/gpu",
					NodeSelector: map[string]string{
						"cloud.google.com/gke-accelerator": "nvidia-tesla-t4",
					},
					MinMonthlyCostPerGPU: 189.07,
					MaxMonthlyCostPerGPU: 189.07,
				},
				{
					Name:         "NVIDIA T4 with Time Sharing",
					Values:       []string{"None", "1"},
					ResourceType: "nvidia.com/gpu",
					NodeSelector: map[string]string{
						"cloud.google.com/gke-accelerator":                "nvidia-tesla-t4",
						"cloud.google.com/gke-max-shared-clients-per-gpu": "8",
						"cloud.google.com/gke-gpu-sharing-strategy":       "time-sharing",
					},
					MinMonthlyCostPerGPU: 23.63,
					MaxMonthlyCostPerGPU: 189.07,
				},
				{
					Name:         "NVIDIA P4",
					Values:       []string{"None", "1", "2"},
					ResourceType: "nvidia.com/gpu",
					NodeSelector: map[string]string{
						"cloud.google.com/gke-accelerator": "nvidia-tesla-p4",
					},
					Tolerations: []corev1.Toleration{
						{
							Key:      "caraml/nvidia-tesla-p4",
							Operator: corev1.TolerationOpEqual,
							Value:    "enabled",
							Effect:   corev1.TaintEffectNoSchedule,
						},
						{
							Key:      "nvidia.com/gpu",
							Operator: corev1.TolerationOpEqual,
							Value:    "present",
							Effect:   corev1.TaintEffectNoSchedule,
						},
					},
					MinMonthlyCostPerGPU: 332.15,
					MaxMonthlyCostPerGPU: 332.15,
				},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			envConfigs, err := InitEnvironmentConfigs(tC.envConfigPath)
			if tC.validateErr == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tC.validateErr.Error(), err.Error())
				return
			}

			cfg := &Config{
				ClusterConfig: ClusterConfig{
					EnvironmentConfigs: envConfigs,
				},
			}

			for _, envCfg := range cfg.ClusterConfig.EnvironmentConfigs {
				deploymentCfg := ParseDeploymentConfig(envCfg, "")
				assert.Equal(t, tC.expectedGPUsConfig, deploymentCfg.GPUs)
			}
		})
	}
}
