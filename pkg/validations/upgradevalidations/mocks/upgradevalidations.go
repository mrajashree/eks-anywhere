// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/validations/upgradevalidations/upgradevalidations.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	executables "github.com/aws/eks-anywhere/pkg/executables"
	types "github.com/aws/eks-anywhere/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// MockValidationsKubectlClient is a mock of ValidationsKubectlClient interface.
type MockValidationsKubectlClient struct {
	ctrl     *gomock.Controller
	recorder *MockValidationsKubectlClientMockRecorder
}

// MockValidationsKubectlClientMockRecorder is the mock recorder for MockValidationsKubectlClient.
type MockValidationsKubectlClientMockRecorder struct {
	mock *MockValidationsKubectlClient
}

// NewMockValidationsKubectlClient creates a new mock instance.
func NewMockValidationsKubectlClient(ctrl *gomock.Controller) *MockValidationsKubectlClient {
	mock := &MockValidationsKubectlClient{ctrl: ctrl}
	mock.recorder = &MockValidationsKubectlClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidationsKubectlClient) EXPECT() *MockValidationsKubectlClientMockRecorder {
	return m.recorder
}

// GetClusters mocks base method.
func (m *MockValidationsKubectlClient) GetClusters(ctx context.Context, cluster *types.Cluster) ([]types.CAPICluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusters", ctx, cluster)
	ret0, _ := ret[0].([]types.CAPICluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusters indicates an expected call of GetClusters.
func (mr *MockValidationsKubectlClientMockRecorder) GetClusters(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusters", reflect.TypeOf((*MockValidationsKubectlClient)(nil).GetClusters), ctx, cluster)
}

// GetEksaCluster mocks base method.
func (m *MockValidationsKubectlClient) GetEksaCluster(ctx context.Context, cluster *types.Cluster) (*v1alpha1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaCluster", ctx, cluster)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaCluster indicates an expected call of GetEksaCluster.
func (mr *MockValidationsKubectlClientMockRecorder) GetEksaCluster(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaCluster", reflect.TypeOf((*MockValidationsKubectlClient)(nil).GetEksaCluster), ctx, cluster)
}

// GetEksaGitOpsConfig mocks base method.
func (m *MockValidationsKubectlClient) GetEksaGitOpsConfig(ctx context.Context, gitOpsConfigName, kubeconfigFile, namespace string) (*v1alpha1.GitOpsConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaGitOpsConfig", ctx, gitOpsConfigName, kubeconfigFile, namespace)
	ret0, _ := ret[0].(*v1alpha1.GitOpsConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaGitOpsConfig indicates an expected call of GetEksaGitOpsConfig.
func (mr *MockValidationsKubectlClientMockRecorder) GetEksaGitOpsConfig(ctx, gitOpsConfigName, kubeconfigFile, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaGitOpsConfig", reflect.TypeOf((*MockValidationsKubectlClient)(nil).GetEksaGitOpsConfig), ctx, gitOpsConfigName, kubeconfigFile, namespace)
}

// GetEksaOIDCConfig mocks base method.
func (m *MockValidationsKubectlClient) GetEksaOIDCConfig(ctx context.Context, oidcConfigName, kubeconfigFile, namespace string) (*v1alpha1.OIDCConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaOIDCConfig", ctx, oidcConfigName, kubeconfigFile, namespace)
	ret0, _ := ret[0].(*v1alpha1.OIDCConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaOIDCConfig indicates an expected call of GetEksaOIDCConfig.
func (mr *MockValidationsKubectlClientMockRecorder) GetEksaOIDCConfig(ctx, oidcConfigName, kubeconfigFile, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaOIDCConfig", reflect.TypeOf((*MockValidationsKubectlClient)(nil).GetEksaOIDCConfig), ctx, oidcConfigName, kubeconfigFile, namespace)
}

// GetEksaVSphereDatacenterConfig mocks base method.
func (m *MockValidationsKubectlClient) GetEksaVSphereDatacenterConfig(ctx context.Context, vsphereDatacenterConfigName, kubeconfigFile, namespace string) (*v1alpha1.VSphereDatacenterConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaVSphereDatacenterConfig", ctx, vsphereDatacenterConfigName, kubeconfigFile, namespace)
	ret0, _ := ret[0].(*v1alpha1.VSphereDatacenterConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaVSphereDatacenterConfig indicates an expected call of GetEksaVSphereDatacenterConfig.
func (mr *MockValidationsKubectlClientMockRecorder) GetEksaVSphereDatacenterConfig(ctx, vsphereDatacenterConfigName, kubeconfigFile, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaVSphereDatacenterConfig", reflect.TypeOf((*MockValidationsKubectlClient)(nil).GetEksaVSphereDatacenterConfig), ctx, vsphereDatacenterConfigName, kubeconfigFile, namespace)
}

// ValidateClustersCRD mocks base method.
func (m *MockValidationsKubectlClient) ValidateClustersCRD(ctx context.Context, cluster *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateClustersCRD", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateClustersCRD indicates an expected call of ValidateClustersCRD.
func (mr *MockValidationsKubectlClientMockRecorder) ValidateClustersCRD(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateClustersCRD", reflect.TypeOf((*MockValidationsKubectlClient)(nil).ValidateClustersCRD), ctx, cluster)
}

// ValidateControlPlaneNodes mocks base method.
func (m *MockValidationsKubectlClient) ValidateControlPlaneNodes(ctx context.Context, cluster *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateControlPlaneNodes", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateControlPlaneNodes indicates an expected call of ValidateControlPlaneNodes.
func (mr *MockValidationsKubectlClientMockRecorder) ValidateControlPlaneNodes(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateControlPlaneNodes", reflect.TypeOf((*MockValidationsKubectlClient)(nil).ValidateControlPlaneNodes), ctx, cluster)
}

// ValidateNodes mocks base method.
func (m *MockValidationsKubectlClient) ValidateNodes(ctx context.Context, kubeconfig string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateNodes", ctx, kubeconfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateNodes indicates an expected call of ValidateNodes.
func (mr *MockValidationsKubectlClientMockRecorder) ValidateNodes(ctx, kubeconfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateNodes", reflect.TypeOf((*MockValidationsKubectlClient)(nil).ValidateNodes), ctx, kubeconfig)
}

// ValidateWorkerNodes mocks base method.
func (m *MockValidationsKubectlClient) ValidateWorkerNodes(ctx context.Context, cluster *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateWorkerNodes", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateWorkerNodes indicates an expected call of ValidateWorkerNodes.
func (mr *MockValidationsKubectlClientMockRecorder) ValidateWorkerNodes(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateWorkerNodes", reflect.TypeOf((*MockValidationsKubectlClient)(nil).ValidateWorkerNodes), ctx, cluster)
}

// Version mocks base method.
func (m *MockValidationsKubectlClient) Version(ctx context.Context, cluster *types.Cluster) (*executables.VersionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", ctx, cluster)
	ret0, _ := ret[0].(*executables.VersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockValidationsKubectlClientMockRecorder) Version(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockValidationsKubectlClient)(nil).Version), ctx, cluster)
}
