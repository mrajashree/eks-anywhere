// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/support/interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	supportbundle "github.com/aws/eks-anywhere/pkg/support"
	gomock "github.com/golang/mock/gomock"
	v1beta2 "github.com/replicatedhq/troubleshoot/pkg/apis/troubleshoot/v1beta2"
)

// MockDiagnosticBundle is a mock of DiagnosticBundle interface.
type MockDiagnosticBundle struct {
	ctrl     *gomock.Controller
	recorder *MockDiagnosticBundleMockRecorder
}

// MockDiagnosticBundleMockRecorder is the mock recorder for MockDiagnosticBundle.
type MockDiagnosticBundleMockRecorder struct {
	mock *MockDiagnosticBundle
}

// NewMockDiagnosticBundle creates a new mock instance.
func NewMockDiagnosticBundle(ctrl *gomock.Controller) *MockDiagnosticBundle {
	mock := &MockDiagnosticBundle{ctrl: ctrl}
	mock.recorder = &MockDiagnosticBundleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiagnosticBundle) EXPECT() *MockDiagnosticBundleMockRecorder {
	return m.recorder
}

// PrintBundleConfig mocks base method.
func (m *MockDiagnosticBundle) PrintBundleConfig() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrintBundleConfig")
	ret0, _ := ret[0].(error)
	return ret0
}

// PrintBundleConfig indicates an expected call of PrintBundleConfig.
func (mr *MockDiagnosticBundleMockRecorder) PrintBundleConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintBundleConfig", reflect.TypeOf((*MockDiagnosticBundle)(nil).PrintBundleConfig))
}

// WithDatacenterConfig mocks base method.
func (m *MockDiagnosticBundle) WithDatacenterConfig(config v1alpha1.Ref) *supportbundle.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDatacenterConfig", config)
	ret0, _ := ret[0].(*supportbundle.EksaDiagnosticBundle)
	return ret0
}

// WithDatacenterConfig indicates an expected call of WithDatacenterConfig.
func (mr *MockDiagnosticBundleMockRecorder) WithDatacenterConfig(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDatacenterConfig", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithDatacenterConfig), config)
}

// WithDefaultAnalyzers mocks base method.
func (m *MockDiagnosticBundle) WithDefaultAnalyzers() *supportbundle.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDefaultAnalyzers")
	ret0, _ := ret[0].(*supportbundle.EksaDiagnosticBundle)
	return ret0
}

// WithDefaultAnalyzers indicates an expected call of WithDefaultAnalyzers.
func (mr *MockDiagnosticBundleMockRecorder) WithDefaultAnalyzers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDefaultAnalyzers", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithDefaultAnalyzers))
}

// WithDefaultCollectors mocks base method.
func (m *MockDiagnosticBundle) WithDefaultCollectors() *supportbundle.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDefaultCollectors")
	ret0, _ := ret[0].(*supportbundle.EksaDiagnosticBundle)
	return ret0
}

// WithDefaultCollectors indicates an expected call of WithDefaultCollectors.
func (mr *MockDiagnosticBundleMockRecorder) WithDefaultCollectors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDefaultCollectors", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithDefaultCollectors))
}

// WithExternalEtcd mocks base method.
func (m *MockDiagnosticBundle) WithExternalEtcd(config *v1alpha1.ExternalEtcdConfiguration) *supportbundle.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithExternalEtcd", config)
	ret0, _ := ret[0].(*supportbundle.EksaDiagnosticBundle)
	return ret0
}

// WithExternalEtcd indicates an expected call of WithExternalEtcd.
func (mr *MockDiagnosticBundleMockRecorder) WithExternalEtcd(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithExternalEtcd", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithExternalEtcd), config)
}

// WithGitOpsConfig mocks base method.
func (m *MockDiagnosticBundle) WithGitOpsConfig(config *v1alpha1.GitOpsConfig) *supportbundle.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithGitOpsConfig", config)
	ret0, _ := ret[0].(*supportbundle.EksaDiagnosticBundle)
	return ret0
}

// WithGitOpsConfig indicates an expected call of WithGitOpsConfig.
func (mr *MockDiagnosticBundleMockRecorder) WithGitOpsConfig(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithGitOpsConfig", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithGitOpsConfig), config)
}

// WithOidcConfig mocks base method.
func (m *MockDiagnosticBundle) WithOidcConfig(config *v1alpha1.OIDCConfig) *supportbundle.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithOidcConfig", config)
	ret0, _ := ret[0].(*supportbundle.EksaDiagnosticBundle)
	return ret0
}

// WithOidcConfig indicates an expected call of WithOidcConfig.
func (mr *MockDiagnosticBundleMockRecorder) WithOidcConfig(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithOidcConfig", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithOidcConfig), config)
}

// MockAnalyzerFactory is a mock of AnalyzerFactory interface.
type MockAnalyzerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyzerFactoryMockRecorder
}

// MockAnalyzerFactoryMockRecorder is the mock recorder for MockAnalyzerFactory.
type MockAnalyzerFactoryMockRecorder struct {
	mock *MockAnalyzerFactory
}

// NewMockAnalyzerFactory creates a new mock instance.
func NewMockAnalyzerFactory(ctrl *gomock.Controller) *MockAnalyzerFactory {
	mock := &MockAnalyzerFactory{ctrl: ctrl}
	mock.recorder = &MockAnalyzerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnalyzerFactory) EXPECT() *MockAnalyzerFactoryMockRecorder {
	return m.recorder
}

// DataCenterConfigAnalyzers mocks base method.
func (m *MockAnalyzerFactory) DataCenterConfigAnalyzers(datacenter v1alpha1.Ref) []*v1beta2.Analyze {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataCenterConfigAnalyzers", datacenter)
	ret0, _ := ret[0].([]*v1beta2.Analyze)
	return ret0
}

// DataCenterConfigAnalyzers indicates an expected call of DataCenterConfigAnalyzers.
func (mr *MockAnalyzerFactoryMockRecorder) DataCenterConfigAnalyzers(datacenter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataCenterConfigAnalyzers", reflect.TypeOf((*MockAnalyzerFactory)(nil).DataCenterConfigAnalyzers), datacenter)
}

// DefaultAnalyzers mocks base method.
func (m *MockAnalyzerFactory) DefaultAnalyzers() []*v1beta2.Analyze {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultAnalyzers")
	ret0, _ := ret[0].([]*v1beta2.Analyze)
	return ret0
}

// DefaultAnalyzers indicates an expected call of DefaultAnalyzers.
func (mr *MockAnalyzerFactoryMockRecorder) DefaultAnalyzers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultAnalyzers", reflect.TypeOf((*MockAnalyzerFactory)(nil).DefaultAnalyzers))
}

// EksaExternalEtcdAnalyzers mocks base method.
func (m *MockAnalyzerFactory) EksaExternalEtcdAnalyzers() []*v1beta2.Analyze {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EksaExternalEtcdAnalyzers")
	ret0, _ := ret[0].([]*v1beta2.Analyze)
	return ret0
}

// EksaExternalEtcdAnalyzers indicates an expected call of EksaExternalEtcdAnalyzers.
func (mr *MockAnalyzerFactoryMockRecorder) EksaExternalEtcdAnalyzers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EksaExternalEtcdAnalyzers", reflect.TypeOf((*MockAnalyzerFactory)(nil).EksaExternalEtcdAnalyzers))
}

// EksaGitopsAnalyzers mocks base method.
func (m *MockAnalyzerFactory) EksaGitopsAnalyzers() []*v1beta2.Analyze {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EksaGitopsAnalyzers")
	ret0, _ := ret[0].([]*v1beta2.Analyze)
	return ret0
}

// EksaGitopsAnalyzers indicates an expected call of EksaGitopsAnalyzers.
func (mr *MockAnalyzerFactoryMockRecorder) EksaGitopsAnalyzers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EksaGitopsAnalyzers", reflect.TypeOf((*MockAnalyzerFactory)(nil).EksaGitopsAnalyzers))
}

// EksaOidcAnalyzers mocks base method.
func (m *MockAnalyzerFactory) EksaOidcAnalyzers() []*v1beta2.Analyze {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EksaOidcAnalyzers")
	ret0, _ := ret[0].([]*v1beta2.Analyze)
	return ret0
}

// EksaOidcAnalyzers indicates an expected call of EksaOidcAnalyzers.
func (mr *MockAnalyzerFactoryMockRecorder) EksaOidcAnalyzers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EksaOidcAnalyzers", reflect.TypeOf((*MockAnalyzerFactory)(nil).EksaOidcAnalyzers))
}

// MockCollectorFactory is a mock of CollectorFactory interface.
type MockCollectorFactory struct {
	ctrl     *gomock.Controller
	recorder *MockCollectorFactoryMockRecorder
}

// MockCollectorFactoryMockRecorder is the mock recorder for MockCollectorFactory.
type MockCollectorFactoryMockRecorder struct {
	mock *MockCollectorFactory
}

// NewMockCollectorFactory creates a new mock instance.
func NewMockCollectorFactory(ctrl *gomock.Controller) *MockCollectorFactory {
	mock := &MockCollectorFactory{ctrl: ctrl}
	mock.recorder = &MockCollectorFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCollectorFactory) EXPECT() *MockCollectorFactoryMockRecorder {
	return m.recorder
}

// DefaultCollectors mocks base method.
func (m *MockCollectorFactory) DefaultCollectors() []*v1beta2.Collect {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultCollectors")
	ret0, _ := ret[0].([]*v1beta2.Collect)
	return ret0
}

// DefaultCollectors indicates an expected call of DefaultCollectors.
func (mr *MockCollectorFactoryMockRecorder) DefaultCollectors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultCollectors", reflect.TypeOf((*MockCollectorFactory)(nil).DefaultCollectors))
}
