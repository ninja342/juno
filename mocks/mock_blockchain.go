// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NethermindEth/juno/blockchain (interfaces: Reader)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	core "github.com/NethermindEth/juno/core"
	felt "github.com/NethermindEth/juno/core/felt"
	gomock "github.com/golang/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// GetBlockByHash mocks base method.
func (m *MockReader) GetBlockByHash(arg0 *felt.Felt) (*core.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByHash", arg0)
	ret0, _ := ret[0].(*core.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHash indicates an expected call of GetBlockByHash.
func (mr *MockReaderMockRecorder) GetBlockByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHash", reflect.TypeOf((*MockReader)(nil).GetBlockByHash), arg0)
}

// GetBlockByNumber mocks base method.
func (m *MockReader) GetBlockByNumber(arg0 uint64) (*core.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByNumber", arg0)
	ret0, _ := ret[0].(*core.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByNumber indicates an expected call of GetBlockByNumber.
func (mr *MockReaderMockRecorder) GetBlockByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByNumber", reflect.TypeOf((*MockReader)(nil).GetBlockByNumber), arg0)
}

// GetReceipt mocks base method.
func (m *MockReader) GetReceipt(arg0 *felt.Felt) (*core.TransactionReceipt, *felt.Felt, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceipt", arg0)
	ret0, _ := ret[0].(*core.TransactionReceipt)
	ret1, _ := ret[1].(*felt.Felt)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetReceipt indicates an expected call of GetReceipt.
func (mr *MockReaderMockRecorder) GetReceipt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipt", reflect.TypeOf((*MockReader)(nil).GetReceipt), arg0)
}

// GetStateUpdateByHash mocks base method.
func (m *MockReader) GetStateUpdateByHash(arg0 *felt.Felt) (*core.StateUpdate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateUpdateByHash", arg0)
	ret0, _ := ret[0].(*core.StateUpdate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateUpdateByHash indicates an expected call of GetStateUpdateByHash.
func (mr *MockReaderMockRecorder) GetStateUpdateByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateUpdateByHash", reflect.TypeOf((*MockReader)(nil).GetStateUpdateByHash), arg0)
}

// GetStateUpdateByNumber mocks base method.
func (m *MockReader) GetStateUpdateByNumber(arg0 uint64) (*core.StateUpdate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateUpdateByNumber", arg0)
	ret0, _ := ret[0].(*core.StateUpdate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateUpdateByNumber indicates an expected call of GetStateUpdateByNumber.
func (mr *MockReaderMockRecorder) GetStateUpdateByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateUpdateByNumber", reflect.TypeOf((*MockReader)(nil).GetStateUpdateByNumber), arg0)
}

// GetTransactionByHash mocks base method.
func (m *MockReader) GetTransactionByHash(arg0 *felt.Felt) (core.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByHash", arg0)
	ret0, _ := ret[0].(core.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByHash indicates an expected call of GetTransactionByHash.
func (mr *MockReaderMockRecorder) GetTransactionByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByHash", reflect.TypeOf((*MockReader)(nil).GetTransactionByHash), arg0)
}

// Head mocks base method.
func (m *MockReader) Head() (*core.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Head")
	ret0, _ := ret[0].(*core.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Head indicates an expected call of Head.
func (mr *MockReaderMockRecorder) Head() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Head", reflect.TypeOf((*MockReader)(nil).Head))
}

// Height mocks base method.
func (m *MockReader) Height() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Height indicates an expected call of Height.
func (mr *MockReaderMockRecorder) Height() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockReader)(nil).Height))
}
