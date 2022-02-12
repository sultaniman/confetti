// Code generated by MockGen. DO NOT EDIT.
// Source: cards.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	entities "github.com/sultaniman/confetti/platform/entities"
	repo "github.com/sultaniman/confetti/platform/repo"
)

// MockCardRepo is a mock of CardRepo interface.
type MockCardRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCardRepoMockRecorder
}

// MockCardRepoMockRecorder is the mock recorder for MockCardRepo.
type MockCardRepoMockRecorder struct {
	mock *MockCardRepo
}

// NewMockCardRepo creates a new mock instance.
func NewMockCardRepo(ctrl *gomock.Controller) *MockCardRepo {
	mock := &MockCardRepo{ctrl: ctrl}
	mock.recorder = &MockCardRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCardRepo) EXPECT() *MockCardRepoMockRecorder {
	return m.recorder
}

// ClaimExists mocks base method.
func (m *MockCardRepo) ClaimExists(cardId, userId uuid.UUID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClaimExists", cardId, userId)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ClaimExists indicates an expected call of ClaimExists.
func (mr *MockCardRepoMockRecorder) ClaimExists(cardId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClaimExists", reflect.TypeOf((*MockCardRepo)(nil).ClaimExists), cardId, userId)
}

// Create mocks base method.
func (m *MockCardRepo) Create(card *entities.NewCard) (*entities.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", card)
	ret0, _ := ret[0].(*entities.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCardRepoMockRecorder) Create(card interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCardRepo)(nil).Create), card)
}

// Delete mocks base method.
func (m *MockCardRepo) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCardRepoMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCardRepo)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockCardRepo) Get(id uuid.UUID) (*entities.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entities.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCardRepoMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCardRepo)(nil).Get), id)
}

// List mocks base method.
func (m *MockCardRepo) List(filterSpec *repo.FilterSpec) ([]entities.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", filterSpec)
	ret0, _ := ret[0].([]entities.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCardRepoMockRecorder) List(filterSpec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCardRepo)(nil).List), filterSpec)
}

// Update mocks base method.
func (m *MockCardRepo) Update(cardId uuid.UUID, newTitle string) (*entities.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", cardId, newTitle)
	ret0, _ := ret[0].(*entities.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCardRepoMockRecorder) Update(cardId, newTitle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCardRepo)(nil).Update), cardId, newTitle)
}
