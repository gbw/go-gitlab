// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: PagesServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=pages_mock.go -package=testing gitlab.com/gitlab-org/api/client-go PagesServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockPagesServiceInterface is a mock of PagesServiceInterface interface.
type MockPagesServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPagesServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockPagesServiceInterfaceMockRecorder is the mock recorder for MockPagesServiceInterface.
type MockPagesServiceInterfaceMockRecorder struct {
	mock *MockPagesServiceInterface
}

// NewMockPagesServiceInterface creates a new mock instance.
func NewMockPagesServiceInterface(ctrl *gomock.Controller) *MockPagesServiceInterface {
	mock := &MockPagesServiceInterface{ctrl: ctrl}
	mock.recorder = &MockPagesServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPagesServiceInterface) EXPECT() *MockPagesServiceInterfaceMockRecorder {
	return m.recorder
}

// GetPages mocks base method.
func (m *MockPagesServiceInterface) GetPages(gid any, options ...gitlab.RequestOptionFunc) (*gitlab.Pages, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPages", varargs...)
	ret0, _ := ret[0].(*gitlab.Pages)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPages indicates an expected call of GetPages.
func (mr *MockPagesServiceInterfaceMockRecorder) GetPages(gid any, options ...any) *MockPagesServiceInterfaceGetPagesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPages", reflect.TypeOf((*MockPagesServiceInterface)(nil).GetPages), varargs...)
	return &MockPagesServiceInterfaceGetPagesCall{Call: call}
}

// MockPagesServiceInterfaceGetPagesCall wrap *gomock.Call
type MockPagesServiceInterfaceGetPagesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPagesServiceInterfaceGetPagesCall) Return(arg0 *gitlab.Pages, arg1 *gitlab.Response, arg2 error) *MockPagesServiceInterfaceGetPagesCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPagesServiceInterfaceGetPagesCall) Do(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.Pages, *gitlab.Response, error)) *MockPagesServiceInterfaceGetPagesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPagesServiceInterfaceGetPagesCall) DoAndReturn(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.Pages, *gitlab.Response, error)) *MockPagesServiceInterfaceGetPagesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UnpublishPages mocks base method.
func (m *MockPagesServiceInterface) UnpublishPages(gid any, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnpublishPages", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpublishPages indicates an expected call of UnpublishPages.
func (mr *MockPagesServiceInterfaceMockRecorder) UnpublishPages(gid any, options ...any) *MockPagesServiceInterfaceUnpublishPagesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpublishPages", reflect.TypeOf((*MockPagesServiceInterface)(nil).UnpublishPages), varargs...)
	return &MockPagesServiceInterfaceUnpublishPagesCall{Call: call}
}

// MockPagesServiceInterfaceUnpublishPagesCall wrap *gomock.Call
type MockPagesServiceInterfaceUnpublishPagesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPagesServiceInterfaceUnpublishPagesCall) Return(arg0 *gitlab.Response, arg1 error) *MockPagesServiceInterfaceUnpublishPagesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPagesServiceInterfaceUnpublishPagesCall) Do(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockPagesServiceInterfaceUnpublishPagesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPagesServiceInterfaceUnpublishPagesCall) DoAndReturn(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockPagesServiceInterfaceUnpublishPagesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdatePages mocks base method.
func (m *MockPagesServiceInterface) UpdatePages(pid any, opt gitlab.UpdatePagesOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Pages, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdatePages", varargs...)
	ret0, _ := ret[0].(*gitlab.Pages)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdatePages indicates an expected call of UpdatePages.
func (mr *MockPagesServiceInterfaceMockRecorder) UpdatePages(pid, opt any, options ...any) *MockPagesServiceInterfaceUpdatePagesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePages", reflect.TypeOf((*MockPagesServiceInterface)(nil).UpdatePages), varargs...)
	return &MockPagesServiceInterfaceUpdatePagesCall{Call: call}
}

// MockPagesServiceInterfaceUpdatePagesCall wrap *gomock.Call
type MockPagesServiceInterfaceUpdatePagesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPagesServiceInterfaceUpdatePagesCall) Return(arg0 *gitlab.Pages, arg1 *gitlab.Response, arg2 error) *MockPagesServiceInterfaceUpdatePagesCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPagesServiceInterfaceUpdatePagesCall) Do(f func(any, gitlab.UpdatePagesOptions, ...gitlab.RequestOptionFunc) (*gitlab.Pages, *gitlab.Response, error)) *MockPagesServiceInterfaceUpdatePagesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPagesServiceInterfaceUpdatePagesCall) DoAndReturn(f func(any, gitlab.UpdatePagesOptions, ...gitlab.RequestOptionFunc) (*gitlab.Pages, *gitlab.Response, error)) *MockPagesServiceInterfaceUpdatePagesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
