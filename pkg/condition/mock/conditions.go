// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/condition/conditions.go

// Package condition is a generated GoMock package.
package condition

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/openshift/gcp-project-operator/pkg/apis/gcp/v1alpha1"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
)

// MockConditions is a mock of Conditions interface
type MockConditions struct {
	ctrl     *gomock.Controller
	recorder *MockConditionsMockRecorder
}

// MockConditionsMockRecorder is the mock recorder for MockConditions
type MockConditionsMockRecorder struct {
	mock *MockConditions
}

// NewMockConditions creates a new mock instance
func NewMockConditions(ctrl *gomock.Controller) *MockConditions {
	mock := &MockConditions{ctrl: ctrl}
	mock.recorder = &MockConditionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConditions) EXPECT() *MockConditionsMockRecorder {
	return m.recorder
}

// SetCondition mocks base method
func (m *MockConditions) SetCondition(conditions *[]v1alpha1.Condition, conditionType v1alpha1.ConditionType, status v1.ConditionStatus, reason, message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCondition", conditions, conditionType, status, reason, message)
}

// SetCondition indicates an expected call of SetCondition
func (mr *MockConditionsMockRecorder) SetCondition(conditions, conditionType, status, reason, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCondition", reflect.TypeOf((*MockConditions)(nil).SetCondition), conditions, conditionType, status, reason, message)
}
