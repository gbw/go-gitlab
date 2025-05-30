// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: GroupMarkdownUploadsServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=group_markdown_uploads_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GroupMarkdownUploadsServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	io "io"
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockGroupMarkdownUploadsServiceInterface is a mock of GroupMarkdownUploadsServiceInterface interface.
type MockGroupMarkdownUploadsServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockGroupMarkdownUploadsServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockGroupMarkdownUploadsServiceInterfaceMockRecorder is the mock recorder for MockGroupMarkdownUploadsServiceInterface.
type MockGroupMarkdownUploadsServiceInterfaceMockRecorder struct {
	mock *MockGroupMarkdownUploadsServiceInterface
}

// NewMockGroupMarkdownUploadsServiceInterface creates a new mock instance.
func NewMockGroupMarkdownUploadsServiceInterface(ctrl *gomock.Controller) *MockGroupMarkdownUploadsServiceInterface {
	mock := &MockGroupMarkdownUploadsServiceInterface{ctrl: ctrl}
	mock.recorder = &MockGroupMarkdownUploadsServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupMarkdownUploadsServiceInterface) EXPECT() *MockGroupMarkdownUploadsServiceInterfaceMockRecorder {
	return m.recorder
}

// DeleteGroupMarkdownUploadByID mocks base method.
func (m *MockGroupMarkdownUploadsServiceInterface) DeleteGroupMarkdownUploadByID(gid any, uploadID int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, uploadID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteGroupMarkdownUploadByID", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGroupMarkdownUploadByID indicates an expected call of DeleteGroupMarkdownUploadByID.
func (mr *MockGroupMarkdownUploadsServiceInterfaceMockRecorder) DeleteGroupMarkdownUploadByID(gid, uploadID any, options ...any) *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadByIDCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, uploadID}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupMarkdownUploadByID", reflect.TypeOf((*MockGroupMarkdownUploadsServiceInterface)(nil).DeleteGroupMarkdownUploadByID), varargs...)
	return &MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadByIDCall{Call: call}
}

// MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadByIDCall wrap *gomock.Call
type MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadByIDCall) Return(arg0 *gitlab.Response, arg1 error) *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadByIDCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadByIDCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteGroupMarkdownUploadBySecretAndFilename mocks base method.
func (m *MockGroupMarkdownUploadsServiceInterface) DeleteGroupMarkdownUploadBySecretAndFilename(gid any, secret, filename string, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, secret, filename}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteGroupMarkdownUploadBySecretAndFilename", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGroupMarkdownUploadBySecretAndFilename indicates an expected call of DeleteGroupMarkdownUploadBySecretAndFilename.
func (mr *MockGroupMarkdownUploadsServiceInterfaceMockRecorder) DeleteGroupMarkdownUploadBySecretAndFilename(gid, secret, filename any, options ...any) *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadBySecretAndFilenameCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, secret, filename}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupMarkdownUploadBySecretAndFilename", reflect.TypeOf((*MockGroupMarkdownUploadsServiceInterface)(nil).DeleteGroupMarkdownUploadBySecretAndFilename), varargs...)
	return &MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadBySecretAndFilenameCall{Call: call}
}

// MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadBySecretAndFilenameCall wrap *gomock.Call
type MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadBySecretAndFilenameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadBySecretAndFilenameCall) Return(arg0 *gitlab.Response, arg1 error) *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadBySecretAndFilenameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadBySecretAndFilenameCall) Do(f func(any, string, string, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadBySecretAndFilenameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadBySecretAndFilenameCall) DoAndReturn(f func(any, string, string, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockGroupMarkdownUploadsServiceInterfaceDeleteGroupMarkdownUploadBySecretAndFilenameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DownloadGroupMarkdownUploadByID mocks base method.
func (m *MockGroupMarkdownUploadsServiceInterface) DownloadGroupMarkdownUploadByID(gid any, uploadID int, options ...gitlab.RequestOptionFunc) (io.Reader, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, uploadID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadGroupMarkdownUploadByID", varargs...)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadGroupMarkdownUploadByID indicates an expected call of DownloadGroupMarkdownUploadByID.
func (mr *MockGroupMarkdownUploadsServiceInterfaceMockRecorder) DownloadGroupMarkdownUploadByID(gid, uploadID any, options ...any) *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadByIDCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, uploadID}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadGroupMarkdownUploadByID", reflect.TypeOf((*MockGroupMarkdownUploadsServiceInterface)(nil).DownloadGroupMarkdownUploadByID), varargs...)
	return &MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadByIDCall{Call: call}
}

// MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadByIDCall wrap *gomock.Call
type MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadByIDCall) Return(arg0 io.Reader, arg1 *gitlab.Response, arg2 error) *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadByIDCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadByIDCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (io.Reader, *gitlab.Response, error)) *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadByIDCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (io.Reader, *gitlab.Response, error)) *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DownloadGroupMarkdownUploadBySecretAndFilename mocks base method.
func (m *MockGroupMarkdownUploadsServiceInterface) DownloadGroupMarkdownUploadBySecretAndFilename(gid any, secret, filename string, options ...gitlab.RequestOptionFunc) (io.Reader, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, secret, filename}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadGroupMarkdownUploadBySecretAndFilename", varargs...)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadGroupMarkdownUploadBySecretAndFilename indicates an expected call of DownloadGroupMarkdownUploadBySecretAndFilename.
func (mr *MockGroupMarkdownUploadsServiceInterfaceMockRecorder) DownloadGroupMarkdownUploadBySecretAndFilename(gid, secret, filename any, options ...any) *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadBySecretAndFilenameCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, secret, filename}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadGroupMarkdownUploadBySecretAndFilename", reflect.TypeOf((*MockGroupMarkdownUploadsServiceInterface)(nil).DownloadGroupMarkdownUploadBySecretAndFilename), varargs...)
	return &MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadBySecretAndFilenameCall{Call: call}
}

// MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadBySecretAndFilenameCall wrap *gomock.Call
type MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadBySecretAndFilenameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadBySecretAndFilenameCall) Return(arg0 io.Reader, arg1 *gitlab.Response, arg2 error) *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadBySecretAndFilenameCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadBySecretAndFilenameCall) Do(f func(any, string, string, ...gitlab.RequestOptionFunc) (io.Reader, *gitlab.Response, error)) *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadBySecretAndFilenameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadBySecretAndFilenameCall) DoAndReturn(f func(any, string, string, ...gitlab.RequestOptionFunc) (io.Reader, *gitlab.Response, error)) *MockGroupMarkdownUploadsServiceInterfaceDownloadGroupMarkdownUploadBySecretAndFilenameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListGroupMarkdownUploads mocks base method.
func (m *MockGroupMarkdownUploadsServiceInterface) ListGroupMarkdownUploads(gid any, opt *gitlab.ListMarkdownUploadsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.GroupMarkdownUpload, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupMarkdownUploads", varargs...)
	ret0, _ := ret[0].([]*gitlab.GroupMarkdownUpload)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGroupMarkdownUploads indicates an expected call of ListGroupMarkdownUploads.
func (mr *MockGroupMarkdownUploadsServiceInterfaceMockRecorder) ListGroupMarkdownUploads(gid, opt any, options ...any) *MockGroupMarkdownUploadsServiceInterfaceListGroupMarkdownUploadsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupMarkdownUploads", reflect.TypeOf((*MockGroupMarkdownUploadsServiceInterface)(nil).ListGroupMarkdownUploads), varargs...)
	return &MockGroupMarkdownUploadsServiceInterfaceListGroupMarkdownUploadsCall{Call: call}
}

// MockGroupMarkdownUploadsServiceInterfaceListGroupMarkdownUploadsCall wrap *gomock.Call
type MockGroupMarkdownUploadsServiceInterfaceListGroupMarkdownUploadsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGroupMarkdownUploadsServiceInterfaceListGroupMarkdownUploadsCall) Return(arg0 []*gitlab.GroupMarkdownUpload, arg1 *gitlab.Response, arg2 error) *MockGroupMarkdownUploadsServiceInterfaceListGroupMarkdownUploadsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGroupMarkdownUploadsServiceInterfaceListGroupMarkdownUploadsCall) Do(f func(any, *gitlab.ListMarkdownUploadsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.GroupMarkdownUpload, *gitlab.Response, error)) *MockGroupMarkdownUploadsServiceInterfaceListGroupMarkdownUploadsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGroupMarkdownUploadsServiceInterfaceListGroupMarkdownUploadsCall) DoAndReturn(f func(any, *gitlab.ListMarkdownUploadsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.GroupMarkdownUpload, *gitlab.Response, error)) *MockGroupMarkdownUploadsServiceInterfaceListGroupMarkdownUploadsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
