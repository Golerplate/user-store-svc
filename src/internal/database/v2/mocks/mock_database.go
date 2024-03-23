// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source interface.go -destination mocks/mock_database.go -package database_mocks
//

// Package database_mocks is a generated GoMock package.
package database_mocks

import (
	context "context"
	reflect "reflect"

	entities_user_v2 "github.com/golerplate/user-store-svc/internal/entities/user/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockDatabase) CreateUser(ctx context.Context, req *entities_user_v2.CreateUserRequest) (*entities_user_v2.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, req)
	ret0, _ := ret[0].(*entities_user_v2.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDatabaseMockRecorder) CreateUser(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDatabase)(nil).CreateUser), ctx, req)
}

// GetUserByEmail mocks base method.
func (m *MockDatabase) GetUserByEmail(ctx context.Context, email string) (*entities_user_v2.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*entities_user_v2.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockDatabaseMockRecorder) GetUserByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockDatabase)(nil).GetUserByEmail), ctx, email)
}

// GetUserByID mocks base method.
func (m *MockDatabase) GetUserByID(ctx context.Context, id string) (*entities_user_v2.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*entities_user_v2.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockDatabaseMockRecorder) GetUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockDatabase)(nil).GetUserByID), ctx, id)
}

// GetUserByUsername mocks base method.
func (m *MockDatabase) GetUserByUsername(ctx context.Context, username string) (*entities_user_v2.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", ctx, username)
	ret0, _ := ret[0].(*entities_user_v2.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockDatabaseMockRecorder) GetUserByUsername(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockDatabase)(nil).GetUserByUsername), ctx, username)
}

// UpdateUsername mocks base method.
func (m *MockDatabase) UpdateUsername(ctx context.Context, userID, username string) (*entities_user_v2.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUsername", ctx, userID, username)
	ret0, _ := ret[0].(*entities_user_v2.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUsername indicates an expected call of UpdateUsername.
func (mr *MockDatabaseMockRecorder) UpdateUsername(ctx, userID, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUsername", reflect.TypeOf((*MockDatabase)(nil).UpdateUsername), ctx, userID, username)
}
