// Code generated by MockGen. DO NOT EDIT.
// Source: organization_token.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/eranelbaz/go-tfe"
	gomock "github.com/golang/mock/gomock"
)

// MockOrganizationTokens is a mock of OrganizationTokens interface.
type MockOrganizationTokens struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationTokensMockRecorder
}

// MockOrganizationTokensMockRecorder is the mock recorder for MockOrganizationTokens.
type MockOrganizationTokensMockRecorder struct {
	mock *MockOrganizationTokens
}

// NewMockOrganizationTokens creates a new mock instance.
func NewMockOrganizationTokens(ctrl *gomock.Controller) *MockOrganizationTokens {
	mock := &MockOrganizationTokens{ctrl: ctrl}
	mock.recorder = &MockOrganizationTokensMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationTokens) EXPECT() *MockOrganizationTokensMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrganizationTokens) Create(ctx context.Context, organization string) (*tfe.OrganizationToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, organization)
	ret0, _ := ret[0].(*tfe.OrganizationToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrganizationTokensMockRecorder) Create(ctx, organization interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrganizationTokens)(nil).Create), ctx, organization)
}

// CreateWithOptions mocks base method.
func (m *MockOrganizationTokens) CreateWithOptions(ctx context.Context, organization string, options tfe.OrganizationTokenCreateOptions) (*tfe.OrganizationToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithOptions", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.OrganizationToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWithOptions indicates an expected call of CreateWithOptions.
func (mr *MockOrganizationTokensMockRecorder) CreateWithOptions(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithOptions", reflect.TypeOf((*MockOrganizationTokens)(nil).CreateWithOptions), ctx, organization, options)
}

// Delete mocks base method.
func (m *MockOrganizationTokens) Delete(ctx context.Context, organization string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, organization)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationTokensMockRecorder) Delete(ctx, organization interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationTokens)(nil).Delete), ctx, organization)
}

// Read mocks base method.
func (m *MockOrganizationTokens) Read(ctx context.Context, organization string) (*tfe.OrganizationToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, organization)
	ret0, _ := ret[0].(*tfe.OrganizationToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockOrganizationTokensMockRecorder) Read(ctx, organization interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockOrganizationTokens)(nil).Read), ctx, organization)
}
