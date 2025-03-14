// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: PipelineTriggersServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=pipeline_triggers_mock.go -package=testing gitlab.com/gitlab-org/api/client-go PipelineTriggersServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockPipelineTriggersServiceInterface is a mock of PipelineTriggersServiceInterface interface.
type MockPipelineTriggersServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPipelineTriggersServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockPipelineTriggersServiceInterfaceMockRecorder is the mock recorder for MockPipelineTriggersServiceInterface.
type MockPipelineTriggersServiceInterfaceMockRecorder struct {
	mock *MockPipelineTriggersServiceInterface
}

// NewMockPipelineTriggersServiceInterface creates a new mock instance.
func NewMockPipelineTriggersServiceInterface(ctrl *gomock.Controller) *MockPipelineTriggersServiceInterface {
	mock := &MockPipelineTriggersServiceInterface{ctrl: ctrl}
	mock.recorder = &MockPipelineTriggersServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelineTriggersServiceInterface) EXPECT() *MockPipelineTriggersServiceInterfaceMockRecorder {
	return m.recorder
}

// AddPipelineTrigger mocks base method.
func (m *MockPipelineTriggersServiceInterface) AddPipelineTrigger(pid any, opt *gitlab.AddPipelineTriggerOptions, options ...gitlab.RequestOptionFunc) (*gitlab.PipelineTrigger, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddPipelineTrigger", varargs...)
	ret0, _ := ret[0].(*gitlab.PipelineTrigger)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddPipelineTrigger indicates an expected call of AddPipelineTrigger.
func (mr *MockPipelineTriggersServiceInterfaceMockRecorder) AddPipelineTrigger(pid, opt any, options ...any) *MockPipelineTriggersServiceInterfaceAddPipelineTriggerCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPipelineTrigger", reflect.TypeOf((*MockPipelineTriggersServiceInterface)(nil).AddPipelineTrigger), varargs...)
	return &MockPipelineTriggersServiceInterfaceAddPipelineTriggerCall{Call: call}
}

// MockPipelineTriggersServiceInterfaceAddPipelineTriggerCall wrap *gomock.Call
type MockPipelineTriggersServiceInterfaceAddPipelineTriggerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelineTriggersServiceInterfaceAddPipelineTriggerCall) Return(arg0 *gitlab.PipelineTrigger, arg1 *gitlab.Response, arg2 error) *MockPipelineTriggersServiceInterfaceAddPipelineTriggerCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelineTriggersServiceInterfaceAddPipelineTriggerCall) Do(f func(any, *gitlab.AddPipelineTriggerOptions, ...gitlab.RequestOptionFunc) (*gitlab.PipelineTrigger, *gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceAddPipelineTriggerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelineTriggersServiceInterfaceAddPipelineTriggerCall) DoAndReturn(f func(any, *gitlab.AddPipelineTriggerOptions, ...gitlab.RequestOptionFunc) (*gitlab.PipelineTrigger, *gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceAddPipelineTriggerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeletePipelineTrigger mocks base method.
func (m *MockPipelineTriggersServiceInterface) DeletePipelineTrigger(pid any, trigger int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, trigger}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePipelineTrigger", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePipelineTrigger indicates an expected call of DeletePipelineTrigger.
func (mr *MockPipelineTriggersServiceInterfaceMockRecorder) DeletePipelineTrigger(pid, trigger any, options ...any) *MockPipelineTriggersServiceInterfaceDeletePipelineTriggerCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, trigger}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipelineTrigger", reflect.TypeOf((*MockPipelineTriggersServiceInterface)(nil).DeletePipelineTrigger), varargs...)
	return &MockPipelineTriggersServiceInterfaceDeletePipelineTriggerCall{Call: call}
}

// MockPipelineTriggersServiceInterfaceDeletePipelineTriggerCall wrap *gomock.Call
type MockPipelineTriggersServiceInterfaceDeletePipelineTriggerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelineTriggersServiceInterfaceDeletePipelineTriggerCall) Return(arg0 *gitlab.Response, arg1 error) *MockPipelineTriggersServiceInterfaceDeletePipelineTriggerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelineTriggersServiceInterfaceDeletePipelineTriggerCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceDeletePipelineTriggerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelineTriggersServiceInterfaceDeletePipelineTriggerCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceDeletePipelineTriggerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// EditPipelineTrigger mocks base method.
func (m *MockPipelineTriggersServiceInterface) EditPipelineTrigger(pid any, trigger int, opt *gitlab.EditPipelineTriggerOptions, options ...gitlab.RequestOptionFunc) (*gitlab.PipelineTrigger, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, trigger, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditPipelineTrigger", varargs...)
	ret0, _ := ret[0].(*gitlab.PipelineTrigger)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditPipelineTrigger indicates an expected call of EditPipelineTrigger.
func (mr *MockPipelineTriggersServiceInterfaceMockRecorder) EditPipelineTrigger(pid, trigger, opt any, options ...any) *MockPipelineTriggersServiceInterfaceEditPipelineTriggerCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, trigger, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditPipelineTrigger", reflect.TypeOf((*MockPipelineTriggersServiceInterface)(nil).EditPipelineTrigger), varargs...)
	return &MockPipelineTriggersServiceInterfaceEditPipelineTriggerCall{Call: call}
}

// MockPipelineTriggersServiceInterfaceEditPipelineTriggerCall wrap *gomock.Call
type MockPipelineTriggersServiceInterfaceEditPipelineTriggerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelineTriggersServiceInterfaceEditPipelineTriggerCall) Return(arg0 *gitlab.PipelineTrigger, arg1 *gitlab.Response, arg2 error) *MockPipelineTriggersServiceInterfaceEditPipelineTriggerCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelineTriggersServiceInterfaceEditPipelineTriggerCall) Do(f func(any, int, *gitlab.EditPipelineTriggerOptions, ...gitlab.RequestOptionFunc) (*gitlab.PipelineTrigger, *gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceEditPipelineTriggerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelineTriggersServiceInterfaceEditPipelineTriggerCall) DoAndReturn(f func(any, int, *gitlab.EditPipelineTriggerOptions, ...gitlab.RequestOptionFunc) (*gitlab.PipelineTrigger, *gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceEditPipelineTriggerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetPipelineTrigger mocks base method.
func (m *MockPipelineTriggersServiceInterface) GetPipelineTrigger(pid any, trigger int, options ...gitlab.RequestOptionFunc) (*gitlab.PipelineTrigger, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, trigger}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPipelineTrigger", varargs...)
	ret0, _ := ret[0].(*gitlab.PipelineTrigger)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPipelineTrigger indicates an expected call of GetPipelineTrigger.
func (mr *MockPipelineTriggersServiceInterfaceMockRecorder) GetPipelineTrigger(pid, trigger any, options ...any) *MockPipelineTriggersServiceInterfaceGetPipelineTriggerCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, trigger}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipelineTrigger", reflect.TypeOf((*MockPipelineTriggersServiceInterface)(nil).GetPipelineTrigger), varargs...)
	return &MockPipelineTriggersServiceInterfaceGetPipelineTriggerCall{Call: call}
}

// MockPipelineTriggersServiceInterfaceGetPipelineTriggerCall wrap *gomock.Call
type MockPipelineTriggersServiceInterfaceGetPipelineTriggerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelineTriggersServiceInterfaceGetPipelineTriggerCall) Return(arg0 *gitlab.PipelineTrigger, arg1 *gitlab.Response, arg2 error) *MockPipelineTriggersServiceInterfaceGetPipelineTriggerCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelineTriggersServiceInterfaceGetPipelineTriggerCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.PipelineTrigger, *gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceGetPipelineTriggerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelineTriggersServiceInterfaceGetPipelineTriggerCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.PipelineTrigger, *gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceGetPipelineTriggerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListPipelineTriggers mocks base method.
func (m *MockPipelineTriggersServiceInterface) ListPipelineTriggers(pid any, opt *gitlab.ListPipelineTriggersOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.PipelineTrigger, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPipelineTriggers", varargs...)
	ret0, _ := ret[0].([]*gitlab.PipelineTrigger)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListPipelineTriggers indicates an expected call of ListPipelineTriggers.
func (mr *MockPipelineTriggersServiceInterfaceMockRecorder) ListPipelineTriggers(pid, opt any, options ...any) *MockPipelineTriggersServiceInterfaceListPipelineTriggersCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipelineTriggers", reflect.TypeOf((*MockPipelineTriggersServiceInterface)(nil).ListPipelineTriggers), varargs...)
	return &MockPipelineTriggersServiceInterfaceListPipelineTriggersCall{Call: call}
}

// MockPipelineTriggersServiceInterfaceListPipelineTriggersCall wrap *gomock.Call
type MockPipelineTriggersServiceInterfaceListPipelineTriggersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelineTriggersServiceInterfaceListPipelineTriggersCall) Return(arg0 []*gitlab.PipelineTrigger, arg1 *gitlab.Response, arg2 error) *MockPipelineTriggersServiceInterfaceListPipelineTriggersCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelineTriggersServiceInterfaceListPipelineTriggersCall) Do(f func(any, *gitlab.ListPipelineTriggersOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.PipelineTrigger, *gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceListPipelineTriggersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelineTriggersServiceInterfaceListPipelineTriggersCall) DoAndReturn(f func(any, *gitlab.ListPipelineTriggersOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.PipelineTrigger, *gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceListPipelineTriggersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RunPipelineTrigger mocks base method.
func (m *MockPipelineTriggersServiceInterface) RunPipelineTrigger(pid any, opt *gitlab.RunPipelineTriggerOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunPipelineTrigger", varargs...)
	ret0, _ := ret[0].(*gitlab.Pipeline)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RunPipelineTrigger indicates an expected call of RunPipelineTrigger.
func (mr *MockPipelineTriggersServiceInterfaceMockRecorder) RunPipelineTrigger(pid, opt any, options ...any) *MockPipelineTriggersServiceInterfaceRunPipelineTriggerCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunPipelineTrigger", reflect.TypeOf((*MockPipelineTriggersServiceInterface)(nil).RunPipelineTrigger), varargs...)
	return &MockPipelineTriggersServiceInterfaceRunPipelineTriggerCall{Call: call}
}

// MockPipelineTriggersServiceInterfaceRunPipelineTriggerCall wrap *gomock.Call
type MockPipelineTriggersServiceInterfaceRunPipelineTriggerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelineTriggersServiceInterfaceRunPipelineTriggerCall) Return(arg0 *gitlab.Pipeline, arg1 *gitlab.Response, arg2 error) *MockPipelineTriggersServiceInterfaceRunPipelineTriggerCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelineTriggersServiceInterfaceRunPipelineTriggerCall) Do(f func(any, *gitlab.RunPipelineTriggerOptions, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceRunPipelineTriggerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelineTriggersServiceInterfaceRunPipelineTriggerCall) DoAndReturn(f func(any, *gitlab.RunPipelineTriggerOptions, ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceRunPipelineTriggerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// TakeOwnershipOfPipelineTrigger mocks base method.
func (m *MockPipelineTriggersServiceInterface) TakeOwnershipOfPipelineTrigger(pid any, trigger int, options ...gitlab.RequestOptionFunc) (*gitlab.PipelineTrigger, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, trigger}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TakeOwnershipOfPipelineTrigger", varargs...)
	ret0, _ := ret[0].(*gitlab.PipelineTrigger)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TakeOwnershipOfPipelineTrigger indicates an expected call of TakeOwnershipOfPipelineTrigger.
func (mr *MockPipelineTriggersServiceInterfaceMockRecorder) TakeOwnershipOfPipelineTrigger(pid, trigger any, options ...any) *MockPipelineTriggersServiceInterfaceTakeOwnershipOfPipelineTriggerCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, trigger}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeOwnershipOfPipelineTrigger", reflect.TypeOf((*MockPipelineTriggersServiceInterface)(nil).TakeOwnershipOfPipelineTrigger), varargs...)
	return &MockPipelineTriggersServiceInterfaceTakeOwnershipOfPipelineTriggerCall{Call: call}
}

// MockPipelineTriggersServiceInterfaceTakeOwnershipOfPipelineTriggerCall wrap *gomock.Call
type MockPipelineTriggersServiceInterfaceTakeOwnershipOfPipelineTriggerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPipelineTriggersServiceInterfaceTakeOwnershipOfPipelineTriggerCall) Return(arg0 *gitlab.PipelineTrigger, arg1 *gitlab.Response, arg2 error) *MockPipelineTriggersServiceInterfaceTakeOwnershipOfPipelineTriggerCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPipelineTriggersServiceInterfaceTakeOwnershipOfPipelineTriggerCall) Do(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.PipelineTrigger, *gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceTakeOwnershipOfPipelineTriggerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPipelineTriggersServiceInterfaceTakeOwnershipOfPipelineTriggerCall) DoAndReturn(f func(any, int, ...gitlab.RequestOptionFunc) (*gitlab.PipelineTrigger, *gitlab.Response, error)) *MockPipelineTriggersServiceInterfaceTakeOwnershipOfPipelineTriggerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
