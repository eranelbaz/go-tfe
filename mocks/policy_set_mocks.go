// Code generated by MockGen. DO NOT EDIT.
// Source: policy_set.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
)

// MockPolicySets is a mock of PolicySets interface.
type MockPolicySets struct {
	ctrl     *gomock.Controller
	recorder *MockPolicySetsMockRecorder
}

// MockPolicySetsMockRecorder is the mock recorder for MockPolicySets.
type MockPolicySetsMockRecorder struct {
	mock *MockPolicySets
}

// NewMockPolicySets creates a new mock instance.
func NewMockPolicySets(ctrl *gomock.Controller) *MockPolicySets {
	mock := &MockPolicySets{ctrl: ctrl}
	mock.recorder = &MockPolicySetsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPolicySets) EXPECT() *MockPolicySetsMockRecorder {
	return m.recorder
}

// AddPolicies mocks base method.
func (m *MockPolicySets) AddPolicies(ctx context.Context, policySetID string, options tfe.PolicySetAddPoliciesOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPolicies", ctx, policySetID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPolicies indicates an expected call of AddPolicies.
func (mr *MockPolicySetsMockRecorder) AddPolicies(ctx, policySetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPolicies", reflect.TypeOf((*MockPolicySets)(nil).AddPolicies), ctx, policySetID, options)
}

// AddProjects mocks base method.
func (m *MockPolicySets) AddProjects(ctx context.Context, policySetID string, options tfe.PolicySetAddProjectsOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProjects", ctx, policySetID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProjects indicates an expected call of AddProjects.
func (mr *MockPolicySetsMockRecorder) AddProjects(ctx, policySetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProjects", reflect.TypeOf((*MockPolicySets)(nil).AddProjects), ctx, policySetID, options)
}

// AddWorkspaceExclusions mocks base method.
func (m *MockPolicySets) AddWorkspaceExclusions(ctx context.Context, policySetID string, options tfe.PolicySetAddWorkspaceExclusionsOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkspaceExclusions", ctx, policySetID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWorkspaceExclusions indicates an expected call of AddWorkspaceExclusions.
func (mr *MockPolicySetsMockRecorder) AddWorkspaceExclusions(ctx, policySetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkspaceExclusions", reflect.TypeOf((*MockPolicySets)(nil).AddWorkspaceExclusions), ctx, policySetID, options)
}

// AddWorkspaces mocks base method.
func (m *MockPolicySets) AddWorkspaces(ctx context.Context, policySetID string, options tfe.PolicySetAddWorkspacesOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkspaces", ctx, policySetID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWorkspaces indicates an expected call of AddWorkspaces.
func (mr *MockPolicySetsMockRecorder) AddWorkspaces(ctx, policySetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkspaces", reflect.TypeOf((*MockPolicySets)(nil).AddWorkspaces), ctx, policySetID, options)
}

// Create mocks base method.
func (m *MockPolicySets) Create(ctx context.Context, organization string, options tfe.PolicySetCreateOptions) (*tfe.PolicySet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.PolicySet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPolicySetsMockRecorder) Create(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPolicySets)(nil).Create), ctx, organization, options)
}

// Delete mocks base method.
func (m *MockPolicySets) Delete(ctx context.Context, policyID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, policyID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPolicySetsMockRecorder) Delete(ctx, policyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPolicySets)(nil).Delete), ctx, policyID)
}

// List mocks base method.
func (m *MockPolicySets) List(ctx context.Context, organization string, options *tfe.PolicySetListOptions) (*tfe.PolicySetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.PolicySetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPolicySetsMockRecorder) List(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPolicySets)(nil).List), ctx, organization, options)
}

// Read mocks base method.
func (m *MockPolicySets) Read(ctx context.Context, policySetID string) (*tfe.PolicySet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, policySetID)
	ret0, _ := ret[0].(*tfe.PolicySet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockPolicySetsMockRecorder) Read(ctx, policySetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockPolicySets)(nil).Read), ctx, policySetID)
}

// ReadWithOptions mocks base method.
func (m *MockPolicySets) ReadWithOptions(ctx context.Context, policySetID string, options *tfe.PolicySetReadOptions) (*tfe.PolicySet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWithOptions", ctx, policySetID, options)
	ret0, _ := ret[0].(*tfe.PolicySet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadWithOptions indicates an expected call of ReadWithOptions.
func (mr *MockPolicySetsMockRecorder) ReadWithOptions(ctx, policySetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWithOptions", reflect.TypeOf((*MockPolicySets)(nil).ReadWithOptions), ctx, policySetID, options)
}

// RemovePolicies mocks base method.
func (m *MockPolicySets) RemovePolicies(ctx context.Context, policySetID string, options tfe.PolicySetRemovePoliciesOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePolicies", ctx, policySetID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePolicies indicates an expected call of RemovePolicies.
func (mr *MockPolicySetsMockRecorder) RemovePolicies(ctx, policySetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePolicies", reflect.TypeOf((*MockPolicySets)(nil).RemovePolicies), ctx, policySetID, options)
}

// RemoveProjects mocks base method.
func (m *MockPolicySets) RemoveProjects(ctx context.Context, policySetID string, options tfe.PolicySetRemoveProjectsOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProjects", ctx, policySetID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProjects indicates an expected call of RemoveProjects.
func (mr *MockPolicySetsMockRecorder) RemoveProjects(ctx, policySetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProjects", reflect.TypeOf((*MockPolicySets)(nil).RemoveProjects), ctx, policySetID, options)
}

// RemoveWorkspaceExclusions mocks base method.
func (m *MockPolicySets) RemoveWorkspaceExclusions(ctx context.Context, policySetID string, options tfe.PolicySetRemoveWorkspaceExclusionsOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveWorkspaceExclusions", ctx, policySetID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveWorkspaceExclusions indicates an expected call of RemoveWorkspaceExclusions.
func (mr *MockPolicySetsMockRecorder) RemoveWorkspaceExclusions(ctx, policySetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveWorkspaceExclusions", reflect.TypeOf((*MockPolicySets)(nil).RemoveWorkspaceExclusions), ctx, policySetID, options)
}

// RemoveWorkspaces mocks base method.
func (m *MockPolicySets) RemoveWorkspaces(ctx context.Context, policySetID string, options tfe.PolicySetRemoveWorkspacesOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveWorkspaces", ctx, policySetID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveWorkspaces indicates an expected call of RemoveWorkspaces.
func (mr *MockPolicySetsMockRecorder) RemoveWorkspaces(ctx, policySetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveWorkspaces", reflect.TypeOf((*MockPolicySets)(nil).RemoveWorkspaces), ctx, policySetID, options)
}

// Update mocks base method.
func (m *MockPolicySets) Update(ctx context.Context, policySetID string, options tfe.PolicySetUpdateOptions) (*tfe.PolicySet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, policySetID, options)
	ret0, _ := ret[0].(*tfe.PolicySet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPolicySetsMockRecorder) Update(ctx, policySetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPolicySets)(nil).Update), ctx, policySetID, options)
}
