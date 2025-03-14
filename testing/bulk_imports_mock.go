// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/gitlab-org/api/client-go (interfaces: BulkImportsServiceInterface)
//
// Generated by this command:
//
//	mockgen -typed -destination=bulk_imports_mock.go -package=testing gitlab.com/gitlab-org/api/client-go BulkImportsServiceInterface
//

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockBulkImportsServiceInterface is a mock of BulkImportsServiceInterface interface.
type MockBulkImportsServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockBulkImportsServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockBulkImportsServiceInterfaceMockRecorder is the mock recorder for MockBulkImportsServiceInterface.
type MockBulkImportsServiceInterfaceMockRecorder struct {
	mock *MockBulkImportsServiceInterface
}

// NewMockBulkImportsServiceInterface creates a new mock instance.
func NewMockBulkImportsServiceInterface(ctrl *gomock.Controller) *MockBulkImportsServiceInterface {
	mock := &MockBulkImportsServiceInterface{ctrl: ctrl}
	mock.recorder = &MockBulkImportsServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBulkImportsServiceInterface) EXPECT() *MockBulkImportsServiceInterfaceMockRecorder {
	return m.recorder
}

// StartMigration mocks base method.
func (m *MockBulkImportsServiceInterface) StartMigration(startMigrationOptions *gitlab.BulkImportStartMigrationOptions, options ...gitlab.RequestOptionFunc) (*gitlab.BulkImportStartMigrationResponse, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{startMigrationOptions}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartMigration", varargs...)
	ret0, _ := ret[0].(*gitlab.BulkImportStartMigrationResponse)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StartMigration indicates an expected call of StartMigration.
func (mr *MockBulkImportsServiceInterfaceMockRecorder) StartMigration(startMigrationOptions any, options ...any) *MockBulkImportsServiceInterfaceStartMigrationCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{startMigrationOptions}, options...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMigration", reflect.TypeOf((*MockBulkImportsServiceInterface)(nil).StartMigration), varargs...)
	return &MockBulkImportsServiceInterfaceStartMigrationCall{Call: call}
}

// MockBulkImportsServiceInterfaceStartMigrationCall wrap *gomock.Call
type MockBulkImportsServiceInterfaceStartMigrationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBulkImportsServiceInterfaceStartMigrationCall) Return(arg0 *gitlab.BulkImportStartMigrationResponse, arg1 *gitlab.Response, arg2 error) *MockBulkImportsServiceInterfaceStartMigrationCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBulkImportsServiceInterfaceStartMigrationCall) Do(f func(*gitlab.BulkImportStartMigrationOptions, ...gitlab.RequestOptionFunc) (*gitlab.BulkImportStartMigrationResponse, *gitlab.Response, error)) *MockBulkImportsServiceInterfaceStartMigrationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBulkImportsServiceInterfaceStartMigrationCall) DoAndReturn(f func(*gitlab.BulkImportStartMigrationOptions, ...gitlab.RequestOptionFunc) (*gitlab.BulkImportStartMigrationResponse, *gitlab.Response, error)) *MockBulkImportsServiceInterfaceStartMigrationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
