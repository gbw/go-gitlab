// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: AvatarRequestsServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=avatar_mock.go -package=testing gitlab.com/gitlab-org/api/client-go AvatarRequestsServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockAvatarRequestsServiceInterface is a mock of AvatarRequestsServiceInterface interface.
type MockAvatarRequestsServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAvatarRequestsServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockAvatarRequestsServiceInterfaceMockRecorder is the mock recorder for MockAvatarRequestsServiceInterface.
type MockAvatarRequestsServiceInterfaceMockRecorder struct {
	mock *MockAvatarRequestsServiceInterface
}

// NewMockAvatarRequestsServiceInterface creates a new mock instance.
func NewMockAvatarRequestsServiceInterface(ctrl *gomock.Controller) *MockAvatarRequestsServiceInterface {
	mock := &MockAvatarRequestsServiceInterface{ctrl: ctrl}
	mock.recorder = &MockAvatarRequestsServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvatarRequestsServiceInterface) EXPECT() *MockAvatarRequestsServiceInterfaceMockRecorder {
	return m.recorder
}

// GetAvatar mocks base method.
func (m *MockAvatarRequestsServiceInterface) GetAvatar(opt *gitlab.GetAvatarOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Avatar, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAvatar", varargs...)
	ret0, _ := ret[0].(*gitlab.Avatar)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAvatar indicates an expected call of GetAvatar.
func (mr *MockAvatarRequestsServiceInterfaceMockRecorder) GetAvatar(opt any, options ...any) *MockAvatarRequestsServiceInterfaceGetAvatarCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvatar", reflect.TypeOf((*MockAvatarRequestsServiceInterface)(nil).GetAvatar), varargs...)
	return &MockAvatarRequestsServiceInterfaceGetAvatarCall{Call: call}
}

// MockAvatarRequestsServiceInterfaceGetAvatarCall wrap *gomock.Call
type MockAvatarRequestsServiceInterfaceGetAvatarCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAvatarRequestsServiceInterfaceGetAvatarCall) Return(arg0 *gitlab.Avatar, arg1 *gitlab.Response, arg2 error) *MockAvatarRequestsServiceInterfaceGetAvatarCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAvatarRequestsServiceInterfaceGetAvatarCall) Do(f func(*gitlab.GetAvatarOptions, ...gitlab.RequestOptionFunc) (*gitlab.Avatar, *gitlab.Response, error)) *MockAvatarRequestsServiceInterfaceGetAvatarCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAvatarRequestsServiceInterfaceGetAvatarCall) DoAndReturn(f func(*gitlab.GetAvatarOptions, ...gitlab.RequestOptionFunc) (*gitlab.Avatar, *gitlab.Response, error)) *MockAvatarRequestsServiceInterfaceGetAvatarCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
