// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: ProjectVulnerabilitiesServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=project_vulnerabilities_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectVulnerabilitiesServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectVulnerabilitiesServiceInterface is a mock of ProjectVulnerabilitiesServiceInterface interface.
type MockProjectVulnerabilitiesServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProjectVulnerabilitiesServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockProjectVulnerabilitiesServiceInterfaceMockRecorder is the mock recorder for MockProjectVulnerabilitiesServiceInterface.
type MockProjectVulnerabilitiesServiceInterfaceMockRecorder struct {
	mock *MockProjectVulnerabilitiesServiceInterface
}

// NewMockProjectVulnerabilitiesServiceInterface creates a new mock instance.
func NewMockProjectVulnerabilitiesServiceInterface(ctrl *gomock.Controller) *MockProjectVulnerabilitiesServiceInterface {
	mock := &MockProjectVulnerabilitiesServiceInterface{ctrl: ctrl}
	mock.recorder = &MockProjectVulnerabilitiesServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectVulnerabilitiesServiceInterface) EXPECT() *MockProjectVulnerabilitiesServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateVulnerability mocks base method.
func (m *MockProjectVulnerabilitiesServiceInterface) CreateVulnerability(pid any, opt *gitlab.CreateVulnerabilityOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectVulnerability, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateVulnerability", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectVulnerability)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateVulnerability indicates an expected call of CreateVulnerability.
func (mr *MockProjectVulnerabilitiesServiceInterfaceMockRecorder) CreateVulnerability(pid, opt any, options ...any) *MockProjectVulnerabilitiesServiceInterfaceCreateVulnerabilityCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVulnerability", reflect.TypeOf((*MockProjectVulnerabilitiesServiceInterface)(nil).CreateVulnerability), varargs...)
	return &MockProjectVulnerabilitiesServiceInterfaceCreateVulnerabilityCall{Call: call}
}

// MockProjectVulnerabilitiesServiceInterfaceCreateVulnerabilityCall wrap *gomock.Call
type MockProjectVulnerabilitiesServiceInterfaceCreateVulnerabilityCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectVulnerabilitiesServiceInterfaceCreateVulnerabilityCall) Return(arg0 *gitlab.ProjectVulnerability, arg1 *gitlab.Response, arg2 error) *MockProjectVulnerabilitiesServiceInterfaceCreateVulnerabilityCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectVulnerabilitiesServiceInterfaceCreateVulnerabilityCall) Do(f func(any, *gitlab.CreateVulnerabilityOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectVulnerability, *gitlab.Response, error)) *MockProjectVulnerabilitiesServiceInterfaceCreateVulnerabilityCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectVulnerabilitiesServiceInterfaceCreateVulnerabilityCall) DoAndReturn(f func(any, *gitlab.CreateVulnerabilityOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectVulnerability, *gitlab.Response, error)) *MockProjectVulnerabilitiesServiceInterfaceCreateVulnerabilityCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListProjectVulnerabilities mocks base method.
func (m *MockProjectVulnerabilitiesServiceInterface) ListProjectVulnerabilities(pid any, opt *gitlab.ListProjectVulnerabilitiesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectVulnerability, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectVulnerabilities", varargs...)
	ret0, _ := ret[0].([]*gitlab.ProjectVulnerability)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectVulnerabilities indicates an expected call of ListProjectVulnerabilities.
func (mr *MockProjectVulnerabilitiesServiceInterfaceMockRecorder) ListProjectVulnerabilities(pid, opt any, options ...any) *MockProjectVulnerabilitiesServiceInterfaceListProjectVulnerabilitiesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectVulnerabilities", reflect.TypeOf((*MockProjectVulnerabilitiesServiceInterface)(nil).ListProjectVulnerabilities), varargs...)
	return &MockProjectVulnerabilitiesServiceInterfaceListProjectVulnerabilitiesCall{Call: call}
}

// MockProjectVulnerabilitiesServiceInterfaceListProjectVulnerabilitiesCall wrap *gomock.Call
type MockProjectVulnerabilitiesServiceInterfaceListProjectVulnerabilitiesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectVulnerabilitiesServiceInterfaceListProjectVulnerabilitiesCall) Return(arg0 []*gitlab.ProjectVulnerability, arg1 *gitlab.Response, arg2 error) *MockProjectVulnerabilitiesServiceInterfaceListProjectVulnerabilitiesCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectVulnerabilitiesServiceInterfaceListProjectVulnerabilitiesCall) Do(f func(any, *gitlab.ListProjectVulnerabilitiesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectVulnerability, *gitlab.Response, error)) *MockProjectVulnerabilitiesServiceInterfaceListProjectVulnerabilitiesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectVulnerabilitiesServiceInterfaceListProjectVulnerabilitiesCall) DoAndReturn(f func(any, *gitlab.ListProjectVulnerabilitiesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectVulnerability, *gitlab.Response, error)) *MockProjectVulnerabilitiesServiceInterfaceListProjectVulnerabilitiesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
