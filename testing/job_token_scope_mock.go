// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: JobTokenScopeServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=job_token_scope_mock.go -package=testing gitlab.com/gitlab-org/api/client-go JobTokenScopeServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockJobTokenScopeServiceInterface is a mock of JobTokenScopeServiceInterface interface.
type MockJobTokenScopeServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockJobTokenScopeServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockJobTokenScopeServiceInterfaceMockRecorder is the mock recorder for MockJobTokenScopeServiceInterface.
type MockJobTokenScopeServiceInterfaceMockRecorder struct {
	mock *MockJobTokenScopeServiceInterface
}

// NewMockJobTokenScopeServiceInterface creates a new mock instance.
func NewMockJobTokenScopeServiceInterface(ctrl *gomock.Controller) *MockJobTokenScopeServiceInterface {
	mock := &MockJobTokenScopeServiceInterface{ctrl: ctrl}
	mock.recorder = &MockJobTokenScopeServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobTokenScopeServiceInterface) EXPECT() *MockJobTokenScopeServiceInterfaceMockRecorder {
	return m.recorder
}

// AddGroupToJobTokenAllowlist mocks base method.
func (m *MockJobTokenScopeServiceInterface) AddGroupToJobTokenAllowlist(pid any, opt *gitlab.AddGroupToJobTokenAllowlistOptions, options ...gitlab.RequestOptionFunc) (*gitlab.JobTokenAllowlistItem, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddGroupToJobTokenAllowlist", varargs...)
	ret0, _ := ret[0].(*gitlab.JobTokenAllowlistItem)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddGroupToJobTokenAllowlist indicates an expected call of AddGroupToJobTokenAllowlist.
func (mr *MockJobTokenScopeServiceInterfaceMockRecorder) AddGroupToJobTokenAllowlist(pid, opt any, options ...any) *MockJobTokenScopeServiceInterfaceAddGroupToJobTokenAllowlistCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGroupToJobTokenAllowlist", reflect.TypeOf((*MockJobTokenScopeServiceInterface)(nil).AddGroupToJobTokenAllowlist), varargs...)
	return &MockJobTokenScopeServiceInterfaceAddGroupToJobTokenAllowlistCall{Call: call}
}

// MockJobTokenScopeServiceInterfaceAddGroupToJobTokenAllowlistCall wrap *gomock.Call
type MockJobTokenScopeServiceInterfaceAddGroupToJobTokenAllowlistCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobTokenScopeServiceInterfaceAddGroupToJobTokenAllowlistCall) Return(arg0 *gitlab.JobTokenAllowlistItem, arg1 *gitlab.Response, arg2 error) *MockJobTokenScopeServiceInterfaceAddGroupToJobTokenAllowlistCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobTokenScopeServiceInterfaceAddGroupToJobTokenAllowlistCall) Do(f func(any, *gitlab.AddGroupToJobTokenAllowlistOptions, ...gitlab.RequestOptionFunc) (*gitlab.JobTokenAllowlistItem, *gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceAddGroupToJobTokenAllowlistCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobTokenScopeServiceInterfaceAddGroupToJobTokenAllowlistCall) DoAndReturn(f func(any, *gitlab.AddGroupToJobTokenAllowlistOptions, ...gitlab.RequestOptionFunc) (*gitlab.JobTokenAllowlistItem, *gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceAddGroupToJobTokenAllowlistCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AddProjectToJobScopeAllowList mocks base method.
func (m *MockJobTokenScopeServiceInterface) AddProjectToJobScopeAllowList(pid any, opt *gitlab.JobTokenInboundAllowOptions, options ...gitlab.RequestOptionFunc) (*gitlab.JobTokenInboundAllowItem, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddProjectToJobScopeAllowList", varargs...)
	ret0, _ := ret[0].(*gitlab.JobTokenInboundAllowItem)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddProjectToJobScopeAllowList indicates an expected call of AddProjectToJobScopeAllowList.
func (mr *MockJobTokenScopeServiceInterfaceMockRecorder) AddProjectToJobScopeAllowList(pid, opt any, options ...any) *MockJobTokenScopeServiceInterfaceAddProjectToJobScopeAllowListCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProjectToJobScopeAllowList", reflect.TypeOf((*MockJobTokenScopeServiceInterface)(nil).AddProjectToJobScopeAllowList), varargs...)
	return &MockJobTokenScopeServiceInterfaceAddProjectToJobScopeAllowListCall{Call: call}
}

// MockJobTokenScopeServiceInterfaceAddProjectToJobScopeAllowListCall wrap *gomock.Call
type MockJobTokenScopeServiceInterfaceAddProjectToJobScopeAllowListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobTokenScopeServiceInterfaceAddProjectToJobScopeAllowListCall) Return(arg0 *gitlab.JobTokenInboundAllowItem, arg1 *gitlab.Response, arg2 error) *MockJobTokenScopeServiceInterfaceAddProjectToJobScopeAllowListCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobTokenScopeServiceInterfaceAddProjectToJobScopeAllowListCall) Do(f func(any, *gitlab.JobTokenInboundAllowOptions, ...gitlab.RequestOptionFunc) (*gitlab.JobTokenInboundAllowItem, *gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceAddProjectToJobScopeAllowListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobTokenScopeServiceInterfaceAddProjectToJobScopeAllowListCall) DoAndReturn(f func(any, *gitlab.JobTokenInboundAllowOptions, ...gitlab.RequestOptionFunc) (*gitlab.JobTokenInboundAllowItem, *gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceAddProjectToJobScopeAllowListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetJobTokenAllowlistGroups mocks base method.
func (m *MockJobTokenScopeServiceInterface) GetJobTokenAllowlistGroups(pid any, opt *gitlab.GetJobTokenAllowlistGroupsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Group, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetJobTokenAllowlistGroups", varargs...)
	ret0, _ := ret[0].([]*gitlab.Group)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetJobTokenAllowlistGroups indicates an expected call of GetJobTokenAllowlistGroups.
func (mr *MockJobTokenScopeServiceInterfaceMockRecorder) GetJobTokenAllowlistGroups(pid, opt any, options ...any) *MockJobTokenScopeServiceInterfaceGetJobTokenAllowlistGroupsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobTokenAllowlistGroups", reflect.TypeOf((*MockJobTokenScopeServiceInterface)(nil).GetJobTokenAllowlistGroups), varargs...)
	return &MockJobTokenScopeServiceInterfaceGetJobTokenAllowlistGroupsCall{Call: call}
}

// MockJobTokenScopeServiceInterfaceGetJobTokenAllowlistGroupsCall wrap *gomock.Call
type MockJobTokenScopeServiceInterfaceGetJobTokenAllowlistGroupsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobTokenScopeServiceInterfaceGetJobTokenAllowlistGroupsCall) Return(arg0 []*gitlab.Group, arg1 *gitlab.Response, arg2 error) *MockJobTokenScopeServiceInterfaceGetJobTokenAllowlistGroupsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobTokenScopeServiceInterfaceGetJobTokenAllowlistGroupsCall) Do(f func(any, *gitlab.GetJobTokenAllowlistGroupsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Group, *gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceGetJobTokenAllowlistGroupsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobTokenScopeServiceInterfaceGetJobTokenAllowlistGroupsCall) DoAndReturn(f func(any, *gitlab.GetJobTokenAllowlistGroupsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Group, *gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceGetJobTokenAllowlistGroupsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetProjectJobTokenAccessSettings mocks base method.
func (m *MockJobTokenScopeServiceInterface) GetProjectJobTokenAccessSettings(pid any, options ...gitlab.RequestOptionFunc) (*gitlab.JobTokenAccessSettings, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProjectJobTokenAccessSettings", varargs...)
	ret0, _ := ret[0].(*gitlab.JobTokenAccessSettings)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProjectJobTokenAccessSettings indicates an expected call of GetProjectJobTokenAccessSettings.
func (mr *MockJobTokenScopeServiceInterfaceMockRecorder) GetProjectJobTokenAccessSettings(pid any, options ...any) *MockJobTokenScopeServiceInterfaceGetProjectJobTokenAccessSettingsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectJobTokenAccessSettings", reflect.TypeOf((*MockJobTokenScopeServiceInterface)(nil).GetProjectJobTokenAccessSettings), varargs...)
	return &MockJobTokenScopeServiceInterfaceGetProjectJobTokenAccessSettingsCall{Call: call}
}

// MockJobTokenScopeServiceInterfaceGetProjectJobTokenAccessSettingsCall wrap *gomock.Call
type MockJobTokenScopeServiceInterfaceGetProjectJobTokenAccessSettingsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobTokenScopeServiceInterfaceGetProjectJobTokenAccessSettingsCall) Return(arg0 *gitlab.JobTokenAccessSettings, arg1 *gitlab.Response, arg2 error) *MockJobTokenScopeServiceInterfaceGetProjectJobTokenAccessSettingsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobTokenScopeServiceInterfaceGetProjectJobTokenAccessSettingsCall) Do(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.JobTokenAccessSettings, *gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceGetProjectJobTokenAccessSettingsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobTokenScopeServiceInterfaceGetProjectJobTokenAccessSettingsCall) DoAndReturn(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.JobTokenAccessSettings, *gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceGetProjectJobTokenAccessSettingsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetProjectJobTokenInboundAllowList mocks base method.
func (m *MockJobTokenScopeServiceInterface) GetProjectJobTokenInboundAllowList(pid any, opt *gitlab.GetJobTokenInboundAllowListOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Project, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProjectJobTokenInboundAllowList", varargs...)
	ret0, _ := ret[0].([]*gitlab.Project)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProjectJobTokenInboundAllowList indicates an expected call of GetProjectJobTokenInboundAllowList.
func (mr *MockJobTokenScopeServiceInterfaceMockRecorder) GetProjectJobTokenInboundAllowList(pid, opt any, options ...any) *MockJobTokenScopeServiceInterfaceGetProjectJobTokenInboundAllowListCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectJobTokenInboundAllowList", reflect.TypeOf((*MockJobTokenScopeServiceInterface)(nil).GetProjectJobTokenInboundAllowList), varargs...)
	return &MockJobTokenScopeServiceInterfaceGetProjectJobTokenInboundAllowListCall{Call: call}
}

// MockJobTokenScopeServiceInterfaceGetProjectJobTokenInboundAllowListCall wrap *gomock.Call
type MockJobTokenScopeServiceInterfaceGetProjectJobTokenInboundAllowListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobTokenScopeServiceInterfaceGetProjectJobTokenInboundAllowListCall) Return(arg0 []*gitlab.Project, arg1 *gitlab.Response, arg2 error) *MockJobTokenScopeServiceInterfaceGetProjectJobTokenInboundAllowListCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobTokenScopeServiceInterfaceGetProjectJobTokenInboundAllowListCall) Do(f func(any, *gitlab.GetJobTokenInboundAllowListOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Project, *gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceGetProjectJobTokenInboundAllowListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobTokenScopeServiceInterfaceGetProjectJobTokenInboundAllowListCall) DoAndReturn(f func(any, *gitlab.GetJobTokenInboundAllowListOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Project, *gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceGetProjectJobTokenInboundAllowListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PatchProjectJobTokenAccessSettings mocks base method.
func (m *MockJobTokenScopeServiceInterface) PatchProjectJobTokenAccessSettings(pid any, opt *gitlab.PatchProjectJobTokenAccessSettingsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchProjectJobTokenAccessSettings", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchProjectJobTokenAccessSettings indicates an expected call of PatchProjectJobTokenAccessSettings.
func (mr *MockJobTokenScopeServiceInterfaceMockRecorder) PatchProjectJobTokenAccessSettings(pid, opt any, options ...any) *MockJobTokenScopeServiceInterfacePatchProjectJobTokenAccessSettingsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchProjectJobTokenAccessSettings", reflect.TypeOf((*MockJobTokenScopeServiceInterface)(nil).PatchProjectJobTokenAccessSettings), varargs...)
	return &MockJobTokenScopeServiceInterfacePatchProjectJobTokenAccessSettingsCall{Call: call}
}

// MockJobTokenScopeServiceInterfacePatchProjectJobTokenAccessSettingsCall wrap *gomock.Call
type MockJobTokenScopeServiceInterfacePatchProjectJobTokenAccessSettingsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobTokenScopeServiceInterfacePatchProjectJobTokenAccessSettingsCall) Return(arg0 *gitlab.Response, arg1 error) *MockJobTokenScopeServiceInterfacePatchProjectJobTokenAccessSettingsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobTokenScopeServiceInterfacePatchProjectJobTokenAccessSettingsCall) Do(f func(any, *gitlab.PatchProjectJobTokenAccessSettingsOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockJobTokenScopeServiceInterfacePatchProjectJobTokenAccessSettingsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobTokenScopeServiceInterfacePatchProjectJobTokenAccessSettingsCall) DoAndReturn(f func(any, *gitlab.PatchProjectJobTokenAccessSettingsOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockJobTokenScopeServiceInterfacePatchProjectJobTokenAccessSettingsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoveGroupFromJobTokenAllowlist mocks base method.
func (m *MockJobTokenScopeServiceInterface) RemoveGroupFromJobTokenAllowlist(pid any, targetGroup int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, targetGroup}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveGroupFromJobTokenAllowlist", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveGroupFromJobTokenAllowlist indicates an expected call of RemoveGroupFromJobTokenAllowlist.
func (mr *MockJobTokenScopeServiceInterfaceMockRecorder) RemoveGroupFromJobTokenAllowlist(pid, targetGroup any, options ...any) *MockJobTokenScopeServiceInterfaceRemoveGroupFromJobTokenAllowlistCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, targetGroup}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveGroupFromJobTokenAllowlist", reflect.TypeOf((*MockJobTokenScopeServiceInterface)(nil).RemoveGroupFromJobTokenAllowlist), varargs...)
	return &MockJobTokenScopeServiceInterfaceRemoveGroupFromJobTokenAllowlistCall{Call: call}
}

// MockJobTokenScopeServiceInterfaceRemoveGroupFromJobTokenAllowlistCall wrap *gomock.Call
type MockJobTokenScopeServiceInterfaceRemoveGroupFromJobTokenAllowlistCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobTokenScopeServiceInterfaceRemoveGroupFromJobTokenAllowlistCall) Return(arg0 *gitlab.Response, arg1 error) *MockJobTokenScopeServiceInterfaceRemoveGroupFromJobTokenAllowlistCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobTokenScopeServiceInterfaceRemoveGroupFromJobTokenAllowlistCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceRemoveGroupFromJobTokenAllowlistCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobTokenScopeServiceInterfaceRemoveGroupFromJobTokenAllowlistCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceRemoveGroupFromJobTokenAllowlistCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoveProjectFromJobScopeAllowList mocks base method.
func (m *MockJobTokenScopeServiceInterface) RemoveProjectFromJobScopeAllowList(pid any, targetProject int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, targetProject}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveProjectFromJobScopeAllowList", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveProjectFromJobScopeAllowList indicates an expected call of RemoveProjectFromJobScopeAllowList.
func (mr *MockJobTokenScopeServiceInterfaceMockRecorder) RemoveProjectFromJobScopeAllowList(pid, targetProject any, options ...any) *MockJobTokenScopeServiceInterfaceRemoveProjectFromJobScopeAllowListCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, targetProject}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProjectFromJobScopeAllowList", reflect.TypeOf((*MockJobTokenScopeServiceInterface)(nil).RemoveProjectFromJobScopeAllowList), varargs...)
	return &MockJobTokenScopeServiceInterfaceRemoveProjectFromJobScopeAllowListCall{Call: call}
}

// MockJobTokenScopeServiceInterfaceRemoveProjectFromJobScopeAllowListCall wrap *gomock.Call
type MockJobTokenScopeServiceInterfaceRemoveProjectFromJobScopeAllowListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobTokenScopeServiceInterfaceRemoveProjectFromJobScopeAllowListCall) Return(arg0 *gitlab.Response, arg1 error) *MockJobTokenScopeServiceInterfaceRemoveProjectFromJobScopeAllowListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobTokenScopeServiceInterfaceRemoveProjectFromJobScopeAllowListCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceRemoveProjectFromJobScopeAllowListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobTokenScopeServiceInterfaceRemoveProjectFromJobScopeAllowListCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockJobTokenScopeServiceInterfaceRemoveProjectFromJobScopeAllowListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
