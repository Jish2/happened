// Code generated by mockery v2.49.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// S3PresignClient is an autogenerated mock type for the S3PresignClient type
type S3PresignClient struct {
	mock.Mock
}

// PresignPutObject provides a mock function with given fields: ctx, params, optFns
func (_m *S3PresignClient) PresignPutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PresignPutObject")
	}

	var r0 *v4.PresignedHTTPRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.PutObjectInput, ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.PutObjectInput, ...func(*s3.PresignOptions)) *v4.PresignedHTTPRequest); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v4.PresignedHTTPRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.PutObjectInput, ...func(*s3.PresignOptions)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewS3PresignClient creates a new instance of S3PresignClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewS3PresignClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *S3PresignClient {
	mock := &S3PresignClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
