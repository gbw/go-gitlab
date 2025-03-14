// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: LicenseServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=license_mock.go -package=testing gitlab.com/gitlab-org/api/client-go LicenseServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockLicenseServiceInterface is a mock of LicenseServiceInterface interface.
type MockLicenseServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockLicenseServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockLicenseServiceInterfaceMockRecorder is the mock recorder for MockLicenseServiceInterface.
type MockLicenseServiceInterfaceMockRecorder struct {
	mock *MockLicenseServiceInterface
}

// NewMockLicenseServiceInterface creates a new mock instance.
func NewMockLicenseServiceInterface(ctrl *gomock.Controller) *MockLicenseServiceInterface {
	mock := &MockLicenseServiceInterface{ctrl: ctrl}
	mock.recorder = &MockLicenseServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLicenseServiceInterface) EXPECT() *MockLicenseServiceInterfaceMockRecorder {
	return m.recorder
}

// AddLicense mocks base method.
func (m *MockLicenseServiceInterface) AddLicense(opt *gitlab.AddLicenseOptions, options ...gitlab.RequestOptionFunc) (*gitlab.License, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddLicense", varargs...)
	ret0, _ := ret[0].(*gitlab.License)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddLicense indicates an expected call of AddLicense.
func (mr *MockLicenseServiceInterfaceMockRecorder) AddLicense(opt any, options ...any) *MockLicenseServiceInterfaceAddLicenseCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLicense", reflect.TypeOf((*MockLicenseServiceInterface)(nil).AddLicense), varargs...)
	return &MockLicenseServiceInterfaceAddLicenseCall{Call: call}
}

// MockLicenseServiceInterfaceAddLicenseCall wrap *gomock.Call
type MockLicenseServiceInterfaceAddLicenseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLicenseServiceInterfaceAddLicenseCall) Return(arg0 *gitlab.License, arg1 *gitlab.Response, arg2 error) *MockLicenseServiceInterfaceAddLicenseCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLicenseServiceInterfaceAddLicenseCall) Do(f func(*gitlab.AddLicenseOptions, ...gitlab.RequestOptionFunc) (*gitlab.License, *gitlab.Response, error)) *MockLicenseServiceInterfaceAddLicenseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLicenseServiceInterfaceAddLicenseCall) DoAndReturn(f func(*gitlab.AddLicenseOptions, ...gitlab.RequestOptionFunc) (*gitlab.License, *gitlab.Response, error)) *MockLicenseServiceInterfaceAddLicenseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteLicense mocks base method.
func (m *MockLicenseServiceInterface) DeleteLicense(licenseID int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{licenseID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteLicense", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLicense indicates an expected call of DeleteLicense.
func (mr *MockLicenseServiceInterfaceMockRecorder) DeleteLicense(licenseID any, options ...any) *MockLicenseServiceInterfaceDeleteLicenseCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{licenseID}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLicense", reflect.TypeOf((*MockLicenseServiceInterface)(nil).DeleteLicense), varargs...)
	return &MockLicenseServiceInterfaceDeleteLicenseCall{Call: call}
}

// MockLicenseServiceInterfaceDeleteLicenseCall wrap *gomock.Call
type MockLicenseServiceInterfaceDeleteLicenseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLicenseServiceInterfaceDeleteLicenseCall) Return(arg0 *gitlab.Response, arg1 error) *MockLicenseServiceInterfaceDeleteLicenseCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLicenseServiceInterfaceDeleteLicenseCall) Do(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockLicenseServiceInterfaceDeleteLicenseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLicenseServiceInterfaceDeleteLicenseCall) DoAndReturn(f func(int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MockLicenseServiceInterfaceDeleteLicenseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetLicense mocks base method.
func (m *MockLicenseServiceInterface) GetLicense(options ...gitlab.RequestOptionFunc) (*gitlab.License, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLicense", varargs...)
	ret0, _ := ret[0].(*gitlab.License)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLicense indicates an expected call of GetLicense.
func (mr *MockLicenseServiceInterfaceMockRecorder) GetLicense(options ...any) *MockLicenseServiceInterfaceGetLicenseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLicense", reflect.TypeOf((*MockLicenseServiceInterface)(nil).GetLicense), options...)
	return &MockLicenseServiceInterfaceGetLicenseCall{Call: call}
}

// MockLicenseServiceInterfaceGetLicenseCall wrap *gomock.Call
type MockLicenseServiceInterfaceGetLicenseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLicenseServiceInterfaceGetLicenseCall) Return(arg0 *gitlab.License, arg1 *gitlab.Response, arg2 error) *MockLicenseServiceInterfaceGetLicenseCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLicenseServiceInterfaceGetLicenseCall) Do(f func(...gitlab.RequestOptionFunc) (*gitlab.License, *gitlab.Response, error)) *MockLicenseServiceInterfaceGetLicenseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLicenseServiceInterfaceGetLicenseCall) DoAndReturn(f func(...gitlab.RequestOptionFunc) (*gitlab.License, *gitlab.Response, error)) *MockLicenseServiceInterfaceGetLicenseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
