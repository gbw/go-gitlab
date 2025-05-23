// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: DependencyListExportServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=dependency_list_export_mock.go -package=testing gitlab.com/gitlab-org/api/client-go DependencyListExportServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	io "io"
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockDependencyListExportServiceInterface is a mock of DependencyListExportServiceInterface interface.
type MockDependencyListExportServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDependencyListExportServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockDependencyListExportServiceInterfaceMockRecorder is the mock recorder for MockDependencyListExportServiceInterface.
type MockDependencyListExportServiceInterfaceMockRecorder struct {
	mock *MockDependencyListExportServiceInterface
}

// NewMockDependencyListExportServiceInterface creates a new mock instance.
func NewMockDependencyListExportServiceInterface(ctrl *gomock.Controller) *MockDependencyListExportServiceInterface {
	mock := &MockDependencyListExportServiceInterface{ctrl: ctrl}
	mock.recorder = &MockDependencyListExportServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDependencyListExportServiceInterface) EXPECT() *MockDependencyListExportServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateDependencyListExport mocks base method.
func (m *MockDependencyListExportServiceInterface) CreateDependencyListExport(pipelineID int, opt *gitlab.CreateDependencyListExportOptions, options ...gitlab.RequestOptionFunc) (*gitlab.DependencyListExport, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pipelineID, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateDependencyListExport", varargs...)
	ret0, _ := ret[0].(*gitlab.DependencyListExport)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateDependencyListExport indicates an expected call of CreateDependencyListExport.
func (mr *MockDependencyListExportServiceInterfaceMockRecorder) CreateDependencyListExport(pipelineID, opt any, options ...any) *MockDependencyListExportServiceInterfaceCreateDependencyListExportCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pipelineID, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDependencyListExport", reflect.TypeOf((*MockDependencyListExportServiceInterface)(nil).CreateDependencyListExport), varargs...)
	return &MockDependencyListExportServiceInterfaceCreateDependencyListExportCall{Call: call}
}

// MockDependencyListExportServiceInterfaceCreateDependencyListExportCall wrap *gomock.Call
type MockDependencyListExportServiceInterfaceCreateDependencyListExportCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDependencyListExportServiceInterfaceCreateDependencyListExportCall) Return(arg0 *gitlab.DependencyListExport, arg1 *gitlab.Response, arg2 error) *MockDependencyListExportServiceInterfaceCreateDependencyListExportCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDependencyListExportServiceInterfaceCreateDependencyListExportCall) Do(f func(int, *gitlab.CreateDependencyListExportOptions, ...gitlab.RequestOptionFunc) (*gitlab.DependencyListExport, *gitlab.Response, error)) *MockDependencyListExportServiceInterfaceCreateDependencyListExportCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDependencyListExportServiceInterfaceCreateDependencyListExportCall) DoAndReturn(f func(int, *gitlab.CreateDependencyListExportOptions, ...gitlab.RequestOptionFunc) (*gitlab.DependencyListExport, *gitlab.Response, error)) *MockDependencyListExportServiceInterfaceCreateDependencyListExportCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DownloadDependencyListExport mocks base method.
func (m *MockDependencyListExportServiceInterface) DownloadDependencyListExport(id int, options ...gitlab.RequestOptionFunc) (io.Reader, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadDependencyListExport", varargs...)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadDependencyListExport indicates an expected call of DownloadDependencyListExport.
func (mr *MockDependencyListExportServiceInterfaceMockRecorder) DownloadDependencyListExport(id any, options ...any) *MockDependencyListExportServiceInterfaceDownloadDependencyListExportCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadDependencyListExport", reflect.TypeOf((*MockDependencyListExportServiceInterface)(nil).DownloadDependencyListExport), varargs...)
	return &MockDependencyListExportServiceInterfaceDownloadDependencyListExportCall{Call: call}
}

// MockDependencyListExportServiceInterfaceDownloadDependencyListExportCall wrap *gomock.Call
type MockDependencyListExportServiceInterfaceDownloadDependencyListExportCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDependencyListExportServiceInterfaceDownloadDependencyListExportCall) Return(arg0 io.Reader, arg1 *gitlab.Response, arg2 error) *MockDependencyListExportServiceInterfaceDownloadDependencyListExportCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDependencyListExportServiceInterfaceDownloadDependencyListExportCall) Do(f func(int, ...gitlab.RequestOptionFunc) (io.Reader, *gitlab.Response, error)) *MockDependencyListExportServiceInterfaceDownloadDependencyListExportCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDependencyListExportServiceInterfaceDownloadDependencyListExportCall) DoAndReturn(f func(int, ...gitlab.RequestOptionFunc) (io.Reader, *gitlab.Response, error)) *MockDependencyListExportServiceInterfaceDownloadDependencyListExportCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetDependencyListExport mocks base method.
func (m *MockDependencyListExportServiceInterface) GetDependencyListExport(id int, options ...gitlab.RequestOptionFunc) (*gitlab.DependencyListExport, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDependencyListExport", varargs...)
	ret0, _ := ret[0].(*gitlab.DependencyListExport)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDependencyListExport indicates an expected call of GetDependencyListExport.
func (mr *MockDependencyListExportServiceInterfaceMockRecorder) GetDependencyListExport(id any, options ...any) *MockDependencyListExportServiceInterfaceGetDependencyListExportCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDependencyListExport", reflect.TypeOf((*MockDependencyListExportServiceInterface)(nil).GetDependencyListExport), varargs...)
	return &MockDependencyListExportServiceInterfaceGetDependencyListExportCall{Call: call}
}

// MockDependencyListExportServiceInterfaceGetDependencyListExportCall wrap *gomock.Call
type MockDependencyListExportServiceInterfaceGetDependencyListExportCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDependencyListExportServiceInterfaceGetDependencyListExportCall) Return(arg0 *gitlab.DependencyListExport, arg1 *gitlab.Response, arg2 error) *MockDependencyListExportServiceInterfaceGetDependencyListExportCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDependencyListExportServiceInterfaceGetDependencyListExportCall) Do(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.DependencyListExport, *gitlab.Response, error)) *MockDependencyListExportServiceInterfaceGetDependencyListExportCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDependencyListExportServiceInterfaceGetDependencyListExportCall) DoAndReturn(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.DependencyListExport, *gitlab.Response, error)) *MockDependencyListExportServiceInterfaceGetDependencyListExportCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
