// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: SystemHooksServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=system_hooks_mock.go -package=testing gitlab.com/gitlab-org/api/client-go SystemHooksServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockSystemHooksServiceInterface is a mock of SystemHooksServiceInterface interface.
type MockSystemHooksServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSystemHooksServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockSystemHooksServiceInterfaceMockRecorder is the mock recorder for MockSystemHooksServiceInterface.
type MockSystemHooksServiceInterfaceMockRecorder struct {
	mock *MockSystemHooksServiceInterface
}

// NewMockSystemHooksServiceInterface creates a new mock instance.
func NewMockSystemHooksServiceInterface(ctrl *gomock.Controller) *MockSystemHooksServiceInterface {
	mock := &MockSystemHooksServiceInterface{ctrl: ctrl}
	mock.recorder = &MockSystemHooksServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemHooksServiceInterface) EXPECT() *MockSystemHooksServiceInterfaceMockRecorder {
	return m.recorder
}

// AddHook mocks base method.
func (m *MockSystemHooksServiceInterface) AddHook(opt *gitlab.AddHookOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Hook, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddHook", varargs...)
	ret0, _ := ret[0].(*gitlab.Hook)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddHook indicates an expected call of AddHook.
func (mr *MockSystemHooksServiceInterfaceMockRecorder) AddHook(opt any, options ...any) *MockSystemHooksServiceInterfaceAddHookCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHook", reflect.TypeOf((*MockSystemHooksServiceInterface)(nil).AddHook), varargs...)
	return &MockSystemHooksServiceInterfaceAddHookCall{Call: call}
}

// MockSystemHooksServiceInterfaceAddHookCall wrap *gomock.Call
type MockSystemHooksServiceInterfaceAddHookCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSystemHooksServiceInterfaceAddHookCall) Return(arg0 *gitlab.Hook, arg1 *gitlab.Response, arg2 error) *MockSystemHooksServiceInterfaceAddHookCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSystemHooksServiceInterfaceAddHookCall) Do(f func(*gitlab.AddHookOptions, ...gitlab.RequestOptionFunc) (*gitlab.Hook, *gitlab.Response, error)) *MockSystemHooksServiceInterfaceAddHookCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSystemHooksServiceInterfaceAddHookCall) DoAndReturn(f func(*gitlab.AddHookOptions, ...gitlab.RequestOptionFunc) (*gitlab.Hook, *gitlab.Response, error)) *MockSystemHooksServiceInterfaceAddHookCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteHook mocks base method.
func (m *MockSystemHooksServiceInterface) DeleteHook(hook int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{hook}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteHook", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteHook indicates an expected call of DeleteHook.
func (mr *MockSystemHooksServiceInterfaceMockRecorder) DeleteHook(hook any, options ...any) *MockSystemHooksServiceInterfaceDeleteHookCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{hook}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHook", reflect.TypeOf((*MockSystemHooksServiceInterface)(nil).DeleteHook), varargs...)
	return &MockSystemHooksServiceInterfaceDeleteHookCall{Call: call}
}

// MockSystemHooksServiceInterfaceDeleteHookCall wrap *gomock.Call
type MockSystemHooksServiceInterfaceDeleteHookCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSystemHooksServiceInterfaceDeleteHookCall) Return(arg0 *gitlab.Response, arg1 error) *MockSystemHooksServiceInterfaceDeleteHookCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSystemHooksServiceInterfaceDeleteHookCall) Do(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockSystemHooksServiceInterfaceDeleteHookCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSystemHooksServiceInterfaceDeleteHookCall) DoAndReturn(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockSystemHooksServiceInterfaceDeleteHookCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetHook mocks base method.
func (m *MockSystemHooksServiceInterface) GetHook(hook int, options ...gitlab.RequestOptionFunc) (*gitlab.Hook, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{hook}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetHook", varargs...)
	ret0, _ := ret[0].(*gitlab.Hook)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHook indicates an expected call of GetHook.
func (mr *MockSystemHooksServiceInterfaceMockRecorder) GetHook(hook any, options ...any) *MockSystemHooksServiceInterfaceGetHookCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{hook}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHook", reflect.TypeOf((*MockSystemHooksServiceInterface)(nil).GetHook), varargs...)
	return &MockSystemHooksServiceInterfaceGetHookCall{Call: call}
}

// MockSystemHooksServiceInterfaceGetHookCall wrap *gomock.Call
type MockSystemHooksServiceInterfaceGetHookCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSystemHooksServiceInterfaceGetHookCall) Return(arg0 *gitlab.Hook, arg1 *gitlab.Response, arg2 error) *MockSystemHooksServiceInterfaceGetHookCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSystemHooksServiceInterfaceGetHookCall) Do(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Hook, *gitlab.Response, error)) *MockSystemHooksServiceInterfaceGetHookCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSystemHooksServiceInterfaceGetHookCall) DoAndReturn(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Hook, *gitlab.Response, error)) *MockSystemHooksServiceInterfaceGetHookCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListHooks mocks base method.
func (m *MockSystemHooksServiceInterface) ListHooks(options ...gitlab.RequestOptionFunc) ([]*gitlab.Hook, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListHooks", varargs...)
	ret0, _ := ret[0].([]*gitlab.Hook)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListHooks indicates an expected call of ListHooks.
func (mr *MockSystemHooksServiceInterfaceMockRecorder) ListHooks(options ...any) *MockSystemHooksServiceInterfaceListHooksCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHooks", reflect.TypeOf((*MockSystemHooksServiceInterface)(nil).ListHooks), options...)
	return &MockSystemHooksServiceInterfaceListHooksCall{Call: call}
}

// MockSystemHooksServiceInterfaceListHooksCall wrap *gomock.Call
type MockSystemHooksServiceInterfaceListHooksCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSystemHooksServiceInterfaceListHooksCall) Return(arg0 []*gitlab.Hook, arg1 *gitlab.Response, arg2 error) *MockSystemHooksServiceInterfaceListHooksCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSystemHooksServiceInterfaceListHooksCall) Do(f func(...gitlab.RequestOptionFunc) ([]*gitlab.Hook, *gitlab.Response, error)) *MockSystemHooksServiceInterfaceListHooksCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSystemHooksServiceInterfaceListHooksCall) DoAndReturn(f func(...gitlab.RequestOptionFunc) ([]*gitlab.Hook, *gitlab.Response, error)) *MockSystemHooksServiceInterfaceListHooksCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// TestHook mocks base method.
func (m *MockSystemHooksServiceInterface) TestHook(hook int, options ...gitlab.RequestOptionFunc) (*gitlab.HookEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{hook}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TestHook", varargs...)
	ret0, _ := ret[0].(*gitlab.HookEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TestHook indicates an expected call of TestHook.
func (mr *MockSystemHooksServiceInterfaceMockRecorder) TestHook(hook any, options ...any) *MockSystemHooksServiceInterfaceTestHookCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{hook}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestHook", reflect.TypeOf((*MockSystemHooksServiceInterface)(nil).TestHook), varargs...)
	return &MockSystemHooksServiceInterfaceTestHookCall{Call: call}
}

// MockSystemHooksServiceInterfaceTestHookCall wrap *gomock.Call
type MockSystemHooksServiceInterfaceTestHookCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSystemHooksServiceInterfaceTestHookCall) Return(arg0 *gitlab.HookEvent, arg1 *gitlab.Response, arg2 error) *MockSystemHooksServiceInterfaceTestHookCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSystemHooksServiceInterfaceTestHookCall) Do(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.HookEvent, *gitlab.Response, error)) *MockSystemHooksServiceInterfaceTestHookCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSystemHooksServiceInterfaceTestHookCall) DoAndReturn(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.HookEvent, *gitlab.Response, error)) *MockSystemHooksServiceInterfaceTestHookCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
