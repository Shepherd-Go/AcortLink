// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	models "acortlink/core/domain/models"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ShortenRepoPostgres is an autogenerated mock type for the ShortenRepoPostgres type
type ShortenRepoPostgres struct {
	mock.Mock
}

// CreateShorten provides a mock function with given fields: ctx, url
func (_m *ShortenRepoPostgres) CreateShorten(ctx context.Context, url models.URL) error {
	ret := _m.Called(ctx, url)

	if len(ret) == 0 {
		panic("no return value specified for CreateShorten")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.URL) error); ok {
		r0 = rf(ctx, url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchUrl provides a mock function with given fields: ctx, path
func (_m *ShortenRepoPostgres) SearchUrl(ctx context.Context, path string) (models.URL, error) {
	ret := _m.Called(ctx, path)

	if len(ret) == 0 {
		panic("no return value specified for SearchUrl")
	}

	var r0 models.URL
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.URL, error)); ok {
		return rf(ctx, path)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.URL); ok {
		r0 = rf(ctx, path)
	} else {
		r0 = ret.Get(0).(models.URL)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewShortenRepoPostgres creates a new instance of ShortenRepoPostgres. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewShortenRepoPostgres(t interface {
	mock.TestingT
	Cleanup(func())
}) *ShortenRepoPostgres {
	mock := &ShortenRepoPostgres{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
