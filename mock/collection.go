// Code generated by mockery 2.9.4. DO NOT EDIT.

package mock

import (
	context "context"
	dq "doublequote"

	mock "github.com/stretchr/testify/mock"
)

// CollectionService is an autogenerated mock type for the CollectionService type
type CollectionService struct {
	mock.Mock
}

// CreateCollection provides a mock function with given fields: ctx, col
func (_m *CollectionService) CreateCollection(ctx context.Context, col *dq.Collection) (*dq.Collection, error) {
	ret := _m.Called(ctx, col)

	var r0 *dq.Collection
	if rf, ok := ret.Get(0).(func(context.Context, *dq.Collection) *dq.Collection); ok {
		r0 = rf(ctx, col)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dq.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dq.Collection) error); ok {
		r1 = rf(ctx, col)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCollection provides a mock function with given fields: ctx, id
func (_m *CollectionService) DeleteCollection(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindCollection provides a mock function with given fields: ctx, filter, include
func (_m *CollectionService) FindCollection(ctx context.Context, filter dq.CollectionFilter, include dq.CollectionInclude) (*dq.Collection, error) {
	ret := _m.Called(ctx, filter, include)

	var r0 *dq.Collection
	if rf, ok := ret.Get(0).(func(context.Context, dq.CollectionFilter, dq.CollectionInclude) *dq.Collection); ok {
		r0 = rf(ctx, filter, include)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dq.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dq.CollectionFilter, dq.CollectionInclude) error); ok {
		r1 = rf(ctx, filter, include)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCollectionByID provides a mock function with given fields: ctx, id, include
func (_m *CollectionService) FindCollectionByID(ctx context.Context, id int, include dq.CollectionInclude) (*dq.Collection, error) {
	ret := _m.Called(ctx, id, include)

	var r0 *dq.Collection
	if rf, ok := ret.Get(0).(func(context.Context, int, dq.CollectionInclude) *dq.Collection); ok {
		r0 = rf(ctx, id, include)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dq.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, dq.CollectionInclude) error); ok {
		r1 = rf(ctx, id, include)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCollections provides a mock function with given fields: ctx, filter, include
func (_m *CollectionService) FindCollections(ctx context.Context, filter dq.CollectionFilter, include dq.CollectionInclude) ([]*dq.Collection, int, error) {
	ret := _m.Called(ctx, filter, include)

	var r0 []*dq.Collection
	if rf, ok := ret.Get(0).(func(context.Context, dq.CollectionFilter, dq.CollectionInclude) []*dq.Collection); ok {
		r0 = rf(ctx, filter, include)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dq.Collection)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, dq.CollectionFilter, dq.CollectionInclude) int); ok {
		r1 = rf(ctx, filter, include)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, dq.CollectionFilter, dq.CollectionInclude) error); ok {
		r2 = rf(ctx, filter, include)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateCollection provides a mock function with given fields: ctx, id, upd
func (_m *CollectionService) UpdateCollection(ctx context.Context, id int, upd dq.CollectionUpdate) (*dq.Collection, error) {
	ret := _m.Called(ctx, id, upd)

	var r0 *dq.Collection
	if rf, ok := ret.Get(0).(func(context.Context, int, dq.CollectionUpdate) *dq.Collection); ok {
		r0 = rf(ctx, id, upd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dq.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, dq.CollectionUpdate) error); ok {
		r1 = rf(ctx, id, upd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
