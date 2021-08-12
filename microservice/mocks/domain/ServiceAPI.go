// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/bibit/microservice/model"
)

// ServiceAPI is an autogenerated mock type for the ServiceAPI type
type ServiceAPI struct {
	mock.Mock
}

// GetListMovie provides a mock function with given fields: ctx, req
func (_m *ServiceAPI) GetListMovie(ctx context.Context, req *model.GetMovieRequest) (*model.GetMovieResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *model.GetMovieResponse
	if rf, ok := ret.Get(0).(func(context.Context, *model.GetMovieRequest) *model.GetMovieResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GetMovieResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.GetMovieRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
