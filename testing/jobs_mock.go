// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: JobsServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=jobs_mock.go -package=testing gitlab.com/gitlab-org/api/client-go JobsServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	bytes "bytes"
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockJobsServiceInterface is a mock of JobsServiceInterface interface.
type MockJobsServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockJobsServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockJobsServiceInterfaceMockRecorder is the mock recorder for MockJobsServiceInterface.
type MockJobsServiceInterfaceMockRecorder struct {
	mock *MockJobsServiceInterface
}

// NewMockJobsServiceInterface creates a new mock instance.
func NewMockJobsServiceInterface(ctrl *gomock.Controller) *MockJobsServiceInterface {
	mock := &MockJobsServiceInterface{ctrl: ctrl}
	mock.recorder = &MockJobsServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobsServiceInterface) EXPECT() *MockJobsServiceInterfaceMockRecorder {
	return m.recorder
}

// CancelJob mocks base method.
func (m *MockJobsServiceInterface) CancelJob(pid any, jobID int, options ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, jobID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelJob", varargs...)
	ret0, _ := ret[0].(*gitlab.Job)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CancelJob indicates an expected call of CancelJob.
func (mr *MockJobsServiceInterfaceMockRecorder) CancelJob(pid, jobID any, options ...any) *MockJobsServiceInterfaceCancelJobCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, jobID}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelJob", reflect.TypeOf((*MockJobsServiceInterface)(nil).CancelJob), varargs...)
	return &MockJobsServiceInterfaceCancelJobCall{Call: call}
}

// MockJobsServiceInterfaceCancelJobCall wrap *gomock.Call
type MockJobsServiceInterfaceCancelJobCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceCancelJobCall) Return(arg0 *gitlab.Job, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceCancelJobCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceCancelJobCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceCancelJobCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceCancelJobCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceCancelJobCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteArtifacts mocks base method.
func (m *MockJobsServiceInterface) DeleteArtifacts(pid any, jobID int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, jobID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteArtifacts", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteArtifacts indicates an expected call of DeleteArtifacts.
func (mr *MockJobsServiceInterfaceMockRecorder) DeleteArtifacts(pid, jobID any, options ...any) *MockJobsServiceInterfaceDeleteArtifactsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, jobID}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArtifacts", reflect.TypeOf((*MockJobsServiceInterface)(nil).DeleteArtifacts), varargs...)
	return &MockJobsServiceInterfaceDeleteArtifactsCall{Call: call}
}

// MockJobsServiceInterfaceDeleteArtifactsCall wrap *gomock.Call
type MockJobsServiceInterfaceDeleteArtifactsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceDeleteArtifactsCall) Return(arg0 *gitlab.Response, arg1 error) *MockJobsServiceInterfaceDeleteArtifactsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceDeleteArtifactsCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockJobsServiceInterfaceDeleteArtifactsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceDeleteArtifactsCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockJobsServiceInterfaceDeleteArtifactsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteProjectArtifacts mocks base method.
func (m *MockJobsServiceInterface) DeleteProjectArtifacts(pid any, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteProjectArtifacts", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProjectArtifacts indicates an expected call of DeleteProjectArtifacts.
func (mr *MockJobsServiceInterfaceMockRecorder) DeleteProjectArtifacts(pid any, options ...any) *MockJobsServiceInterfaceDeleteProjectArtifactsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectArtifacts", reflect.TypeOf((*MockJobsServiceInterface)(nil).DeleteProjectArtifacts), varargs...)
	return &MockJobsServiceInterfaceDeleteProjectArtifactsCall{Call: call}
}

// MockJobsServiceInterfaceDeleteProjectArtifactsCall wrap *gomock.Call
type MockJobsServiceInterfaceDeleteProjectArtifactsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceDeleteProjectArtifactsCall) Return(arg0 *gitlab.Response, arg1 error) *MockJobsServiceInterfaceDeleteProjectArtifactsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceDeleteProjectArtifactsCall) Do(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockJobsServiceInterfaceDeleteProjectArtifactsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceDeleteProjectArtifactsCall) DoAndReturn(f func(any, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockJobsServiceInterfaceDeleteProjectArtifactsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DownloadArtifactsFile mocks base method.
func (m *MockJobsServiceInterface) DownloadArtifactsFile(pid any, refName string, opt *gitlab.DownloadArtifactsFileOptions, options ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, refName, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadArtifactsFile", varargs...)
	ret0, _ := ret[0].(*bytes.Reader)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadArtifactsFile indicates an expected call of DownloadArtifactsFile.
func (mr *MockJobsServiceInterfaceMockRecorder) DownloadArtifactsFile(pid, refName, opt any, options ...any) *MockJobsServiceInterfaceDownloadArtifactsFileCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, refName, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadArtifactsFile", reflect.TypeOf((*MockJobsServiceInterface)(nil).DownloadArtifactsFile), varargs...)
	return &MockJobsServiceInterfaceDownloadArtifactsFileCall{Call: call}
}

// MockJobsServiceInterfaceDownloadArtifactsFileCall wrap *gomock.Call
type MockJobsServiceInterfaceDownloadArtifactsFileCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceDownloadArtifactsFileCall) Return(arg0 *bytes.Reader, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceDownloadArtifactsFileCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceDownloadArtifactsFileCall) Do(f func(any, string, *gitlab.DownloadArtifactsFileOptions, ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error)) *MockJobsServiceInterfaceDownloadArtifactsFileCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceDownloadArtifactsFileCall) DoAndReturn(f func(any, string, *gitlab.DownloadArtifactsFileOptions, ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error)) *MockJobsServiceInterfaceDownloadArtifactsFileCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DownloadSingleArtifactsFile mocks base method.
func (m *MockJobsServiceInterface) DownloadSingleArtifactsFile(pid any, jobID int, artifactPath string, options ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, jobID, artifactPath}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadSingleArtifactsFile", varargs...)
	ret0, _ := ret[0].(*bytes.Reader)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadSingleArtifactsFile indicates an expected call of DownloadSingleArtifactsFile.
func (mr *MockJobsServiceInterfaceMockRecorder) DownloadSingleArtifactsFile(pid, jobID, artifactPath any, options ...any) *MockJobsServiceInterfaceDownloadSingleArtifactsFileCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, jobID, artifactPath}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadSingleArtifactsFile", reflect.TypeOf((*MockJobsServiceInterface)(nil).DownloadSingleArtifactsFile), varargs...)
	return &MockJobsServiceInterfaceDownloadSingleArtifactsFileCall{Call: call}
}

// MockJobsServiceInterfaceDownloadSingleArtifactsFileCall wrap *gomock.Call
type MockJobsServiceInterfaceDownloadSingleArtifactsFileCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceDownloadSingleArtifactsFileCall) Return(arg0 *bytes.Reader, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceDownloadSingleArtifactsFileCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceDownloadSingleArtifactsFileCall) Do(f func(any, int, string, ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error)) *MockJobsServiceInterfaceDownloadSingleArtifactsFileCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceDownloadSingleArtifactsFileCall) DoAndReturn(f func(any, int, string, ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error)) *MockJobsServiceInterfaceDownloadSingleArtifactsFileCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DownloadSingleArtifactsFileByTagOrBranch mocks base method.
func (m *MockJobsServiceInterface) DownloadSingleArtifactsFileByTagOrBranch(pid any, refName, artifactPath string, opt *gitlab.DownloadArtifactsFileOptions, options ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, refName, artifactPath, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadSingleArtifactsFileByTagOrBranch", varargs...)
	ret0, _ := ret[0].(*bytes.Reader)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadSingleArtifactsFileByTagOrBranch indicates an expected call of DownloadSingleArtifactsFileByTagOrBranch.
func (mr *MockJobsServiceInterfaceMockRecorder) DownloadSingleArtifactsFileByTagOrBranch(pid, refName, artifactPath, opt any, options ...any) *MockJobsServiceInterfaceDownloadSingleArtifactsFileByTagOrBranchCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, refName, artifactPath, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadSingleArtifactsFileByTagOrBranch", reflect.TypeOf((*MockJobsServiceInterface)(nil).DownloadSingleArtifactsFileByTagOrBranch), varargs...)
	return &MockJobsServiceInterfaceDownloadSingleArtifactsFileByTagOrBranchCall{Call: call}
}

// MockJobsServiceInterfaceDownloadSingleArtifactsFileByTagOrBranchCall wrap *gomock.Call
type MockJobsServiceInterfaceDownloadSingleArtifactsFileByTagOrBranchCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceDownloadSingleArtifactsFileByTagOrBranchCall) Return(arg0 *bytes.Reader, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceDownloadSingleArtifactsFileByTagOrBranchCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceDownloadSingleArtifactsFileByTagOrBranchCall) Do(f func(any, string, string, *gitlab.DownloadArtifactsFileOptions, ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error)) *MockJobsServiceInterfaceDownloadSingleArtifactsFileByTagOrBranchCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceDownloadSingleArtifactsFileByTagOrBranchCall) DoAndReturn(f func(any, string, string, *gitlab.DownloadArtifactsFileOptions, ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error)) *MockJobsServiceInterfaceDownloadSingleArtifactsFileByTagOrBranchCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EraseJob mocks base method.
func (m *MockJobsServiceInterface) EraseJob(pid any, jobID int, options ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, jobID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EraseJob", varargs...)
	ret0, _ := ret[0].(*gitlab.Job)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EraseJob indicates an expected call of EraseJob.
func (mr *MockJobsServiceInterfaceMockRecorder) EraseJob(pid, jobID any, options ...any) *MockJobsServiceInterfaceEraseJobCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, jobID}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EraseJob", reflect.TypeOf((*MockJobsServiceInterface)(nil).EraseJob), varargs...)
	return &MockJobsServiceInterfaceEraseJobCall{Call: call}
}

// MockJobsServiceInterfaceEraseJobCall wrap *gomock.Call
type MockJobsServiceInterfaceEraseJobCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceEraseJobCall) Return(arg0 *gitlab.Job, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceEraseJobCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceEraseJobCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceEraseJobCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceEraseJobCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceEraseJobCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetJob mocks base method.
func (m *MockJobsServiceInterface) GetJob(pid any, jobID int, options ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, jobID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetJob", varargs...)
	ret0, _ := ret[0].(*gitlab.Job)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetJob indicates an expected call of GetJob.
func (mr *MockJobsServiceInterfaceMockRecorder) GetJob(pid, jobID any, options ...any) *MockJobsServiceInterfaceGetJobCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, jobID}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJob", reflect.TypeOf((*MockJobsServiceInterface)(nil).GetJob), varargs...)
	return &MockJobsServiceInterfaceGetJobCall{Call: call}
}

// MockJobsServiceInterfaceGetJobCall wrap *gomock.Call
type MockJobsServiceInterfaceGetJobCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceGetJobCall) Return(arg0 *gitlab.Job, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceGetJobCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceGetJobCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceGetJobCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceGetJobCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceGetJobCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetJobArtifacts mocks base method.
func (m *MockJobsServiceInterface) GetJobArtifacts(pid any, jobID int, options ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, jobID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetJobArtifacts", varargs...)
	ret0, _ := ret[0].(*bytes.Reader)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetJobArtifacts indicates an expected call of GetJobArtifacts.
func (mr *MockJobsServiceInterfaceMockRecorder) GetJobArtifacts(pid, jobID any, options ...any) *MockJobsServiceInterfaceGetJobArtifactsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, jobID}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobArtifacts", reflect.TypeOf((*MockJobsServiceInterface)(nil).GetJobArtifacts), varargs...)
	return &MockJobsServiceInterfaceGetJobArtifactsCall{Call: call}
}

// MockJobsServiceInterfaceGetJobArtifactsCall wrap *gomock.Call
type MockJobsServiceInterfaceGetJobArtifactsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceGetJobArtifactsCall) Return(arg0 *bytes.Reader, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceGetJobArtifactsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceGetJobArtifactsCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error)) *MockJobsServiceInterfaceGetJobArtifactsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceGetJobArtifactsCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error)) *MockJobsServiceInterfaceGetJobArtifactsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetJobTokensJob mocks base method.
func (m *MockJobsServiceInterface) GetJobTokensJob(opts *gitlab.GetJobTokensJobOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetJobTokensJob", varargs...)
	ret0, _ := ret[0].(*gitlab.Job)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetJobTokensJob indicates an expected call of GetJobTokensJob.
func (mr *MockJobsServiceInterfaceMockRecorder) GetJobTokensJob(opts any, options ...any) *MockJobsServiceInterfaceGetJobTokensJobCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opts}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobTokensJob", reflect.TypeOf((*MockJobsServiceInterface)(nil).GetJobTokensJob), varargs...)
	return &MockJobsServiceInterfaceGetJobTokensJobCall{Call: call}
}

// MockJobsServiceInterfaceGetJobTokensJobCall wrap *gomock.Call
type MockJobsServiceInterfaceGetJobTokensJobCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceGetJobTokensJobCall) Return(arg0 *gitlab.Job, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceGetJobTokensJobCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceGetJobTokensJobCall) Do(f func(*gitlab.GetJobTokensJobOptions, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceGetJobTokensJobCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceGetJobTokensJobCall) DoAndReturn(f func(*gitlab.GetJobTokensJobOptions, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceGetJobTokensJobCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetTraceFile mocks base method.
func (m *MockJobsServiceInterface) GetTraceFile(pid any, jobID int, options ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, jobID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTraceFile", varargs...)
	ret0, _ := ret[0].(*bytes.Reader)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTraceFile indicates an expected call of GetTraceFile.
func (mr *MockJobsServiceInterfaceMockRecorder) GetTraceFile(pid, jobID any, options ...any) *MockJobsServiceInterfaceGetTraceFileCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, jobID}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTraceFile", reflect.TypeOf((*MockJobsServiceInterface)(nil).GetTraceFile), varargs...)
	return &MockJobsServiceInterfaceGetTraceFileCall{Call: call}
}

// MockJobsServiceInterfaceGetTraceFileCall wrap *gomock.Call
type MockJobsServiceInterfaceGetTraceFileCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceGetTraceFileCall) Return(arg0 *bytes.Reader, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceGetTraceFileCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceGetTraceFileCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error)) *MockJobsServiceInterfaceGetTraceFileCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceGetTraceFileCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*bytes.Reader, *gitlab.Response, error)) *MockJobsServiceInterfaceGetTraceFileCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// KeepArtifacts mocks base method.
func (m *MockJobsServiceInterface) KeepArtifacts(pid any, jobID int, options ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, jobID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "KeepArtifacts", varargs...)
	ret0, _ := ret[0].(*gitlab.Job)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// KeepArtifacts indicates an expected call of KeepArtifacts.
func (mr *MockJobsServiceInterfaceMockRecorder) KeepArtifacts(pid, jobID any, options ...any) *MockJobsServiceInterfaceKeepArtifactsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, jobID}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeepArtifacts", reflect.TypeOf((*MockJobsServiceInterface)(nil).KeepArtifacts), varargs...)
	return &MockJobsServiceInterfaceKeepArtifactsCall{Call: call}
}

// MockJobsServiceInterfaceKeepArtifactsCall wrap *gomock.Call
type MockJobsServiceInterfaceKeepArtifactsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceKeepArtifactsCall) Return(arg0 *gitlab.Job, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceKeepArtifactsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceKeepArtifactsCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceKeepArtifactsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceKeepArtifactsCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceKeepArtifactsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListPipelineBridges mocks base method.
func (m *MockJobsServiceInterface) ListPipelineBridges(pid any, pipelineID int, opts *gitlab.ListJobsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Bridge, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, pipelineID, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPipelineBridges", varargs...)
	ret0, _ := ret[0].([]*gitlab.Bridge)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListPipelineBridges indicates an expected call of ListPipelineBridges.
func (mr *MockJobsServiceInterfaceMockRecorder) ListPipelineBridges(pid, pipelineID, opts any, options ...any) *MockJobsServiceInterfaceListPipelineBridgesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, pipelineID, opts}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipelineBridges", reflect.TypeOf((*MockJobsServiceInterface)(nil).ListPipelineBridges), varargs...)
	return &MockJobsServiceInterfaceListPipelineBridgesCall{Call: call}
}

// MockJobsServiceInterfaceListPipelineBridgesCall wrap *gomock.Call
type MockJobsServiceInterfaceListPipelineBridgesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceListPipelineBridgesCall) Return(arg0 []*gitlab.Bridge, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceListPipelineBridgesCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceListPipelineBridgesCall) Do(f func(any, int, *gitlab.ListJobsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Bridge, *gitlab.Response, error)) *MockJobsServiceInterfaceListPipelineBridgesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceListPipelineBridgesCall) DoAndReturn(f func(any, int, *gitlab.ListJobsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Bridge, *gitlab.Response, error)) *MockJobsServiceInterfaceListPipelineBridgesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListPipelineJobs mocks base method.
func (m *MockJobsServiceInterface) ListPipelineJobs(pid any, pipelineID int, opts *gitlab.ListJobsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, pipelineID, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPipelineJobs", varargs...)
	ret0, _ := ret[0].([]*gitlab.Job)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListPipelineJobs indicates an expected call of ListPipelineJobs.
func (mr *MockJobsServiceInterfaceMockRecorder) ListPipelineJobs(pid, pipelineID, opts any, options ...any) *MockJobsServiceInterfaceListPipelineJobsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, pipelineID, opts}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipelineJobs", reflect.TypeOf((*MockJobsServiceInterface)(nil).ListPipelineJobs), varargs...)
	return &MockJobsServiceInterfaceListPipelineJobsCall{Call: call}
}

// MockJobsServiceInterfaceListPipelineJobsCall wrap *gomock.Call
type MockJobsServiceInterfaceListPipelineJobsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceListPipelineJobsCall) Return(arg0 []*gitlab.Job, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceListPipelineJobsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceListPipelineJobsCall) Do(f func(any, int, *gitlab.ListJobsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceListPipelineJobsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceListPipelineJobsCall) DoAndReturn(f func(any, int, *gitlab.ListJobsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceListPipelineJobsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListProjectJobs mocks base method.
func (m *MockJobsServiceInterface) ListProjectJobs(pid any, opts *gitlab.ListJobsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectJobs", varargs...)
	ret0, _ := ret[0].([]*gitlab.Job)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectJobs indicates an expected call of ListProjectJobs.
func (mr *MockJobsServiceInterfaceMockRecorder) ListProjectJobs(pid, opts any, options ...any) *MockJobsServiceInterfaceListProjectJobsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opts}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectJobs", reflect.TypeOf((*MockJobsServiceInterface)(nil).ListProjectJobs), varargs...)
	return &MockJobsServiceInterfaceListProjectJobsCall{Call: call}
}

// MockJobsServiceInterfaceListProjectJobsCall wrap *gomock.Call
type MockJobsServiceInterfaceListProjectJobsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceListProjectJobsCall) Return(arg0 []*gitlab.Job, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceListProjectJobsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceListProjectJobsCall) Do(f func(any, *gitlab.ListJobsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceListProjectJobsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceListProjectJobsCall) DoAndReturn(f func(any, *gitlab.ListJobsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceListProjectJobsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PlayJob mocks base method.
func (m *MockJobsServiceInterface) PlayJob(pid any, jobID int, opt *gitlab.PlayJobOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, jobID, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PlayJob", varargs...)
	ret0, _ := ret[0].(*gitlab.Job)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PlayJob indicates an expected call of PlayJob.
func (mr *MockJobsServiceInterfaceMockRecorder) PlayJob(pid, jobID, opt any, options ...any) *MockJobsServiceInterfacePlayJobCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, jobID, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlayJob", reflect.TypeOf((*MockJobsServiceInterface)(nil).PlayJob), varargs...)
	return &MockJobsServiceInterfacePlayJobCall{Call: call}
}

// MockJobsServiceInterfacePlayJobCall wrap *gomock.Call
type MockJobsServiceInterfacePlayJobCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfacePlayJobCall) Return(arg0 *gitlab.Job, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfacePlayJobCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfacePlayJobCall) Do(f func(any, int, *gitlab.PlayJobOptions, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfacePlayJobCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfacePlayJobCall) DoAndReturn(f func(any, int, *gitlab.PlayJobOptions, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfacePlayJobCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RetryJob mocks base method.
func (m *MockJobsServiceInterface) RetryJob(pid any, jobID int, options ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, jobID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetryJob", varargs...)
	ret0, _ := ret[0].(*gitlab.Job)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RetryJob indicates an expected call of RetryJob.
func (mr *MockJobsServiceInterfaceMockRecorder) RetryJob(pid, jobID any, options ...any) *MockJobsServiceInterfaceRetryJobCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, jobID}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetryJob", reflect.TypeOf((*MockJobsServiceInterface)(nil).RetryJob), varargs...)
	return &MockJobsServiceInterfaceRetryJobCall{Call: call}
}

// MockJobsServiceInterfaceRetryJobCall wrap *gomock.Call
type MockJobsServiceInterfaceRetryJobCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockJobsServiceInterfaceRetryJobCall) Return(arg0 *gitlab.Job, arg1 *gitlab.Response, arg2 error) *MockJobsServiceInterfaceRetryJobCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockJobsServiceInterfaceRetryJobCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceRetryJobCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockJobsServiceInterfaceRetryJobCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Job, *gitlab.Response, error)) *MockJobsServiceInterfaceRetryJobCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
