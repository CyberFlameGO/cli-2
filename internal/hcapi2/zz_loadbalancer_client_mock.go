// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hetznercloud/cli/internal/hcapi2 (interfaces: LoadBalancerClient)

// Package hcapi2 is a generated GoMock package.
package hcapi2

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hcloud "github.com/hetznercloud/hcloud-go/hcloud"
)

// MockLoadBalancerClient is a mock of LoadBalancerClient interface.
type MockLoadBalancerClient struct {
	ctrl     *gomock.Controller
	recorder *MockLoadBalancerClientMockRecorder
}

// MockLoadBalancerClientMockRecorder is the mock recorder for MockLoadBalancerClient.
type MockLoadBalancerClientMockRecorder struct {
	mock *MockLoadBalancerClient
}

// NewMockLoadBalancerClient creates a new mock instance.
func NewMockLoadBalancerClient(ctrl *gomock.Controller) *MockLoadBalancerClient {
	mock := &MockLoadBalancerClient{ctrl: ctrl}
	mock.recorder = &MockLoadBalancerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoadBalancerClient) EXPECT() *MockLoadBalancerClientMockRecorder {
	return m.recorder
}

// AddIPTarget mocks base method.
func (m *MockLoadBalancerClient) AddIPTarget(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 hcloud.LoadBalancerAddIPTargetOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddIPTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddIPTarget indicates an expected call of AddIPTarget.
func (mr *MockLoadBalancerClientMockRecorder) AddIPTarget(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIPTarget", reflect.TypeOf((*MockLoadBalancerClient)(nil).AddIPTarget), arg0, arg1, arg2)
}

// AddLabelSelectorTarget mocks base method.
func (m *MockLoadBalancerClient) AddLabelSelectorTarget(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 hcloud.LoadBalancerAddLabelSelectorTargetOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLabelSelectorTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddLabelSelectorTarget indicates an expected call of AddLabelSelectorTarget.
func (mr *MockLoadBalancerClientMockRecorder) AddLabelSelectorTarget(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLabelSelectorTarget", reflect.TypeOf((*MockLoadBalancerClient)(nil).AddLabelSelectorTarget), arg0, arg1, arg2)
}

// AddServerTarget mocks base method.
func (m *MockLoadBalancerClient) AddServerTarget(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 hcloud.LoadBalancerAddServerTargetOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddServerTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddServerTarget indicates an expected call of AddServerTarget.
func (mr *MockLoadBalancerClientMockRecorder) AddServerTarget(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddServerTarget", reflect.TypeOf((*MockLoadBalancerClient)(nil).AddServerTarget), arg0, arg1, arg2)
}

// AddService mocks base method.
func (m *MockLoadBalancerClient) AddService(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 hcloud.LoadBalancerAddServiceOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddService", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddService indicates an expected call of AddService.
func (mr *MockLoadBalancerClientMockRecorder) AddService(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddService", reflect.TypeOf((*MockLoadBalancerClient)(nil).AddService), arg0, arg1, arg2)
}

// All mocks base method.
func (m *MockLoadBalancerClient) All(arg0 context.Context) ([]*hcloud.LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", arg0)
	ret0, _ := ret[0].([]*hcloud.LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockLoadBalancerClientMockRecorder) All(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockLoadBalancerClient)(nil).All), arg0)
}

// AllWithOpts mocks base method.
func (m *MockLoadBalancerClient) AllWithOpts(arg0 context.Context, arg1 hcloud.LoadBalancerListOpts) ([]*hcloud.LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllWithOpts", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllWithOpts indicates an expected call of AllWithOpts.
func (mr *MockLoadBalancerClientMockRecorder) AllWithOpts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllWithOpts", reflect.TypeOf((*MockLoadBalancerClient)(nil).AllWithOpts), arg0, arg1)
}

// AttachToNetwork mocks base method.
func (m *MockLoadBalancerClient) AttachToNetwork(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 hcloud.LoadBalancerAttachToNetworkOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachToNetwork", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AttachToNetwork indicates an expected call of AttachToNetwork.
func (mr *MockLoadBalancerClientMockRecorder) AttachToNetwork(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachToNetwork", reflect.TypeOf((*MockLoadBalancerClient)(nil).AttachToNetwork), arg0, arg1, arg2)
}

// ChangeAlgorithm mocks base method.
func (m *MockLoadBalancerClient) ChangeAlgorithm(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 hcloud.LoadBalancerChangeAlgorithmOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeAlgorithm", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeAlgorithm indicates an expected call of ChangeAlgorithm.
func (mr *MockLoadBalancerClientMockRecorder) ChangeAlgorithm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeAlgorithm", reflect.TypeOf((*MockLoadBalancerClient)(nil).ChangeAlgorithm), arg0, arg1, arg2)
}

// ChangeDNSPtr mocks base method.
func (m *MockLoadBalancerClient) ChangeDNSPtr(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 string, arg3 *string) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeDNSPtr", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeDNSPtr indicates an expected call of ChangeDNSPtr.
func (mr *MockLoadBalancerClientMockRecorder) ChangeDNSPtr(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeDNSPtr", reflect.TypeOf((*MockLoadBalancerClient)(nil).ChangeDNSPtr), arg0, arg1, arg2, arg3)
}

// ChangeProtection mocks base method.
func (m *MockLoadBalancerClient) ChangeProtection(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 hcloud.LoadBalancerChangeProtectionOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeProtection", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeProtection indicates an expected call of ChangeProtection.
func (mr *MockLoadBalancerClientMockRecorder) ChangeProtection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeProtection", reflect.TypeOf((*MockLoadBalancerClient)(nil).ChangeProtection), arg0, arg1, arg2)
}

// ChangeType mocks base method.
func (m *MockLoadBalancerClient) ChangeType(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 hcloud.LoadBalancerChangeTypeOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeType", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeType indicates an expected call of ChangeType.
func (mr *MockLoadBalancerClientMockRecorder) ChangeType(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeType", reflect.TypeOf((*MockLoadBalancerClient)(nil).ChangeType), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockLoadBalancerClient) Create(arg0 context.Context, arg1 hcloud.LoadBalancerCreateOpts) (hcloud.LoadBalancerCreateResult, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(hcloud.LoadBalancerCreateResult)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockLoadBalancerClientMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLoadBalancerClient)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockLoadBalancerClient) Delete(arg0 context.Context, arg1 *hcloud.LoadBalancer) (*hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockLoadBalancerClientMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLoadBalancerClient)(nil).Delete), arg0, arg1)
}

// DeleteService mocks base method.
func (m *MockLoadBalancerClient) DeleteService(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 int) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockLoadBalancerClientMockRecorder) DeleteService(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockLoadBalancerClient)(nil).DeleteService), arg0, arg1, arg2)
}

// DetachFromNetwork mocks base method.
func (m *MockLoadBalancerClient) DetachFromNetwork(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 hcloud.LoadBalancerDetachFromNetworkOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachFromNetwork", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DetachFromNetwork indicates an expected call of DetachFromNetwork.
func (mr *MockLoadBalancerClientMockRecorder) DetachFromNetwork(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachFromNetwork", reflect.TypeOf((*MockLoadBalancerClient)(nil).DetachFromNetwork), arg0, arg1, arg2)
}

// DisablePublicInterface mocks base method.
func (m *MockLoadBalancerClient) DisablePublicInterface(arg0 context.Context, arg1 *hcloud.LoadBalancer) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisablePublicInterface", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DisablePublicInterface indicates an expected call of DisablePublicInterface.
func (mr *MockLoadBalancerClientMockRecorder) DisablePublicInterface(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisablePublicInterface", reflect.TypeOf((*MockLoadBalancerClient)(nil).DisablePublicInterface), arg0, arg1)
}

// EnablePublicInterface mocks base method.
func (m *MockLoadBalancerClient) EnablePublicInterface(arg0 context.Context, arg1 *hcloud.LoadBalancer) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnablePublicInterface", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EnablePublicInterface indicates an expected call of EnablePublicInterface.
func (mr *MockLoadBalancerClientMockRecorder) EnablePublicInterface(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePublicInterface", reflect.TypeOf((*MockLoadBalancerClient)(nil).EnablePublicInterface), arg0, arg1)
}

// Get mocks base method.
func (m *MockLoadBalancerClient) Get(arg0 context.Context, arg1 string) (*hcloud.LoadBalancer, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.LoadBalancer)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockLoadBalancerClientMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLoadBalancerClient)(nil).Get), arg0, arg1)
}

// GetByID mocks base method.
func (m *MockLoadBalancerClient) GetByID(arg0 context.Context, arg1 int) (*hcloud.LoadBalancer, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.LoadBalancer)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockLoadBalancerClientMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockLoadBalancerClient)(nil).GetByID), arg0, arg1)
}

// GetByName mocks base method.
func (m *MockLoadBalancerClient) GetByName(arg0 context.Context, arg1 string) (*hcloud.LoadBalancer, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.LoadBalancer)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByName indicates an expected call of GetByName.
func (mr *MockLoadBalancerClientMockRecorder) GetByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockLoadBalancerClient)(nil).GetByName), arg0, arg1)
}

// GetMetrics mocks base method.
func (m *MockLoadBalancerClient) GetMetrics(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 hcloud.LoadBalancerGetMetricsOpts) (*hcloud.LoadBalancerMetrics, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetrics", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.LoadBalancerMetrics)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMetrics indicates an expected call of GetMetrics.
func (mr *MockLoadBalancerClientMockRecorder) GetMetrics(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetrics", reflect.TypeOf((*MockLoadBalancerClient)(nil).GetMetrics), arg0, arg1, arg2)
}

// LabelKeys mocks base method.
func (m *MockLoadBalancerClient) LabelKeys(arg0 string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LabelKeys", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// LabelKeys indicates an expected call of LabelKeys.
func (mr *MockLoadBalancerClientMockRecorder) LabelKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LabelKeys", reflect.TypeOf((*MockLoadBalancerClient)(nil).LabelKeys), arg0)
}

// List mocks base method.
func (m *MockLoadBalancerClient) List(arg0 context.Context, arg1 hcloud.LoadBalancerListOpts) ([]*hcloud.LoadBalancer, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.LoadBalancer)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockLoadBalancerClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLoadBalancerClient)(nil).List), arg0, arg1)
}

// LoadBalancerName mocks base method.
func (m *MockLoadBalancerClient) LoadBalancerName(arg0 int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadBalancerName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// LoadBalancerName indicates an expected call of LoadBalancerName.
func (mr *MockLoadBalancerClientMockRecorder) LoadBalancerName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadBalancerName", reflect.TypeOf((*MockLoadBalancerClient)(nil).LoadBalancerName), arg0)
}

// Names mocks base method.
func (m *MockLoadBalancerClient) Names() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Names")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Names indicates an expected call of Names.
func (mr *MockLoadBalancerClientMockRecorder) Names() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Names", reflect.TypeOf((*MockLoadBalancerClient)(nil).Names))
}

// RemoveIPTarget mocks base method.
func (m *MockLoadBalancerClient) RemoveIPTarget(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 net.IP) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveIPTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RemoveIPTarget indicates an expected call of RemoveIPTarget.
func (mr *MockLoadBalancerClientMockRecorder) RemoveIPTarget(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIPTarget", reflect.TypeOf((*MockLoadBalancerClient)(nil).RemoveIPTarget), arg0, arg1, arg2)
}

// RemoveLabelSelectorTarget mocks base method.
func (m *MockLoadBalancerClient) RemoveLabelSelectorTarget(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 string) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLabelSelectorTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RemoveLabelSelectorTarget indicates an expected call of RemoveLabelSelectorTarget.
func (mr *MockLoadBalancerClientMockRecorder) RemoveLabelSelectorTarget(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLabelSelectorTarget", reflect.TypeOf((*MockLoadBalancerClient)(nil).RemoveLabelSelectorTarget), arg0, arg1, arg2)
}

// RemoveServerTarget mocks base method.
func (m *MockLoadBalancerClient) RemoveServerTarget(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 *hcloud.Server) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveServerTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RemoveServerTarget indicates an expected call of RemoveServerTarget.
func (mr *MockLoadBalancerClientMockRecorder) RemoveServerTarget(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveServerTarget", reflect.TypeOf((*MockLoadBalancerClient)(nil).RemoveServerTarget), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockLoadBalancerClient) Update(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 hcloud.LoadBalancerUpdateOpts) (*hcloud.LoadBalancer, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*hcloud.LoadBalancer)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockLoadBalancerClientMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLoadBalancerClient)(nil).Update), arg0, arg1, arg2)
}

// UpdateService mocks base method.
func (m *MockLoadBalancerClient) UpdateService(arg0 context.Context, arg1 *hcloud.LoadBalancer, arg2 int, arg3 hcloud.LoadBalancerUpdateServiceOpts) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockLoadBalancerClientMockRecorder) UpdateService(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockLoadBalancerClient)(nil).UpdateService), arg0, arg1, arg2, arg3)
}
