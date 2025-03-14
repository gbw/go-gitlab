// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: TopicsServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=topics_mock.go -package=testing gitlab.com/gitlab-org/api/client-go TopicsServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockTopicsServiceInterface is a mock of TopicsServiceInterface interface.
type MockTopicsServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTopicsServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockTopicsServiceInterfaceMockRecorder is the mock recorder for MockTopicsServiceInterface.
type MockTopicsServiceInterfaceMockRecorder struct {
	mock *MockTopicsServiceInterface
}

// NewMockTopicsServiceInterface creates a new mock instance.
func NewMockTopicsServiceInterface(ctrl *gomock.Controller) *MockTopicsServiceInterface {
	mock := &MockTopicsServiceInterface{ctrl: ctrl}
	mock.recorder = &MockTopicsServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTopicsServiceInterface) EXPECT() *MockTopicsServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateTopic mocks base method.
func (m *MockTopicsServiceInterface) CreateTopic(opt *gitlab.CreateTopicOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Topic, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTopic", varargs...)
	ret0, _ := ret[0].(*gitlab.Topic)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateTopic indicates an expected call of CreateTopic.
func (mr *MockTopicsServiceInterfaceMockRecorder) CreateTopic(opt any, options ...any) *MockTopicsServiceInterfaceCreateTopicCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopic", reflect.TypeOf((*MockTopicsServiceInterface)(nil).CreateTopic), varargs...)
	return &MockTopicsServiceInterfaceCreateTopicCall{Call: call}
}

// MockTopicsServiceInterfaceCreateTopicCall wrap *gomock.Call
type MockTopicsServiceInterfaceCreateTopicCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTopicsServiceInterfaceCreateTopicCall) Return(arg0 *gitlab.Topic, arg1 *gitlab.Response, arg2 error) *MockTopicsServiceInterfaceCreateTopicCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTopicsServiceInterfaceCreateTopicCall) Do(f func(*gitlab.CreateTopicOptions, ...gitlab.RequestOptionFunc) (*gitlab.Topic, *gitlab.Response, error)) *MockTopicsServiceInterfaceCreateTopicCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTopicsServiceInterfaceCreateTopicCall) DoAndReturn(f func(*gitlab.CreateTopicOptions, ...gitlab.RequestOptionFunc) (*gitlab.Topic, *gitlab.Response, error)) *MockTopicsServiceInterfaceCreateTopicCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteTopic mocks base method.
func (m *MockTopicsServiceInterface) DeleteTopic(topic int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{topic}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTopic", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTopic indicates an expected call of DeleteTopic.
func (mr *MockTopicsServiceInterfaceMockRecorder) DeleteTopic(topic any, options ...any) *MockTopicsServiceInterfaceDeleteTopicCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{topic}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTopic", reflect.TypeOf((*MockTopicsServiceInterface)(nil).DeleteTopic), varargs...)
	return &MockTopicsServiceInterfaceDeleteTopicCall{Call: call}
}

// MockTopicsServiceInterfaceDeleteTopicCall wrap *gomock.Call
type MockTopicsServiceInterfaceDeleteTopicCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTopicsServiceInterfaceDeleteTopicCall) Return(arg0 *gitlab.Response, arg1 error) *MockTopicsServiceInterfaceDeleteTopicCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTopicsServiceInterfaceDeleteTopicCall) Do(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockTopicsServiceInterfaceDeleteTopicCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTopicsServiceInterfaceDeleteTopicCall) DoAndReturn(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockTopicsServiceInterfaceDeleteTopicCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetTopic mocks base method.
func (m *MockTopicsServiceInterface) GetTopic(topic int, options ...gitlab.RequestOptionFunc) (*gitlab.Topic, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{topic}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTopic", varargs...)
	ret0, _ := ret[0].(*gitlab.Topic)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTopic indicates an expected call of GetTopic.
func (mr *MockTopicsServiceInterfaceMockRecorder) GetTopic(topic any, options ...any) *MockTopicsServiceInterfaceGetTopicCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{topic}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopic", reflect.TypeOf((*MockTopicsServiceInterface)(nil).GetTopic), varargs...)
	return &MockTopicsServiceInterfaceGetTopicCall{Call: call}
}

// MockTopicsServiceInterfaceGetTopicCall wrap *gomock.Call
type MockTopicsServiceInterfaceGetTopicCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTopicsServiceInterfaceGetTopicCall) Return(arg0 *gitlab.Topic, arg1 *gitlab.Response, arg2 error) *MockTopicsServiceInterfaceGetTopicCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTopicsServiceInterfaceGetTopicCall) Do(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Topic, *gitlab.Response, error)) *MockTopicsServiceInterfaceGetTopicCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTopicsServiceInterfaceGetTopicCall) DoAndReturn(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Topic, *gitlab.Response, error)) *MockTopicsServiceInterfaceGetTopicCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListTopics mocks base method.
func (m *MockTopicsServiceInterface) ListTopics(opt *gitlab.ListTopicsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Topic, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTopics", varargs...)
	ret0, _ := ret[0].([]*gitlab.Topic)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTopics indicates an expected call of ListTopics.
func (mr *MockTopicsServiceInterfaceMockRecorder) ListTopics(opt any, options ...any) *MockTopicsServiceInterfaceListTopicsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTopics", reflect.TypeOf((*MockTopicsServiceInterface)(nil).ListTopics), varargs...)
	return &MockTopicsServiceInterfaceListTopicsCall{Call: call}
}

// MockTopicsServiceInterfaceListTopicsCall wrap *gomock.Call
type MockTopicsServiceInterfaceListTopicsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTopicsServiceInterfaceListTopicsCall) Return(arg0 []*gitlab.Topic, arg1 *gitlab.Response, arg2 error) *MockTopicsServiceInterfaceListTopicsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTopicsServiceInterfaceListTopicsCall) Do(f func(*gitlab.ListTopicsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Topic, *gitlab.Response, error)) *MockTopicsServiceInterfaceListTopicsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTopicsServiceInterfaceListTopicsCall) DoAndReturn(f func(*gitlab.ListTopicsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Topic, *gitlab.Response, error)) *MockTopicsServiceInterfaceListTopicsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateTopic mocks base method.
func (m *MockTopicsServiceInterface) UpdateTopic(topic int, opt *gitlab.UpdateTopicOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Topic, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{topic, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTopic", varargs...)
	ret0, _ := ret[0].(*gitlab.Topic)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateTopic indicates an expected call of UpdateTopic.
func (mr *MockTopicsServiceInterfaceMockRecorder) UpdateTopic(topic, opt any, options ...any) *MockTopicsServiceInterfaceUpdateTopicCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{topic, opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTopic", reflect.TypeOf((*MockTopicsServiceInterface)(nil).UpdateTopic), varargs...)
	return &MockTopicsServiceInterfaceUpdateTopicCall{Call: call}
}

// MockTopicsServiceInterfaceUpdateTopicCall wrap *gomock.Call
type MockTopicsServiceInterfaceUpdateTopicCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTopicsServiceInterfaceUpdateTopicCall) Return(arg0 *gitlab.Topic, arg1 *gitlab.Response, arg2 error) *MockTopicsServiceInterfaceUpdateTopicCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTopicsServiceInterfaceUpdateTopicCall) Do(f func(int, *gitlab.UpdateTopicOptions, ...gitlab.RequestOptionFunc) (*gitlab.Topic, *gitlab.Response, error)) *MockTopicsServiceInterfaceUpdateTopicCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTopicsServiceInterfaceUpdateTopicCall) DoAndReturn(f func(int, *gitlab.UpdateTopicOptions, ...gitlab.RequestOptionFunc) (*gitlab.Topic, *gitlab.Response, error)) *MockTopicsServiceInterfaceUpdateTopicCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
