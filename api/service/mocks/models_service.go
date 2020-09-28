// Code generated by mockery v2.0.4. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/gojek/merlin/models"
	mock "github.com/stretchr/testify/mock"
)

// ModelsService is an autogenerated mock type for the ModelsService type
type ModelsService struct {
	mock.Mock
}

// FindById provides a mock function with given fields: ctx, modelId
func (_m *ModelsService) FindById(ctx context.Context, modelId models.Id) (*models.Model, error) {
	ret := _m.Called(ctx, modelId)

	var r0 *models.Model
	if rf, ok := ret.Get(0).(func(context.Context, models.Id) *models.Model); ok {
		r0 = rf(ctx, modelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Model)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Id) error); ok {
		r1 = rf(ctx, modelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListModels provides a mock function with given fields: ctx, projectId, name
func (_m *ModelsService) ListModels(ctx context.Context, projectId models.Id, name string) ([]*models.Model, error) {
	ret := _m.Called(ctx, projectId, name)

	var r0 []*models.Model
	if rf, ok := ret.Get(0).(func(context.Context, models.Id, string) []*models.Model); ok {
		r0 = rf(ctx, projectId, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Model)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Id, string) error); ok {
		r1 = rf(ctx, projectId, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, model
func (_m *ModelsService) Save(ctx context.Context, model *models.Model) (*models.Model, error) {
	ret := _m.Called(ctx, model)

	var r0 *models.Model
	if rf, ok := ret.Get(0).(func(context.Context, *models.Model) *models.Model); ok {
		r0 = rf(ctx, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Model)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Model) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, model
func (_m *ModelsService) Update(ctx context.Context, model *models.Model) (*models.Model, error) {
	ret := _m.Called(ctx, model)

	var r0 *models.Model
	if rf, ok := ret.Get(0).(func(context.Context, *models.Model) *models.Model); ok {
		r0 = rf(ctx, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Model)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Model) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
