// Automatically generated by MockGen. DO NOT EDIT!
// Source: ../src/github.com/seantcanavan/threads/threadmanager.go

package threads

import (
	"fmt"
	"github.com/golang/mock/gomock"

)

// Mock of thread_actions interface
type Mockthread_actions struct {
	ctrl     *gomock.Controller
	recorder *_Mockthread_actionsRecorder
}

// Recorder for Mockthread_actions (not exported)
type _Mockthread_actionsRecorder struct {
	mock *Mockthread_actions
}

func NewMockthread_actions(ctrl *gomock.Controller) *Mockthread_actions {
	fmt.Println("func NewMockthread_actions(ctrl *gomock.Controller) *Mockthread_actions")
	mock := &Mockthread_actions{ctrl: ctrl}
	mock.recorder = &_Mockthread_actionsRecorder{mock}
	return mock
}

func (_m *Mockthread_actions) EXPECT() *_Mockthread_actionsRecorder {
	fmt.Println("(_m *Mockthread_actions) EXPECT() *_Mockthread_actionsRecorder")
	return _m.recorder
}

func (_m *Mockthread_actions) Start() {
	fmt.Println("(_m *Mockthread_actions) Start()")
	_m.ctrl.Call(_m, "Start")
}

func (_mr *_Mockthread_actionsRecorder) Start() *gomock.Call {
	fmt.Println("(_mr *_Mockthread_actionsRecorder) Start() *gomock.Call")
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *Mockthread_actions) Stop() {
	fmt.Println("(_m *Mockthread_actions) Stop()")
	_m.ctrl.Call(_m, "Stop")
}

func (_mr *_Mockthread_actionsRecorder) Stop() *gomock.Call {
	fmt.Println("(_mr *_Mockthread_actionsRecorder) Stop() *gomock.Call")
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}

func (_m *Mockthread_actions) ChangeJob() {
	fmt.Println("(_m *Mockthread_actions) ChangeJob()")
	_m.ctrl.Call(_m, "ChangeJob")
}

func (_mr *_Mockthread_actionsRecorder) ChangeJob() *gomock.Call {
	fmt.Println("(_mr *_Mockthread_actionsRecorder) ChangeJob() *gomock.Call")
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ChangeJob")
}

func (_m *Mockthread_actions) PrintDetails() {
	fmt.Println("(_m *Mockthread_actions) PrintDetails()")
	_m.ctrl.Call(_m, "PrintDetails")
}

func (_mr *_Mockthread_actionsRecorder) PrintDetails() *gomock.Call {
	fmt.Println("(_mr *_Mockthread_actionsRecorder) PrintDetails() *gomock.Call")
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PrintDetails")
}
