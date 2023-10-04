// Code generated by MockGen. DO NOT EDIT.
// Source: admin_setting_cost_estimation.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/eranelbaz/go-tfe"
	gomock "github.com/golang/mock/gomock"
)

// MockCostEstimationSettings is a mock of CostEstimationSettings interface.
type MockCostEstimationSettings struct {
	ctrl     *gomock.Controller
	recorder *MockCostEstimationSettingsMockRecorder
}

// MockCostEstimationSettingsMockRecorder is the mock recorder for MockCostEstimationSettings.
type MockCostEstimationSettingsMockRecorder struct {
	mock *MockCostEstimationSettings
}

// NewMockCostEstimationSettings creates a new mock instance.
func NewMockCostEstimationSettings(ctrl *gomock.Controller) *MockCostEstimationSettings {
	mock := &MockCostEstimationSettings{ctrl: ctrl}
	mock.recorder = &MockCostEstimationSettingsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCostEstimationSettings) EXPECT() *MockCostEstimationSettingsMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockCostEstimationSettings) Read(ctx context.Context) (*tfe.AdminCostEstimationSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx)
	ret0, _ := ret[0].(*tfe.AdminCostEstimationSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockCostEstimationSettingsMockRecorder) Read(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockCostEstimationSettings)(nil).Read), ctx)
}

// Update mocks base method.
func (m *MockCostEstimationSettings) Update(ctx context.Context, options tfe.AdminCostEstimationSettingOptions) (*tfe.AdminCostEstimationSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, options)
	ret0, _ := ret[0].(*tfe.AdminCostEstimationSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCostEstimationSettingsMockRecorder) Update(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCostEstimationSettings)(nil).Update), ctx, options)
}
