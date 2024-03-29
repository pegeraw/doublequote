// Code generated by mockery 2.9.4. DO NOT EDIT.

package mock

import (
	context "context"
	dq "doublequote"

	mock "github.com/stretchr/testify/mock"
)

// FeedService is an autogenerated mock type for the FeedService type
type FeedService struct {
	mock.Mock
}

// CreateFeed provides a mock function with given fields: ctx, feed
func (_m *FeedService) CreateFeed(ctx context.Context, feed *dq.Feed) (*dq.Feed, error) {
	ret := _m.Called(ctx, feed)

	var r0 *dq.Feed
	if rf, ok := ret.Get(0).(func(context.Context, *dq.Feed) *dq.Feed); ok {
		r0 = rf(ctx, feed)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dq.Feed)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dq.Feed) error); ok {
		r1 = rf(ctx, feed)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFeed provides a mock function with given fields: ctx, id
func (_m *FeedService) DeleteFeed(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindFeed provides a mock function with given fields: ctx, filter, include
func (_m *FeedService) FindFeed(ctx context.Context, filter dq.FeedFilter, include dq.FeedInclude) (*dq.Feed, error) {
	ret := _m.Called(ctx, filter, include)

	var r0 *dq.Feed
	if rf, ok := ret.Get(0).(func(context.Context, dq.FeedFilter, dq.FeedInclude) *dq.Feed); ok {
		r0 = rf(ctx, filter, include)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dq.Feed)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dq.FeedFilter, dq.FeedInclude) error); ok {
		r1 = rf(ctx, filter, include)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFeedByID provides a mock function with given fields: ctx, id, include
func (_m *FeedService) FindFeedByID(ctx context.Context, id int, include dq.FeedInclude) (*dq.Feed, error) {
	ret := _m.Called(ctx, id, include)

	var r0 *dq.Feed
	if rf, ok := ret.Get(0).(func(context.Context, int, dq.FeedInclude) *dq.Feed); ok {
		r0 = rf(ctx, id, include)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dq.Feed)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, dq.FeedInclude) error); ok {
		r1 = rf(ctx, id, include)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFeeds provides a mock function with given fields: ctx, filter, include
func (_m *FeedService) FindFeeds(ctx context.Context, filter dq.FeedFilter, include dq.FeedInclude) ([]*dq.Feed, int, error) {
	ret := _m.Called(ctx, filter, include)

	var r0 []*dq.Feed
	if rf, ok := ret.Get(0).(func(context.Context, dq.FeedFilter, dq.FeedInclude) []*dq.Feed); ok {
		r0 = rf(ctx, filter, include)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dq.Feed)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, dq.FeedFilter, dq.FeedInclude) int); ok {
		r1 = rf(ctx, filter, include)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, dq.FeedFilter, dq.FeedInclude) error); ok {
		r2 = rf(ctx, filter, include)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateFeed provides a mock function with given fields: ctx, id, upd
func (_m *FeedService) UpdateFeed(ctx context.Context, id int, upd dq.FeedUpdate) (*dq.Feed, error) {
	ret := _m.Called(ctx, id, upd)

	var r0 *dq.Feed
	if rf, ok := ret.Get(0).(func(context.Context, int, dq.FeedUpdate) *dq.Feed); ok {
		r0 = rf(ctx, id, upd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dq.Feed)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, dq.FeedUpdate) error); ok {
		r1 = rf(ctx, id, upd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
