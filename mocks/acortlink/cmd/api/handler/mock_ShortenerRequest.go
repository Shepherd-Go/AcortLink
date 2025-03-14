// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// ShortenerRequest is an autogenerated mock type for the ShortenerRequest type
type ShortenerRequest struct {
	mock.Mock
}

// CreateShortURL provides a mock function with given fields: c
func (_m *ShortenerRequest) CreateShortURL(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for CreateShortURL")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchOriginalUrl provides a mock function with given fields: c
func (_m *ShortenerRequest) SearchOriginalUrl(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for SearchOriginalUrl")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewShortenerRequest creates a new instance of ShortenerRequest. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewShortenerRequest(t interface {
	mock.TestingT
	Cleanup(func())
}) *ShortenerRequest {
	mock := &ShortenerRequest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
