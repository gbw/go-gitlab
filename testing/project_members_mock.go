// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: ProjectMembersServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=project_members_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectMembersServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectMembersServiceInterface is a mock of ProjectMembersServiceInterface interface.
type MockProjectMembersServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProjectMembersServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockProjectMembersServiceInterfaceMockRecorder is the mock recorder for MockProjectMembersServiceInterface.
type MockProjectMembersServiceInterfaceMockRecorder struct {
	mock *MockProjectMembersServiceInterface
}

// NewMockProjectMembersServiceInterface creates a new mock instance.
func NewMockProjectMembersServiceInterface(ctrl *gomock.Controller) *MockProjectMembersServiceInterface {
	mock := &MockProjectMembersServiceInterface{ctrl: ctrl}
	mock.recorder = &MockProjectMembersServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectMembersServiceInterface) EXPECT() *MockProjectMembersServiceInterfaceMockRecorder {
	return m.recorder
}

// AddProjectMember mocks base method.
func (m *MockProjectMembersServiceInterface) AddProjectMember(pid any, opt *gitlab.AddProjectMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddProjectMember", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddProjectMember indicates an expected call of AddProjectMember.
func (mr *MockProjectMembersServiceInterfaceMockRecorder) AddProjectMember(pid, opt any, options ...any) *MockProjectMembersServiceInterfaceAddProjectMemberCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProjectMember", reflect.TypeOf((*MockProjectMembersServiceInterface)(nil).AddProjectMember), varargs...)
	return &MockProjectMembersServiceInterfaceAddProjectMemberCall{Call: call}
}

// MockProjectMembersServiceInterfaceAddProjectMemberCall wrap *gomock.Call
type MockProjectMembersServiceInterfaceAddProjectMemberCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectMembersServiceInterfaceAddProjectMemberCall) Return(arg0 *gitlab.ProjectMember, arg1 *gitlab.Response, arg2 error) *MockProjectMembersServiceInterfaceAddProjectMemberCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectMembersServiceInterfaceAddProjectMemberCall) Do(f func(any, *gitlab.AddProjectMemberOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error)) *MockProjectMembersServiceInterfaceAddProjectMemberCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectMembersServiceInterfaceAddProjectMemberCall) DoAndReturn(f func(any, *gitlab.AddProjectMemberOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error)) *MockProjectMembersServiceInterfaceAddProjectMemberCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteProjectMember mocks base method.
func (m *MockProjectMembersServiceInterface) DeleteProjectMember(pid any, user int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, user}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteProjectMember", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProjectMember indicates an expected call of DeleteProjectMember.
func (mr *MockProjectMembersServiceInterfaceMockRecorder) DeleteProjectMember(pid, user any, options ...any) *MockProjectMembersServiceInterfaceDeleteProjectMemberCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, user}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectMember", reflect.TypeOf((*MockProjectMembersServiceInterface)(nil).DeleteProjectMember), varargs...)
	return &MockProjectMembersServiceInterfaceDeleteProjectMemberCall{Call: call}
}

// MockProjectMembersServiceInterfaceDeleteProjectMemberCall wrap *gomock.Call
type MockProjectMembersServiceInterfaceDeleteProjectMemberCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectMembersServiceInterfaceDeleteProjectMemberCall) Return(arg0 *gitlab.Response, arg1 error) *MockProjectMembersServiceInterfaceDeleteProjectMemberCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectMembersServiceInterfaceDeleteProjectMemberCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockProjectMembersServiceInterfaceDeleteProjectMemberCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectMembersServiceInterfaceDeleteProjectMemberCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockProjectMembersServiceInterfaceDeleteProjectMemberCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EditProjectMember mocks base method.
func (m *MockProjectMembersServiceInterface) EditProjectMember(pid any, user int, opt *gitlab.EditProjectMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, user, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditProjectMember", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditProjectMember indicates an expected call of EditProjectMember.
func (mr *MockProjectMembersServiceInterfaceMockRecorder) EditProjectMember(pid, user, opt any, options ...any) *MockProjectMembersServiceInterfaceEditProjectMemberCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, user, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditProjectMember", reflect.TypeOf((*MockProjectMembersServiceInterface)(nil).EditProjectMember), varargs...)
	return &MockProjectMembersServiceInterfaceEditProjectMemberCall{Call: call}
}

// MockProjectMembersServiceInterfaceEditProjectMemberCall wrap *gomock.Call
type MockProjectMembersServiceInterfaceEditProjectMemberCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectMembersServiceInterfaceEditProjectMemberCall) Return(arg0 *gitlab.ProjectMember, arg1 *gitlab.Response, arg2 error) *MockProjectMembersServiceInterfaceEditProjectMemberCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectMembersServiceInterfaceEditProjectMemberCall) Do(f func(any, int, *gitlab.EditProjectMemberOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error)) *MockProjectMembersServiceInterfaceEditProjectMemberCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectMembersServiceInterfaceEditProjectMemberCall) DoAndReturn(f func(any, int, *gitlab.EditProjectMemberOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error)) *MockProjectMembersServiceInterfaceEditProjectMemberCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetInheritedProjectMember mocks base method.
func (m *MockProjectMembersServiceInterface) GetInheritedProjectMember(pid any, user int, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, user}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInheritedProjectMember", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetInheritedProjectMember indicates an expected call of GetInheritedProjectMember.
func (mr *MockProjectMembersServiceInterfaceMockRecorder) GetInheritedProjectMember(pid, user any, options ...any) *MockProjectMembersServiceInterfaceGetInheritedProjectMemberCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, user}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInheritedProjectMember", reflect.TypeOf((*MockProjectMembersServiceInterface)(nil).GetInheritedProjectMember), varargs...)
	return &MockProjectMembersServiceInterfaceGetInheritedProjectMemberCall{Call: call}
}

// MockProjectMembersServiceInterfaceGetInheritedProjectMemberCall wrap *gomock.Call
type MockProjectMembersServiceInterfaceGetInheritedProjectMemberCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectMembersServiceInterfaceGetInheritedProjectMemberCall) Return(arg0 *gitlab.ProjectMember, arg1 *gitlab.Response, arg2 error) *MockProjectMembersServiceInterfaceGetInheritedProjectMemberCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectMembersServiceInterfaceGetInheritedProjectMemberCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error)) *MockProjectMembersServiceInterfaceGetInheritedProjectMemberCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectMembersServiceInterfaceGetInheritedProjectMemberCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error)) *MockProjectMembersServiceInterfaceGetInheritedProjectMemberCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetProjectMember mocks base method.
func (m *MockProjectMembersServiceInterface) GetProjectMember(pid any, user int, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, user}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProjectMember", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProjectMember indicates an expected call of GetProjectMember.
func (mr *MockProjectMembersServiceInterfaceMockRecorder) GetProjectMember(pid, user any, options ...any) *MockProjectMembersServiceInterfaceGetProjectMemberCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, user}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectMember", reflect.TypeOf((*MockProjectMembersServiceInterface)(nil).GetProjectMember), varargs...)
	return &MockProjectMembersServiceInterfaceGetProjectMemberCall{Call: call}
}

// MockProjectMembersServiceInterfaceGetProjectMemberCall wrap *gomock.Call
type MockProjectMembersServiceInterfaceGetProjectMemberCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectMembersServiceInterfaceGetProjectMemberCall) Return(arg0 *gitlab.ProjectMember, arg1 *gitlab.Response, arg2 error) *MockProjectMembersServiceInterfaceGetProjectMemberCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectMembersServiceInterfaceGetProjectMemberCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error)) *MockProjectMembersServiceInterfaceGetProjectMemberCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectMembersServiceInterfaceGetProjectMemberCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error)) *MockProjectMembersServiceInterfaceGetProjectMemberCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListAllProjectMembers mocks base method.
func (m *MockProjectMembersServiceInterface) ListAllProjectMembers(pid any, opt *gitlab.ListProjectMembersOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAllProjectMembers", varargs...)
	ret0, _ := ret[0].([]*gitlab.ProjectMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAllProjectMembers indicates an expected call of ListAllProjectMembers.
func (mr *MockProjectMembersServiceInterfaceMockRecorder) ListAllProjectMembers(pid, opt any, options ...any) *MockProjectMembersServiceInterfaceListAllProjectMembersCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllProjectMembers", reflect.TypeOf((*MockProjectMembersServiceInterface)(nil).ListAllProjectMembers), varargs...)
	return &MockProjectMembersServiceInterfaceListAllProjectMembersCall{Call: call}
}

// MockProjectMembersServiceInterfaceListAllProjectMembersCall wrap *gomock.Call
type MockProjectMembersServiceInterfaceListAllProjectMembersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectMembersServiceInterfaceListAllProjectMembersCall) Return(arg0 []*gitlab.ProjectMember, arg1 *gitlab.Response, arg2 error) *MockProjectMembersServiceInterfaceListAllProjectMembersCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectMembersServiceInterfaceListAllProjectMembersCall) Do(f func(any, *gitlab.ListProjectMembersOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectMember, *gitlab.Response, error)) *MockProjectMembersServiceInterfaceListAllProjectMembersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectMembersServiceInterfaceListAllProjectMembersCall) DoAndReturn(f func(any, *gitlab.ListProjectMembersOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectMember, *gitlab.Response, error)) *MockProjectMembersServiceInterfaceListAllProjectMembersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListProjectMembers mocks base method.
func (m *MockProjectMembersServiceInterface) ListProjectMembers(pid any, opt *gitlab.ListProjectMembersOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectMembers", varargs...)
	ret0, _ := ret[0].([]*gitlab.ProjectMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectMembers indicates an expected call of ListProjectMembers.
func (mr *MockProjectMembersServiceInterfaceMockRecorder) ListProjectMembers(pid, opt any, options ...any) *MockProjectMembersServiceInterfaceListProjectMembersCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectMembers", reflect.TypeOf((*MockProjectMembersServiceInterface)(nil).ListProjectMembers), varargs...)
	return &MockProjectMembersServiceInterfaceListProjectMembersCall{Call: call}
}

// MockProjectMembersServiceInterfaceListProjectMembersCall wrap *gomock.Call
type MockProjectMembersServiceInterfaceListProjectMembersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectMembersServiceInterfaceListProjectMembersCall) Return(arg0 []*gitlab.ProjectMember, arg1 *gitlab.Response, arg2 error) *MockProjectMembersServiceInterfaceListProjectMembersCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectMembersServiceInterfaceListProjectMembersCall) Do(f func(any, *gitlab.ListProjectMembersOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectMember, *gitlab.Response, error)) *MockProjectMembersServiceInterfaceListProjectMembersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectMembersServiceInterfaceListProjectMembersCall) DoAndReturn(f func(any, *gitlab.ListProjectMembersOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectMember, *gitlab.Response, error)) *MockProjectMembersServiceInterfaceListProjectMembersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
