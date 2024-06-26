// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/service/file/file.service.go

// Package mock_file is a generated GoMock package.
package mock_file

import (
	reflect "reflect"

	apperrors "github.com/bookpanda/minio-api/apperrors"
	dto "github.com/bookpanda/minio-api/internal/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockService) Delete(req *dto.DeleteFileRequest) (*dto.DeleteFileResponse, *apperrors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", req)
	ret0, _ := ret[0].(*dto.DeleteFileResponse)
	ret1, _ := ret[1].(*apperrors.AppError)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), req)
}

// Get mocks base method.
func (m *MockService) Get(req *dto.GetFileRequest) (*dto.GetFileResponse, *apperrors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", req)
	ret0, _ := ret[0].(*dto.GetFileResponse)
	ret1, _ := ret[1].(*apperrors.AppError)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServiceMockRecorder) Get(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockService)(nil).Get), req)
}

// Upload mocks base method.
func (m *MockService) Upload(req *dto.UploadFileRequest) (*dto.UploadFileResponse, *apperrors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", req)
	ret0, _ := ret[0].(*dto.UploadFileResponse)
	ret1, _ := ret[1].(*apperrors.AppError)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockServiceMockRecorder) Upload(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockService)(nil).Upload), req)
}
