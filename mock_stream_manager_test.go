// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dragonbreath2000/quic-go (interfaces: StreamManager)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/dragonbreath2000/quic-go -destination mock_stream_manager_test.go github.com/dragonbreath2000/quic-go StreamManager
//

// Package quic is a generated GoMock package.
package quic

import (
	context "context"
	reflect "reflect"

	protocol "github.com/dragonbreath2000/quic-go/internal/protocol"
	wire "github.com/dragonbreath2000/quic-go/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockStreamManager is a mock of StreamManager interface.
type MockStreamManager struct {
	ctrl     *gomock.Controller
	recorder *MockStreamManagerMockRecorder
}

// MockStreamManagerMockRecorder is the mock recorder for MockStreamManager.
type MockStreamManagerMockRecorder struct {
	mock *MockStreamManager
}

// NewMockStreamManager creates a new mock instance.
func NewMockStreamManager(ctrl *gomock.Controller) *MockStreamManager {
	mock := &MockStreamManager{ctrl: ctrl}
	mock.recorder = &MockStreamManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamManager) EXPECT() *MockStreamManagerMockRecorder {
	return m.recorder
}

// AcceptStream mocks base method.
func (m *MockStreamManager) AcceptStream(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptStream", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptStream indicates an expected call of AcceptStream.
func (mr *MockStreamManagerMockRecorder) AcceptStream(arg0 any) *MockStreamManagerAcceptStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptStream", reflect.TypeOf((*MockStreamManager)(nil).AcceptStream), arg0)
	return &MockStreamManagerAcceptStreamCall{Call: call}
}

// MockStreamManagerAcceptStreamCall wrap *gomock.Call
type MockStreamManagerAcceptStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerAcceptStreamCall) Return(arg0 Stream, arg1 error) *MockStreamManagerAcceptStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerAcceptStreamCall) Do(f func(context.Context) (Stream, error)) *MockStreamManagerAcceptStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerAcceptStreamCall) DoAndReturn(f func(context.Context) (Stream, error)) *MockStreamManagerAcceptStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AcceptUniStream mocks base method.
func (m *MockStreamManager) AcceptUniStream(arg0 context.Context) (ReceiveStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptUniStream", arg0)
	ret0, _ := ret[0].(ReceiveStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptUniStream indicates an expected call of AcceptUniStream.
func (mr *MockStreamManagerMockRecorder) AcceptUniStream(arg0 any) *MockStreamManagerAcceptUniStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptUniStream", reflect.TypeOf((*MockStreamManager)(nil).AcceptUniStream), arg0)
	return &MockStreamManagerAcceptUniStreamCall{Call: call}
}

// MockStreamManagerAcceptUniStreamCall wrap *gomock.Call
type MockStreamManagerAcceptUniStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerAcceptUniStreamCall) Return(arg0 ReceiveStream, arg1 error) *MockStreamManagerAcceptUniStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerAcceptUniStreamCall) Do(f func(context.Context) (ReceiveStream, error)) *MockStreamManagerAcceptUniStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerAcceptUniStreamCall) DoAndReturn(f func(context.Context) (ReceiveStream, error)) *MockStreamManagerAcceptUniStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CloseWithError mocks base method.
func (m *MockStreamManager) CloseWithError(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseWithError", arg0)
}

// CloseWithError indicates an expected call of CloseWithError.
func (mr *MockStreamManagerMockRecorder) CloseWithError(arg0 any) *MockStreamManagerCloseWithErrorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseWithError", reflect.TypeOf((*MockStreamManager)(nil).CloseWithError), arg0)
	return &MockStreamManagerCloseWithErrorCall{Call: call}
}

// MockStreamManagerCloseWithErrorCall wrap *gomock.Call
type MockStreamManagerCloseWithErrorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerCloseWithErrorCall) Return() *MockStreamManagerCloseWithErrorCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerCloseWithErrorCall) Do(f func(error)) *MockStreamManagerCloseWithErrorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerCloseWithErrorCall) DoAndReturn(f func(error)) *MockStreamManagerCloseWithErrorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteStream mocks base method.
func (m *MockStreamManager) DeleteStream(arg0 protocol.StreamID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStream", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStream indicates an expected call of DeleteStream.
func (mr *MockStreamManagerMockRecorder) DeleteStream(arg0 any) *MockStreamManagerDeleteStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStream", reflect.TypeOf((*MockStreamManager)(nil).DeleteStream), arg0)
	return &MockStreamManagerDeleteStreamCall{Call: call}
}

// MockStreamManagerDeleteStreamCall wrap *gomock.Call
type MockStreamManagerDeleteStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerDeleteStreamCall) Return(arg0 error) *MockStreamManagerDeleteStreamCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerDeleteStreamCall) Do(f func(protocol.StreamID) error) *MockStreamManagerDeleteStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerDeleteStreamCall) DoAndReturn(f func(protocol.StreamID) error) *MockStreamManagerDeleteStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetOrOpenReceiveStream mocks base method.
func (m *MockStreamManager) GetOrOpenReceiveStream(arg0 protocol.StreamID) (receiveStreamI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrOpenReceiveStream", arg0)
	ret0, _ := ret[0].(receiveStreamI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrOpenReceiveStream indicates an expected call of GetOrOpenReceiveStream.
func (mr *MockStreamManagerMockRecorder) GetOrOpenReceiveStream(arg0 any) *MockStreamManagerGetOrOpenReceiveStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrOpenReceiveStream", reflect.TypeOf((*MockStreamManager)(nil).GetOrOpenReceiveStream), arg0)
	return &MockStreamManagerGetOrOpenReceiveStreamCall{Call: call}
}

// MockStreamManagerGetOrOpenReceiveStreamCall wrap *gomock.Call
type MockStreamManagerGetOrOpenReceiveStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerGetOrOpenReceiveStreamCall) Return(arg0 receiveStreamI, arg1 error) *MockStreamManagerGetOrOpenReceiveStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerGetOrOpenReceiveStreamCall) Do(f func(protocol.StreamID) (receiveStreamI, error)) *MockStreamManagerGetOrOpenReceiveStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerGetOrOpenReceiveStreamCall) DoAndReturn(f func(protocol.StreamID) (receiveStreamI, error)) *MockStreamManagerGetOrOpenReceiveStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetOrOpenSendStream mocks base method.
func (m *MockStreamManager) GetOrOpenSendStream(arg0 protocol.StreamID) (sendStreamI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrOpenSendStream", arg0)
	ret0, _ := ret[0].(sendStreamI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrOpenSendStream indicates an expected call of GetOrOpenSendStream.
func (mr *MockStreamManagerMockRecorder) GetOrOpenSendStream(arg0 any) *MockStreamManagerGetOrOpenSendStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrOpenSendStream", reflect.TypeOf((*MockStreamManager)(nil).GetOrOpenSendStream), arg0)
	return &MockStreamManagerGetOrOpenSendStreamCall{Call: call}
}

// MockStreamManagerGetOrOpenSendStreamCall wrap *gomock.Call
type MockStreamManagerGetOrOpenSendStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerGetOrOpenSendStreamCall) Return(arg0 sendStreamI, arg1 error) *MockStreamManagerGetOrOpenSendStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerGetOrOpenSendStreamCall) Do(f func(protocol.StreamID) (sendStreamI, error)) *MockStreamManagerGetOrOpenSendStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerGetOrOpenSendStreamCall) DoAndReturn(f func(protocol.StreamID) (sendStreamI, error)) *MockStreamManagerGetOrOpenSendStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HandleMaxStreamsFrame mocks base method.
func (m *MockStreamManager) HandleMaxStreamsFrame(arg0 *wire.MaxStreamsFrame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleMaxStreamsFrame", arg0)
}

// HandleMaxStreamsFrame indicates an expected call of HandleMaxStreamsFrame.
func (mr *MockStreamManagerMockRecorder) HandleMaxStreamsFrame(arg0 any) *MockStreamManagerHandleMaxStreamsFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMaxStreamsFrame", reflect.TypeOf((*MockStreamManager)(nil).HandleMaxStreamsFrame), arg0)
	return &MockStreamManagerHandleMaxStreamsFrameCall{Call: call}
}

// MockStreamManagerHandleMaxStreamsFrameCall wrap *gomock.Call
type MockStreamManagerHandleMaxStreamsFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerHandleMaxStreamsFrameCall) Return() *MockStreamManagerHandleMaxStreamsFrameCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerHandleMaxStreamsFrameCall) Do(f func(*wire.MaxStreamsFrame)) *MockStreamManagerHandleMaxStreamsFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerHandleMaxStreamsFrameCall) DoAndReturn(f func(*wire.MaxStreamsFrame)) *MockStreamManagerHandleMaxStreamsFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenStream mocks base method.
func (m *MockStreamManager) OpenStream() (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStream")
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStream indicates an expected call of OpenStream.
func (mr *MockStreamManagerMockRecorder) OpenStream() *MockStreamManagerOpenStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStream", reflect.TypeOf((*MockStreamManager)(nil).OpenStream))
	return &MockStreamManagerOpenStreamCall{Call: call}
}

// MockStreamManagerOpenStreamCall wrap *gomock.Call
type MockStreamManagerOpenStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerOpenStreamCall) Return(arg0 Stream, arg1 error) *MockStreamManagerOpenStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerOpenStreamCall) Do(f func() (Stream, error)) *MockStreamManagerOpenStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerOpenStreamCall) DoAndReturn(f func() (Stream, error)) *MockStreamManagerOpenStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenStreamSync mocks base method.
func (m *MockStreamManager) OpenStreamSync(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStreamSync", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStreamSync indicates an expected call of OpenStreamSync.
func (mr *MockStreamManagerMockRecorder) OpenStreamSync(arg0 any) *MockStreamManagerOpenStreamSyncCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStreamSync", reflect.TypeOf((*MockStreamManager)(nil).OpenStreamSync), arg0)
	return &MockStreamManagerOpenStreamSyncCall{Call: call}
}

// MockStreamManagerOpenStreamSyncCall wrap *gomock.Call
type MockStreamManagerOpenStreamSyncCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerOpenStreamSyncCall) Return(arg0 Stream, arg1 error) *MockStreamManagerOpenStreamSyncCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerOpenStreamSyncCall) Do(f func(context.Context) (Stream, error)) *MockStreamManagerOpenStreamSyncCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerOpenStreamSyncCall) DoAndReturn(f func(context.Context) (Stream, error)) *MockStreamManagerOpenStreamSyncCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenUniStream mocks base method.
func (m *MockStreamManager) OpenUniStream() (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStream")
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStream indicates an expected call of OpenUniStream.
func (mr *MockStreamManagerMockRecorder) OpenUniStream() *MockStreamManagerOpenUniStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStream", reflect.TypeOf((*MockStreamManager)(nil).OpenUniStream))
	return &MockStreamManagerOpenUniStreamCall{Call: call}
}

// MockStreamManagerOpenUniStreamCall wrap *gomock.Call
type MockStreamManagerOpenUniStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerOpenUniStreamCall) Return(arg0 SendStream, arg1 error) *MockStreamManagerOpenUniStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerOpenUniStreamCall) Do(f func() (SendStream, error)) *MockStreamManagerOpenUniStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerOpenUniStreamCall) DoAndReturn(f func() (SendStream, error)) *MockStreamManagerOpenUniStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// OpenUniStreamSync mocks base method.
func (m *MockStreamManager) OpenUniStreamSync(arg0 context.Context) (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStreamSync", arg0)
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStreamSync indicates an expected call of OpenUniStreamSync.
func (mr *MockStreamManagerMockRecorder) OpenUniStreamSync(arg0 any) *MockStreamManagerOpenUniStreamSyncCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStreamSync", reflect.TypeOf((*MockStreamManager)(nil).OpenUniStreamSync), arg0)
	return &MockStreamManagerOpenUniStreamSyncCall{Call: call}
}

// MockStreamManagerOpenUniStreamSyncCall wrap *gomock.Call
type MockStreamManagerOpenUniStreamSyncCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerOpenUniStreamSyncCall) Return(arg0 SendStream, arg1 error) *MockStreamManagerOpenUniStreamSyncCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerOpenUniStreamSyncCall) Do(f func(context.Context) (SendStream, error)) *MockStreamManagerOpenUniStreamSyncCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerOpenUniStreamSyncCall) DoAndReturn(f func(context.Context) (SendStream, error)) *MockStreamManagerOpenUniStreamSyncCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ResetFor0RTT mocks base method.
func (m *MockStreamManager) ResetFor0RTT() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetFor0RTT")
}

// ResetFor0RTT indicates an expected call of ResetFor0RTT.
func (mr *MockStreamManagerMockRecorder) ResetFor0RTT() *MockStreamManagerResetFor0RTTCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetFor0RTT", reflect.TypeOf((*MockStreamManager)(nil).ResetFor0RTT))
	return &MockStreamManagerResetFor0RTTCall{Call: call}
}

// MockStreamManagerResetFor0RTTCall wrap *gomock.Call
type MockStreamManagerResetFor0RTTCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerResetFor0RTTCall) Return() *MockStreamManagerResetFor0RTTCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerResetFor0RTTCall) Do(f func()) *MockStreamManagerResetFor0RTTCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerResetFor0RTTCall) DoAndReturn(f func()) *MockStreamManagerResetFor0RTTCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateLimits mocks base method.
func (m *MockStreamManager) UpdateLimits(arg0 *wire.TransportParameters) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateLimits", arg0)
}

// UpdateLimits indicates an expected call of UpdateLimits.
func (mr *MockStreamManagerMockRecorder) UpdateLimits(arg0 any) *MockStreamManagerUpdateLimitsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLimits", reflect.TypeOf((*MockStreamManager)(nil).UpdateLimits), arg0)
	return &MockStreamManagerUpdateLimitsCall{Call: call}
}

// MockStreamManagerUpdateLimitsCall wrap *gomock.Call
type MockStreamManagerUpdateLimitsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerUpdateLimitsCall) Return() *MockStreamManagerUpdateLimitsCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerUpdateLimitsCall) Do(f func(*wire.TransportParameters)) *MockStreamManagerUpdateLimitsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerUpdateLimitsCall) DoAndReturn(f func(*wire.TransportParameters)) *MockStreamManagerUpdateLimitsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UseResetMaps mocks base method.
func (m *MockStreamManager) UseResetMaps() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UseResetMaps")
}

// UseResetMaps indicates an expected call of UseResetMaps.
func (mr *MockStreamManagerMockRecorder) UseResetMaps() *MockStreamManagerUseResetMapsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseResetMaps", reflect.TypeOf((*MockStreamManager)(nil).UseResetMaps))
	return &MockStreamManagerUseResetMapsCall{Call: call}
}

// MockStreamManagerUseResetMapsCall wrap *gomock.Call
type MockStreamManagerUseResetMapsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamManagerUseResetMapsCall) Return() *MockStreamManagerUseResetMapsCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamManagerUseResetMapsCall) Do(f func()) *MockStreamManagerUseResetMapsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamManagerUseResetMapsCall) DoAndReturn(f func()) *MockStreamManagerUseResetMapsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
