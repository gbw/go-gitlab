// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: DORAMetricsServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=dora_metrics_mock.go -package=testing gitlab.com/gitlab-org/api/client-go DORAMetricsServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockDORAMetricsServiceInterface is a mock of DORAMetricsServiceInterface interface.
type MockDORAMetricsServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDORAMetricsServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockDORAMetricsServiceInterfaceMockRecorder is the mock recorder for MockDORAMetricsServiceInterface.
type MockDORAMetricsServiceInterfaceMockRecorder struct {
	mock *MockDORAMetricsServiceInterface
}

// NewMockDORAMetricsServiceInterface creates a new mock instance.
func NewMockDORAMetricsServiceInterface(ctrl *gomock.Controller) *MockDORAMetricsServiceInterface {
	mock := &MockDORAMetricsServiceInterface{ctrl: ctrl}
	mock.recorder = &MockDORAMetricsServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDORAMetricsServiceInterface) EXPECT() *MockDORAMetricsServiceInterfaceMockRecorder {
	return m.recorder
}

// GetGroupDORAMetrics mocks base method.
func (m *MockDORAMetricsServiceInterface) GetGroupDORAMetrics(gid any, opt gitlab.GetDORAMetricsOptions, options ...gitlab.RequestOptionFunc) ([]gitlab.DORAMetric, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupDORAMetrics", varargs...)
	ret0, _ := ret[0].([]gitlab.DORAMetric)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroupDORAMetrics indicates an expected call of GetGroupDORAMetrics.
func (mr *MockDORAMetricsServiceInterfaceMockRecorder) GetGroupDORAMetrics(gid, opt any, options ...any) *MockDORAMetricsServiceInterfaceGetGroupDORAMetricsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupDORAMetrics", reflect.TypeOf((*MockDORAMetricsServiceInterface)(nil).GetGroupDORAMetrics), varargs...)
	return &MockDORAMetricsServiceInterfaceGetGroupDORAMetricsCall{Call: call}
}

// MockDORAMetricsServiceInterfaceGetGroupDORAMetricsCall wrap *gomock.Call
type MockDORAMetricsServiceInterfaceGetGroupDORAMetricsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDORAMetricsServiceInterfaceGetGroupDORAMetricsCall) Return(arg0 []gitlab.DORAMetric, arg1 *gitlab.Response, arg2 error) *MockDORAMetricsServiceInterfaceGetGroupDORAMetricsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDORAMetricsServiceInterfaceGetGroupDORAMetricsCall) Do(f func(any, gitlab.GetDORAMetricsOptions, ...gitlab.RequestOptionFunc) ([]gitlab.DORAMetric, *gitlab.Response, error)) *MockDORAMetricsServiceInterfaceGetGroupDORAMetricsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDORAMetricsServiceInterfaceGetGroupDORAMetricsCall) DoAndReturn(f func(any, gitlab.GetDORAMetricsOptions, ...gitlab.RequestOptionFunc) ([]gitlab.DORAMetric, *gitlab.Response, error)) *MockDORAMetricsServiceInterfaceGetGroupDORAMetricsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetProjectDORAMetrics mocks base method.
func (m *MockDORAMetricsServiceInterface) GetProjectDORAMetrics(pid any, opt gitlab.GetDORAMetricsOptions, options ...gitlab.RequestOptionFunc) ([]gitlab.DORAMetric, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProjectDORAMetrics", varargs...)
	ret0, _ := ret[0].([]gitlab.DORAMetric)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProjectDORAMetrics indicates an expected call of GetProjectDORAMetrics.
func (mr *MockDORAMetricsServiceInterfaceMockRecorder) GetProjectDORAMetrics(pid, opt any, options ...any) *MockDORAMetricsServiceInterfaceGetProjectDORAMetricsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectDORAMetrics", reflect.TypeOf((*MockDORAMetricsServiceInterface)(nil).GetProjectDORAMetrics), varargs...)
	return &MockDORAMetricsServiceInterfaceGetProjectDORAMetricsCall{Call: call}
}

// MockDORAMetricsServiceInterfaceGetProjectDORAMetricsCall wrap *gomock.Call
type MockDORAMetricsServiceInterfaceGetProjectDORAMetricsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDORAMetricsServiceInterfaceGetProjectDORAMetricsCall) Return(arg0 []gitlab.DORAMetric, arg1 *gitlab.Response, arg2 error) *MockDORAMetricsServiceInterfaceGetProjectDORAMetricsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDORAMetricsServiceInterfaceGetProjectDORAMetricsCall) Do(f func(any, gitlab.GetDORAMetricsOptions, ...gitlab.RequestOptionFunc) ([]gitlab.DORAMetric, *gitlab.Response, error)) *MockDORAMetricsServiceInterfaceGetProjectDORAMetricsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDORAMetricsServiceInterfaceGetProjectDORAMetricsCall) DoAndReturn(f func(any, gitlab.GetDORAMetricsOptions, ...gitlab.RequestOptionFunc) ([]gitlab.DORAMetric, *gitlab.Response, error)) *MockDORAMetricsServiceInterfaceGetProjectDORAMetricsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
