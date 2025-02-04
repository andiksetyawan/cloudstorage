// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	cloudstorage "cloudstorage"
	context "context"

	io "io"

	mock "github.com/stretchr/testify/mock"

	time "time"

	url "net/url"
)

// CloudStorage is an autogenerated mock type for the CloudStorage type
type CloudStorage struct {
	mock.Mock
}

// Download provides a mock function with given fields: ctx, bucketName, objectName
func (_m *CloudStorage) Download(ctx context.Context, bucketName string, objectName string) (io.Reader, error) {
	ret := _m.Called(ctx, bucketName, objectName)

	if len(ret) == 0 {
		panic("no return value specified for Download")
	}

	var r0 io.Reader
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (io.Reader, error)); ok {
		return rf(ctx, bucketName, objectName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) io.Reader); ok {
		r0 = rf(ctx, bucketName, objectName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, bucketName, objectName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Presign provides a mock function with given fields: ctx, bucketName, objectName, expires
func (_m *CloudStorage) Presign(ctx context.Context, bucketName string, objectName string, expires time.Duration) (*url.URL, error) {
	ret := _m.Called(ctx, bucketName, objectName, expires)

	if len(ret) == 0 {
		panic("no return value specified for Presign")
	}

	var r0 *url.URL
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Duration) (*url.URL, error)); ok {
		return rf(ctx, bucketName, objectName, expires)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Duration) *url.URL); ok {
		r0 = rf(ctx, bucketName, objectName, expires)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, time.Duration) error); ok {
		r1 = rf(ctx, bucketName, objectName, expires)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upload provides a mock function with given fields: ctx, bucketName, file
func (_m *CloudStorage) Upload(ctx context.Context, bucketName string, file cloudstorage.UploadOption) (cloudstorage.FileInfo, error) {
	ret := _m.Called(ctx, bucketName, file)

	if len(ret) == 0 {
		panic("no return value specified for Upload")
	}

	var r0 cloudstorage.FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, cloudstorage.UploadOption) (cloudstorage.FileInfo, error)); ok {
		return rf(ctx, bucketName, file)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, cloudstorage.UploadOption) cloudstorage.FileInfo); ok {
		r0 = rf(ctx, bucketName, file)
	} else {
		r0 = ret.Get(0).(cloudstorage.FileInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, cloudstorage.UploadOption) error); ok {
		r1 = rf(ctx, bucketName, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadPresign provides a mock function with given fields: ctx, bucketName, expires, file
func (_m *CloudStorage) UploadPresign(ctx context.Context, bucketName string, expires time.Duration, file cloudstorage.UploadOption) (*url.URL, error) {
	ret := _m.Called(ctx, bucketName, expires, file)

	if len(ret) == 0 {
		panic("no return value specified for UploadPresign")
	}

	var r0 *url.URL
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration, cloudstorage.UploadOption) (*url.URL, error)); ok {
		return rf(ctx, bucketName, expires, file)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration, cloudstorage.UploadOption) *url.URL); ok {
		r0 = rf(ctx, bucketName, expires, file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, time.Duration, cloudstorage.UploadOption) error); ok {
		r1 = rf(ctx, bucketName, expires, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCloudStorage creates a new instance of CloudStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCloudStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *CloudStorage {
	mock := &CloudStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
