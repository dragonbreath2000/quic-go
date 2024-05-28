// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dragonbreath2000/quic-go (interfaces: FrameSource)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/dragonbreath2000/quic-go -destination mock_frame_source_test.go github.com/dragonbreath2000/quic-go FrameSource
//

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	ackhandler "github.com/dragonbreath2000/quic-go/internal/ackhandler"
	protocol "github.com/dragonbreath2000/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockFrameSource is a mock of FrameSource interface.
type MockFrameSource struct {
	ctrl     *gomock.Controller
	recorder *MockFrameSourceMockRecorder
}

// MockFrameSourceMockRecorder is the mock recorder for MockFrameSource.
type MockFrameSourceMockRecorder struct {
	mock *MockFrameSource
}

// NewMockFrameSource creates a new mock instance.
func NewMockFrameSource(ctrl *gomock.Controller) *MockFrameSource {
	mock := &MockFrameSource{ctrl: ctrl}
	mock.recorder = &MockFrameSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFrameSource) EXPECT() *MockFrameSourceMockRecorder {
	return m.recorder
}

// AppendControlFrames mocks base method.
func (m *MockFrameSource) AppendControlFrames(arg0 []ackhandler.Frame, arg1 protocol.ByteCount, arg2 protocol.Version) ([]ackhandler.Frame, protocol.ByteCount) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendControlFrames", arg0, arg1, arg2)
	ret0, _ := ret[0].([]ackhandler.Frame)
	ret1, _ := ret[1].(protocol.ByteCount)
	return ret0, ret1
}

// AppendControlFrames indicates an expected call of AppendControlFrames.
func (mr *MockFrameSourceMockRecorder) AppendControlFrames(arg0, arg1, arg2 any) *MockFrameSourceAppendControlFramesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendControlFrames", reflect.TypeOf((*MockFrameSource)(nil).AppendControlFrames), arg0, arg1, arg2)
	return &MockFrameSourceAppendControlFramesCall{Call: call}
}

// MockFrameSourceAppendControlFramesCall wrap *gomock.Call
type MockFrameSourceAppendControlFramesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFrameSourceAppendControlFramesCall) Return(arg0 []ackhandler.Frame, arg1 protocol.ByteCount) *MockFrameSourceAppendControlFramesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFrameSourceAppendControlFramesCall) Do(f func([]ackhandler.Frame, protocol.ByteCount, protocol.Version) ([]ackhandler.Frame, protocol.ByteCount)) *MockFrameSourceAppendControlFramesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFrameSourceAppendControlFramesCall) DoAndReturn(f func([]ackhandler.Frame, protocol.ByteCount, protocol.Version) ([]ackhandler.Frame, protocol.ByteCount)) *MockFrameSourceAppendControlFramesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AppendStreamFrames mocks base method.
func (m *MockFrameSource) AppendStreamFrames(arg0 []ackhandler.StreamFrame, arg1 protocol.ByteCount, arg2 protocol.Version) ([]ackhandler.StreamFrame, protocol.ByteCount) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendStreamFrames", arg0, arg1, arg2)
	ret0, _ := ret[0].([]ackhandler.StreamFrame)
	ret1, _ := ret[1].(protocol.ByteCount)
	return ret0, ret1
}

// AppendStreamFrames indicates an expected call of AppendStreamFrames.
func (mr *MockFrameSourceMockRecorder) AppendStreamFrames(arg0, arg1, arg2 any) *MockFrameSourceAppendStreamFramesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendStreamFrames", reflect.TypeOf((*MockFrameSource)(nil).AppendStreamFrames), arg0, arg1, arg2)
	return &MockFrameSourceAppendStreamFramesCall{Call: call}
}

// MockFrameSourceAppendStreamFramesCall wrap *gomock.Call
type MockFrameSourceAppendStreamFramesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFrameSourceAppendStreamFramesCall) Return(arg0 []ackhandler.StreamFrame, arg1 protocol.ByteCount) *MockFrameSourceAppendStreamFramesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFrameSourceAppendStreamFramesCall) Do(f func([]ackhandler.StreamFrame, protocol.ByteCount, protocol.Version) ([]ackhandler.StreamFrame, protocol.ByteCount)) *MockFrameSourceAppendStreamFramesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFrameSourceAppendStreamFramesCall) DoAndReturn(f func([]ackhandler.StreamFrame, protocol.ByteCount, protocol.Version) ([]ackhandler.StreamFrame, protocol.ByteCount)) *MockFrameSourceAppendStreamFramesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HasData mocks base method.
func (m *MockFrameSource) HasData() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasData")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasData indicates an expected call of HasData.
func (mr *MockFrameSourceMockRecorder) HasData() *MockFrameSourceHasDataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasData", reflect.TypeOf((*MockFrameSource)(nil).HasData))
	return &MockFrameSourceHasDataCall{Call: call}
}

// MockFrameSourceHasDataCall wrap *gomock.Call
type MockFrameSourceHasDataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFrameSourceHasDataCall) Return(arg0 bool) *MockFrameSourceHasDataCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFrameSourceHasDataCall) Do(f func() bool) *MockFrameSourceHasDataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFrameSourceHasDataCall) DoAndReturn(f func() bool) *MockFrameSourceHasDataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
