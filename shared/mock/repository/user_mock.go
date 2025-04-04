// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/user.go
//
// Generated by this command:
//
//	mockgen -source=./internal/repository/user.go -destination=./shared/mock/repository/user_mock.go -package repository
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	domain "github.com/imanudd/inventorySvc-clean-architecture/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepositoryImpl is a mock of UserRepositoryImpl interface.
type MockUserRepositoryImpl struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryImplMockRecorder
	isgomock struct{}
}

// MockUserRepositoryImplMockRecorder is the mock recorder for MockUserRepositoryImpl.
type MockUserRepositoryImplMockRecorder struct {
	mock *MockUserRepositoryImpl
}

// NewMockUserRepositoryImpl creates a new mock instance.
func NewMockUserRepositoryImpl(ctrl *gomock.Controller) *MockUserRepositoryImpl {
	mock := &MockUserRepositoryImpl{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryImpl) EXPECT() *MockUserRepositoryImplMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockUserRepositoryImpl) GetByID(ctx context.Context, id int) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUserRepositoryImplMockRecorder) GetByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserRepositoryImpl)(nil).GetByID), ctx, id)
}

// GetByUsernameOrEmail mocks base method.
func (m *MockUserRepositoryImpl) GetByUsernameOrEmail(ctx context.Context, req *domain.GetByUsernameOrEmail) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsernameOrEmail", ctx, req)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsernameOrEmail indicates an expected call of GetByUsernameOrEmail.
func (mr *MockUserRepositoryImplMockRecorder) GetByUsernameOrEmail(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsernameOrEmail", reflect.TypeOf((*MockUserRepositoryImpl)(nil).GetByUsernameOrEmail), ctx, req)
}

// RegisterUser mocks base method.
func (m *MockUserRepositoryImpl) RegisterUser(ctx context.Context, req *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUserRepositoryImplMockRecorder) RegisterUser(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUserRepositoryImpl)(nil).RegisterUser), ctx, req)
}
