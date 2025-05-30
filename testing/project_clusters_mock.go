// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: ProjectClustersServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=project_clusters_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectClustersServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectClustersServiceInterface is a mock of ProjectClustersServiceInterface interface.
type MockProjectClustersServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProjectClustersServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockProjectClustersServiceInterfaceMockRecorder is the mock recorder for MockProjectClustersServiceInterface.
type MockProjectClustersServiceInterfaceMockRecorder struct {
	mock *MockProjectClustersServiceInterface
}

// NewMockProjectClustersServiceInterface creates a new mock instance.
func NewMockProjectClustersServiceInterface(ctrl *gomock.Controller) *MockProjectClustersServiceInterface {
	mock := &MockProjectClustersServiceInterface{ctrl: ctrl}
	mock.recorder = &MockProjectClustersServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectClustersServiceInterface) EXPECT() *MockProjectClustersServiceInterfaceMockRecorder {
	return m.recorder
}

// AddCluster mocks base method.
func (m *MockProjectClustersServiceInterface) AddCluster(pid any, opt *gitlab.AddClusterOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectCluster, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddCluster", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectCluster)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddCluster indicates an expected call of AddCluster.
func (mr *MockProjectClustersServiceInterfaceMockRecorder) AddCluster(pid, opt any, options ...any) *MockProjectClustersServiceInterfaceAddClusterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCluster", reflect.TypeOf((*MockProjectClustersServiceInterface)(nil).AddCluster), varargs...)
	return &MockProjectClustersServiceInterfaceAddClusterCall{Call: call}
}

// MockProjectClustersServiceInterfaceAddClusterCall wrap *gomock.Call
type MockProjectClustersServiceInterfaceAddClusterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectClustersServiceInterfaceAddClusterCall) Return(arg0 *gitlab.ProjectCluster, arg1 *gitlab.Response, arg2 error) *MockProjectClustersServiceInterfaceAddClusterCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectClustersServiceInterfaceAddClusterCall) Do(f func(any, *gitlab.AddClusterOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectCluster, *gitlab.Response, error)) *MockProjectClustersServiceInterfaceAddClusterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectClustersServiceInterfaceAddClusterCall) DoAndReturn(f func(any, *gitlab.AddClusterOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectCluster, *gitlab.Response, error)) *MockProjectClustersServiceInterfaceAddClusterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteCluster mocks base method.
func (m *MockProjectClustersServiceInterface) DeleteCluster(pid any, cluster int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, cluster}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCluster", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCluster indicates an expected call of DeleteCluster.
func (mr *MockProjectClustersServiceInterfaceMockRecorder) DeleteCluster(pid, cluster any, options ...any) *MockProjectClustersServiceInterfaceDeleteClusterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, cluster}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockProjectClustersServiceInterface)(nil).DeleteCluster), varargs...)
	return &MockProjectClustersServiceInterfaceDeleteClusterCall{Call: call}
}

// MockProjectClustersServiceInterfaceDeleteClusterCall wrap *gomock.Call
type MockProjectClustersServiceInterfaceDeleteClusterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectClustersServiceInterfaceDeleteClusterCall) Return(arg0 *gitlab.Response, arg1 error) *MockProjectClustersServiceInterfaceDeleteClusterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectClustersServiceInterfaceDeleteClusterCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockProjectClustersServiceInterfaceDeleteClusterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectClustersServiceInterfaceDeleteClusterCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockProjectClustersServiceInterfaceDeleteClusterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EditCluster mocks base method.
func (m *MockProjectClustersServiceInterface) EditCluster(pid any, cluster int, opt *gitlab.EditClusterOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectCluster, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, cluster, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditCluster", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectCluster)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditCluster indicates an expected call of EditCluster.
func (mr *MockProjectClustersServiceInterfaceMockRecorder) EditCluster(pid, cluster, opt any, options ...any) *MockProjectClustersServiceInterfaceEditClusterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, cluster, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditCluster", reflect.TypeOf((*MockProjectClustersServiceInterface)(nil).EditCluster), varargs...)
	return &MockProjectClustersServiceInterfaceEditClusterCall{Call: call}
}

// MockProjectClustersServiceInterfaceEditClusterCall wrap *gomock.Call
type MockProjectClustersServiceInterfaceEditClusterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectClustersServiceInterfaceEditClusterCall) Return(arg0 *gitlab.ProjectCluster, arg1 *gitlab.Response, arg2 error) *MockProjectClustersServiceInterfaceEditClusterCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectClustersServiceInterfaceEditClusterCall) Do(f func(any, int, *gitlab.EditClusterOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectCluster, *gitlab.Response, error)) *MockProjectClustersServiceInterfaceEditClusterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectClustersServiceInterfaceEditClusterCall) DoAndReturn(f func(any, int, *gitlab.EditClusterOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectCluster, *gitlab.Response, error)) *MockProjectClustersServiceInterfaceEditClusterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCluster mocks base method.
func (m *MockProjectClustersServiceInterface) GetCluster(pid any, cluster int, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectCluster, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, cluster}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCluster", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectCluster)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCluster indicates an expected call of GetCluster.
func (mr *MockProjectClustersServiceInterfaceMockRecorder) GetCluster(pid, cluster any, options ...any) *MockProjectClustersServiceInterfaceGetClusterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, cluster}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockProjectClustersServiceInterface)(nil).GetCluster), varargs...)
	return &MockProjectClustersServiceInterfaceGetClusterCall{Call: call}
}

// MockProjectClustersServiceInterfaceGetClusterCall wrap *gomock.Call
type MockProjectClustersServiceInterfaceGetClusterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectClustersServiceInterfaceGetClusterCall) Return(arg0 *gitlab.ProjectCluster, arg1 *gitlab.Response, arg2 error) *MockProjectClustersServiceInterfaceGetClusterCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectClustersServiceInterfaceGetClusterCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.ProjectCluster, *gitlab.Response, error)) *MockProjectClustersServiceInterfaceGetClusterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectClustersServiceInterfaceGetClusterCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.ProjectCluster, *gitlab.Response, error)) *MockProjectClustersServiceInterfaceGetClusterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListClusters mocks base method.
func (m *MockProjectClustersServiceInterface) ListClusters(pid any, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectCluster, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].([]*gitlab.ProjectCluster)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockProjectClustersServiceInterfaceMockRecorder) ListClusters(pid any, options ...any) *MockProjectClustersServiceInterfaceListClustersCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockProjectClustersServiceInterface)(nil).ListClusters), varargs...)
	return &MockProjectClustersServiceInterfaceListClustersCall{Call: call}
}

// MockProjectClustersServiceInterfaceListClustersCall wrap *gomock.Call
type MockProjectClustersServiceInterfaceListClustersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectClustersServiceInterfaceListClustersCall) Return(arg0 []*gitlab.ProjectCluster, arg1 *gitlab.Response, arg2 error) *MockProjectClustersServiceInterfaceListClustersCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectClustersServiceInterfaceListClustersCall) Do(f func(any, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectCluster, *gitlab.Response, error)) *MockProjectClustersServiceInterfaceListClustersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectClustersServiceInterfaceListClustersCall) DoAndReturn(f func(any, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectCluster, *gitlab.Response, error)) *MockProjectClustersServiceInterfaceListClustersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
