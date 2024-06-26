// Code generated by MockGen. DO NOT EDIT.
// Source: userrepo.go
//
// Generated by this command:
//
//	mockgen -destination=../../../mocks/repository/userrepo/userrepo.go -package=userrepo -source=userrepo.go
//

// Package userrepo is a generated GoMock package.
package userrepo

import (
	context "context"
	models "insider_case_study/db/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockUserRepository) FindAll(filter map[string]any) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", filter)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockUserRepositoryMockRecorder) FindAll(filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockUserRepository)(nil).FindAll), filter)
}

// GetMessageResponse mocks base method.
func (m *MockUserRepository) GetMessageResponse(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageResponse", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageResponse indicates an expected call of GetMessageResponse.
func (mr *MockUserRepositoryMockRecorder) GetMessageResponse(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageResponse", reflect.TypeOf((*MockUserRepository)(nil).GetMessageResponse), ctx, key)
}

// SetMessageResponse mocks base method.
func (m *MockUserRepository) SetMessageResponse(ctx context.Context, key, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMessageResponse", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMessageResponse indicates an expected call of SetMessageResponse.
func (mr *MockUserRepositoryMockRecorder) SetMessageResponse(ctx, key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMessageResponse", reflect.TypeOf((*MockUserRepository)(nil).SetMessageResponse), ctx, key, value)
}

// Update mocks base method.
func (m *MockUserRepository) Update(user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), user)
}
