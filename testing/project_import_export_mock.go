// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: ProjectImportExportServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=project_import_export_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ProjectImportExportServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	io "io"
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectImportExportServiceInterface is a mock of ProjectImportExportServiceInterface interface.
type MockProjectImportExportServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProjectImportExportServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockProjectImportExportServiceInterfaceMockRecorder is the mock recorder for MockProjectImportExportServiceInterface.
type MockProjectImportExportServiceInterfaceMockRecorder struct {
	mock *MockProjectImportExportServiceInterface
}

// NewMockProjectImportExportServiceInterface creates a new mock instance.
func NewMockProjectImportExportServiceInterface(ctrl *gomock.Controller) *MockProjectImportExportServiceInterface {
	mock := &MockProjectImportExportServiceInterface{ctrl: ctrl}
	mock.recorder = &MockProjectImportExportServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectImportExportServiceInterface) EXPECT() *MockProjectImportExportServiceInterfaceMockRecorder {
	return m.recorder
}

// ExportDownload mocks base method.
func (m *MockProjectImportExportServiceInterface) ExportDownload(pid any, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExportDownload", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExportDownload indicates an expected call of ExportDownload.
func (mr *MockProjectImportExportServiceInterfaceMockRecorder) ExportDownload(pid any, options ...any) *MockProjectImportExportServiceInterfaceExportDownloadCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportDownload", reflect.TypeOf((*MockProjectImportExportServiceInterface)(nil).ExportDownload), varargs...)
	return &MockProjectImportExportServiceInterfaceExportDownloadCall{Call: call}
}

// MockProjectImportExportServiceInterfaceExportDownloadCall wrap *gomock.Call
type MockProjectImportExportServiceInterfaceExportDownloadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectImportExportServiceInterfaceExportDownloadCall) Return(arg0 []byte, arg1 *gitlab.Response, arg2 error) *MockProjectImportExportServiceInterfaceExportDownloadCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectImportExportServiceInterfaceExportDownloadCall) Do(f func(any, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)) *MockProjectImportExportServiceInterfaceExportDownloadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectImportExportServiceInterfaceExportDownloadCall) DoAndReturn(f func(any, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)) *MockProjectImportExportServiceInterfaceExportDownloadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ExportStatus mocks base method.
func (m *MockProjectImportExportServiceInterface) ExportStatus(pid any, options ...gitlab.RequestOptionFunc) (*gitlab.ExportStatus, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExportStatus", varargs...)
	ret0, _ := ret[0].(*gitlab.ExportStatus)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExportStatus indicates an expected call of ExportStatus.
func (mr *MockProjectImportExportServiceInterfaceMockRecorder) ExportStatus(pid any, options ...any) *MockProjectImportExportServiceInterfaceExportStatusCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportStatus", reflect.TypeOf((*MockProjectImportExportServiceInterface)(nil).ExportStatus), varargs...)
	return &MockProjectImportExportServiceInterfaceExportStatusCall{Call: call}
}

// MockProjectImportExportServiceInterfaceExportStatusCall wrap *gomock.Call
type MockProjectImportExportServiceInterfaceExportStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectImportExportServiceInterfaceExportStatusCall) Return(arg0 *gitlab.ExportStatus, arg1 *gitlab.Response, arg2 error) *MockProjectImportExportServiceInterfaceExportStatusCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectImportExportServiceInterfaceExportStatusCall) Do(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.ExportStatus, *gitlab.Response, error)) *MockProjectImportExportServiceInterfaceExportStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectImportExportServiceInterfaceExportStatusCall) DoAndReturn(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.ExportStatus, *gitlab.Response, error)) *MockProjectImportExportServiceInterfaceExportStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ImportFromFile mocks base method.
func (m *MockProjectImportExportServiceInterface) ImportFromFile(archive io.Reader, opt *gitlab.ImportFileOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ImportStatus, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{archive, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportFromFile", varargs...)
	ret0, _ := ret[0].(*gitlab.ImportStatus)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImportFromFile indicates an expected call of ImportFromFile.
func (mr *MockProjectImportExportServiceInterfaceMockRecorder) ImportFromFile(archive, opt any, options ...any) *MockProjectImportExportServiceInterfaceImportFromFileCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{archive, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportFromFile", reflect.TypeOf((*MockProjectImportExportServiceInterface)(nil).ImportFromFile), varargs...)
	return &MockProjectImportExportServiceInterfaceImportFromFileCall{Call: call}
}

// MockProjectImportExportServiceInterfaceImportFromFileCall wrap *gomock.Call
type MockProjectImportExportServiceInterfaceImportFromFileCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectImportExportServiceInterfaceImportFromFileCall) Return(arg0 *gitlab.ImportStatus, arg1 *gitlab.Response, arg2 error) *MockProjectImportExportServiceInterfaceImportFromFileCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectImportExportServiceInterfaceImportFromFileCall) Do(f func(io.Reader, *gitlab.ImportFileOptions, ...gitlab.RequestOptionFunc) (*gitlab.ImportStatus, *gitlab.Response, error)) *MockProjectImportExportServiceInterfaceImportFromFileCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectImportExportServiceInterfaceImportFromFileCall) DoAndReturn(f func(io.Reader, *gitlab.ImportFileOptions, ...gitlab.RequestOptionFunc) (*gitlab.ImportStatus, *gitlab.Response, error)) *MockProjectImportExportServiceInterfaceImportFromFileCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ImportStatus mocks base method.
func (m *MockProjectImportExportServiceInterface) ImportStatus(pid any, options ...gitlab.RequestOptionFunc) (*gitlab.ImportStatus, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportStatus", varargs...)
	ret0, _ := ret[0].(*gitlab.ImportStatus)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImportStatus indicates an expected call of ImportStatus.
func (mr *MockProjectImportExportServiceInterfaceMockRecorder) ImportStatus(pid any, options ...any) *MockProjectImportExportServiceInterfaceImportStatusCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportStatus", reflect.TypeOf((*MockProjectImportExportServiceInterface)(nil).ImportStatus), varargs...)
	return &MockProjectImportExportServiceInterfaceImportStatusCall{Call: call}
}

// MockProjectImportExportServiceInterfaceImportStatusCall wrap *gomock.Call
type MockProjectImportExportServiceInterfaceImportStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectImportExportServiceInterfaceImportStatusCall) Return(arg0 *gitlab.ImportStatus, arg1 *gitlab.Response, arg2 error) *MockProjectImportExportServiceInterfaceImportStatusCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectImportExportServiceInterfaceImportStatusCall) Do(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.ImportStatus, *gitlab.Response, error)) *MockProjectImportExportServiceInterfaceImportStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectImportExportServiceInterfaceImportStatusCall) DoAndReturn(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.ImportStatus, *gitlab.Response, error)) *MockProjectImportExportServiceInterfaceImportStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ScheduleExport mocks base method.
func (m *MockProjectImportExportServiceInterface) ScheduleExport(pid any, opt *gitlab.ScheduleExportOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScheduleExport", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleExport indicates an expected call of ScheduleExport.
func (mr *MockProjectImportExportServiceInterfaceMockRecorder) ScheduleExport(pid, opt any, options ...any) *MockProjectImportExportServiceInterfaceScheduleExportCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleExport", reflect.TypeOf((*MockProjectImportExportServiceInterface)(nil).ScheduleExport), varargs...)
	return &MockProjectImportExportServiceInterfaceScheduleExportCall{Call: call}
}

// MockProjectImportExportServiceInterfaceScheduleExportCall wrap *gomock.Call
type MockProjectImportExportServiceInterfaceScheduleExportCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectImportExportServiceInterfaceScheduleExportCall) Return(arg0 *gitlab.Response, arg1 error) *MockProjectImportExportServiceInterfaceScheduleExportCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectImportExportServiceInterfaceScheduleExportCall) Do(f func(any, *gitlab.ScheduleExportOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockProjectImportExportServiceInterfaceScheduleExportCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectImportExportServiceInterfaceScheduleExportCall) DoAndReturn(f func(any, *gitlab.ScheduleExportOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockProjectImportExportServiceInterfaceScheduleExportCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
