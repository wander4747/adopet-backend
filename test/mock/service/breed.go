// Code generated by MockGen. DO NOT EDIT.
// Source: breed.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/wander4747/adopet-backend/pkg/entity"
)

// MockBreed is a mock of Breed interface.
type MockBreed struct {
	ctrl     *gomock.Controller
	recorder *MockBreedMockRecorder
}

// MockBreedMockRecorder is the mock recorder for MockBreed.
type MockBreedMockRecorder struct {
	mock *MockBreed
}

// NewMockBreed creates a new mock instance.
func NewMockBreed(ctrl *gomock.Controller) *MockBreed {
	mock := &MockBreed{ctrl: ctrl}
	mock.recorder = &MockBreedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBreed) EXPECT() *MockBreedMockRecorder {
	return m.recorder
}

// FindByAnimalID mocks base method.
func (m *MockBreed) FindByAnimalID(ctx context.Context, stateID int) ([]*entity.Breed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAnimalID", ctx, stateID)
	ret0, _ := ret[0].([]*entity.Breed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByAnimalID indicates an expected call of FindByAnimalID.
func (mr *MockBreedMockRecorder) FindByAnimalID(ctx, stateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAnimalID", reflect.TypeOf((*MockBreed)(nil).FindByAnimalID), ctx, stateID)
}
