// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockClient) Create(ctx context.Context, obj *unstructured.Unstructured, options v1.CreateOptions, subresources ...string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockClientMockRecorder) Create(ctx, obj, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClient)(nil).Create), varargs...)
}

// Update mocks base method
func (m *MockClient) Update(ctx context.Context, obj *unstructured.Unstructured, options v1.UpdateOptions, subresources ...string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockClientMockRecorder) Update(ctx, obj, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClient)(nil).Update), varargs...)
}

// UpdateStatus mocks base method
func (m *MockClient) UpdateStatus(ctx context.Context, obj *unstructured.Unstructured, options v1.UpdateOptions) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, obj, options)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockClientMockRecorder) UpdateStatus(ctx, obj, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockClient)(nil).UpdateStatus), ctx, obj, options)
}

// Delete mocks base method
func (m *MockClient) Delete(ctx context.Context, name string, options v1.DeleteOptions, subresources ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, name, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockClientMockRecorder) Delete(ctx, name, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, name, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), varargs...)
}

// DeleteCollection mocks base method
func (m *MockClient) DeleteCollection(ctx context.Context, options v1.DeleteOptions, listOptions v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", ctx, options, listOptions)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockClientMockRecorder) DeleteCollection(ctx, options, listOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockClient)(nil).DeleteCollection), ctx, options, listOptions)
}

// Get mocks base method
func (m *MockClient) Get(ctx context.Context, name string, options v1.GetOptions, subresources ...string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, name, options}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockClientMockRecorder) Get(ctx, name, options interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, name, options}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), varargs...)
}

// List mocks base method
func (m *MockClient) List(ctx context.Context, opts v1.ListOptions) (*unstructured.UnstructuredList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, opts)
	ret0, _ := ret[0].(*unstructured.UnstructuredList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockClientMockRecorder) List(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClient)(nil).List), ctx, opts)
}
