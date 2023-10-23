// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "github.com/caraml-dev/merlin/models"
	mock "github.com/stretchr/testify/mock"
)

// DeploymentService is an autogenerated mock type for the DeploymentService type
type DeploymentService struct {
	mock.Mock
}

// ListDeployments provides a mock function with given fields: modelID, versionID, endpointUUID
func (_m *DeploymentService) ListDeployments(modelID string, versionID string, endpointUUID string) ([]*models.Deployment, error) {
	ret := _m.Called(modelID, versionID, endpointUUID)

	var r0 []*models.Deployment
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) ([]*models.Deployment, error)); ok {
		return rf(modelID, versionID, endpointUUID)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) []*models.Deployment); ok {
		r0 = rf(modelID, versionID, endpointUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Deployment)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(modelID, versionID, endpointUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDeploymentService interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeploymentService creates a new instance of DeploymentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeploymentService(t mockConstructorTestingTNewDeploymentService) *DeploymentService {
	mock := &DeploymentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}