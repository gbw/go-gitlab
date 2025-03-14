// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: PipelinesServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=pipelines_mock.go -package=testing gitlab.com/gitlab-org/api/client-go PipelinesServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockPipelinesServiceInterface is a mock of PipelinesServiceInterface interface.
type MockPipelinesServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockPipelinesServiceInterfaceMockRecorder is the mock recorder for MockPipelinesServiceInterface.
type MockPipelinesServiceInterfaceMockRecorder struct {
	mock *MockPipelinesServiceInterface
}

// NewMockPipelinesServiceInterface creates a new mock instance.
func NewMockPipelinesServiceInterface(ctrl *gomock.Controller) *MockPipelinesServiceInterface {
	mock := &MockPipelinesServiceInterface{ctrl: ctrl}
	mock.recorder = &MockPipelinesServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesServiceInterface) EXPECT() *MockPipelinesServiceInterfaceMockRecorder {
	return m.recorder
}

// CancelPipelineBuild mocks base method.
func (m *MockPipelinesServiceInterface) CancelPipelineBuild(pid any, pipeline int, options ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, pipeline}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelPipelineBuild", varargs...)
	ret0, _ := ret[0].(*gitlab.Pipeline)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CancelPipelineBuild indicates an expected call of CancelPipelineBuild.
func (mr *MockPipelinesServiceInterfaceMockRecorder) CancelPipelineBuild(pid, pipeline any, options ...any) *MockPipelinesServiceInterfaceCancelPipelineBuildCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, pipeline}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelPipelineBuild", reflect.TypeOf((*MockPipelinesServiceInterface)(nil).CancelPipelineBuild), varargs...)
	return &MockPipelinesServiceInterfaceCancelPipelineBuildCall{Call: call}
}

// MockPipelinesServiceInterfaceCancelPipelineBuildCall wrap *gomock.Call
type MockPipelinesServiceInterfaceCancelPipelineBuildCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelinesServiceInterfaceCancelPipelineBuildCall) Return(arg0 *gitlab.Pipeline, arg1 *gitlab.Response, arg2 error) *MockPipelinesServiceInterfaceCancelPipelineBuildCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelinesServiceInterfaceCancelPipelineBuildCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelinesServiceInterfaceCancelPipelineBuildCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelinesServiceInterfaceCancelPipelineBuildCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelinesServiceInterfaceCancelPipelineBuildCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreatePipeline mocks base method.
func (m *MockPipelinesServiceInterface) CreatePipeline(pid any, opt *gitlab.CreatePipelineOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreatePipeline", varargs...)
	ret0, _ := ret[0].(*gitlab.Pipeline)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreatePipeline indicates an expected call of CreatePipeline.
func (mr *MockPipelinesServiceInterfaceMockRecorder) CreatePipeline(pid, opt any, options ...any) *MockPipelinesServiceInterfaceCreatePipelineCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipeline", reflect.TypeOf((*MockPipelinesServiceInterface)(nil).CreatePipeline), varargs...)
	return &MockPipelinesServiceInterfaceCreatePipelineCall{Call: call}
}

// MockPipelinesServiceInterfaceCreatePipelineCall wrap *gomock.Call
type MockPipelinesServiceInterfaceCreatePipelineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelinesServiceInterfaceCreatePipelineCall) Return(arg0 *gitlab.Pipeline, arg1 *gitlab.Response, arg2 error) *MockPipelinesServiceInterfaceCreatePipelineCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelinesServiceInterfaceCreatePipelineCall) Do(f func(any, *gitlab.CreatePipelineOptions, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelinesServiceInterfaceCreatePipelineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelinesServiceInterfaceCreatePipelineCall) DoAndReturn(f func(any, *gitlab.CreatePipelineOptions, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelinesServiceInterfaceCreatePipelineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeletePipeline mocks base method.
func (m *MockPipelinesServiceInterface) DeletePipeline(pid any, pipeline int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, pipeline}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePipeline", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePipeline indicates an expected call of DeletePipeline.
func (mr *MockPipelinesServiceInterfaceMockRecorder) DeletePipeline(pid, pipeline any, options ...any) *MockPipelinesServiceInterfaceDeletePipelineCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, pipeline}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipeline", reflect.TypeOf((*MockPipelinesServiceInterface)(nil).DeletePipeline), varargs...)
	return &MockPipelinesServiceInterfaceDeletePipelineCall{Call: call}
}

// MockPipelinesServiceInterfaceDeletePipelineCall wrap *gomock.Call
type MockPipelinesServiceInterfaceDeletePipelineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelinesServiceInterfaceDeletePipelineCall) Return(arg0 *gitlab.Response, arg1 error) *MockPipelinesServiceInterfaceDeletePipelineCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelinesServiceInterfaceDeletePipelineCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockPipelinesServiceInterfaceDeletePipelineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelinesServiceInterfaceDeletePipelineCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockPipelinesServiceInterfaceDeletePipelineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetLatestPipeline mocks base method.
func (m *MockPipelinesServiceInterface) GetLatestPipeline(pid any, opt *gitlab.GetLatestPipelineOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLatestPipeline", varargs...)
	ret0, _ := ret[0].(*gitlab.Pipeline)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLatestPipeline indicates an expected call of GetLatestPipeline.
func (mr *MockPipelinesServiceInterfaceMockRecorder) GetLatestPipeline(pid, opt any, options ...any) *MockPipelinesServiceInterfaceGetLatestPipelineCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestPipeline", reflect.TypeOf((*MockPipelinesServiceInterface)(nil).GetLatestPipeline), varargs...)
	return &MockPipelinesServiceInterfaceGetLatestPipelineCall{Call: call}
}

// MockPipelinesServiceInterfaceGetLatestPipelineCall wrap *gomock.Call
type MockPipelinesServiceInterfaceGetLatestPipelineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelinesServiceInterfaceGetLatestPipelineCall) Return(arg0 *gitlab.Pipeline, arg1 *gitlab.Response, arg2 error) *MockPipelinesServiceInterfaceGetLatestPipelineCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelinesServiceInterfaceGetLatestPipelineCall) Do(f func(any, *gitlab.GetLatestPipelineOptions, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelinesServiceInterfaceGetLatestPipelineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelinesServiceInterfaceGetLatestPipelineCall) DoAndReturn(f func(any, *gitlab.GetLatestPipelineOptions, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelinesServiceInterfaceGetLatestPipelineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetPipeline mocks base method.
func (m *MockPipelinesServiceInterface) GetPipeline(pid any, pipeline int, options ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, pipeline}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPipeline", varargs...)
	ret0, _ := ret[0].(*gitlab.Pipeline)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPipeline indicates an expected call of GetPipeline.
func (mr *MockPipelinesServiceInterfaceMockRecorder) GetPipeline(pid, pipeline any, options ...any) *MockPipelinesServiceInterfaceGetPipelineCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, pipeline}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipeline", reflect.TypeOf((*MockPipelinesServiceInterface)(nil).GetPipeline), varargs...)
	return &MockPipelinesServiceInterfaceGetPipelineCall{Call: call}
}

// MockPipelinesServiceInterfaceGetPipelineCall wrap *gomock.Call
type MockPipelinesServiceInterfaceGetPipelineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelinesServiceInterfaceGetPipelineCall) Return(arg0 *gitlab.Pipeline, arg1 *gitlab.Response, arg2 error) *MockPipelinesServiceInterfaceGetPipelineCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelinesServiceInterfaceGetPipelineCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelinesServiceInterfaceGetPipelineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelinesServiceInterfaceGetPipelineCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelinesServiceInterfaceGetPipelineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetPipelineTestReport mocks base method.
func (m *MockPipelinesServiceInterface) GetPipelineTestReport(pid any, pipeline int, options ...gitlab.RequestOptionFunc) (*gitlab.PipelineTestReport, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, pipeline}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPipelineTestReport", varargs...)
	ret0, _ := ret[0].(*gitlab.PipelineTestReport)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPipelineTestReport indicates an expected call of GetPipelineTestReport.
func (mr *MockPipelinesServiceInterfaceMockRecorder) GetPipelineTestReport(pid, pipeline any, options ...any) *MockPipelinesServiceInterfaceGetPipelineTestReportCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, pipeline}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipelineTestReport", reflect.TypeOf((*MockPipelinesServiceInterface)(nil).GetPipelineTestReport), varargs...)
	return &MockPipelinesServiceInterfaceGetPipelineTestReportCall{Call: call}
}

// MockPipelinesServiceInterfaceGetPipelineTestReportCall wrap *gomock.Call
type MockPipelinesServiceInterfaceGetPipelineTestReportCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelinesServiceInterfaceGetPipelineTestReportCall) Return(arg0 *gitlab.PipelineTestReport, arg1 *gitlab.Response, arg2 error) *MockPipelinesServiceInterfaceGetPipelineTestReportCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelinesServiceInterfaceGetPipelineTestReportCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.PipelineTestReport, *gitlab.Response, error)) *MockPipelinesServiceInterfaceGetPipelineTestReportCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelinesServiceInterfaceGetPipelineTestReportCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.PipelineTestReport, *gitlab.Response, error)) *MockPipelinesServiceInterfaceGetPipelineTestReportCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetPipelineVariables mocks base method.
func (m *MockPipelinesServiceInterface) GetPipelineVariables(pid any, pipeline int, options ...gitlab.RequestOptionFunc) ([]*gitlab.PipelineVariable, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, pipeline}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPipelineVariables", varargs...)
	ret0, _ := ret[0].([]*gitlab.PipelineVariable)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPipelineVariables indicates an expected call of GetPipelineVariables.
func (mr *MockPipelinesServiceInterfaceMockRecorder) GetPipelineVariables(pid, pipeline any, options ...any) *MockPipelinesServiceInterfaceGetPipelineVariablesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, pipeline}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipelineVariables", reflect.TypeOf((*MockPipelinesServiceInterface)(nil).GetPipelineVariables), varargs...)
	return &MockPipelinesServiceInterfaceGetPipelineVariablesCall{Call: call}
}

// MockPipelinesServiceInterfaceGetPipelineVariablesCall wrap *gomock.Call
type MockPipelinesServiceInterfaceGetPipelineVariablesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelinesServiceInterfaceGetPipelineVariablesCall) Return(arg0 []*gitlab.PipelineVariable, arg1 *gitlab.Response, arg2 error) *MockPipelinesServiceInterfaceGetPipelineVariablesCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelinesServiceInterfaceGetPipelineVariablesCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) ([]*gitlab.PipelineVariable, *gitlab.Response, error)) *MockPipelinesServiceInterfaceGetPipelineVariablesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelinesServiceInterfaceGetPipelineVariablesCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) ([]*gitlab.PipelineVariable, *gitlab.Response, error)) *MockPipelinesServiceInterfaceGetPipelineVariablesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListProjectPipelines mocks base method.
func (m *MockPipelinesServiceInterface) ListProjectPipelines(pid any, opt *gitlab.ListProjectPipelinesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.PipelineInfo, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectPipelines", varargs...)
	ret0, _ := ret[0].([]*gitlab.PipelineInfo)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectPipelines indicates an expected call of ListProjectPipelines.
func (mr *MockPipelinesServiceInterfaceMockRecorder) ListProjectPipelines(pid, opt any, options ...any) *MockPipelinesServiceInterfaceListProjectPipelinesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectPipelines", reflect.TypeOf((*MockPipelinesServiceInterface)(nil).ListProjectPipelines), varargs...)
	return &MockPipelinesServiceInterfaceListProjectPipelinesCall{Call: call}
}

// MockPipelinesServiceInterfaceListProjectPipelinesCall wrap *gomock.Call
type MockPipelinesServiceInterfaceListProjectPipelinesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelinesServiceInterfaceListProjectPipelinesCall) Return(arg0 []*gitlab.PipelineInfo, arg1 *gitlab.Response, arg2 error) *MockPipelinesServiceInterfaceListProjectPipelinesCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelinesServiceInterfaceListProjectPipelinesCall) Do(f func(any, *gitlab.ListProjectPipelinesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.PipelineInfo, *gitlab.Response, error)) *MockPipelinesServiceInterfaceListProjectPipelinesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelinesServiceInterfaceListProjectPipelinesCall) DoAndReturn(f func(any, *gitlab.ListProjectPipelinesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.PipelineInfo, *gitlab.Response, error)) *MockPipelinesServiceInterfaceListProjectPipelinesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RetryPipelineBuild mocks base method.
func (m *MockPipelinesServiceInterface) RetryPipelineBuild(pid any, pipeline int, options ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, pipeline}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetryPipelineBuild", varargs...)
	ret0, _ := ret[0].(*gitlab.Pipeline)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RetryPipelineBuild indicates an expected call of RetryPipelineBuild.
func (mr *MockPipelinesServiceInterfaceMockRecorder) RetryPipelineBuild(pid, pipeline any, options ...any) *MockPipelinesServiceInterfaceRetryPipelineBuildCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, pipeline}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetryPipelineBuild", reflect.TypeOf((*MockPipelinesServiceInterface)(nil).RetryPipelineBuild), varargs...)
	return &MockPipelinesServiceInterfaceRetryPipelineBuildCall{Call: call}
}

// MockPipelinesServiceInterfaceRetryPipelineBuildCall wrap *gomock.Call
type MockPipelinesServiceInterfaceRetryPipelineBuildCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelinesServiceInterfaceRetryPipelineBuildCall) Return(arg0 *gitlab.Pipeline, arg1 *gitlab.Response, arg2 error) *MockPipelinesServiceInterfaceRetryPipelineBuildCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelinesServiceInterfaceRetryPipelineBuildCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelinesServiceInterfaceRetryPipelineBuildCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelinesServiceInterfaceRetryPipelineBuildCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelinesServiceInterfaceRetryPipelineBuildCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdatePipelineMetadata mocks base method.
func (m *MockPipelinesServiceInterface) UpdatePipelineMetadata(pid any, pipeline int, opt *gitlab.UpdatePipelineMetadataOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, pipeline, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdatePipelineMetadata", varargs...)
	ret0, _ := ret[0].(*gitlab.Pipeline)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdatePipelineMetadata indicates an expected call of UpdatePipelineMetadata.
func (mr *MockPipelinesServiceInterfaceMockRecorder) UpdatePipelineMetadata(pid, pipeline, opt any, options ...any) *MockPipelinesServiceInterfaceUpdatePipelineMetadataCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, pipeline, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipelineMetadata", reflect.TypeOf((*MockPipelinesServiceInterface)(nil).UpdatePipelineMetadata), varargs...)
	return &MockPipelinesServiceInterfaceUpdatePipelineMetadataCall{Call: call}
}

// MockPipelinesServiceInterfaceUpdatePipelineMetadataCall wrap *gomock.Call
type MockPipelinesServiceInterfaceUpdatePipelineMetadataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelinesServiceInterfaceUpdatePipelineMetadataCall) Return(arg0 *gitlab.Pipeline, arg1 *gitlab.Response, arg2 error) *MockPipelinesServiceInterfaceUpdatePipelineMetadataCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelinesServiceInterfaceUpdatePipelineMetadataCall) Do(f func(any, int, *gitlab.UpdatePipelineMetadataOptions, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelinesServiceInterfaceUpdatePipelineMetadataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelinesServiceInterfaceUpdatePipelineMetadataCall) DoAndReturn(f func(any, int, *gitlab.UpdatePipelineMetadataOptions, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelinesServiceInterfaceUpdatePipelineMetadataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
