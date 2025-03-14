// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: ResourceStateEventsServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=resource_state_events_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ResourceStateEventsServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockResourceStateEventsServiceInterface is a mock of ResourceStateEventsServiceInterface interface.
type MockResourceStateEventsServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockResourceStateEventsServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockResourceStateEventsServiceInterfaceMockRecorder is the mock recorder for MockResourceStateEventsServiceInterface.
type MockResourceStateEventsServiceInterfaceMockRecorder struct {
	mock *MockResourceStateEventsServiceInterface
}

// NewMockResourceStateEventsServiceInterface creates a new mock instance.
func NewMockResourceStateEventsServiceInterface(ctrl *gomock.Controller) *MockResourceStateEventsServiceInterface {
	mock := &MockResourceStateEventsServiceInterface{ctrl: ctrl}
	mock.recorder = &MockResourceStateEventsServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceStateEventsServiceInterface) EXPECT() *MockResourceStateEventsServiceInterfaceMockRecorder {
	return m.recorder
}

// GetIssueStateEvent mocks base method.
func (m *MockResourceStateEventsServiceInterface) GetIssueStateEvent(pid any, issue, event int, options ...gitlab.RequestOptionFunc) (*gitlab.StateEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, event}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIssueStateEvent", varargs...)
	ret0, _ := ret[0].(*gitlab.StateEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssueStateEvent indicates an expected call of GetIssueStateEvent.
func (mr *MockResourceStateEventsServiceInterfaceMockRecorder) GetIssueStateEvent(pid, issue, event any, options ...any) *MockResourceStateEventsServiceInterfaceGetIssueStateEventCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, event}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueStateEvent", reflect.TypeOf((*MockResourceStateEventsServiceInterface)(nil).GetIssueStateEvent), varargs...)
	return &MockResourceStateEventsServiceInterfaceGetIssueStateEventCall{Call: call}
}

// MockResourceStateEventsServiceInterfaceGetIssueStateEventCall wrap *gomock.Call
type MockResourceStateEventsServiceInterfaceGetIssueStateEventCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceStateEventsServiceInterfaceGetIssueStateEventCall) Return(arg0 *gitlab.StateEvent, arg1 *gitlab.Response, arg2 error) *MockResourceStateEventsServiceInterfaceGetIssueStateEventCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceStateEventsServiceInterfaceGetIssueStateEventCall) Do(f func(any, int, int, ...gitlab.RequestOptionFunc) (*gitlab.StateEvent, *gitlab.Response, error)) *MockResourceStateEventsServiceInterfaceGetIssueStateEventCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceStateEventsServiceInterfaceGetIssueStateEventCall) DoAndReturn(f func(any, int, int, ...gitlab.RequestOptionFunc) (*gitlab.StateEvent, *gitlab.Response, error)) *MockResourceStateEventsServiceInterfaceGetIssueStateEventCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetMergeRequestStateEvent mocks base method.
func (m *MockResourceStateEventsServiceInterface) GetMergeRequestStateEvent(pid any, request, event int, options ...gitlab.RequestOptionFunc) (*gitlab.StateEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, request, event}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeRequestStateEvent", varargs...)
	ret0, _ := ret[0].(*gitlab.StateEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMergeRequestStateEvent indicates an expected call of GetMergeRequestStateEvent.
func (mr *MockResourceStateEventsServiceInterfaceMockRecorder) GetMergeRequestStateEvent(pid, request, event any, options ...any) *MockResourceStateEventsServiceInterfaceGetMergeRequestStateEventCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, request, event}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestStateEvent", reflect.TypeOf((*MockResourceStateEventsServiceInterface)(nil).GetMergeRequestStateEvent), varargs...)
	return &MockResourceStateEventsServiceInterfaceGetMergeRequestStateEventCall{Call: call}
}

// MockResourceStateEventsServiceInterfaceGetMergeRequestStateEventCall wrap *gomock.Call
type MockResourceStateEventsServiceInterfaceGetMergeRequestStateEventCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceStateEventsServiceInterfaceGetMergeRequestStateEventCall) Return(arg0 *gitlab.StateEvent, arg1 *gitlab.Response, arg2 error) *MockResourceStateEventsServiceInterfaceGetMergeRequestStateEventCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceStateEventsServiceInterfaceGetMergeRequestStateEventCall) Do(f func(any, int, int, ...gitlab.RequestOptionFunc) (*gitlab.StateEvent, *gitlab.Response, error)) *MockResourceStateEventsServiceInterfaceGetMergeRequestStateEventCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceStateEventsServiceInterfaceGetMergeRequestStateEventCall) DoAndReturn(f func(any, int, int, ...gitlab.RequestOptionFunc) (*gitlab.StateEvent, *gitlab.Response, error)) *MockResourceStateEventsServiceInterfaceGetMergeRequestStateEventCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListIssueStateEvents mocks base method.
func (m *MockResourceStateEventsServiceInterface) ListIssueStateEvents(pid any, issue int, opt *gitlab.ListStateEventsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.StateEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIssueStateEvents", varargs...)
	ret0, _ := ret[0].([]*gitlab.StateEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListIssueStateEvents indicates an expected call of ListIssueStateEvents.
func (mr *MockResourceStateEventsServiceInterfaceMockRecorder) ListIssueStateEvents(pid, issue, opt any, options ...any) *MockResourceStateEventsServiceInterfaceListIssueStateEventsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIssueStateEvents", reflect.TypeOf((*MockResourceStateEventsServiceInterface)(nil).ListIssueStateEvents), varargs...)
	return &MockResourceStateEventsServiceInterfaceListIssueStateEventsCall{Call: call}
}

// MockResourceStateEventsServiceInterfaceListIssueStateEventsCall wrap *gomock.Call
type MockResourceStateEventsServiceInterfaceListIssueStateEventsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceStateEventsServiceInterfaceListIssueStateEventsCall) Return(arg0 []*gitlab.StateEvent, arg1 *gitlab.Response, arg2 error) *MockResourceStateEventsServiceInterfaceListIssueStateEventsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceStateEventsServiceInterfaceListIssueStateEventsCall) Do(f func(any, int, *gitlab.ListStateEventsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.StateEvent, *gitlab.Response, error)) *MockResourceStateEventsServiceInterfaceListIssueStateEventsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceStateEventsServiceInterfaceListIssueStateEventsCall) DoAndReturn(f func(any, int, *gitlab.ListStateEventsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.StateEvent, *gitlab.Response, error)) *MockResourceStateEventsServiceInterfaceListIssueStateEventsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListMergeStateEvents mocks base method.
func (m *MockResourceStateEventsServiceInterface) ListMergeStateEvents(pid any, request int, opt *gitlab.ListStateEventsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.StateEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, request, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMergeStateEvents", varargs...)
	ret0, _ := ret[0].([]*gitlab.StateEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMergeStateEvents indicates an expected call of ListMergeStateEvents.
func (mr *MockResourceStateEventsServiceInterfaceMockRecorder) ListMergeStateEvents(pid, request, opt any, options ...any) *MockResourceStateEventsServiceInterfaceListMergeStateEventsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, request, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMergeStateEvents", reflect.TypeOf((*MockResourceStateEventsServiceInterface)(nil).ListMergeStateEvents), varargs...)
	return &MockResourceStateEventsServiceInterfaceListMergeStateEventsCall{Call: call}
}

// MockResourceStateEventsServiceInterfaceListMergeStateEventsCall wrap *gomock.Call
type MockResourceStateEventsServiceInterfaceListMergeStateEventsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceStateEventsServiceInterfaceListMergeStateEventsCall) Return(arg0 []*gitlab.StateEvent, arg1 *gitlab.Response, arg2 error) *MockResourceStateEventsServiceInterfaceListMergeStateEventsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceStateEventsServiceInterfaceListMergeStateEventsCall) Do(f func(any, int, *gitlab.ListStateEventsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.StateEvent, *gitlab.Response, error)) *MockResourceStateEventsServiceInterfaceListMergeStateEventsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceStateEventsServiceInterfaceListMergeStateEventsCall) DoAndReturn(f func(any, int, *gitlab.ListStateEventsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.StateEvent, *gitlab.Response, error)) *MockResourceStateEventsServiceInterfaceListMergeStateEventsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
