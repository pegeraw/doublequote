// Code generated by mockery 2.9.4. DO NOT EDIT.

package mock

import (
	context "context"
	dq "doublequote"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *UserService) CreateUser(ctx context.Context, user *dq.User) (*dq.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *dq.User
	if rf, ok := ret.Get(0).(func(context.Context, *dq.User) *dq.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dq.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dq.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *UserService) DeleteUser(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUser provides a mock function with given fields: ctx, filter, include
func (_m *UserService) FindUser(ctx context.Context, filter dq.UserFilter, include dq.UserInclude) (*dq.User, error) {
	ret := _m.Called(ctx, filter, include)

	var r0 *dq.User
	if rf, ok := ret.Get(0).(func(context.Context, dq.UserFilter, dq.UserInclude) *dq.User); ok {
		r0 = rf(ctx, filter, include)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dq.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dq.UserFilter, dq.UserInclude) error); ok {
		r1 = rf(ctx, filter, include)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByID provides a mock function with given fields: ctx, id, include
func (_m *UserService) FindUserByID(ctx context.Context, id int, include dq.UserInclude) (*dq.User, error) {
	ret := _m.Called(ctx, id, include)

	var r0 *dq.User
	if rf, ok := ret.Get(0).(func(context.Context, int, dq.UserInclude) *dq.User); ok {
		r0 = rf(ctx, id, include)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dq.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, dq.UserInclude) error); ok {
		r1 = rf(ctx, id, include)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUsers provides a mock function with given fields: ctx, filter, include
func (_m *UserService) FindUsers(ctx context.Context, filter dq.UserFilter, include dq.UserInclude) ([]*dq.User, int, error) {
	ret := _m.Called(ctx, filter, include)

	var r0 []*dq.User
	if rf, ok := ret.Get(0).(func(context.Context, dq.UserFilter, dq.UserInclude) []*dq.User); ok {
		r0 = rf(ctx, filter, include)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dq.User)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, dq.UserFilter, dq.UserInclude) int); ok {
		r1 = rf(ctx, filter, include)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, dq.UserFilter, dq.UserInclude) error); ok {
		r2 = rf(ctx, filter, include)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateUser provides a mock function with given fields: ctx, id, upd
func (_m *UserService) UpdateUser(ctx context.Context, id int, upd dq.UserUpdate) (*dq.User, error) {
	ret := _m.Called(ctx, id, upd)

	var r0 *dq.User
	if rf, ok := ret.Get(0).(func(context.Context, int, dq.UserUpdate) *dq.User); ok {
		r0 = rf(ctx, id, upd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dq.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, dq.UserUpdate) error); ok {
		r1 = rf(ctx, id, upd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
