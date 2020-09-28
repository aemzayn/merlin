// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/gojek/merlin/models"

// AlertStorage is an autogenerated mock type for the AlertStorage type
type AlertStorage struct {
	mock.Mock
}

// CreateModelEndpointAlert provides a mock function with given fields: alert
func (_m *AlertStorage) CreateModelEndpointAlert(alert *models.ModelEndpointAlert) error {
	ret := _m.Called(alert)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.ModelEndpointAlert) error); ok {
		r0 = rf(alert)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteModelEndpointAlert provides a mock function with given fields: modelId, modelEndpointId
func (_m *AlertStorage) DeleteModelEndpointAlert(modelId models.Id, modelEndpointId models.Id) error {
	ret := _m.Called(modelId, modelEndpointId)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Id, models.Id) error); ok {
		r0 = rf(modelId, modelEndpointId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetModelEndpointAlert provides a mock function with given fields: modelId, modelEndpointId
func (_m *AlertStorage) GetModelEndpointAlert(modelId models.Id, modelEndpointId models.Id) (*models.ModelEndpointAlert, error) {
	ret := _m.Called(modelId, modelEndpointId)

	var r0 *models.ModelEndpointAlert
	if rf, ok := ret.Get(0).(func(models.Id, models.Id) *models.ModelEndpointAlert); ok {
		r0 = rf(modelId, modelEndpointId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ModelEndpointAlert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Id, models.Id) error); ok {
		r1 = rf(modelId, modelEndpointId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListModelEndpointAlerts provides a mock function with given fields: modelId
func (_m *AlertStorage) ListModelEndpointAlerts(modelId models.Id) ([]*models.ModelEndpointAlert, error) {
	ret := _m.Called(modelId)

	var r0 []*models.ModelEndpointAlert
	if rf, ok := ret.Get(0).(func(models.Id) []*models.ModelEndpointAlert); ok {
		r0 = rf(modelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ModelEndpointAlert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Id) error); ok {
		r1 = rf(modelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateModelEndpointAlert provides a mock function with given fields: alert
func (_m *AlertStorage) UpdateModelEndpointAlert(alert *models.ModelEndpointAlert) error {
	ret := _m.Called(alert)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.ModelEndpointAlert) error); ok {
		r0 = rf(alert)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
