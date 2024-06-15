// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/file/file.repository.go

// Package mock_file is a generated GoMock package.
package mock_file

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRepository) Delete(bucketName, objectKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", bucketName, objectKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), bucketName, objectKey)
}

// Get mocks base method.
func (m *MockRepository) Get(bucketName, objectKey string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", bucketName, objectKey)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), bucketName, objectKey)
}

// GetURL mocks base method.
func (m *MockRepository) GetURL(bucketName, objectKey string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL", bucketName, objectKey)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetURL indicates an expected call of GetURL.
func (mr *MockRepositoryMockRecorder) GetURL(bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockRepository)(nil).GetURL), bucketName, objectKey)
}

// Upload mocks base method.
func (m *MockRepository) Upload(file []byte, bucketName, objectKey string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", file, bucketName, objectKey)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Upload indicates an expected call of Upload.
func (mr *MockRepositoryMockRecorder) Upload(file, bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockRepository)(nil).Upload), file, bucketName, objectKey)
}