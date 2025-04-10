// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: ProjectSnippetsServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=project_snippets_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectSnippetsServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectSnippetsServiceInterface is a mock of ProjectSnippetsServiceInterface interface.
type MockProjectSnippetsServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProjectSnippetsServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockProjectSnippetsServiceInterfaceMockRecorder is the mock recorder for MockProjectSnippetsServiceInterface.
type MockProjectSnippetsServiceInterfaceMockRecorder struct {
	mock *MockProjectSnippetsServiceInterface
}

// NewMockProjectSnippetsServiceInterface creates a new mock instance.
func NewMockProjectSnippetsServiceInterface(ctrl *gomock.Controller) *MockProjectSnippetsServiceInterface {
	mock := &MockProjectSnippetsServiceInterface{ctrl: ctrl}
	mock.recorder = &MockProjectSnippetsServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectSnippetsServiceInterface) EXPECT() *MockProjectSnippetsServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateSnippet mocks base method.
func (m *MockProjectSnippetsServiceInterface) CreateSnippet(pid any, opt *gitlab.CreateProjectSnippetOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSnippet indicates an expected call of CreateSnippet.
func (mr *MockProjectSnippetsServiceInterfaceMockRecorder) CreateSnippet(pid, opt any, options ...any) *MockProjectSnippetsServiceInterfaceCreateSnippetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnippet", reflect.TypeOf((*MockProjectSnippetsServiceInterface)(nil).CreateSnippet), varargs...)
	return &MockProjectSnippetsServiceInterfaceCreateSnippetCall{Call: call}
}

// MockProjectSnippetsServiceInterfaceCreateSnippetCall wrap *gomock.Call
type MockProjectSnippetsServiceInterfaceCreateSnippetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectSnippetsServiceInterfaceCreateSnippetCall) Return(arg0 *gitlab.Snippet, arg1 *gitlab.Response, arg2 error) *MockProjectSnippetsServiceInterfaceCreateSnippetCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectSnippetsServiceInterfaceCreateSnippetCall) Do(f func(any, *gitlab.CreateProjectSnippetOptions, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *MockProjectSnippetsServiceInterfaceCreateSnippetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectSnippetsServiceInterfaceCreateSnippetCall) DoAndReturn(f func(any, *gitlab.CreateProjectSnippetOptions, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *MockProjectSnippetsServiceInterfaceCreateSnippetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteSnippet mocks base method.
func (m *MockProjectSnippetsServiceInterface) DeleteSnippet(pid any, snippet int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, snippet}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSnippet indicates an expected call of DeleteSnippet.
func (mr *MockProjectSnippetsServiceInterfaceMockRecorder) DeleteSnippet(pid, snippet any, options ...any) *MockProjectSnippetsServiceInterfaceDeleteSnippetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, snippet}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnippet", reflect.TypeOf((*MockProjectSnippetsServiceInterface)(nil).DeleteSnippet), varargs...)
	return &MockProjectSnippetsServiceInterfaceDeleteSnippetCall{Call: call}
}

// MockProjectSnippetsServiceInterfaceDeleteSnippetCall wrap *gomock.Call
type MockProjectSnippetsServiceInterfaceDeleteSnippetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectSnippetsServiceInterfaceDeleteSnippetCall) Return(arg0 *gitlab.Response, arg1 error) *MockProjectSnippetsServiceInterfaceDeleteSnippetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectSnippetsServiceInterfaceDeleteSnippetCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockProjectSnippetsServiceInterfaceDeleteSnippetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectSnippetsServiceInterfaceDeleteSnippetCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockProjectSnippetsServiceInterfaceDeleteSnippetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSnippet mocks base method.
func (m *MockProjectSnippetsServiceInterface) GetSnippet(pid any, snippet int, options ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, snippet}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSnippet indicates an expected call of GetSnippet.
func (mr *MockProjectSnippetsServiceInterfaceMockRecorder) GetSnippet(pid, snippet any, options ...any) *MockProjectSnippetsServiceInterfaceGetSnippetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, snippet}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnippet", reflect.TypeOf((*MockProjectSnippetsServiceInterface)(nil).GetSnippet), varargs...)
	return &MockProjectSnippetsServiceInterfaceGetSnippetCall{Call: call}
}

// MockProjectSnippetsServiceInterfaceGetSnippetCall wrap *gomock.Call
type MockProjectSnippetsServiceInterfaceGetSnippetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectSnippetsServiceInterfaceGetSnippetCall) Return(arg0 *gitlab.Snippet, arg1 *gitlab.Response, arg2 error) *MockProjectSnippetsServiceInterfaceGetSnippetCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectSnippetsServiceInterfaceGetSnippetCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *MockProjectSnippetsServiceInterfaceGetSnippetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectSnippetsServiceInterfaceGetSnippetCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *MockProjectSnippetsServiceInterfaceGetSnippetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListSnippets mocks base method.
func (m *MockProjectSnippetsServiceInterface) ListSnippets(pid any, opt *gitlab.ListProjectSnippetsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSnippets", varargs...)
	ret0, _ := ret[0].([]*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListSnippets indicates an expected call of ListSnippets.
func (mr *MockProjectSnippetsServiceInterfaceMockRecorder) ListSnippets(pid, opt any, options ...any) *MockProjectSnippetsServiceInterfaceListSnippetsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSnippets", reflect.TypeOf((*MockProjectSnippetsServiceInterface)(nil).ListSnippets), varargs...)
	return &MockProjectSnippetsServiceInterfaceListSnippetsCall{Call: call}
}

// MockProjectSnippetsServiceInterfaceListSnippetsCall wrap *gomock.Call
type MockProjectSnippetsServiceInterfaceListSnippetsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectSnippetsServiceInterfaceListSnippetsCall) Return(arg0 []*gitlab.Snippet, arg1 *gitlab.Response, arg2 error) *MockProjectSnippetsServiceInterfaceListSnippetsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectSnippetsServiceInterfaceListSnippetsCall) Do(f func(any, *gitlab.ListProjectSnippetsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error)) *MockProjectSnippetsServiceInterfaceListSnippetsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectSnippetsServiceInterfaceListSnippetsCall) DoAndReturn(f func(any, *gitlab.ListProjectSnippetsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error)) *MockProjectSnippetsServiceInterfaceListSnippetsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SnippetContent mocks base method.
func (m *MockProjectSnippetsServiceInterface) SnippetContent(pid any, snippet int, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, snippet}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SnippetContent", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SnippetContent indicates an expected call of SnippetContent.
func (mr *MockProjectSnippetsServiceInterfaceMockRecorder) SnippetContent(pid, snippet any, options ...any) *MockProjectSnippetsServiceInterfaceSnippetContentCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, snippet}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnippetContent", reflect.TypeOf((*MockProjectSnippetsServiceInterface)(nil).SnippetContent), varargs...)
	return &MockProjectSnippetsServiceInterfaceSnippetContentCall{Call: call}
}

// MockProjectSnippetsServiceInterfaceSnippetContentCall wrap *gomock.Call
type MockProjectSnippetsServiceInterfaceSnippetContentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectSnippetsServiceInterfaceSnippetContentCall) Return(arg0 []byte, arg1 *gitlab.Response, arg2 error) *MockProjectSnippetsServiceInterfaceSnippetContentCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectSnippetsServiceInterfaceSnippetContentCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)) *MockProjectSnippetsServiceInterfaceSnippetContentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectSnippetsServiceInterfaceSnippetContentCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)) *MockProjectSnippetsServiceInterfaceSnippetContentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateSnippet mocks base method.
func (m *MockProjectSnippetsServiceInterface) UpdateSnippet(pid any, snippet int, opt *gitlab.UpdateProjectSnippetOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, snippet, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateSnippet indicates an expected call of UpdateSnippet.
func (mr *MockProjectSnippetsServiceInterfaceMockRecorder) UpdateSnippet(pid, snippet, opt any, options ...any) *MockProjectSnippetsServiceInterfaceUpdateSnippetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, snippet, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSnippet", reflect.TypeOf((*MockProjectSnippetsServiceInterface)(nil).UpdateSnippet), varargs...)
	return &MockProjectSnippetsServiceInterfaceUpdateSnippetCall{Call: call}
}

// MockProjectSnippetsServiceInterfaceUpdateSnippetCall wrap *gomock.Call
type MockProjectSnippetsServiceInterfaceUpdateSnippetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectSnippetsServiceInterfaceUpdateSnippetCall) Return(arg0 *gitlab.Snippet, arg1 *gitlab.Response, arg2 error) *MockProjectSnippetsServiceInterfaceUpdateSnippetCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectSnippetsServiceInterfaceUpdateSnippetCall) Do(f func(any, int, *gitlab.UpdateProjectSnippetOptions, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *MockProjectSnippetsServiceInterfaceUpdateSnippetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectSnippetsServiceInterfaceUpdateSnippetCall) DoAndReturn(f func(any, int, *gitlab.UpdateProjectSnippetOptions, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *MockProjectSnippetsServiceInterfaceUpdateSnippetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
