// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: DockerfileTemplatesServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=dockerfile_templates_mock.go -package=testing gitlab.com/gitlab-org/api/client-go DockerfileTemplatesServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockDockerfileTemplatesServiceInterface is a mock of DockerfileTemplatesServiceInterface interface.
type MockDockerfileTemplatesServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDockerfileTemplatesServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockDockerfileTemplatesServiceInterfaceMockRecorder is the mock recorder for MockDockerfileTemplatesServiceInterface.
type MockDockerfileTemplatesServiceInterfaceMockRecorder struct {
	mock *MockDockerfileTemplatesServiceInterface
}

// NewMockDockerfileTemplatesServiceInterface creates a new mock instance.
func NewMockDockerfileTemplatesServiceInterface(ctrl *gomock.Controller) *MockDockerfileTemplatesServiceInterface {
	mock := &MockDockerfileTemplatesServiceInterface{ctrl: ctrl}
	mock.recorder = &MockDockerfileTemplatesServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDockerfileTemplatesServiceInterface) EXPECT() *MockDockerfileTemplatesServiceInterfaceMockRecorder {
	return m.recorder
}

// GetTemplate mocks base method.
func (m *MockDockerfileTemplatesServiceInterface) GetTemplate(key string, options ...gitlab.RequestOptionFunc) (*gitlab.DockerfileTemplate, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{key}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTemplate", varargs...)
	ret0, _ := ret[0].(*gitlab.DockerfileTemplate)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTemplate indicates an expected call of GetTemplate.
func (mr *MockDockerfileTemplatesServiceInterfaceMockRecorder) GetTemplate(key any, options ...any) *MockDockerfileTemplatesServiceInterfaceGetTemplateCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{key}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplate", reflect.TypeOf((*MockDockerfileTemplatesServiceInterface)(nil).GetTemplate), varargs...)
	return &MockDockerfileTemplatesServiceInterfaceGetTemplateCall{Call: call}
}

// MockDockerfileTemplatesServiceInterfaceGetTemplateCall wrap *gomock.Call
type MockDockerfileTemplatesServiceInterfaceGetTemplateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDockerfileTemplatesServiceInterfaceGetTemplateCall) Return(arg0 *gitlab.DockerfileTemplate, arg1 *gitlab.Response, arg2 error) *MockDockerfileTemplatesServiceInterfaceGetTemplateCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDockerfileTemplatesServiceInterfaceGetTemplateCall) Do(f func(string, ...gitlab.RequestOptionFunc) (*gitlab.DockerfileTemplate, *gitlab.Response, error)) *MockDockerfileTemplatesServiceInterfaceGetTemplateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDockerfileTemplatesServiceInterfaceGetTemplateCall) DoAndReturn(f func(string, ...gitlab.RequestOptionFunc) (*gitlab.DockerfileTemplate, *gitlab.Response, error)) *MockDockerfileTemplatesServiceInterfaceGetTemplateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListTemplates mocks base method.
func (m *MockDockerfileTemplatesServiceInterface) ListTemplates(opt *gitlab.ListDockerfileTemplatesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.DockerfileTemplateListItem, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTemplates", varargs...)
	ret0, _ := ret[0].([]*gitlab.DockerfileTemplateListItem)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTemplates indicates an expected call of ListTemplates.
func (mr *MockDockerfileTemplatesServiceInterfaceMockRecorder) ListTemplates(opt any, options ...any) *MockDockerfileTemplatesServiceInterfaceListTemplatesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTemplates", reflect.TypeOf((*MockDockerfileTemplatesServiceInterface)(nil).ListTemplates), varargs...)
	return &MockDockerfileTemplatesServiceInterfaceListTemplatesCall{Call: call}
}

// MockDockerfileTemplatesServiceInterfaceListTemplatesCall wrap *gomock.Call
type MockDockerfileTemplatesServiceInterfaceListTemplatesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDockerfileTemplatesServiceInterfaceListTemplatesCall) Return(arg0 []*gitlab.DockerfileTemplateListItem, arg1 *gitlab.Response, arg2 error) *MockDockerfileTemplatesServiceInterfaceListTemplatesCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDockerfileTemplatesServiceInterfaceListTemplatesCall) Do(f func(*gitlab.ListDockerfileTemplatesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.DockerfileTemplateListItem, *gitlab.Response, error)) *MockDockerfileTemplatesServiceInterfaceListTemplatesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDockerfileTemplatesServiceInterfaceListTemplatesCall) DoAndReturn(f func(*gitlab.ListDockerfileTemplatesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.DockerfileTemplateListItem, *gitlab.Response, error)) *MockDockerfileTemplatesServiceInterfaceListTemplatesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
