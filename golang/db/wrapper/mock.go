// Code generated by MockGen. DO NOT EDIT.
// Source: db/wrapper/wrapper.go

// Package wrapper is a generated GoMock package.
package wrapper

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/lukejmann/blockshot/golang/db"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// GetFlawedMintsForBlock mocks base method.
func (m *MockQuerier) GetFlawedMintsForBlock(ctx context.Context, blockNum int32) ([]db.Mint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlawedMintsForBlock", ctx, blockNum)
	ret0, _ := ret[0].([]db.Mint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlawedMintsForBlock indicates an expected call of GetFlawedMintsForBlock.
func (mr *MockQuerierMockRecorder) GetFlawedMintsForBlock(ctx, blockNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlawedMintsForBlock", reflect.TypeOf((*MockQuerier)(nil).GetFlawedMintsForBlock), ctx, blockNum)
}

// GetHighestBlock mocks base method.
func (m *MockQuerier) GetHighestBlock(ctx context.Context) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHighestBlock", ctx)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHighestBlock indicates an expected call of GetHighestBlock.
func (mr *MockQuerierMockRecorder) GetHighestBlock(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHighestBlock", reflect.TypeOf((*MockQuerier)(nil).GetHighestBlock), ctx)
}

// GetMintsForBlock mocks base method.
func (m *MockQuerier) GetMintsForBlock(ctx context.Context, blockNum int32) ([]db.Mint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMintsForBlock", ctx, blockNum)
	ret0, _ := ret[0].([]db.Mint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMintsForBlock indicates an expected call of GetMintsForBlock.
func (mr *MockQuerierMockRecorder) GetMintsForBlock(ctx, blockNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMintsForBlock", reflect.TypeOf((*MockQuerier)(nil).GetMintsForBlock), ctx, blockNum)
}

// InsertMintWithImageData mocks base method.
func (m *MockQuerier) InsertMintWithImageData(ctx context.Context, arg db.InsertMintWithImageDataParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMintWithImageData", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertMintWithImageData indicates an expected call of InsertMintWithImageData.
func (mr *MockQuerierMockRecorder) InsertMintWithImageData(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMintWithImageData", reflect.TypeOf((*MockQuerier)(nil).InsertMintWithImageData), ctx, arg)
}

// InsertMintWithImageURL mocks base method.
func (m *MockQuerier) InsertMintWithImageURL(ctx context.Context, arg db.InsertMintWithImageURLParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMintWithImageURL", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertMintWithImageURL indicates an expected call of InsertMintWithImageURL.
func (mr *MockQuerierMockRecorder) InsertMintWithImageURL(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMintWithImageURL", reflect.TypeOf((*MockQuerier)(nil).InsertMintWithImageURL), ctx, arg)
}

// WithTx mocks base method.
func (m *MockQuerier) WithTx(arg0 context.Context, arg1 func(db.Querier) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithTx indicates an expected call of WithTx.
func (mr *MockQuerierMockRecorder) WithTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockQuerier)(nil).WithTx), arg0, arg1)
}
