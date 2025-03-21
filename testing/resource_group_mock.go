// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: ResourceGroupServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=resource_group_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ResourceGroupServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockResourceGroupServiceInterface is a mock of ResourceGroupServiceInterface interface.
type MockResourceGroupServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockResourceGroupServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockResourceGroupServiceInterfaceMockRecorder is the mock recorder for MockResourceGroupServiceInterface.
type MockResourceGroupServiceInterfaceMockRecorder struct {
	mock *MockResourceGroupServiceInterface
}

// NewMockResourceGroupServiceInterface creates a new mock instance.
func NewMockResourceGroupServiceInterface(ctrl *gomock.Controller) *MockResourceGroupServiceInterface {
	mock := &MockResourceGroupServiceInterface{ctrl: ctrl}
	mock.recorder = &MockResourceGroupServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceGroupServiceInterface) EXPECT() *MockResourceGroupServiceInterfaceMockRecorder {
	return m.recorder
}

// EditAnExistingResourceGroup mocks base method.
func (m *MockResourceGroupServiceInterface) EditAnExistingResourceGroup(pid any, key string, opts *gitlab.EditAnExistingResourceGroupOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, key, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditAnExistingResourceGroup", varargs...)
	ret0, _ := ret[0].(*gitlab.ResourceGroup)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditAnExistingResourceGroup indicates an expected call of EditAnExistingResourceGroup.
func (mr *MockResourceGroupServiceInterfaceMockRecorder) EditAnExistingResourceGroup(pid, key, opts any, options ...any) *MockResourceGroupServiceInterfaceEditAnExistingResourceGroupCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, key, opts}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditAnExistingResourceGroup", reflect.TypeOf((*MockResourceGroupServiceInterface)(nil).EditAnExistingResourceGroup), varargs...)
	return &MockResourceGroupServiceInterfaceEditAnExistingResourceGroupCall{Call: call}
}

// MockResourceGroupServiceInterfaceEditAnExistingResourceGroupCall wrap *gomock.Call
type MockResourceGroupServiceInterfaceEditAnExistingResourceGroupCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceGroupServiceInterfaceEditAnExistingResourceGroupCall) Return(arg0 *gitlab.ResourceGroup, arg1 *gitlab.Response, arg2 error) *MockResourceGroupServiceInterfaceEditAnExistingResourceGroupCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceGroupServiceInterfaceEditAnExistingResourceGroupCall) Do(f func(any, string, *gitlab.EditAnExistingResourceGroupOptions, ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error)) *MockResourceGroupServiceInterfaceEditAnExistingResourceGroupCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceGroupServiceInterfaceEditAnExistingResourceGroupCall) DoAndReturn(f func(any, string, *gitlab.EditAnExistingResourceGroupOptions, ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error)) *MockResourceGroupServiceInterfaceEditAnExistingResourceGroupCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetASpecificResourceGroup mocks base method.
func (m *MockResourceGroupServiceInterface) GetASpecificResourceGroup(pid any, key string, options ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, key}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetASpecificResourceGroup", varargs...)
	ret0, _ := ret[0].(*gitlab.ResourceGroup)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetASpecificResourceGroup indicates an expected call of GetASpecificResourceGroup.
func (mr *MockResourceGroupServiceInterfaceMockRecorder) GetASpecificResourceGroup(pid, key any, options ...any) *MockResourceGroupServiceInterfaceGetASpecificResourceGroupCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, key}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetASpecificResourceGroup", reflect.TypeOf((*MockResourceGroupServiceInterface)(nil).GetASpecificResourceGroup), varargs...)
	return &MockResourceGroupServiceInterfaceGetASpecificResourceGroupCall{Call: call}
}

// MockResourceGroupServiceInterfaceGetASpecificResourceGroupCall wrap *gomock.Call
type MockResourceGroupServiceInterfaceGetASpecificResourceGroupCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceGroupServiceInterfaceGetASpecificResourceGroupCall) Return(arg0 *gitlab.ResourceGroup, arg1 *gitlab.Response, arg2 error) *MockResourceGroupServiceInterfaceGetASpecificResourceGroupCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceGroupServiceInterfaceGetASpecificResourceGroupCall) Do(f func(any, string, ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error)) *MockResourceGroupServiceInterfaceGetASpecificResourceGroupCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceGroupServiceInterfaceGetASpecificResourceGroupCall) DoAndReturn(f func(any, string, ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error)) *MockResourceGroupServiceInterfaceGetASpecificResourceGroupCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAllResourceGroupsForAProject mocks base method.
func (m *MockResourceGroupServiceInterface) GetAllResourceGroupsForAProject(pid any, options ...gitlab.RequestOptionFunc) ([]*gitlab.ResourceGroup, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllResourceGroupsForAProject", varargs...)
	ret0, _ := ret[0].([]*gitlab.ResourceGroup)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllResourceGroupsForAProject indicates an expected call of GetAllResourceGroupsForAProject.
func (mr *MockResourceGroupServiceInterfaceMockRecorder) GetAllResourceGroupsForAProject(pid any, options ...any) *MockResourceGroupServiceInterfaceGetAllResourceGroupsForAProjectCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllResourceGroupsForAProject", reflect.TypeOf((*MockResourceGroupServiceInterface)(nil).GetAllResourceGroupsForAProject), varargs...)
	return &MockResourceGroupServiceInterfaceGetAllResourceGroupsForAProjectCall{Call: call}
}

// MockResourceGroupServiceInterfaceGetAllResourceGroupsForAProjectCall wrap *gomock.Call
type MockResourceGroupServiceInterfaceGetAllResourceGroupsForAProjectCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceGroupServiceInterfaceGetAllResourceGroupsForAProjectCall) Return(arg0 []*gitlab.ResourceGroup, arg1 *gitlab.Response, arg2 error) *MockResourceGroupServiceInterfaceGetAllResourceGroupsForAProjectCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceGroupServiceInterfaceGetAllResourceGroupsForAProjectCall) Do(f func(any, ...gitlab.RequestOptionFunc) ([]*gitlab.ResourceGroup, *gitlab.Response, error)) *MockResourceGroupServiceInterfaceGetAllResourceGroupsForAProjectCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceGroupServiceInterfaceGetAllResourceGroupsForAProjectCall) DoAndReturn(f func(any, ...gitlab.RequestOptionFunc) ([]*gitlab.ResourceGroup, *gitlab.Response, error)) *MockResourceGroupServiceInterfaceGetAllResourceGroupsForAProjectCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListUpcomingJobsForASpecificResourceGroup mocks base method.
func (m *MockResourceGroupServiceInterface) ListUpcomingJobsForASpecificResourceGroup(pid any, key string, options ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, key}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUpcomingJobsForASpecificResourceGroup", varargs...)
	ret0, _ := ret[0].([]*gitlab.Job)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListUpcomingJobsForASpecificResourceGroup indicates an expected call of ListUpcomingJobsForASpecificResourceGroup.
func (mr *MockResourceGroupServiceInterfaceMockRecorder) ListUpcomingJobsForASpecificResourceGroup(pid, key any, options ...any) *MockResourceGroupServiceInterfaceListUpcomingJobsForASpecificResourceGroupCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, key}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUpcomingJobsForASpecificResourceGroup", reflect.TypeOf((*MockResourceGroupServiceInterface)(nil).ListUpcomingJobsForASpecificResourceGroup), varargs...)
	return &MockResourceGroupServiceInterfaceListUpcomingJobsForASpecificResourceGroupCall{Call: call}
}

// MockResourceGroupServiceInterfaceListUpcomingJobsForASpecificResourceGroupCall wrap *gomock.Call
type MockResourceGroupServiceInterfaceListUpcomingJobsForASpecificResourceGroupCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceGroupServiceInterfaceListUpcomingJobsForASpecificResourceGroupCall) Return(arg0 []*gitlab.Job, arg1 *gitlab.Response, arg2 error) *MockResourceGroupServiceInterfaceListUpcomingJobsForASpecificResourceGroupCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceGroupServiceInterfaceListUpcomingJobsForASpecificResourceGroupCall) Do(f func(any, string, ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error)) *MockResourceGroupServiceInterfaceListUpcomingJobsForASpecificResourceGroupCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceGroupServiceInterfaceListUpcomingJobsForASpecificResourceGroupCall) DoAndReturn(f func(any, string, ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error)) *MockResourceGroupServiceInterfaceListUpcomingJobsForASpecificResourceGroupCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
