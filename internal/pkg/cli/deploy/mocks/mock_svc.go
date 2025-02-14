// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/cli/deploy/svc.go

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// Mockuploader is a mock of uploader interface.
type Mockuploader struct {
	ctrl     *gomock.Controller
	recorder *MockuploaderMockRecorder
}

// MockuploaderMockRecorder is the mock recorder for Mockuploader.
type MockuploaderMockRecorder struct {
	mock *Mockuploader
}

// NewMockuploader creates a new mock instance.
func NewMockuploader(ctrl *gomock.Controller) *Mockuploader {
	mock := &Mockuploader{ctrl: ctrl}
	mock.recorder = &MockuploaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockuploader) EXPECT() *MockuploaderMockRecorder {
	return m.recorder
}

// Upload mocks base method.
func (m *Mockuploader) Upload(bucket, key string, data io.Reader) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", bucket, key, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockuploaderMockRecorder) Upload(bucket, key, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*Mockuploader)(nil).Upload), bucket, key, data)
}

// MockversionGetter is a mock of versionGetter interface.
type MockversionGetter struct {
	ctrl     *gomock.Controller
	recorder *MockversionGetterMockRecorder
}

// MockversionGetterMockRecorder is the mock recorder for MockversionGetter.
type MockversionGetterMockRecorder struct {
	mock *MockversionGetter
}

// NewMockversionGetter creates a new mock instance.
func NewMockversionGetter(ctrl *gomock.Controller) *MockversionGetter {
	mock := &MockversionGetter{ctrl: ctrl}
	mock.recorder = &MockversionGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockversionGetter) EXPECT() *MockversionGetterMockRecorder {
	return m.recorder
}

// Version mocks base method.
func (m *MockversionGetter) Version() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockversionGetterMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockversionGetter)(nil).Version))
}

// MockserviceForceUpdater is a mock of serviceForceUpdater interface.
type MockserviceForceUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockserviceForceUpdaterMockRecorder
}

// MockserviceForceUpdaterMockRecorder is the mock recorder for MockserviceForceUpdater.
type MockserviceForceUpdaterMockRecorder struct {
	mock *MockserviceForceUpdater
}

// NewMockserviceForceUpdater creates a new mock instance.
func NewMockserviceForceUpdater(ctrl *gomock.Controller) *MockserviceForceUpdater {
	mock := &MockserviceForceUpdater{ctrl: ctrl}
	mock.recorder = &MockserviceForceUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockserviceForceUpdater) EXPECT() *MockserviceForceUpdaterMockRecorder {
	return m.recorder
}

// ForceUpdateService mocks base method.
func (m *MockserviceForceUpdater) ForceUpdateService(app, env, svc string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceUpdateService", app, env, svc)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForceUpdateService indicates an expected call of ForceUpdateService.
func (mr *MockserviceForceUpdaterMockRecorder) ForceUpdateService(app, env, svc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceUpdateService", reflect.TypeOf((*MockserviceForceUpdater)(nil).ForceUpdateService), app, env, svc)
}

// LastUpdatedAt mocks base method.
func (m *MockserviceForceUpdater) LastUpdatedAt(app, env, svc string) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastUpdatedAt", app, env, svc)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastUpdatedAt indicates an expected call of LastUpdatedAt.
func (mr *MockserviceForceUpdaterMockRecorder) LastUpdatedAt(app, env, svc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastUpdatedAt", reflect.TypeOf((*MockserviceForceUpdater)(nil).LastUpdatedAt), app, env, svc)
}

// MockaliasCertValidator is a mock of aliasCertValidator interface.
type MockaliasCertValidator struct {
	ctrl     *gomock.Controller
	recorder *MockaliasCertValidatorMockRecorder
}

// MockaliasCertValidatorMockRecorder is the mock recorder for MockaliasCertValidator.
type MockaliasCertValidatorMockRecorder struct {
	mock *MockaliasCertValidator
}

// NewMockaliasCertValidator creates a new mock instance.
func NewMockaliasCertValidator(ctrl *gomock.Controller) *MockaliasCertValidator {
	mock := &MockaliasCertValidator{ctrl: ctrl}
	mock.recorder = &MockaliasCertValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockaliasCertValidator) EXPECT() *MockaliasCertValidatorMockRecorder {
	return m.recorder
}

// ValidateCertAliases mocks base method.
func (m *MockaliasCertValidator) ValidateCertAliases(aliases, certs []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateCertAliases", aliases, certs)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateCertAliases indicates an expected call of ValidateCertAliases.
func (mr *MockaliasCertValidatorMockRecorder) ValidateCertAliases(aliases, certs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateCertAliases", reflect.TypeOf((*MockaliasCertValidator)(nil).ValidateCertAliases), aliases, certs)
}
