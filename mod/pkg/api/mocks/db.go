// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jonnylangefeld/go-api/pkg/db (interfaces: ClientInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "github.com/PauloLancao/GOLang/mod/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// MockClientInterface is a mock of ClientInterface interface
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

// MockClientInterfaceMockRecorder is the mock recorder for MockClientInterface
type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

// NewMockClientInterface creates a new mock instance
func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

// Connect mocks base method
func (m *MockClientInterface) Connect(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockClientInterfaceMockRecorder) Connect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockClientInterface)(nil).Connect), arg0)
}

// GetArticleByID mocks base method
func (m *MockClientInterface) GetArticleByID(arg0 int) *types.Article {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleByID", arg0)
	ret0, _ := ret[0].(*types.Article)
	return ret0
}

// GetArticleByID indicates an expected call of GetArticleByID
func (mr *MockClientInterfaceMockRecorder) GetArticleByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleByID", reflect.TypeOf((*MockClientInterface)(nil).GetArticleByID), arg0)
}

// GetArticles mocks base method
func (m *MockClientInterface) GetArticles(arg0 int) *types.ArticleList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticles", arg0)
	ret0, _ := ret[0].(*types.ArticleList)
	return ret0
}

// GetArticles indicates an expected call of GetArticles
func (mr *MockClientInterfaceMockRecorder) GetArticles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockClientInterface)(nil).GetArticles), arg0)
}

// GetOrderByID mocks base method
func (m *MockClientInterface) GetOrderByID(arg0 int) *types.Order {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByID", arg0)
	ret0, _ := ret[0].(*types.Order)
	return ret0
}

// GetOrderByID indicates an expected call of GetOrderByID
func (mr *MockClientInterfaceMockRecorder) GetOrderByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByID", reflect.TypeOf((*MockClientInterface)(nil).GetOrderByID), arg0)
}

// GetOrders mocks base method
func (m *MockClientInterface) GetOrders(arg0 int) *types.OrderList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", arg0)
	ret0, _ := ret[0].(*types.OrderList)
	return ret0
}

// GetOrders indicates an expected call of GetOrders
func (mr *MockClientInterfaceMockRecorder) GetOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockClientInterface)(nil).GetOrders), arg0)
}

// Ping mocks base method
func (m *MockClientInterface) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockClientInterfaceMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockClientInterface)(nil).Ping))
}

// SetArticle mocks base method
func (m *MockClientInterface) SetArticle(arg0 *types.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetArticle", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetArticle indicates an expected call of SetArticle
func (mr *MockClientInterfaceMockRecorder) SetArticle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetArticle", reflect.TypeOf((*MockClientInterface)(nil).SetArticle), arg0)
}

// SetOrder mocks base method
func (m *MockClientInterface) SetOrder(arg0 *types.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOrder", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOrder indicates an expected call of SetOrder
func (mr *MockClientInterfaceMockRecorder) SetOrder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrder", reflect.TypeOf((*MockClientInterface)(nil).SetOrder), arg0)
}
