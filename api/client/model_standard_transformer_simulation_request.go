/*
 * Merlin
 *
 * API Guide for accessing Merlin's model management, deployment, and serving functionalities
 *
 * API version: 0.14.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type StandardTransformerSimulationRequest struct {
	Payload               *ModelMap              `json:"payload,omitempty"`
	Headers               *ModelMap              `json:"headers,omitempty"`
	Config                *ModelMap              `json:"config,omitempty"`
	ModelPredictionConfig *ModelPredictionConfig `json:"model_prediction_config,omitempty"`
	Protocol              *Protocol              `json:"protocol,omitempty"`
}
