// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1 (interfaces: BeaconServiceClient,BeaconService_LatestCrystallizedStateClient,BeaconService_LatestAttestationClient)

package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	v10 "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockBeaconServiceClient is a mock of BeaconServiceClient interface
type MockBeaconServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconServiceClientMockRecorder
}

// MockBeaconServiceClientMockRecorder is the mock recorder for MockBeaconServiceClient
type MockBeaconServiceClientMockRecorder struct {
	mock *MockBeaconServiceClient
}

// NewMockBeaconServiceClient creates a new mock instance
func NewMockBeaconServiceClient(ctrl *gomock.Controller) *MockBeaconServiceClient {
	mock := &MockBeaconServiceClient{ctrl: ctrl}
	mock.recorder = &MockBeaconServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconServiceClient) EXPECT() *MockBeaconServiceClientMockRecorder {
	return m.recorder
}

// CanonicalHead mocks base method
func (m *MockBeaconServiceClient) CanonicalHead(arg0 context.Context, arg1 *empty.Empty, arg2 ...grpc.CallOption) (*v1.BeaconBlock, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CanonicalHead", varargs...)
	ret0, _ := ret[0].(*v1.BeaconBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanonicalHead indicates an expected call of CanonicalHead
func (mr *MockBeaconServiceClientMockRecorder) CanonicalHead(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanonicalHead", reflect.TypeOf((*MockBeaconServiceClient)(nil).CanonicalHead), varargs...)
}

// GenesisTimeAndCanonicalState mocks base method
func (m *MockBeaconServiceClient) GenesisTimeAndCanonicalState(arg0 context.Context, arg1 *empty.Empty, arg2 ...grpc.CallOption) (*v10.GenesisTimeAndStateResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GenesisTimeAndCanonicalState", varargs...)
	ret0, _ := ret[0].(*v10.GenesisTimeAndStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenesisTimeAndCanonicalState indicates an expected call of GenesisTimeAndCanonicalState
func (mr *MockBeaconServiceClientMockRecorder) GenesisTimeAndCanonicalState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenesisTimeAndCanonicalState", reflect.TypeOf((*MockBeaconServiceClient)(nil).GenesisTimeAndCanonicalState), varargs...)
}

// LatestAttestation mocks base method
func (m *MockBeaconServiceClient) LatestAttestation(arg0 context.Context, arg1 *empty.Empty, arg2 ...grpc.CallOption) (v10.BeaconService_LatestAttestationClient, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LatestAttestation", varargs...)
	ret0, _ := ret[0].(v10.BeaconService_LatestAttestationClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestAttestation indicates an expected call of LatestAttestation
func (mr *MockBeaconServiceClientMockRecorder) LatestAttestation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestAttestation", reflect.TypeOf((*MockBeaconServiceClient)(nil).LatestAttestation), varargs...)
}

// LatestCrystallizedState mocks base method
func (m *MockBeaconServiceClient) LatestCrystallizedState(arg0 context.Context, arg1 *empty.Empty, arg2 ...grpc.CallOption) (v10.BeaconService_LatestCrystallizedStateClient, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LatestCrystallizedState", varargs...)
	ret0, _ := ret[0].(v10.BeaconService_LatestCrystallizedStateClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestCrystallizedState indicates an expected call of LatestCrystallizedState
func (mr *MockBeaconServiceClientMockRecorder) LatestCrystallizedState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestCrystallizedState", reflect.TypeOf((*MockBeaconServiceClient)(nil).LatestCrystallizedState), varargs...)
}

// MockBeaconService_LatestCrystallizedStateClient is a mock of BeaconService_LatestCrystallizedStateClient interface
type MockBeaconService_LatestCrystallizedStateClient struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconService_LatestCrystallizedStateClientMockRecorder
}

// MockBeaconService_LatestCrystallizedStateClientMockRecorder is the mock recorder for MockBeaconService_LatestCrystallizedStateClient
type MockBeaconService_LatestCrystallizedStateClientMockRecorder struct {
	mock *MockBeaconService_LatestCrystallizedStateClient
}

// NewMockBeaconService_LatestCrystallizedStateClient creates a new mock instance
func NewMockBeaconService_LatestCrystallizedStateClient(ctrl *gomock.Controller) *MockBeaconService_LatestCrystallizedStateClient {
	mock := &MockBeaconService_LatestCrystallizedStateClient{ctrl: ctrl}
	mock.recorder = &MockBeaconService_LatestCrystallizedStateClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconService_LatestCrystallizedStateClient) EXPECT() *MockBeaconService_LatestCrystallizedStateClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockBeaconService_LatestCrystallizedStateClient) CloseSend() error {
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockBeaconService_LatestCrystallizedStateClientMockRecorder) CloseSend() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockBeaconService_LatestCrystallizedStateClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockBeaconService_LatestCrystallizedStateClient) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBeaconService_LatestCrystallizedStateClientMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBeaconService_LatestCrystallizedStateClient)(nil).Context))
}

// Header mocks base method
func (m *MockBeaconService_LatestCrystallizedStateClient) Header() (metadata.MD, error) {
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockBeaconService_LatestCrystallizedStateClientMockRecorder) Header() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockBeaconService_LatestCrystallizedStateClient)(nil).Header))
}

// Recv mocks base method
func (m *MockBeaconService_LatestCrystallizedStateClient) Recv() (*v1.CrystallizedState, error) {
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.CrystallizedState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockBeaconService_LatestCrystallizedStateClientMockRecorder) Recv() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockBeaconService_LatestCrystallizedStateClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockBeaconService_LatestCrystallizedStateClient) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBeaconService_LatestCrystallizedStateClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBeaconService_LatestCrystallizedStateClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockBeaconService_LatestCrystallizedStateClient) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBeaconService_LatestCrystallizedStateClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBeaconService_LatestCrystallizedStateClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockBeaconService_LatestCrystallizedStateClient) Trailer() metadata.MD {
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockBeaconService_LatestCrystallizedStateClientMockRecorder) Trailer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockBeaconService_LatestCrystallizedStateClient)(nil).Trailer))
}

// MockBeaconService_LatestAttestationClient is a mock of BeaconService_LatestAttestationClient interface
type MockBeaconService_LatestAttestationClient struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconService_LatestAttestationClientMockRecorder
}

// MockBeaconService_LatestAttestationClientMockRecorder is the mock recorder for MockBeaconService_LatestAttestationClient
type MockBeaconService_LatestAttestationClientMockRecorder struct {
	mock *MockBeaconService_LatestAttestationClient
}

// NewMockBeaconService_LatestAttestationClient creates a new mock instance
func NewMockBeaconService_LatestAttestationClient(ctrl *gomock.Controller) *MockBeaconService_LatestAttestationClient {
	mock := &MockBeaconService_LatestAttestationClient{ctrl: ctrl}
	mock.recorder = &MockBeaconService_LatestAttestationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconService_LatestAttestationClient) EXPECT() *MockBeaconService_LatestAttestationClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockBeaconService_LatestAttestationClient) CloseSend() error {
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) CloseSend() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockBeaconService_LatestAttestationClient) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).Context))
}

// Header mocks base method
func (m *MockBeaconService_LatestAttestationClient) Header() (metadata.MD, error) {
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) Header() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).Header))
}

// Recv mocks base method
func (m *MockBeaconService_LatestAttestationClient) Recv() (*v1.AggregatedAttestation, error) {
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.AggregatedAttestation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) Recv() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockBeaconService_LatestAttestationClient) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockBeaconService_LatestAttestationClient) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockBeaconService_LatestAttestationClient) Trailer() metadata.MD {
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockBeaconService_LatestAttestationClientMockRecorder) Trailer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockBeaconService_LatestAttestationClient)(nil).Trailer))
}
