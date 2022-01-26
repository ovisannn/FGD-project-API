// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	reports "disspace/business/reports"

	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, data, id
func (_m *UseCase) Create(ctx context.Context, data *reports.Domain, id string) error {
	ret := _m.Called(ctx, data, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *reports.Domain, string) error); ok {
		r0 = rf(ctx, data, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, sort
func (_m *UseCase) GetAll(ctx context.Context, sort string) ([]reports.Domain, error) {
	ret := _m.Called(ctx, sort)

	var r0 []reports.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) []reports.Domain); ok {
		r0 = rf(ctx, sort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reports.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommentReport provides a mock function with given fields: ctx, sort, q
func (_m *UseCase) GetCommentReport(ctx context.Context, sort string, q string) ([]reports.Domain, error) {
	ret := _m.Called(ctx, sort, q)

	var r0 []reports.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []reports.Domain); ok {
		r0 = rf(ctx, sort, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reports.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, sort, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserReport provides a mock function with given fields: ctx, sort, q
func (_m *UseCase) GetUserReport(ctx context.Context, sort string, q string) ([]reports.Domain, error) {
	ret := _m.Called(ctx, sort, q)

	var r0 []reports.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []reports.Domain); ok {
		r0 = rf(ctx, sort, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reports.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, sort, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
