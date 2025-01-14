// Code generated by MockGen. DO NOT EDIT.
// Source: project.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/eranelbaz/go-tfe"
	gomock "github.com/golang/mock/gomock"
)

// MockProjects is a mock of Projects interface.
type MockProjects struct {
	ctrl     *gomock.Controller
	recorder *MockProjectsMockRecorder
}

// MockProjectsMockRecorder is the mock recorder for MockProjects.
type MockProjectsMockRecorder struct {
	mock *MockProjects
}

// NewMockProjects creates a new mock instance.
func NewMockProjects(ctrl *gomock.Controller) *MockProjects {
	mock := &MockProjects{ctrl: ctrl}
	mock.recorder = &MockProjectsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjects) EXPECT() *MockProjectsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProjects) Create(ctx context.Context, organization string, options tfe.ProjectCreateOptions) (*tfe.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProjectsMockRecorder) Create(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProjects)(nil).Create), ctx, organization, options)
}

// Delete mocks base method.
func (m *MockProjects) Delete(ctx context.Context, projectID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProjectsMockRecorder) Delete(ctx, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProjects)(nil).Delete), ctx, projectID)
}

// List mocks base method.
func (m *MockProjects) List(ctx context.Context, organization string, options *tfe.ProjectListOptions) (*tfe.ProjectList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.ProjectList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProjectsMockRecorder) List(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProjects)(nil).List), ctx, organization, options)
}

// Read mocks base method.
func (m *MockProjects) Read(ctx context.Context, projectID string) (*tfe.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, projectID)
	ret0, _ := ret[0].(*tfe.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockProjectsMockRecorder) Read(ctx, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockProjects)(nil).Read), ctx, projectID)
}

// Update mocks base method.
func (m *MockProjects) Update(ctx context.Context, projectID string, options tfe.ProjectUpdateOptions) (*tfe.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, projectID, options)
	ret0, _ := ret[0].(*tfe.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProjectsMockRecorder) Update(ctx, projectID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProjects)(nil).Update), ctx, projectID, options)
}
