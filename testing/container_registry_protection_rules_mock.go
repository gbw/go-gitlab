// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: ContainerRegistryProtectionRulesServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=container_registry_protection_rules_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ContainerRegistryProtectionRulesServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockContainerRegistryProtectionRulesServiceInterface is a mock of ContainerRegistryProtectionRulesServiceInterface interface.
type MockContainerRegistryProtectionRulesServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockContainerRegistryProtectionRulesServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockContainerRegistryProtectionRulesServiceInterfaceMockRecorder is the mock recorder for MockContainerRegistryProtectionRulesServiceInterface.
type MockContainerRegistryProtectionRulesServiceInterfaceMockRecorder struct {
	mock *MockContainerRegistryProtectionRulesServiceInterface
}

// NewMockContainerRegistryProtectionRulesServiceInterface creates a new mock instance.
func NewMockContainerRegistryProtectionRulesServiceInterface(ctrl *gomock.Controller) *MockContainerRegistryProtectionRulesServiceInterface {
	mock := &MockContainerRegistryProtectionRulesServiceInterface{ctrl: ctrl}
	mock.recorder = &MockContainerRegistryProtectionRulesServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContainerRegistryProtectionRulesServiceInterface) EXPECT() *MockContainerRegistryProtectionRulesServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateContainerRegistryProtectionRule mocks base method.
func (m *MockContainerRegistryProtectionRulesServiceInterface) CreateContainerRegistryProtectionRule(pid any, opt *gitlab.CreateContainerRegistryProtectionRuleOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateContainerRegistryProtectionRule", varargs...)
	ret0, _ := ret[0].(*gitlab.ContainerRegistryProtectionRule)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateContainerRegistryProtectionRule indicates an expected call of CreateContainerRegistryProtectionRule.
func (mr *MockContainerRegistryProtectionRulesServiceInterfaceMockRecorder) CreateContainerRegistryProtectionRule(pid, opt any, options ...any) *MockContainerRegistryProtectionRulesServiceInterfaceCreateContainerRegistryProtectionRuleCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainerRegistryProtectionRule", reflect.TypeOf((*MockContainerRegistryProtectionRulesServiceInterface)(nil).CreateContainerRegistryProtectionRule), varargs...)
	return &MockContainerRegistryProtectionRulesServiceInterfaceCreateContainerRegistryProtectionRuleCall{Call: call}
}

// MockContainerRegistryProtectionRulesServiceInterfaceCreateContainerRegistryProtectionRuleCall wrap *gomock.Call
type MockContainerRegistryProtectionRulesServiceInterfaceCreateContainerRegistryProtectionRuleCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockContainerRegistryProtectionRulesServiceInterfaceCreateContainerRegistryProtectionRuleCall) Return(arg0 *gitlab.ContainerRegistryProtectionRule, arg1 *gitlab.Response, arg2 error) *MockContainerRegistryProtectionRulesServiceInterfaceCreateContainerRegistryProtectionRuleCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockContainerRegistryProtectionRulesServiceInterfaceCreateContainerRegistryProtectionRuleCall) Do(f func(any, *gitlab.CreateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) (*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error)) *MockContainerRegistryProtectionRulesServiceInterfaceCreateContainerRegistryProtectionRuleCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockContainerRegistryProtectionRulesServiceInterfaceCreateContainerRegistryProtectionRuleCall) DoAndReturn(f func(any, *gitlab.CreateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) (*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error)) *MockContainerRegistryProtectionRulesServiceInterfaceCreateContainerRegistryProtectionRuleCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteContainerRegistryProtectionRule mocks base method.
func (m *MockContainerRegistryProtectionRulesServiceInterface) DeleteContainerRegistryProtectionRule(pid any, ruleID int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, ruleID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteContainerRegistryProtectionRule", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteContainerRegistryProtectionRule indicates an expected call of DeleteContainerRegistryProtectionRule.
func (mr *MockContainerRegistryProtectionRulesServiceInterfaceMockRecorder) DeleteContainerRegistryProtectionRule(pid, ruleID any, options ...any) *MockContainerRegistryProtectionRulesServiceInterfaceDeleteContainerRegistryProtectionRuleCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, ruleID}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainerRegistryProtectionRule", reflect.TypeOf((*MockContainerRegistryProtectionRulesServiceInterface)(nil).DeleteContainerRegistryProtectionRule), varargs...)
	return &MockContainerRegistryProtectionRulesServiceInterfaceDeleteContainerRegistryProtectionRuleCall{Call: call}
}

// MockContainerRegistryProtectionRulesServiceInterfaceDeleteContainerRegistryProtectionRuleCall wrap *gomock.Call
type MockContainerRegistryProtectionRulesServiceInterfaceDeleteContainerRegistryProtectionRuleCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockContainerRegistryProtectionRulesServiceInterfaceDeleteContainerRegistryProtectionRuleCall) Return(arg0 *gitlab.Response, arg1 error) *MockContainerRegistryProtectionRulesServiceInterfaceDeleteContainerRegistryProtectionRuleCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockContainerRegistryProtectionRulesServiceInterfaceDeleteContainerRegistryProtectionRuleCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockContainerRegistryProtectionRulesServiceInterfaceDeleteContainerRegistryProtectionRuleCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockContainerRegistryProtectionRulesServiceInterfaceDeleteContainerRegistryProtectionRuleCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockContainerRegistryProtectionRulesServiceInterfaceDeleteContainerRegistryProtectionRuleCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListContainerRegistryProtectionRules mocks base method.
func (m *MockContainerRegistryProtectionRulesServiceInterface) ListContainerRegistryProtectionRules(pid any, options ...gitlab.RequestOptionFunc) ([]*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListContainerRegistryProtectionRules", varargs...)
	ret0, _ := ret[0].([]*gitlab.ContainerRegistryProtectionRule)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListContainerRegistryProtectionRules indicates an expected call of ListContainerRegistryProtectionRules.
func (mr *MockContainerRegistryProtectionRulesServiceInterfaceMockRecorder) ListContainerRegistryProtectionRules(pid any, options ...any) *MockContainerRegistryProtectionRulesServiceInterfaceListContainerRegistryProtectionRulesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainerRegistryProtectionRules", reflect.TypeOf((*MockContainerRegistryProtectionRulesServiceInterface)(nil).ListContainerRegistryProtectionRules), varargs...)
	return &MockContainerRegistryProtectionRulesServiceInterfaceListContainerRegistryProtectionRulesCall{Call: call}
}

// MockContainerRegistryProtectionRulesServiceInterfaceListContainerRegistryProtectionRulesCall wrap *gomock.Call
type MockContainerRegistryProtectionRulesServiceInterfaceListContainerRegistryProtectionRulesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockContainerRegistryProtectionRulesServiceInterfaceListContainerRegistryProtectionRulesCall) Return(arg0 []*gitlab.ContainerRegistryProtectionRule, arg1 *gitlab.Response, arg2 error) *MockContainerRegistryProtectionRulesServiceInterfaceListContainerRegistryProtectionRulesCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockContainerRegistryProtectionRulesServiceInterfaceListContainerRegistryProtectionRulesCall) Do(f func(any, ...gitlab.RequestOptionFunc) ([]*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error)) *MockContainerRegistryProtectionRulesServiceInterfaceListContainerRegistryProtectionRulesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockContainerRegistryProtectionRulesServiceInterfaceListContainerRegistryProtectionRulesCall) DoAndReturn(f func(any, ...gitlab.RequestOptionFunc) ([]*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error)) *MockContainerRegistryProtectionRulesServiceInterfaceListContainerRegistryProtectionRulesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateContainerRegistryProtectionRule mocks base method.
func (m *MockContainerRegistryProtectionRulesServiceInterface) UpdateContainerRegistryProtectionRule(pid any, ruleID int, opt *gitlab.UpdateContainerRegistryProtectionRuleOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, ruleID, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateContainerRegistryProtectionRule", varargs...)
	ret0, _ := ret[0].(*gitlab.ContainerRegistryProtectionRule)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateContainerRegistryProtectionRule indicates an expected call of UpdateContainerRegistryProtectionRule.
func (mr *MockContainerRegistryProtectionRulesServiceInterfaceMockRecorder) UpdateContainerRegistryProtectionRule(pid, ruleID, opt any, options ...any) *MockContainerRegistryProtectionRulesServiceInterfaceUpdateContainerRegistryProtectionRuleCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, ruleID, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContainerRegistryProtectionRule", reflect.TypeOf((*MockContainerRegistryProtectionRulesServiceInterface)(nil).UpdateContainerRegistryProtectionRule), varargs...)
	return &MockContainerRegistryProtectionRulesServiceInterfaceUpdateContainerRegistryProtectionRuleCall{Call: call}
}

// MockContainerRegistryProtectionRulesServiceInterfaceUpdateContainerRegistryProtectionRuleCall wrap *gomock.Call
type MockContainerRegistryProtectionRulesServiceInterfaceUpdateContainerRegistryProtectionRuleCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockContainerRegistryProtectionRulesServiceInterfaceUpdateContainerRegistryProtectionRuleCall) Return(arg0 *gitlab.ContainerRegistryProtectionRule, arg1 *gitlab.Response, arg2 error) *MockContainerRegistryProtectionRulesServiceInterfaceUpdateContainerRegistryProtectionRuleCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockContainerRegistryProtectionRulesServiceInterfaceUpdateContainerRegistryProtectionRuleCall) Do(f func(any, int, *gitlab.UpdateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) (*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error)) *MockContainerRegistryProtectionRulesServiceInterfaceUpdateContainerRegistryProtectionRuleCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockContainerRegistryProtectionRulesServiceInterfaceUpdateContainerRegistryProtectionRuleCall) DoAndReturn(f func(any, int, *gitlab.UpdateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) (*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error)) *MockContainerRegistryProtectionRulesServiceInterfaceUpdateContainerRegistryProtectionRuleCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
