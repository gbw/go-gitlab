// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: IssuesStatisticsServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=issues_statistics_mock.go -package=testing gitlab.com/gitlab-org/api/client-go IssuesStatisticsServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockIssuesStatisticsServiceInterface is a mock of IssuesStatisticsServiceInterface interface.
type MockIssuesStatisticsServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIssuesStatisticsServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockIssuesStatisticsServiceInterfaceMockRecorder is the mock recorder for MockIssuesStatisticsServiceInterface.
type MockIssuesStatisticsServiceInterfaceMockRecorder struct {
	mock *MockIssuesStatisticsServiceInterface
}

// NewMockIssuesStatisticsServiceInterface creates a new mock instance.
func NewMockIssuesStatisticsServiceInterface(ctrl *gomock.Controller) *MockIssuesStatisticsServiceInterface {
	mock := &MockIssuesStatisticsServiceInterface{ctrl: ctrl}
	mock.recorder = &MockIssuesStatisticsServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssuesStatisticsServiceInterface) EXPECT() *MockIssuesStatisticsServiceInterfaceMockRecorder {
	return m.recorder
}

// GetGroupIssuesStatistics mocks base method.
func (m *MockIssuesStatisticsServiceInterface) GetGroupIssuesStatistics(gid any, opt *gitlab.GetGroupIssuesStatisticsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.IssuesStatistics, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupIssuesStatistics", varargs...)
	ret0, _ := ret[0].(*gitlab.IssuesStatistics)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroupIssuesStatistics indicates an expected call of GetGroupIssuesStatistics.
func (mr *MockIssuesStatisticsServiceInterfaceMockRecorder) GetGroupIssuesStatistics(gid, opt any, options ...any) *MockIssuesStatisticsServiceInterfaceGetGroupIssuesStatisticsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupIssuesStatistics", reflect.TypeOf((*MockIssuesStatisticsServiceInterface)(nil).GetGroupIssuesStatistics), varargs...)
	return &MockIssuesStatisticsServiceInterfaceGetGroupIssuesStatisticsCall{Call: call}
}

// MockIssuesStatisticsServiceInterfaceGetGroupIssuesStatisticsCall wrap *gomock.Call
type MockIssuesStatisticsServiceInterfaceGetGroupIssuesStatisticsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIssuesStatisticsServiceInterfaceGetGroupIssuesStatisticsCall) Return(arg0 *gitlab.IssuesStatistics, arg1 *gitlab.Response, arg2 error) *MockIssuesStatisticsServiceInterfaceGetGroupIssuesStatisticsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIssuesStatisticsServiceInterfaceGetGroupIssuesStatisticsCall) Do(f func(any, *gitlab.GetGroupIssuesStatisticsOptions, ...gitlab.RequestOptionFunc) (*gitlab.IssuesStatistics, *gitlab.Response, error)) *MockIssuesStatisticsServiceInterfaceGetGroupIssuesStatisticsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIssuesStatisticsServiceInterfaceGetGroupIssuesStatisticsCall) DoAndReturn(f func(any, *gitlab.GetGroupIssuesStatisticsOptions, ...gitlab.RequestOptionFunc) (*gitlab.IssuesStatistics, *gitlab.Response, error)) *MockIssuesStatisticsServiceInterfaceGetGroupIssuesStatisticsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetIssuesStatistics mocks base method.
func (m *MockIssuesStatisticsServiceInterface) GetIssuesStatistics(opt *gitlab.GetIssuesStatisticsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.IssuesStatistics, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIssuesStatistics", varargs...)
	ret0, _ := ret[0].(*gitlab.IssuesStatistics)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssuesStatistics indicates an expected call of GetIssuesStatistics.
func (mr *MockIssuesStatisticsServiceInterfaceMockRecorder) GetIssuesStatistics(opt any, options ...any) *MockIssuesStatisticsServiceInterfaceGetIssuesStatisticsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssuesStatistics", reflect.TypeOf((*MockIssuesStatisticsServiceInterface)(nil).GetIssuesStatistics), varargs...)
	return &MockIssuesStatisticsServiceInterfaceGetIssuesStatisticsCall{Call: call}
}

// MockIssuesStatisticsServiceInterfaceGetIssuesStatisticsCall wrap *gomock.Call
type MockIssuesStatisticsServiceInterfaceGetIssuesStatisticsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIssuesStatisticsServiceInterfaceGetIssuesStatisticsCall) Return(arg0 *gitlab.IssuesStatistics, arg1 *gitlab.Response, arg2 error) *MockIssuesStatisticsServiceInterfaceGetIssuesStatisticsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIssuesStatisticsServiceInterfaceGetIssuesStatisticsCall) Do(f func(*gitlab.GetIssuesStatisticsOptions, ...gitlab.RequestOptionFunc) (*gitlab.IssuesStatistics, *gitlab.Response, error)) *MockIssuesStatisticsServiceInterfaceGetIssuesStatisticsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIssuesStatisticsServiceInterfaceGetIssuesStatisticsCall) DoAndReturn(f func(*gitlab.GetIssuesStatisticsOptions, ...gitlab.RequestOptionFunc) (*gitlab.IssuesStatistics, *gitlab.Response, error)) *MockIssuesStatisticsServiceInterfaceGetIssuesStatisticsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetProjectIssuesStatistics mocks base method.
func (m *MockIssuesStatisticsServiceInterface) GetProjectIssuesStatistics(pid any, opt *gitlab.GetProjectIssuesStatisticsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.IssuesStatistics, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProjectIssuesStatistics", varargs...)
	ret0, _ := ret[0].(*gitlab.IssuesStatistics)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProjectIssuesStatistics indicates an expected call of GetProjectIssuesStatistics.
func (mr *MockIssuesStatisticsServiceInterfaceMockRecorder) GetProjectIssuesStatistics(pid, opt any, options ...any) *MockIssuesStatisticsServiceInterfaceGetProjectIssuesStatisticsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectIssuesStatistics", reflect.TypeOf((*MockIssuesStatisticsServiceInterface)(nil).GetProjectIssuesStatistics), varargs...)
	return &MockIssuesStatisticsServiceInterfaceGetProjectIssuesStatisticsCall{Call: call}
}

// MockIssuesStatisticsServiceInterfaceGetProjectIssuesStatisticsCall wrap *gomock.Call
type MockIssuesStatisticsServiceInterfaceGetProjectIssuesStatisticsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockIssuesStatisticsServiceInterfaceGetProjectIssuesStatisticsCall) Return(arg0 *gitlab.IssuesStatistics, arg1 *gitlab.Response, arg2 error) *MockIssuesStatisticsServiceInterfaceGetProjectIssuesStatisticsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockIssuesStatisticsServiceInterfaceGetProjectIssuesStatisticsCall) Do(f func(any, *gitlab.GetProjectIssuesStatisticsOptions, ...gitlab.RequestOptionFunc) (*gitlab.IssuesStatistics, *gitlab.Response, error)) *MockIssuesStatisticsServiceInterfaceGetProjectIssuesStatisticsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockIssuesStatisticsServiceInterfaceGetProjectIssuesStatisticsCall) DoAndReturn(f func(any, *gitlab.GetProjectIssuesStatisticsOptions, ...gitlab.RequestOptionFunc) (*gitlab.IssuesStatistics, *gitlab.Response, error)) *MockIssuesStatisticsServiceInterfaceGetProjectIssuesStatisticsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
