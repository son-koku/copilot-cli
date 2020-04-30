// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/archer/app.go

// Package mocks is a generated GoMock package.
package mocks

import (
	archer "github.com/aws/amazon-ecs-cli-v2/internal/pkg/archer"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockApplicationStore is a mock of ApplicationStore interface
type MockApplicationStore struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationStoreMockRecorder
}

// MockApplicationStoreMockRecorder is the mock recorder for MockApplicationStore
type MockApplicationStoreMockRecorder struct {
	mock *MockApplicationStore
}

// NewMockApplicationStore creates a new mock instance
func NewMockApplicationStore(ctrl *gomock.Controller) *MockApplicationStore {
	mock := &MockApplicationStore{ctrl: ctrl}
	mock.recorder = &MockApplicationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationStore) EXPECT() *MockApplicationStoreMockRecorder {
	return m.recorder
}

// ListServices mocks base method
func (m *MockApplicationStore) ListServices(projectName string) ([]*archer.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", projectName)
	ret0, _ := ret[0].([]*archer.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockApplicationStoreMockRecorder) ListServices(projectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockApplicationStore)(nil).ListServices), projectName)
}

// GetService mocks base method
func (m *MockApplicationStore) GetService(projectName, applicationName string) (*archer.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", projectName, applicationName)
	ret0, _ := ret[0].(*archer.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockApplicationStoreMockRecorder) GetService(projectName, applicationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockApplicationStore)(nil).GetService), projectName, applicationName)
}

// CreateService mocks base method
func (m *MockApplicationStore) CreateService(app *archer.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", app)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateService indicates an expected call of CreateService
func (mr *MockApplicationStoreMockRecorder) CreateService(app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockApplicationStore)(nil).CreateService), app)
}

// DeleteService mocks base method
func (m *MockApplicationStore) DeleteService(projectName, appName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", projectName, appName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService
func (mr *MockApplicationStoreMockRecorder) DeleteService(projectName, appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockApplicationStore)(nil).DeleteService), projectName, appName)
}

// MockApplicationLister is a mock of ApplicationLister interface
type MockApplicationLister struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationListerMockRecorder
}

// MockApplicationListerMockRecorder is the mock recorder for MockApplicationLister
type MockApplicationListerMockRecorder struct {
	mock *MockApplicationLister
}

// NewMockApplicationLister creates a new mock instance
func NewMockApplicationLister(ctrl *gomock.Controller) *MockApplicationLister {
	mock := &MockApplicationLister{ctrl: ctrl}
	mock.recorder = &MockApplicationListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationLister) EXPECT() *MockApplicationListerMockRecorder {
	return m.recorder
}

// ListServices mocks base method
func (m *MockApplicationLister) ListServices(projectName string) ([]*archer.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", projectName)
	ret0, _ := ret[0].([]*archer.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockApplicationListerMockRecorder) ListServices(projectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockApplicationLister)(nil).ListServices), projectName)
}

// MockApplicationGetter is a mock of ApplicationGetter interface
type MockApplicationGetter struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationGetterMockRecorder
}

// MockApplicationGetterMockRecorder is the mock recorder for MockApplicationGetter
type MockApplicationGetterMockRecorder struct {
	mock *MockApplicationGetter
}

// NewMockApplicationGetter creates a new mock instance
func NewMockApplicationGetter(ctrl *gomock.Controller) *MockApplicationGetter {
	mock := &MockApplicationGetter{ctrl: ctrl}
	mock.recorder = &MockApplicationGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationGetter) EXPECT() *MockApplicationGetterMockRecorder {
	return m.recorder
}

// GetService mocks base method
func (m *MockApplicationGetter) GetService(projectName, applicationName string) (*archer.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", projectName, applicationName)
	ret0, _ := ret[0].(*archer.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockApplicationGetterMockRecorder) GetService(projectName, applicationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockApplicationGetter)(nil).GetService), projectName, applicationName)
}

// MockApplicationCreator is a mock of ApplicationCreator interface
type MockApplicationCreator struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationCreatorMockRecorder
}

// MockApplicationCreatorMockRecorder is the mock recorder for MockApplicationCreator
type MockApplicationCreatorMockRecorder struct {
	mock *MockApplicationCreator
}

// NewMockApplicationCreator creates a new mock instance
func NewMockApplicationCreator(ctrl *gomock.Controller) *MockApplicationCreator {
	mock := &MockApplicationCreator{ctrl: ctrl}
	mock.recorder = &MockApplicationCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationCreator) EXPECT() *MockApplicationCreatorMockRecorder {
	return m.recorder
}

// CreateService mocks base method
func (m *MockApplicationCreator) CreateService(app *archer.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", app)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateService indicates an expected call of CreateService
func (mr *MockApplicationCreatorMockRecorder) CreateService(app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockApplicationCreator)(nil).CreateService), app)
}

// MockApplicationDeleter is a mock of ApplicationDeleter interface
type MockApplicationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationDeleterMockRecorder
}

// MockApplicationDeleterMockRecorder is the mock recorder for MockApplicationDeleter
type MockApplicationDeleterMockRecorder struct {
	mock *MockApplicationDeleter
}

// NewMockApplicationDeleter creates a new mock instance
func NewMockApplicationDeleter(ctrl *gomock.Controller) *MockApplicationDeleter {
	mock := &MockApplicationDeleter{ctrl: ctrl}
	mock.recorder = &MockApplicationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationDeleter) EXPECT() *MockApplicationDeleterMockRecorder {
	return m.recorder
}

// DeleteService mocks base method
func (m *MockApplicationDeleter) DeleteService(projectName, appName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", projectName, appName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService
func (mr *MockApplicationDeleterMockRecorder) DeleteService(projectName, appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockApplicationDeleter)(nil).DeleteService), projectName, appName)
}
