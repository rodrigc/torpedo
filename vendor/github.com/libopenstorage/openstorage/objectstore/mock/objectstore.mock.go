// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/objectstore (interfaces: ObjectStore)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	reflect "reflect"
)

// MockObjectStore is a mock of ObjectStore interface
type MockObjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStoreMockRecorder
}

// MockObjectStoreMockRecorder is the mock recorder for MockObjectStore
type MockObjectStoreMockRecorder struct {
	mock *MockObjectStore
}

// NewMockObjectStore creates a new mock instance
func NewMockObjectStore(ctrl *gomock.Controller) *MockObjectStore {
	mock := &MockObjectStore{ctrl: ctrl}
	mock.recorder = &MockObjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockObjectStore) EXPECT() *MockObjectStoreMockRecorder {
	return m.recorder
}

// ObjectStoreCreate mocks base method
func (m *MockObjectStore) ObjectStoreCreate(arg0 string) (*api.ObjectstoreInfo, error) {
	ret := m.ctrl.Call(m, "ObjectStoreCreate", arg0)
	ret0, _ := ret[0].(*api.ObjectstoreInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectStoreCreate indicates an expected call of ObjectStoreCreate
func (mr *MockObjectStoreMockRecorder) ObjectStoreCreate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStoreCreate", reflect.TypeOf((*MockObjectStore)(nil).ObjectStoreCreate), arg0)
}

// ObjectStoreDelete mocks base method
func (m *MockObjectStore) ObjectStoreDelete(arg0 string) error {
	ret := m.ctrl.Call(m, "ObjectStoreDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ObjectStoreDelete indicates an expected call of ObjectStoreDelete
func (mr *MockObjectStoreMockRecorder) ObjectStoreDelete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStoreDelete", reflect.TypeOf((*MockObjectStore)(nil).ObjectStoreDelete), arg0)
}

// ObjectStoreInspect mocks base method
func (m *MockObjectStore) ObjectStoreInspect(arg0 string) (*api.ObjectstoreInfo, error) {
	ret := m.ctrl.Call(m, "ObjectStoreInspect", arg0)
	ret0, _ := ret[0].(*api.ObjectstoreInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectStoreInspect indicates an expected call of ObjectStoreInspect
func (mr *MockObjectStoreMockRecorder) ObjectStoreInspect(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStoreInspect", reflect.TypeOf((*MockObjectStore)(nil).ObjectStoreInspect), arg0)
}

// ObjectStoreUpdate mocks base method
func (m *MockObjectStore) ObjectStoreUpdate(arg0 string, arg1 bool) error {
	ret := m.ctrl.Call(m, "ObjectStoreUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ObjectStoreUpdate indicates an expected call of ObjectStoreUpdate
func (mr *MockObjectStoreMockRecorder) ObjectStoreUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStoreUpdate", reflect.TypeOf((*MockObjectStore)(nil).ObjectStoreUpdate), arg0, arg1)
}
