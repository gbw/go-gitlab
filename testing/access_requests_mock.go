// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: AccessRequestsServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=access_requests_mock.go -package=testing gitlab.com/gitlab-org/api/client-go AccessRequestsServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockAccessRequestsServiceInterface is a mock of AccessRequestsServiceInterface interface.
type MockAccessRequestsServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAccessRequestsServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockAccessRequestsServiceInterfaceMockRecorder is the mock recorder for MockAccessRequestsServiceInterface.
type MockAccessRequestsServiceInterfaceMockRecorder struct {
	mock *MockAccessRequestsServiceInterface
}

// NewMockAccessRequestsServiceInterface creates a new mock instance.
func NewMockAccessRequestsServiceInterface(ctrl *gomock.Controller) *MockAccessRequestsServiceInterface {
	mock := &MockAccessRequestsServiceInterface{ctrl: ctrl}
	mock.recorder = &MockAccessRequestsServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessRequestsServiceInterface) EXPECT() *MockAccessRequestsServiceInterfaceMockRecorder {
	return m.recorder
}

// ApproveGroupAccessRequest mocks base method.
func (m *MockAccessRequestsServiceInterface) ApproveGroupAccessRequest(gid any, user int, opt *gitlab.ApproveAccessRequestOptions, options ...gitlab.RequestOptionFunc) (*gitlab.AccessRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, user, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ApproveGroupAccessRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.AccessRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ApproveGroupAccessRequest indicates an expected call of ApproveGroupAccessRequest.
func (mr *MockAccessRequestsServiceInterfaceMockRecorder) ApproveGroupAccessRequest(gid, user, opt any, options ...any) *MockAccessRequestsServiceInterfaceApproveGroupAccessRequestCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, user, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveGroupAccessRequest", reflect.TypeOf((*MockAccessRequestsServiceInterface)(nil).ApproveGroupAccessRequest), varargs...)
	return &MockAccessRequestsServiceInterfaceApproveGroupAccessRequestCall{Call: call}
}

// MockAccessRequestsServiceInterfaceApproveGroupAccessRequestCall wrap *gomock.Call
type MockAccessRequestsServiceInterfaceApproveGroupAccessRequestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAccessRequestsServiceInterfaceApproveGroupAccessRequestCall) Return(arg0 *gitlab.AccessRequest, arg1 *gitlab.Response, arg2 error) *MockAccessRequestsServiceInterfaceApproveGroupAccessRequestCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAccessRequestsServiceInterfaceApproveGroupAccessRequestCall) Do(f func(any, int, *gitlab.ApproveAccessRequestOptions, ...gitlab.RequestOptionFunc) (*gitlab.AccessRequest, *gitlab.Response, error)) *MockAccessRequestsServiceInterfaceApproveGroupAccessRequestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAccessRequestsServiceInterfaceApproveGroupAccessRequestCall) DoAndReturn(f func(any, int, *gitlab.ApproveAccessRequestOptions, ...gitlab.RequestOptionFunc) (*gitlab.AccessRequest, *gitlab.Response, error)) *MockAccessRequestsServiceInterfaceApproveGroupAccessRequestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ApproveProjectAccessRequest mocks base method.
func (m *MockAccessRequestsServiceInterface) ApproveProjectAccessRequest(pid any, user int, opt *gitlab.ApproveAccessRequestOptions, options ...gitlab.RequestOptionFunc) (*gitlab.AccessRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, user, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ApproveProjectAccessRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.AccessRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ApproveProjectAccessRequest indicates an expected call of ApproveProjectAccessRequest.
func (mr *MockAccessRequestsServiceInterfaceMockRecorder) ApproveProjectAccessRequest(pid, user, opt any, options ...any) *MockAccessRequestsServiceInterfaceApproveProjectAccessRequestCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, user, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveProjectAccessRequest", reflect.TypeOf((*MockAccessRequestsServiceInterface)(nil).ApproveProjectAccessRequest), varargs...)
	return &MockAccessRequestsServiceInterfaceApproveProjectAccessRequestCall{Call: call}
}

// MockAccessRequestsServiceInterfaceApproveProjectAccessRequestCall wrap *gomock.Call
type MockAccessRequestsServiceInterfaceApproveProjectAccessRequestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAccessRequestsServiceInterfaceApproveProjectAccessRequestCall) Return(arg0 *gitlab.AccessRequest, arg1 *gitlab.Response, arg2 error) *MockAccessRequestsServiceInterfaceApproveProjectAccessRequestCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAccessRequestsServiceInterfaceApproveProjectAccessRequestCall) Do(f func(any, int, *gitlab.ApproveAccessRequestOptions, ...gitlab.RequestOptionFunc) (*gitlab.AccessRequest, *gitlab.Response, error)) *MockAccessRequestsServiceInterfaceApproveProjectAccessRequestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAccessRequestsServiceInterfaceApproveProjectAccessRequestCall) DoAndReturn(f func(any, int, *gitlab.ApproveAccessRequestOptions, ...gitlab.RequestOptionFunc) (*gitlab.AccessRequest, *gitlab.Response, error)) *MockAccessRequestsServiceInterfaceApproveProjectAccessRequestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DenyGroupAccessRequest mocks base method.
func (m *MockAccessRequestsServiceInterface) DenyGroupAccessRequest(gid any, user int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, user}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DenyGroupAccessRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DenyGroupAccessRequest indicates an expected call of DenyGroupAccessRequest.
func (mr *MockAccessRequestsServiceInterfaceMockRecorder) DenyGroupAccessRequest(gid, user any, options ...any) *MockAccessRequestsServiceInterfaceDenyGroupAccessRequestCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, user}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DenyGroupAccessRequest", reflect.TypeOf((*MockAccessRequestsServiceInterface)(nil).DenyGroupAccessRequest), varargs...)
	return &MockAccessRequestsServiceInterfaceDenyGroupAccessRequestCall{Call: call}
}

// MockAccessRequestsServiceInterfaceDenyGroupAccessRequestCall wrap *gomock.Call
type MockAccessRequestsServiceInterfaceDenyGroupAccessRequestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAccessRequestsServiceInterfaceDenyGroupAccessRequestCall) Return(arg0 *gitlab.Response, arg1 error) *MockAccessRequestsServiceInterfaceDenyGroupAccessRequestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAccessRequestsServiceInterfaceDenyGroupAccessRequestCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockAccessRequestsServiceInterfaceDenyGroupAccessRequestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAccessRequestsServiceInterfaceDenyGroupAccessRequestCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockAccessRequestsServiceInterfaceDenyGroupAccessRequestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DenyProjectAccessRequest mocks base method.
func (m *MockAccessRequestsServiceInterface) DenyProjectAccessRequest(pid any, user int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, user}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DenyProjectAccessRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DenyProjectAccessRequest indicates an expected call of DenyProjectAccessRequest.
func (mr *MockAccessRequestsServiceInterfaceMockRecorder) DenyProjectAccessRequest(pid, user any, options ...any) *MockAccessRequestsServiceInterfaceDenyProjectAccessRequestCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, user}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DenyProjectAccessRequest", reflect.TypeOf((*MockAccessRequestsServiceInterface)(nil).DenyProjectAccessRequest), varargs...)
	return &MockAccessRequestsServiceInterfaceDenyProjectAccessRequestCall{Call: call}
}

// MockAccessRequestsServiceInterfaceDenyProjectAccessRequestCall wrap *gomock.Call
type MockAccessRequestsServiceInterfaceDenyProjectAccessRequestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAccessRequestsServiceInterfaceDenyProjectAccessRequestCall) Return(arg0 *gitlab.Response, arg1 error) *MockAccessRequestsServiceInterfaceDenyProjectAccessRequestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAccessRequestsServiceInterfaceDenyProjectAccessRequestCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockAccessRequestsServiceInterfaceDenyProjectAccessRequestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAccessRequestsServiceInterfaceDenyProjectAccessRequestCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockAccessRequestsServiceInterfaceDenyProjectAccessRequestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListGroupAccessRequests mocks base method.
func (m *MockAccessRequestsServiceInterface) ListGroupAccessRequests(gid any, opt *gitlab.ListAccessRequestsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.AccessRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupAccessRequests", varargs...)
	ret0, _ := ret[0].([]*gitlab.AccessRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGroupAccessRequests indicates an expected call of ListGroupAccessRequests.
func (mr *MockAccessRequestsServiceInterfaceMockRecorder) ListGroupAccessRequests(gid, opt any, options ...any) *MockAccessRequestsServiceInterfaceListGroupAccessRequestsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupAccessRequests", reflect.TypeOf((*MockAccessRequestsServiceInterface)(nil).ListGroupAccessRequests), varargs...)
	return &MockAccessRequestsServiceInterfaceListGroupAccessRequestsCall{Call: call}
}

// MockAccessRequestsServiceInterfaceListGroupAccessRequestsCall wrap *gomock.Call
type MockAccessRequestsServiceInterfaceListGroupAccessRequestsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAccessRequestsServiceInterfaceListGroupAccessRequestsCall) Return(arg0 []*gitlab.AccessRequest, arg1 *gitlab.Response, arg2 error) *MockAccessRequestsServiceInterfaceListGroupAccessRequestsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAccessRequestsServiceInterfaceListGroupAccessRequestsCall) Do(f func(any, *gitlab.ListAccessRequestsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.AccessRequest, *gitlab.Response, error)) *MockAccessRequestsServiceInterfaceListGroupAccessRequestsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAccessRequestsServiceInterfaceListGroupAccessRequestsCall) DoAndReturn(f func(any, *gitlab.ListAccessRequestsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.AccessRequest, *gitlab.Response, error)) *MockAccessRequestsServiceInterfaceListGroupAccessRequestsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListProjectAccessRequests mocks base method.
func (m *MockAccessRequestsServiceInterface) ListProjectAccessRequests(pid any, opt *gitlab.ListAccessRequestsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.AccessRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectAccessRequests", varargs...)
	ret0, _ := ret[0].([]*gitlab.AccessRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectAccessRequests indicates an expected call of ListProjectAccessRequests.
func (mr *MockAccessRequestsServiceInterfaceMockRecorder) ListProjectAccessRequests(pid, opt any, options ...any) *MockAccessRequestsServiceInterfaceListProjectAccessRequestsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectAccessRequests", reflect.TypeOf((*MockAccessRequestsServiceInterface)(nil).ListProjectAccessRequests), varargs...)
	return &MockAccessRequestsServiceInterfaceListProjectAccessRequestsCall{Call: call}
}

// MockAccessRequestsServiceInterfaceListProjectAccessRequestsCall wrap *gomock.Call
type MockAccessRequestsServiceInterfaceListProjectAccessRequestsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAccessRequestsServiceInterfaceListProjectAccessRequestsCall) Return(arg0 []*gitlab.AccessRequest, arg1 *gitlab.Response, arg2 error) *MockAccessRequestsServiceInterfaceListProjectAccessRequestsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAccessRequestsServiceInterfaceListProjectAccessRequestsCall) Do(f func(any, *gitlab.ListAccessRequestsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.AccessRequest, *gitlab.Response, error)) *MockAccessRequestsServiceInterfaceListProjectAccessRequestsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAccessRequestsServiceInterfaceListProjectAccessRequestsCall) DoAndReturn(f func(any, *gitlab.ListAccessRequestsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.AccessRequest, *gitlab.Response, error)) *MockAccessRequestsServiceInterfaceListProjectAccessRequestsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RequestGroupAccess mocks base method.
func (m *MockAccessRequestsServiceInterface) RequestGroupAccess(gid any, options ...gitlab.RequestOptionFunc) (*gitlab.AccessRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RequestGroupAccess", varargs...)
	ret0, _ := ret[0].(*gitlab.AccessRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RequestGroupAccess indicates an expected call of RequestGroupAccess.
func (mr *MockAccessRequestsServiceInterfaceMockRecorder) RequestGroupAccess(gid any, options ...any) *MockAccessRequestsServiceInterfaceRequestGroupAccessCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestGroupAccess", reflect.TypeOf((*MockAccessRequestsServiceInterface)(nil).RequestGroupAccess), varargs...)
	return &MockAccessRequestsServiceInterfaceRequestGroupAccessCall{Call: call}
}

// MockAccessRequestsServiceInterfaceRequestGroupAccessCall wrap *gomock.Call
type MockAccessRequestsServiceInterfaceRequestGroupAccessCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAccessRequestsServiceInterfaceRequestGroupAccessCall) Return(arg0 *gitlab.AccessRequest, arg1 *gitlab.Response, arg2 error) *MockAccessRequestsServiceInterfaceRequestGroupAccessCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAccessRequestsServiceInterfaceRequestGroupAccessCall) Do(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.AccessRequest, *gitlab.Response, error)) *MockAccessRequestsServiceInterfaceRequestGroupAccessCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAccessRequestsServiceInterfaceRequestGroupAccessCall) DoAndReturn(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.AccessRequest, *gitlab.Response, error)) *MockAccessRequestsServiceInterfaceRequestGroupAccessCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RequestProjectAccess mocks base method.
func (m *MockAccessRequestsServiceInterface) RequestProjectAccess(pid any, options ...gitlab.RequestOptionFunc) (*gitlab.AccessRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RequestProjectAccess", varargs...)
	ret0, _ := ret[0].(*gitlab.AccessRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RequestProjectAccess indicates an expected call of RequestProjectAccess.
func (mr *MockAccessRequestsServiceInterfaceMockRecorder) RequestProjectAccess(pid any, options ...any) *MockAccessRequestsServiceInterfaceRequestProjectAccessCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestProjectAccess", reflect.TypeOf((*MockAccessRequestsServiceInterface)(nil).RequestProjectAccess), varargs...)
	return &MockAccessRequestsServiceInterfaceRequestProjectAccessCall{Call: call}
}

// MockAccessRequestsServiceInterfaceRequestProjectAccessCall wrap *gomock.Call
type MockAccessRequestsServiceInterfaceRequestProjectAccessCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAccessRequestsServiceInterfaceRequestProjectAccessCall) Return(arg0 *gitlab.AccessRequest, arg1 *gitlab.Response, arg2 error) *MockAccessRequestsServiceInterfaceRequestProjectAccessCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAccessRequestsServiceInterfaceRequestProjectAccessCall) Do(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.AccessRequest, *gitlab.Response, error)) *MockAccessRequestsServiceInterfaceRequestProjectAccessCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAccessRequestsServiceInterfaceRequestProjectAccessCall) DoAndReturn(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.AccessRequest, *gitlab.Response, error)) *MockAccessRequestsServiceInterfaceRequestProjectAccessCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
