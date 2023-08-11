// Code generated by MockGen. DO NOT EDIT.
// Source: order_repository.go

// Package order is a generated GoMock package.
package order

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockOrderRepository) Save(ctx context.Context, order *Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockOrderRepositoryMockRecorder) Save(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockOrderRepository)(nil).Save), ctx, order)
}
