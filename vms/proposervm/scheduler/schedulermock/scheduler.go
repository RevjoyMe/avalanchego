// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/vms/proposervm/scheduler (interfaces: Scheduler)
//
// Generated by this command:
//
//	mockgen -package=schedulermock -destination=vms/proposervm/scheduler/schedulermock/scheduler.go -mock_names=Scheduler=Scheduler github.com/ava-labs/avalanchego/vms/proposervm/scheduler Scheduler
//

// Package schedulermock is a generated GoMock package.
package schedulermock

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// Scheduler is a mock of Scheduler interface.
type Scheduler struct {
	ctrl     *gomock.Controller
	recorder *SchedulerMockRecorder
}

// SchedulerMockRecorder is the mock recorder for Scheduler.
type SchedulerMockRecorder struct {
	mock *Scheduler
}

// NewScheduler creates a new mock instance.
func NewScheduler(ctrl *gomock.Controller) *Scheduler {
	mock := &Scheduler{ctrl: ctrl}
	mock.recorder = &SchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Scheduler) EXPECT() *SchedulerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *Scheduler) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *SchedulerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*Scheduler)(nil).Close))
}

// Dispatch mocks base method.
func (m *Scheduler) Dispatch(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Dispatch", arg0)
}

// Dispatch indicates an expected call of Dispatch.
func (mr *SchedulerMockRecorder) Dispatch(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*Scheduler)(nil).Dispatch), arg0)
}

// SetBuildBlockTime mocks base method.
func (m *Scheduler) SetBuildBlockTime(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBuildBlockTime", arg0)
}

// SetBuildBlockTime indicates an expected call of SetBuildBlockTime.
func (mr *SchedulerMockRecorder) SetBuildBlockTime(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBuildBlockTime", reflect.TypeOf((*Scheduler)(nil).SetBuildBlockTime), arg0)
}
