// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: InstanceVariablesServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=instance_variables_mock.go -package=testing gitlab.com/gitlab-org/api/client-go InstanceVariablesServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockInstanceVariablesServiceInterface is a mock of InstanceVariablesServiceInterface interface.
type MockInstanceVariablesServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceVariablesServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockInstanceVariablesServiceInterfaceMockRecorder is the mock recorder for MockInstanceVariablesServiceInterface.
type MockInstanceVariablesServiceInterfaceMockRecorder struct {
	mock *MockInstanceVariablesServiceInterface
}

// NewMockInstanceVariablesServiceInterface creates a new mock instance.
func NewMockInstanceVariablesServiceInterface(ctrl *gomock.Controller) *MockInstanceVariablesServiceInterface {
	mock := &MockInstanceVariablesServiceInterface{ctrl: ctrl}
	mock.recorder = &MockInstanceVariablesServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstanceVariablesServiceInterface) EXPECT() *MockInstanceVariablesServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateVariable mocks base method.
func (m *MockInstanceVariablesServiceInterface) CreateVariable(opt *gitlab.CreateInstanceVariableOptions, options ...gitlab.RequestOptionFunc) (*gitlab.InstanceVariable, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateVariable", varargs...)
	ret0, _ := ret[0].(*gitlab.InstanceVariable)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateVariable indicates an expected call of CreateVariable.
func (mr *MockInstanceVariablesServiceInterfaceMockRecorder) CreateVariable(opt any, options ...any) *MockInstanceVariablesServiceInterfaceCreateVariableCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVariable", reflect.TypeOf((*MockInstanceVariablesServiceInterface)(nil).CreateVariable), varargs...)
	return &MockInstanceVariablesServiceInterfaceCreateVariableCall{Call: call}
}

// MockInstanceVariablesServiceInterfaceCreateVariableCall wrap *gomock.Call
type MockInstanceVariablesServiceInterfaceCreateVariableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInstanceVariablesServiceInterfaceCreateVariableCall) Return(arg0 *gitlab.InstanceVariable, arg1 *gitlab.Response, arg2 error) *MockInstanceVariablesServiceInterfaceCreateVariableCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInstanceVariablesServiceInterfaceCreateVariableCall) Do(f func(*gitlab.CreateInstanceVariableOptions, ...gitlab.RequestOptionFunc) (*gitlab.InstanceVariable, *gitlab.Response, error)) *MockInstanceVariablesServiceInterfaceCreateVariableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInstanceVariablesServiceInterfaceCreateVariableCall) DoAndReturn(f func(*gitlab.CreateInstanceVariableOptions, ...gitlab.RequestOptionFunc) (*gitlab.InstanceVariable, *gitlab.Response, error)) *MockInstanceVariablesServiceInterfaceCreateVariableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetVariable mocks base method.
func (m *MockInstanceVariablesServiceInterface) GetVariable(key string, options ...gitlab.RequestOptionFunc) (*gitlab.InstanceVariable, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{key}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVariable", varargs...)
	ret0, _ := ret[0].(*gitlab.InstanceVariable)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetVariable indicates an expected call of GetVariable.
func (mr *MockInstanceVariablesServiceInterfaceMockRecorder) GetVariable(key any, options ...any) *MockInstanceVariablesServiceInterfaceGetVariableCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{key}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVariable", reflect.TypeOf((*MockInstanceVariablesServiceInterface)(nil).GetVariable), varargs...)
	return &MockInstanceVariablesServiceInterfaceGetVariableCall{Call: call}
}

// MockInstanceVariablesServiceInterfaceGetVariableCall wrap *gomock.Call
type MockInstanceVariablesServiceInterfaceGetVariableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInstanceVariablesServiceInterfaceGetVariableCall) Return(arg0 *gitlab.InstanceVariable, arg1 *gitlab.Response, arg2 error) *MockInstanceVariablesServiceInterfaceGetVariableCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInstanceVariablesServiceInterfaceGetVariableCall) Do(f func(string, ...gitlab.RequestOptionFunc) (*gitlab.InstanceVariable, *gitlab.Response, error)) *MockInstanceVariablesServiceInterfaceGetVariableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInstanceVariablesServiceInterfaceGetVariableCall) DoAndReturn(f func(string, ...gitlab.RequestOptionFunc) (*gitlab.InstanceVariable, *gitlab.Response, error)) *MockInstanceVariablesServiceInterfaceGetVariableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListVariables mocks base method.
func (m *MockInstanceVariablesServiceInterface) ListVariables(opt *gitlab.ListInstanceVariablesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.InstanceVariable, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVariables", varargs...)
	ret0, _ := ret[0].([]*gitlab.InstanceVariable)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListVariables indicates an expected call of ListVariables.
func (mr *MockInstanceVariablesServiceInterfaceMockRecorder) ListVariables(opt any, options ...any) *MockInstanceVariablesServiceInterfaceListVariablesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVariables", reflect.TypeOf((*MockInstanceVariablesServiceInterface)(nil).ListVariables), varargs...)
	return &MockInstanceVariablesServiceInterfaceListVariablesCall{Call: call}
}

// MockInstanceVariablesServiceInterfaceListVariablesCall wrap *gomock.Call
type MockInstanceVariablesServiceInterfaceListVariablesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInstanceVariablesServiceInterfaceListVariablesCall) Return(arg0 []*gitlab.InstanceVariable, arg1 *gitlab.Response, arg2 error) *MockInstanceVariablesServiceInterfaceListVariablesCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInstanceVariablesServiceInterfaceListVariablesCall) Do(f func(*gitlab.ListInstanceVariablesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.InstanceVariable, *gitlab.Response, error)) *MockInstanceVariablesServiceInterfaceListVariablesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInstanceVariablesServiceInterfaceListVariablesCall) DoAndReturn(f func(*gitlab.ListInstanceVariablesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.InstanceVariable, *gitlab.Response, error)) *MockInstanceVariablesServiceInterfaceListVariablesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoveVariable mocks base method.
func (m *MockInstanceVariablesServiceInterface) RemoveVariable(key string, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{key}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveVariable", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveVariable indicates an expected call of RemoveVariable.
func (mr *MockInstanceVariablesServiceInterfaceMockRecorder) RemoveVariable(key any, options ...any) *MockInstanceVariablesServiceInterfaceRemoveVariableCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{key}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVariable", reflect.TypeOf((*MockInstanceVariablesServiceInterface)(nil).RemoveVariable), varargs...)
	return &MockInstanceVariablesServiceInterfaceRemoveVariableCall{Call: call}
}

// MockInstanceVariablesServiceInterfaceRemoveVariableCall wrap *gomock.Call
type MockInstanceVariablesServiceInterfaceRemoveVariableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInstanceVariablesServiceInterfaceRemoveVariableCall) Return(arg0 *gitlab.Response, arg1 error) *MockInstanceVariablesServiceInterfaceRemoveVariableCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInstanceVariablesServiceInterfaceRemoveVariableCall) Do(f func(string, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockInstanceVariablesServiceInterfaceRemoveVariableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInstanceVariablesServiceInterfaceRemoveVariableCall) DoAndReturn(f func(string, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockInstanceVariablesServiceInterfaceRemoveVariableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateVariable mocks base method.
func (m *MockInstanceVariablesServiceInterface) UpdateVariable(key string, opt *gitlab.UpdateInstanceVariableOptions, options ...gitlab.RequestOptionFunc) (*gitlab.InstanceVariable, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{key, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateVariable", varargs...)
	ret0, _ := ret[0].(*gitlab.InstanceVariable)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateVariable indicates an expected call of UpdateVariable.
func (mr *MockInstanceVariablesServiceInterfaceMockRecorder) UpdateVariable(key, opt any, options ...any) *MockInstanceVariablesServiceInterfaceUpdateVariableCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{key, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVariable", reflect.TypeOf((*MockInstanceVariablesServiceInterface)(nil).UpdateVariable), varargs...)
	return &MockInstanceVariablesServiceInterfaceUpdateVariableCall{Call: call}
}

// MockInstanceVariablesServiceInterfaceUpdateVariableCall wrap *gomock.Call
type MockInstanceVariablesServiceInterfaceUpdateVariableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInstanceVariablesServiceInterfaceUpdateVariableCall) Return(arg0 *gitlab.InstanceVariable, arg1 *gitlab.Response, arg2 error) *MockInstanceVariablesServiceInterfaceUpdateVariableCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInstanceVariablesServiceInterfaceUpdateVariableCall) Do(f func(string, *gitlab.UpdateInstanceVariableOptions, ...gitlab.RequestOptionFunc) (*gitlab.InstanceVariable, *gitlab.Response, error)) *MockInstanceVariablesServiceInterfaceUpdateVariableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInstanceVariablesServiceInterfaceUpdateVariableCall) DoAndReturn(f func(string, *gitlab.UpdateInstanceVariableOptions, ...gitlab.RequestOptionFunc) (*gitlab.InstanceVariable, *gitlab.Response, error)) *MockInstanceVariablesServiceInterfaceUpdateVariableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
