// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: GraphQLInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=graphql_mock.go -package=testing gitlab.com/gitlab-org/api/client-go GraphQLInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockGraphQLInterface is a mock of GraphQLInterface interface.
type MockGraphQLInterface struct {
	ctrl     *gomock.Controller
	recorder *MockGraphQLInterfaceMockRecorder
	isgomock struct{}
}

// MockGraphQLInterfaceMockRecorder is the mock recorder for MockGraphQLInterface.
type MockGraphQLInterfaceMockRecorder struct {
	mock *MockGraphQLInterface
}

// NewMockGraphQLInterface creates a new mock instance.
func NewMockGraphQLInterface(ctrl *gomock.Controller) *MockGraphQLInterface {
	mock := &MockGraphQLInterface{ctrl: ctrl}
	mock.recorder = &MockGraphQLInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGraphQLInterface) EXPECT() *MockGraphQLInterfaceMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockGraphQLInterface) Do(query gitlab.GraphQLQuery, response any, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{query, response}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Do", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockGraphQLInterfaceMockRecorder) Do(query, response any, options ...any) *MockGraphQLInterfaceDoCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{query, response}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockGraphQLInterface)(nil).Do), varargs...)
	return &MockGraphQLInterfaceDoCall{Call: call}
}

// MockGraphQLInterfaceDoCall wrap *gomock.Call
type MockGraphQLInterfaceDoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockGraphQLInterfaceDoCall) Return(arg0 *gitlab.Response, arg1 error) *MockGraphQLInterfaceDoCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockGraphQLInterfaceDoCall) Do(f func(gitlab.GraphQLQuery, any, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockGraphQLInterfaceDoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockGraphQLInterfaceDoCall) DoAndReturn(f func(gitlab.GraphQLQuery, any, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockGraphQLInterfaceDoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
