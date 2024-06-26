// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/router/context.go

// Package mock_router is a generated GoMock package.
package mock_router

import (
	reflect "reflect"

	apperrors "github.com/bookpanda/minio-api/apperrors"
	dto "github.com/bookpanda/minio-api/internal/dto"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// Bind mocks base method.
func (m *MockContext) Bind(obj interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind.
func (mr *MockContextMockRecorder) Bind(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockContext)(nil).Bind), obj)
}

// FormFile mocks base method.
func (m *MockContext) FormFile(key string, allowedContentType map[string]struct{}, maxFileSize int64) (*dto.DecomposedFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormFile", key, allowedContentType, maxFileSize)
	ret0, _ := ret[0].(*dto.DecomposedFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormFile indicates an expected call of FormFile.
func (mr *MockContextMockRecorder) FormFile(key, allowedContentType, maxFileSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormFile", reflect.TypeOf((*MockContext)(nil).FormFile), key, allowedContentType, maxFileSize)
}

// JSON mocks base method.
func (m *MockContext) JSON(statusCode int, obj interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "JSON", statusCode, obj)
}

// JSON indicates an expected call of JSON.
func (mr *MockContextMockRecorder) JSON(statusCode, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSON", reflect.TypeOf((*MockContext)(nil).JSON), statusCode, obj)
}

// NewUUID mocks base method.
func (m *MockContext) NewUUID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUUID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// NewUUID indicates an expected call of NewUUID.
func (mr *MockContextMockRecorder) NewUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUUID", reflect.TypeOf((*MockContext)(nil).NewUUID))
}

// Param mocks base method.
func (m *MockContext) Param(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Param", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// Param indicates an expected call of Param.
func (mr *MockContextMockRecorder) Param(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param", reflect.TypeOf((*MockContext)(nil).Param), key)
}

// PostForm mocks base method.
func (m *MockContext) PostForm(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostForm", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// PostForm indicates an expected call of PostForm.
func (mr *MockContextMockRecorder) PostForm(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostForm", reflect.TypeOf((*MockContext)(nil).PostForm), key)
}

// Query mocks base method.
func (m *MockContext) Query(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockContextMockRecorder) Query(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockContext)(nil).Query), key)
}

// ResponseError mocks base method.
func (m *MockContext) ResponseError(err *apperrors.AppError) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResponseError", err)
}

// ResponseError indicates an expected call of ResponseError.
func (mr *MockContextMockRecorder) ResponseError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseError", reflect.TypeOf((*MockContext)(nil).ResponseError), err)
}
