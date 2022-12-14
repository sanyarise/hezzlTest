// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecases/cashrepo/cashrepo.go

// Package cashrepo is a generated GoMock package.
package cashrepo

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/sanyarise/hezzl/internal/pb"
)

// MockCash is a mock of Cash interface.
type MockCash struct {
	ctrl     *gomock.Controller
	recorder *MockCashMockRecorder
}

// MockCashMockRecorder is the mock recorder for MockCash.
type MockCashMockRecorder struct {
	mock *MockCash
}

// NewMockCash creates a new mock instance.
func NewMockCash(ctrl *gomock.Controller) *MockCash {
	mock := &MockCash{ctrl: ctrl}
	mock.recorder = &MockCashMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCash) EXPECT() *MockCashMockRecorder {
	return m.recorder
}

// CheckCash mocks base method.
func (m *MockCash) CheckCash(key string) bool {
	//m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCash", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckCash indicates an expected call of CheckCash.
func (mr *MockCashMockRecorder) CheckCash(key interface{}) *gomock.Call {
	//mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCash", reflect.TypeOf((*MockCash)(nil).CheckCash), key)
}

// CreateCash mocks base method.
func (m *MockCash) CreateCash(ctx context.Context, res chan *pb.User, key string) error {
	//m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCash", ctx, res, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCash indicates an expected call of CreateCash.
func (mr *MockCashMockRecorder) CreateCash(ctx, res, key interface{}) *gomock.Call {
	//mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCash", reflect.TypeOf((*MockCash)(nil).CreateCash), ctx, res, key)
}

// GetCash mocks base method.
func (m *MockCash) GetCash(key string) ([]*pb.User, error) {
	//m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCash", key)
	ret0, _ := ret[0].([]*pb.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCash indicates an expected call of GetCash.
func (mr *MockCashMockRecorder) GetCash(key interface{}) *gomock.Call {
	//mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCash", reflect.TypeOf((*MockCash)(nil).GetCash), key)
}
