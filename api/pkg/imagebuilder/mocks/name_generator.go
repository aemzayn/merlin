// Code generated by mockery v2.0.0-alpha.14. DO NOT EDIT.

package mocks

import (
	mlp "github.com/caraml-dev/merlin/mlp"
	mock "github.com/stretchr/testify/mock"

	models "github.com/caraml-dev/merlin/models"
)

// nameGenerator is an autogenerated mock type for the nameGenerator type
type nameGenerator struct {
	mock.Mock
}

// generateBuilderJobName provides a mock function with given fields: project, model, version
func (_m *nameGenerator) generateBuilderJobName(project mlp.Project, model *models.Model, version *models.Version) string {
	ret := _m.Called(project, model, version)

	var r0 string
	if rf, ok := ret.Get(0).(func(mlp.Project, *models.Model, *models.Version) string); ok {
		r0 = rf(project, model, version)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// generateDockerImageName provides a mock function with given fields: project, model
func (_m *nameGenerator) generateDockerImageName(project mlp.Project, model *models.Model) string {
	ret := _m.Called(project, model)

	var r0 string
	if rf, ok := ret.Get(0).(func(mlp.Project, *models.Model) string); ok {
		r0 = rf(project, model)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
