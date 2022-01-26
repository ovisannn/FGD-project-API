// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	comments "disspace/business/comments"

	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, data, id
func (_m *UseCase) Create(ctx context.Context, data *comments.Domain, id string) (comments.Domain, error) {
	ret := _m.Called(ctx, data, id)

	var r0 comments.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *comments.Domain, string) comments.Domain); ok {
		r0 = rf(ctx, data, id)
	} else {
		r0 = ret.Get(0).(comments.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *comments.Domain, string) error); ok {
		r1 = rf(ctx, data, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id, threadId
func (_m *UseCase) Delete(ctx context.Context, id string, threadId string) error {
	ret := _m.Called(ctx, id, threadId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, id, threadId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllInThread provides a mock function with given fields: ctx, threadId, parentId
func (_m *UseCase) GetAllInThread(ctx context.Context, threadId string, parentId string) ([]comments.Domain, error) {
	ret := _m.Called(ctx, threadId, parentId)

	var r0 []comments.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []comments.Domain); ok {
		r0 = rf(ctx, threadId, parentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, threadId, parentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *UseCase) GetByID(ctx context.Context, id string) (comments.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 comments.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) comments.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(comments.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, q, sort
func (_m *UseCase) Search(ctx context.Context, q string, sort string) ([]comments.Domain, error) {
	ret := _m.Called(ctx, q, sort)

	var r0 []comments.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []comments.Domain); ok {
		r0 = rf(ctx, q, sort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, q, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
