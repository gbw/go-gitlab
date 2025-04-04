// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: ProjectAccessTokensServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=project_access_tokens_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectAccessTokensServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectAccessTokensServiceInterface is a mock of ProjectAccessTokensServiceInterface interface.
type MockProjectAccessTokensServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProjectAccessTokensServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockProjectAccessTokensServiceInterfaceMockRecorder is the mock recorder for MockProjectAccessTokensServiceInterface.
type MockProjectAccessTokensServiceInterfaceMockRecorder struct {
	mock *MockProjectAccessTokensServiceInterface
}

// NewMockProjectAccessTokensServiceInterface creates a new mock instance.
func NewMockProjectAccessTokensServiceInterface(ctrl *gomock.Controller) *MockProjectAccessTokensServiceInterface {
	mock := &MockProjectAccessTokensServiceInterface{ctrl: ctrl}
	mock.recorder = &MockProjectAccessTokensServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectAccessTokensServiceInterface) EXPECT() *MockProjectAccessTokensServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateProjectAccessToken mocks base method.
func (m *MockProjectAccessTokensServiceInterface) CreateProjectAccessToken(pid any, opt *gitlab.CreateProjectAccessTokenOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateProjectAccessToken", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectAccessToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateProjectAccessToken indicates an expected call of CreateProjectAccessToken.
func (mr *MockProjectAccessTokensServiceInterfaceMockRecorder) CreateProjectAccessToken(pid, opt any, options ...any) *MockProjectAccessTokensServiceInterfaceCreateProjectAccessTokenCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectAccessToken", reflect.TypeOf((*MockProjectAccessTokensServiceInterface)(nil).CreateProjectAccessToken), varargs...)
	return &MockProjectAccessTokensServiceInterfaceCreateProjectAccessTokenCall{Call: call}
}

// MockProjectAccessTokensServiceInterfaceCreateProjectAccessTokenCall wrap *gomock.Call
type MockProjectAccessTokensServiceInterfaceCreateProjectAccessTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectAccessTokensServiceInterfaceCreateProjectAccessTokenCall) Return(arg0 *gitlab.ProjectAccessToken, arg1 *gitlab.Response, arg2 error) *MockProjectAccessTokensServiceInterfaceCreateProjectAccessTokenCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectAccessTokensServiceInterfaceCreateProjectAccessTokenCall) Do(f func(any, *gitlab.CreateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)) *MockProjectAccessTokensServiceInterfaceCreateProjectAccessTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectAccessTokensServiceInterfaceCreateProjectAccessTokenCall) DoAndReturn(f func(any, *gitlab.CreateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)) *MockProjectAccessTokensServiceInterfaceCreateProjectAccessTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetProjectAccessToken mocks base method.
func (m *MockProjectAccessTokensServiceInterface) GetProjectAccessToken(pid any, id int, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProjectAccessToken", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectAccessToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProjectAccessToken indicates an expected call of GetProjectAccessToken.
func (mr *MockProjectAccessTokensServiceInterfaceMockRecorder) GetProjectAccessToken(pid, id any, options ...any) *MockProjectAccessTokensServiceInterfaceGetProjectAccessTokenCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, id}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectAccessToken", reflect.TypeOf((*MockProjectAccessTokensServiceInterface)(nil).GetProjectAccessToken), varargs...)
	return &MockProjectAccessTokensServiceInterfaceGetProjectAccessTokenCall{Call: call}
}

// MockProjectAccessTokensServiceInterfaceGetProjectAccessTokenCall wrap *gomock.Call
type MockProjectAccessTokensServiceInterfaceGetProjectAccessTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectAccessTokensServiceInterfaceGetProjectAccessTokenCall) Return(arg0 *gitlab.ProjectAccessToken, arg1 *gitlab.Response, arg2 error) *MockProjectAccessTokensServiceInterfaceGetProjectAccessTokenCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectAccessTokensServiceInterfaceGetProjectAccessTokenCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)) *MockProjectAccessTokensServiceInterfaceGetProjectAccessTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectAccessTokensServiceInterfaceGetProjectAccessTokenCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)) *MockProjectAccessTokensServiceInterfaceGetProjectAccessTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListProjectAccessTokens mocks base method.
func (m *MockProjectAccessTokensServiceInterface) ListProjectAccessTokens(pid any, opt *gitlab.ListProjectAccessTokensOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectAccessToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectAccessTokens", varargs...)
	ret0, _ := ret[0].([]*gitlab.ProjectAccessToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectAccessTokens indicates an expected call of ListProjectAccessTokens.
func (mr *MockProjectAccessTokensServiceInterfaceMockRecorder) ListProjectAccessTokens(pid, opt any, options ...any) *MockProjectAccessTokensServiceInterfaceListProjectAccessTokensCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectAccessTokens", reflect.TypeOf((*MockProjectAccessTokensServiceInterface)(nil).ListProjectAccessTokens), varargs...)
	return &MockProjectAccessTokensServiceInterfaceListProjectAccessTokensCall{Call: call}
}

// MockProjectAccessTokensServiceInterfaceListProjectAccessTokensCall wrap *gomock.Call
type MockProjectAccessTokensServiceInterfaceListProjectAccessTokensCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectAccessTokensServiceInterfaceListProjectAccessTokensCall) Return(arg0 []*gitlab.ProjectAccessToken, arg1 *gitlab.Response, arg2 error) *MockProjectAccessTokensServiceInterfaceListProjectAccessTokensCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectAccessTokensServiceInterfaceListProjectAccessTokensCall) Do(f func(any, *gitlab.ListProjectAccessTokensOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectAccessToken, *gitlab.Response, error)) *MockProjectAccessTokensServiceInterfaceListProjectAccessTokensCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectAccessTokensServiceInterfaceListProjectAccessTokensCall) DoAndReturn(f func(any, *gitlab.ListProjectAccessTokensOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectAccessToken, *gitlab.Response, error)) *MockProjectAccessTokensServiceInterfaceListProjectAccessTokensCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RevokeProjectAccessToken mocks base method.
func (m *MockProjectAccessTokensServiceInterface) RevokeProjectAccessToken(pid any, id int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RevokeProjectAccessToken", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeProjectAccessToken indicates an expected call of RevokeProjectAccessToken.
func (mr *MockProjectAccessTokensServiceInterfaceMockRecorder) RevokeProjectAccessToken(pid, id any, options ...any) *MockProjectAccessTokensServiceInterfaceRevokeProjectAccessTokenCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, id}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeProjectAccessToken", reflect.TypeOf((*MockProjectAccessTokensServiceInterface)(nil).RevokeProjectAccessToken), varargs...)
	return &MockProjectAccessTokensServiceInterfaceRevokeProjectAccessTokenCall{Call: call}
}

// MockProjectAccessTokensServiceInterfaceRevokeProjectAccessTokenCall wrap *gomock.Call
type MockProjectAccessTokensServiceInterfaceRevokeProjectAccessTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectAccessTokensServiceInterfaceRevokeProjectAccessTokenCall) Return(arg0 *gitlab.Response, arg1 error) *MockProjectAccessTokensServiceInterfaceRevokeProjectAccessTokenCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectAccessTokensServiceInterfaceRevokeProjectAccessTokenCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockProjectAccessTokensServiceInterfaceRevokeProjectAccessTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectAccessTokensServiceInterfaceRevokeProjectAccessTokenCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockProjectAccessTokensServiceInterfaceRevokeProjectAccessTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RotateProjectAccessToken mocks base method.
func (m *MockProjectAccessTokensServiceInterface) RotateProjectAccessToken(pid any, id int, opt *gitlab.RotateProjectAccessTokenOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, id, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RotateProjectAccessToken", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectAccessToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RotateProjectAccessToken indicates an expected call of RotateProjectAccessToken.
func (mr *MockProjectAccessTokensServiceInterfaceMockRecorder) RotateProjectAccessToken(pid, id, opt any, options ...any) *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, id, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateProjectAccessToken", reflect.TypeOf((*MockProjectAccessTokensServiceInterface)(nil).RotateProjectAccessToken), varargs...)
	return &MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenCall{Call: call}
}

// MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenCall wrap *gomock.Call
type MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenCall) Return(arg0 *gitlab.ProjectAccessToken, arg1 *gitlab.Response, arg2 error) *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenCall) Do(f func(any, int, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)) *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenCall) DoAndReturn(f func(any, int, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)) *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RotateProjectAccessTokenSelf mocks base method.
func (m *MockProjectAccessTokensServiceInterface) RotateProjectAccessTokenSelf(pid any, opt *gitlab.RotateProjectAccessTokenOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RotateProjectAccessTokenSelf", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectAccessToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RotateProjectAccessTokenSelf indicates an expected call of RotateProjectAccessTokenSelf.
func (mr *MockProjectAccessTokensServiceInterfaceMockRecorder) RotateProjectAccessTokenSelf(pid, opt any, options ...any) *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenSelfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateProjectAccessTokenSelf", reflect.TypeOf((*MockProjectAccessTokensServiceInterface)(nil).RotateProjectAccessTokenSelf), varargs...)
	return &MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenSelfCall{Call: call}
}

// MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenSelfCall wrap *gomock.Call
type MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenSelfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenSelfCall) Return(arg0 *gitlab.ProjectAccessToken, arg1 *gitlab.Response, arg2 error) *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenSelfCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenSelfCall) Do(f func(any, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)) *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenSelfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenSelfCall) DoAndReturn(f func(any, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)) *MockProjectAccessTokensServiceInterfaceRotateProjectAccessTokenSelfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
