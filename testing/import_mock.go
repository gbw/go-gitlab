// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: ImportServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=import_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ImportServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockImportServiceInterface is a mock of ImportServiceInterface interface.
type MockImportServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockImportServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockImportServiceInterfaceMockRecorder is the mock recorder for MockImportServiceInterface.
type MockImportServiceInterfaceMockRecorder struct {
	mock *MockImportServiceInterface
}

// NewMockImportServiceInterface creates a new mock instance.
func NewMockImportServiceInterface(ctrl *gomock.Controller) *MockImportServiceInterface {
	mock := &MockImportServiceInterface{ctrl: ctrl}
	mock.recorder = &MockImportServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImportServiceInterface) EXPECT() *MockImportServiceInterfaceMockRecorder {
	return m.recorder
}

// CancelGitHubProjectImport mocks base method.
func (m *MockImportServiceInterface) CancelGitHubProjectImport(opt *gitlab.CancelGitHubProjectImportOptions, options ...gitlab.RequestOptionFunc) (*gitlab.CancelledGitHubImport, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelGitHubProjectImport", varargs...)
	ret0, _ := ret[0].(*gitlab.CancelledGitHubImport)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CancelGitHubProjectImport indicates an expected call of CancelGitHubProjectImport.
func (mr *MockImportServiceInterfaceMockRecorder) CancelGitHubProjectImport(opt any, options ...any) *MockImportServiceInterfaceCancelGitHubProjectImportCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelGitHubProjectImport", reflect.TypeOf((*MockImportServiceInterface)(nil).CancelGitHubProjectImport), varargs...)
	return &MockImportServiceInterfaceCancelGitHubProjectImportCall{Call: call}
}

// MockImportServiceInterfaceCancelGitHubProjectImportCall wrap *gomock.Call
type MockImportServiceInterfaceCancelGitHubProjectImportCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceInterfaceCancelGitHubProjectImportCall) Return(arg0 *gitlab.CancelledGitHubImport, arg1 *gitlab.Response, arg2 error) *MockImportServiceInterfaceCancelGitHubProjectImportCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceInterfaceCancelGitHubProjectImportCall) Do(f func(*gitlab.CancelGitHubProjectImportOptions, ...gitlab.RequestOptionFunc) (*gitlab.CancelledGitHubImport, *gitlab.Response, error)) *MockImportServiceInterfaceCancelGitHubProjectImportCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceInterfaceCancelGitHubProjectImportCall) DoAndReturn(f func(*gitlab.CancelGitHubProjectImportOptions, ...gitlab.RequestOptionFunc) (*gitlab.CancelledGitHubImport, *gitlab.Response, error)) *MockImportServiceInterfaceCancelGitHubProjectImportCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ImportGitHubGistsIntoGitLabSnippets mocks base method.
func (m *MockImportServiceInterface) ImportGitHubGistsIntoGitLabSnippets(opt *gitlab.ImportGitHubGistsIntoGitLabSnippetsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportGitHubGistsIntoGitLabSnippets", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportGitHubGistsIntoGitLabSnippets indicates an expected call of ImportGitHubGistsIntoGitLabSnippets.
func (mr *MockImportServiceInterfaceMockRecorder) ImportGitHubGistsIntoGitLabSnippets(opt any, options ...any) *MockImportServiceInterfaceImportGitHubGistsIntoGitLabSnippetsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportGitHubGistsIntoGitLabSnippets", reflect.TypeOf((*MockImportServiceInterface)(nil).ImportGitHubGistsIntoGitLabSnippets), varargs...)
	return &MockImportServiceInterfaceImportGitHubGistsIntoGitLabSnippetsCall{Call: call}
}

// MockImportServiceInterfaceImportGitHubGistsIntoGitLabSnippetsCall wrap *gomock.Call
type MockImportServiceInterfaceImportGitHubGistsIntoGitLabSnippetsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceInterfaceImportGitHubGistsIntoGitLabSnippetsCall) Return(arg0 *gitlab.Response, arg1 error) *MockImportServiceInterfaceImportGitHubGistsIntoGitLabSnippetsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceInterfaceImportGitHubGistsIntoGitLabSnippetsCall) Do(f func(*gitlab.ImportGitHubGistsIntoGitLabSnippetsOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockImportServiceInterfaceImportGitHubGistsIntoGitLabSnippetsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceInterfaceImportGitHubGistsIntoGitLabSnippetsCall) DoAndReturn(f func(*gitlab.ImportGitHubGistsIntoGitLabSnippetsOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockImportServiceInterfaceImportGitHubGistsIntoGitLabSnippetsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ImportRepositoryFromBitbucketCloud mocks base method.
func (m *MockImportServiceInterface) ImportRepositoryFromBitbucketCloud(opt *gitlab.ImportRepositoryFromBitbucketCloudOptions, options ...gitlab.RequestOptionFunc) (*gitlab.BitbucketCloudImport, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportRepositoryFromBitbucketCloud", varargs...)
	ret0, _ := ret[0].(*gitlab.BitbucketCloudImport)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImportRepositoryFromBitbucketCloud indicates an expected call of ImportRepositoryFromBitbucketCloud.
func (mr *MockImportServiceInterfaceMockRecorder) ImportRepositoryFromBitbucketCloud(opt any, options ...any) *MockImportServiceInterfaceImportRepositoryFromBitbucketCloudCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportRepositoryFromBitbucketCloud", reflect.TypeOf((*MockImportServiceInterface)(nil).ImportRepositoryFromBitbucketCloud), varargs...)
	return &MockImportServiceInterfaceImportRepositoryFromBitbucketCloudCall{Call: call}
}

// MockImportServiceInterfaceImportRepositoryFromBitbucketCloudCall wrap *gomock.Call
type MockImportServiceInterfaceImportRepositoryFromBitbucketCloudCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceInterfaceImportRepositoryFromBitbucketCloudCall) Return(arg0 *gitlab.BitbucketCloudImport, arg1 *gitlab.Response, arg2 error) *MockImportServiceInterfaceImportRepositoryFromBitbucketCloudCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceInterfaceImportRepositoryFromBitbucketCloudCall) Do(f func(*gitlab.ImportRepositoryFromBitbucketCloudOptions, ...gitlab.RequestOptionFunc) (*gitlab.BitbucketCloudImport, *gitlab.Response, error)) *MockImportServiceInterfaceImportRepositoryFromBitbucketCloudCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceInterfaceImportRepositoryFromBitbucketCloudCall) DoAndReturn(f func(*gitlab.ImportRepositoryFromBitbucketCloudOptions, ...gitlab.RequestOptionFunc) (*gitlab.BitbucketCloudImport, *gitlab.Response, error)) *MockImportServiceInterfaceImportRepositoryFromBitbucketCloudCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ImportRepositoryFromBitbucketServer mocks base method.
func (m *MockImportServiceInterface) ImportRepositoryFromBitbucketServer(opt *gitlab.ImportRepositoryFromBitbucketServerOptions, options ...gitlab.RequestOptionFunc) (*gitlab.BitbucketServerImport, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportRepositoryFromBitbucketServer", varargs...)
	ret0, _ := ret[0].(*gitlab.BitbucketServerImport)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImportRepositoryFromBitbucketServer indicates an expected call of ImportRepositoryFromBitbucketServer.
func (mr *MockImportServiceInterfaceMockRecorder) ImportRepositoryFromBitbucketServer(opt any, options ...any) *MockImportServiceInterfaceImportRepositoryFromBitbucketServerCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportRepositoryFromBitbucketServer", reflect.TypeOf((*MockImportServiceInterface)(nil).ImportRepositoryFromBitbucketServer), varargs...)
	return &MockImportServiceInterfaceImportRepositoryFromBitbucketServerCall{Call: call}
}

// MockImportServiceInterfaceImportRepositoryFromBitbucketServerCall wrap *gomock.Call
type MockImportServiceInterfaceImportRepositoryFromBitbucketServerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceInterfaceImportRepositoryFromBitbucketServerCall) Return(arg0 *gitlab.BitbucketServerImport, arg1 *gitlab.Response, arg2 error) *MockImportServiceInterfaceImportRepositoryFromBitbucketServerCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceInterfaceImportRepositoryFromBitbucketServerCall) Do(f func(*gitlab.ImportRepositoryFromBitbucketServerOptions, ...gitlab.RequestOptionFunc) (*gitlab.BitbucketServerImport, *gitlab.Response, error)) *MockImportServiceInterfaceImportRepositoryFromBitbucketServerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceInterfaceImportRepositoryFromBitbucketServerCall) DoAndReturn(f func(*gitlab.ImportRepositoryFromBitbucketServerOptions, ...gitlab.RequestOptionFunc) (*gitlab.BitbucketServerImport, *gitlab.Response, error)) *MockImportServiceInterfaceImportRepositoryFromBitbucketServerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ImportRepositoryFromGitHub mocks base method.
func (m *MockImportServiceInterface) ImportRepositoryFromGitHub(opt *gitlab.ImportRepositoryFromGitHubOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GitHubImport, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportRepositoryFromGitHub", varargs...)
	ret0, _ := ret[0].(*gitlab.GitHubImport)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImportRepositoryFromGitHub indicates an expected call of ImportRepositoryFromGitHub.
func (mr *MockImportServiceInterfaceMockRecorder) ImportRepositoryFromGitHub(opt any, options ...any) *MockImportServiceInterfaceImportRepositoryFromGitHubCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportRepositoryFromGitHub", reflect.TypeOf((*MockImportServiceInterface)(nil).ImportRepositoryFromGitHub), varargs...)
	return &MockImportServiceInterfaceImportRepositoryFromGitHubCall{Call: call}
}

// MockImportServiceInterfaceImportRepositoryFromGitHubCall wrap *gomock.Call
type MockImportServiceInterfaceImportRepositoryFromGitHubCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceInterfaceImportRepositoryFromGitHubCall) Return(arg0 *gitlab.GitHubImport, arg1 *gitlab.Response, arg2 error) *MockImportServiceInterfaceImportRepositoryFromGitHubCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceInterfaceImportRepositoryFromGitHubCall) Do(f func(*gitlab.ImportRepositoryFromGitHubOptions, ...gitlab.RequestOptionFunc) (*gitlab.GitHubImport, *gitlab.Response, error)) *MockImportServiceInterfaceImportRepositoryFromGitHubCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceInterfaceImportRepositoryFromGitHubCall) DoAndReturn(f func(*gitlab.ImportRepositoryFromGitHubOptions, ...gitlab.RequestOptionFunc) (*gitlab.GitHubImport, *gitlab.Response, error)) *MockImportServiceInterfaceImportRepositoryFromGitHubCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
