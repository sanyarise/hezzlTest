// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecases/userrepo/userrepo.go

// Package userrepo is a generated GoMock package.
package userrepo

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/sanyarise/hezzl/internal/pb"
)

// MockUserStore is a mock of UserStore interface.
type MockUserStore struct {
	ctrl     *gomock.Controller
	recorder *MockUserStoreMockRecorder
}

// MockUserStoreMockRecorder is the mock recorder for MockUserStore.
type MockUserStoreMockRecorder struct {
	mock *MockUserStore
}

// NewMockUserStore creates a new mock instance.
func NewMockUserStore(ctrl *gomock.Controller) *MockUserStore {
	mock := &MockUserStore{ctrl: ctrl}
	mock.recorder = &MockUserStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStore) EXPECT() *MockUserStoreMockRecorder {
	return m.recorder
}

// DeleteUser mocks base method.
func (m *MockUserStore) DeleteUser(ctx context.Context, id string) error {
	//m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserStoreMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	//mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserStore)(nil).DeleteUser), ctx, id)
}

// GetAllUsers mocks base method.
func (m *MockUserStore) GetAllUsers(ctx context.Context) (chan *pb.User, error) {
	//m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", ctx)
	ret0, _ := ret[0].(chan *pb.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserStoreMockRecorder) GetAllUsers(ctx interface{}) *gomock.Call {
	//mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserStore)(nil).GetAllUsers), ctx)
}

// SaveUser mocks base method.
func (m *MockUserStore) SaveUser(ctx context.Context, user *pb.User) error {
	//m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockUserStoreMockRecorder) SaveUser(ctx, user interface{}) *gomock.Call {
	//mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockUserStore)(nil).SaveUser), ctx, user)
}
