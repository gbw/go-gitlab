// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: GroupClustersServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=group_clusters_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupClustersServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockGroupClustersServiceInterface is a mock of GroupClustersServiceInterface interface.
type MockGroupClustersServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockGroupClustersServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockGroupClustersServiceInterfaceMockRecorder is the mock recorder for MockGroupClustersServiceInterface.
type MockGroupClustersServiceInterfaceMockRecorder struct {
	mock *MockGroupClustersServiceInterface
}

// NewMockGroupClustersServiceInterface creates a new mock instance.
func NewMockGroupClustersServiceInterface(ctrl *gomock.Controller) *MockGroupClustersServiceInterface {
	mock := &MockGroupClustersServiceInterface{ctrl: ctrl}
	mock.recorder = &MockGroupClustersServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupClustersServiceInterface) EXPECT() *MockGroupClustersServiceInterfaceMockRecorder {
	return m.recorder
}

// AddCluster mocks base method.
func (m *MockGroupClustersServiceInterface) AddCluster(pid any, opt *gitlab.AddGroupClusterOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupCluster, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddCluster", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupCluster)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddCluster indicates an expected call of AddCluster.
func (mr *MockGroupClustersServiceInterfaceMockRecorder) AddCluster(pid, opt any, options ...any) *MockGroupClustersServiceInterfaceAddClusterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCluster", reflect.TypeOf((*MockGroupClustersServiceInterface)(nil).AddCluster), varargs...)
	return &MockGroupClustersServiceInterfaceAddClusterCall{Call: call}
}

// MockGroupClustersServiceInterfaceAddClusterCall wrap *gomock.Call
type MockGroupClustersServiceInterfaceAddClusterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupClustersServiceInterfaceAddClusterCall) Return(arg0 *gitlab.GroupCluster, arg1 *gitlab.Response, arg2 error) *MockGroupClustersServiceInterfaceAddClusterCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupClustersServiceInterfaceAddClusterCall) Do(f func(any, *gitlab.AddGroupClusterOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupCluster, *gitlab.Response, error)) *MockGroupClustersServiceInterfaceAddClusterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupClustersServiceInterfaceAddClusterCall) DoAndReturn(f func(any, *gitlab.AddGroupClusterOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupCluster, *gitlab.Response, error)) *MockGroupClustersServiceInterfaceAddClusterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteCluster mocks base method.
func (m *MockGroupClustersServiceInterface) DeleteCluster(pid any, cluster int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
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
func (mr *MockGroupClustersServiceInterfaceMockRecorder) DeleteCluster(pid, cluster any, options ...any) *MockGroupClustersServiceInterfaceDeleteClusterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, cluster}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockGroupClustersServiceInterface)(nil).DeleteCluster), varargs...)
	return &MockGroupClustersServiceInterfaceDeleteClusterCall{Call: call}
}

// MockGroupClustersServiceInterfaceDeleteClusterCall wrap *gomock.Call
type MockGroupClustersServiceInterfaceDeleteClusterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupClustersServiceInterfaceDeleteClusterCall) Return(arg0 *gitlab.Response, arg1 error) *MockGroupClustersServiceInterfaceDeleteClusterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupClustersServiceInterfaceDeleteClusterCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockGroupClustersServiceInterfaceDeleteClusterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupClustersServiceInterfaceDeleteClusterCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockGroupClustersServiceInterfaceDeleteClusterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EditCluster mocks base method.
func (m *MockGroupClustersServiceInterface) EditCluster(pid any, cluster int, opt *gitlab.EditGroupClusterOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupCluster, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, cluster, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditCluster", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupCluster)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditCluster indicates an expected call of EditCluster.
func (mr *MockGroupClustersServiceInterfaceMockRecorder) EditCluster(pid, cluster, opt any, options ...any) *MockGroupClustersServiceInterfaceEditClusterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, cluster, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditCluster", reflect.TypeOf((*MockGroupClustersServiceInterface)(nil).EditCluster), varargs...)
	return &MockGroupClustersServiceInterfaceEditClusterCall{Call: call}
}

// MockGroupClustersServiceInterfaceEditClusterCall wrap *gomock.Call
type MockGroupClustersServiceInterfaceEditClusterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupClustersServiceInterfaceEditClusterCall) Return(arg0 *gitlab.GroupCluster, arg1 *gitlab.Response, arg2 error) *MockGroupClustersServiceInterfaceEditClusterCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupClustersServiceInterfaceEditClusterCall) Do(f func(any, int, *gitlab.EditGroupClusterOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupCluster, *gitlab.Response, error)) *MockGroupClustersServiceInterfaceEditClusterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupClustersServiceInterfaceEditClusterCall) DoAndReturn(f func(any, int, *gitlab.EditGroupClusterOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupCluster, *gitlab.Response, error)) *MockGroupClustersServiceInterfaceEditClusterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCluster mocks base method.
func (m *MockGroupClustersServiceInterface) GetCluster(pid any, cluster int, options ...gitlab.RequestOptionFunc) (*gitlab.GroupCluster, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, cluster}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCluster", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupCluster)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCluster indicates an expected call of GetCluster.
func (mr *MockGroupClustersServiceInterfaceMockRecorder) GetCluster(pid, cluster any, options ...any) *MockGroupClustersServiceInterfaceGetClusterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, cluster}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockGroupClustersServiceInterface)(nil).GetCluster), varargs...)
	return &MockGroupClustersServiceInterfaceGetClusterCall{Call: call}
}

// MockGroupClustersServiceInterfaceGetClusterCall wrap *gomock.Call
type MockGroupClustersServiceInterfaceGetClusterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupClustersServiceInterfaceGetClusterCall) Return(arg0 *gitlab.GroupCluster, arg1 *gitlab.Response, arg2 error) *MockGroupClustersServiceInterfaceGetClusterCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupClustersServiceInterfaceGetClusterCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.GroupCluster, *gitlab.Response, error)) *MockGroupClustersServiceInterfaceGetClusterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupClustersServiceInterfaceGetClusterCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.GroupCluster, *gitlab.Response, error)) *MockGroupClustersServiceInterfaceGetClusterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListClusters mocks base method.
func (m *MockGroupClustersServiceInterface) ListClusters(pid any, options ...gitlab.RequestOptionFunc) ([]*gitlab.GroupCluster, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].([]*gitlab.GroupCluster)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockGroupClustersServiceInterfaceMockRecorder) ListClusters(pid any, options ...any) *MockGroupClustersServiceInterfaceListClustersCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockGroupClustersServiceInterface)(nil).ListClusters), varargs...)
	return &MockGroupClustersServiceInterfaceListClustersCall{Call: call}
}

// MockGroupClustersServiceInterfaceListClustersCall wrap *gomock.Call
type MockGroupClustersServiceInterfaceListClustersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupClustersServiceInterfaceListClustersCall) Return(arg0 []*gitlab.GroupCluster, arg1 *gitlab.Response, arg2 error) *MockGroupClustersServiceInterfaceListClustersCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupClustersServiceInterfaceListClustersCall) Do(f func(any, ...gitlab.RequestOptionFunc) ([]*gitlab.GroupCluster, *gitlab.Response, error)) *MockGroupClustersServiceInterfaceListClustersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupClustersServiceInterfaceListClustersCall) DoAndReturn(f func(any, ...gitlab.RequestOptionFunc) ([]*gitlab.GroupCluster, *gitlab.Response, error)) *MockGroupClustersServiceInterfaceListClustersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
