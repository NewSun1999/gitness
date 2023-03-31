// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/harness/gitness/internal/store (interfaces: PrincipalStore,SpaceStore,RepoStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/harness/gitness/types"
	enum "github.com/harness/gitness/types/enum"
)

// MockPrincipalStore is a mock of PrincipalStore interface.
type MockPrincipalStore struct {
	ctrl     *gomock.Controller
	recorder *MockPrincipalStoreMockRecorder
}

// MockPrincipalStoreMockRecorder is the mock recorder for MockPrincipalStore.
type MockPrincipalStoreMockRecorder struct {
	mock *MockPrincipalStore
}

// NewMockPrincipalStore creates a new mock instance.
func NewMockPrincipalStore(ctrl *gomock.Controller) *MockPrincipalStore {
	mock := &MockPrincipalStore{ctrl: ctrl}
	mock.recorder = &MockPrincipalStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrincipalStore) EXPECT() *MockPrincipalStoreMockRecorder {
	return m.recorder
}

// CountServiceAccounts mocks base method.
func (m *MockPrincipalStore) CountServiceAccounts(arg0 context.Context, arg1 enum.ParentResourceType, arg2 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountServiceAccounts", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountServiceAccounts indicates an expected call of CountServiceAccounts.
func (mr *MockPrincipalStoreMockRecorder) CountServiceAccounts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountServiceAccounts", reflect.TypeOf((*MockPrincipalStore)(nil).CountServiceAccounts), arg0, arg1, arg2)
}

// CountServices mocks base method.
func (m *MockPrincipalStore) CountServices(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountServices", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountServices indicates an expected call of CountServices.
func (mr *MockPrincipalStoreMockRecorder) CountServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountServices", reflect.TypeOf((*MockPrincipalStore)(nil).CountServices), arg0)
}

// CountUsers mocks base method.
func (m *MockPrincipalStore) CountUsers(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUsers", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUsers indicates an expected call of CountUsers.
func (mr *MockPrincipalStoreMockRecorder) CountUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUsers", reflect.TypeOf((*MockPrincipalStore)(nil).CountUsers), arg0)
}

// CreateService mocks base method.
func (m *MockPrincipalStore) CreateService(arg0 context.Context, arg1 *types.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateService indicates an expected call of CreateService.
func (mr *MockPrincipalStoreMockRecorder) CreateService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockPrincipalStore)(nil).CreateService), arg0, arg1)
}

// CreateServiceAccount mocks base method.
func (m *MockPrincipalStore) CreateServiceAccount(arg0 context.Context, arg1 *types.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServiceAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateServiceAccount indicates an expected call of CreateServiceAccount.
func (mr *MockPrincipalStoreMockRecorder) CreateServiceAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServiceAccount", reflect.TypeOf((*MockPrincipalStore)(nil).CreateServiceAccount), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockPrincipalStore) CreateUser(arg0 context.Context, arg1 *types.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockPrincipalStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockPrincipalStore)(nil).CreateUser), arg0, arg1)
}

// DeleteService mocks base method.
func (m *MockPrincipalStore) DeleteService(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockPrincipalStoreMockRecorder) DeleteService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockPrincipalStore)(nil).DeleteService), arg0, arg1)
}

// DeleteServiceAccount mocks base method.
func (m *MockPrincipalStore) DeleteServiceAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServiceAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServiceAccount indicates an expected call of DeleteServiceAccount.
func (mr *MockPrincipalStoreMockRecorder) DeleteServiceAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServiceAccount", reflect.TypeOf((*MockPrincipalStore)(nil).DeleteServiceAccount), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockPrincipalStore) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockPrincipalStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockPrincipalStore)(nil).DeleteUser), arg0, arg1)
}

// Find mocks base method.
func (m *MockPrincipalStore) Find(arg0 context.Context, arg1 int64) (*types.Principal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*types.Principal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockPrincipalStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPrincipalStore)(nil).Find), arg0, arg1)
}

// FindByEmail mocks base method.
func (m *MockPrincipalStore) FindByEmail(arg0 context.Context, arg1 string) (*types.Principal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", arg0, arg1)
	ret0, _ := ret[0].(*types.Principal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockPrincipalStoreMockRecorder) FindByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockPrincipalStore)(nil).FindByEmail), arg0, arg1)
}

// FindByUID mocks base method.
func (m *MockPrincipalStore) FindByUID(arg0 context.Context, arg1 string) (*types.Principal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUID", arg0, arg1)
	ret0, _ := ret[0].(*types.Principal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUID indicates an expected call of FindByUID.
func (mr *MockPrincipalStoreMockRecorder) FindByUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUID", reflect.TypeOf((*MockPrincipalStore)(nil).FindByUID), arg0, arg1)
}

// FindService mocks base method.
func (m *MockPrincipalStore) FindService(arg0 context.Context, arg1 int64) (*types.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindService", arg0, arg1)
	ret0, _ := ret[0].(*types.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindService indicates an expected call of FindService.
func (mr *MockPrincipalStoreMockRecorder) FindService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindService", reflect.TypeOf((*MockPrincipalStore)(nil).FindService), arg0, arg1)
}

// FindServiceAccount mocks base method.
func (m *MockPrincipalStore) FindServiceAccount(arg0 context.Context, arg1 int64) (*types.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindServiceAccount", arg0, arg1)
	ret0, _ := ret[0].(*types.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceAccount indicates an expected call of FindServiceAccount.
func (mr *MockPrincipalStoreMockRecorder) FindServiceAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceAccount", reflect.TypeOf((*MockPrincipalStore)(nil).FindServiceAccount), arg0, arg1)
}

// FindServiceAccountByUID mocks base method.
func (m *MockPrincipalStore) FindServiceAccountByUID(arg0 context.Context, arg1 string) (*types.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindServiceAccountByUID", arg0, arg1)
	ret0, _ := ret[0].(*types.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceAccountByUID indicates an expected call of FindServiceAccountByUID.
func (mr *MockPrincipalStoreMockRecorder) FindServiceAccountByUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceAccountByUID", reflect.TypeOf((*MockPrincipalStore)(nil).FindServiceAccountByUID), arg0, arg1)
}

// FindServiceByUID mocks base method.
func (m *MockPrincipalStore) FindServiceByUID(arg0 context.Context, arg1 string) (*types.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindServiceByUID", arg0, arg1)
	ret0, _ := ret[0].(*types.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindServiceByUID indicates an expected call of FindServiceByUID.
func (mr *MockPrincipalStoreMockRecorder) FindServiceByUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindServiceByUID", reflect.TypeOf((*MockPrincipalStore)(nil).FindServiceByUID), arg0, arg1)
}

// FindUser mocks base method.
func (m *MockPrincipalStore) FindUser(arg0 context.Context, arg1 int64) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", arg0, arg1)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockPrincipalStoreMockRecorder) FindUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockPrincipalStore)(nil).FindUser), arg0, arg1)
}

// FindUserByEmail mocks base method.
func (m *MockPrincipalStore) FindUserByEmail(arg0 context.Context, arg1 string) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockPrincipalStoreMockRecorder) FindUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockPrincipalStore)(nil).FindUserByEmail), arg0, arg1)
}

// FindUserByUID mocks base method.
func (m *MockPrincipalStore) FindUserByUID(arg0 context.Context, arg1 string) (*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByUID", arg0, arg1)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUID indicates an expected call of FindUserByUID.
func (mr *MockPrincipalStoreMockRecorder) FindUserByUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUID", reflect.TypeOf((*MockPrincipalStore)(nil).FindUserByUID), arg0, arg1)
}

// ListServiceAccounts mocks base method.
func (m *MockPrincipalStore) ListServiceAccounts(arg0 context.Context, arg1 enum.ParentResourceType, arg2 int64) ([]*types.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServiceAccounts", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServiceAccounts indicates an expected call of ListServiceAccounts.
func (mr *MockPrincipalStoreMockRecorder) ListServiceAccounts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServiceAccounts", reflect.TypeOf((*MockPrincipalStore)(nil).ListServiceAccounts), arg0, arg1, arg2)
}

// ListServices mocks base method.
func (m *MockPrincipalStore) ListServices(arg0 context.Context) ([]*types.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", arg0)
	ret0, _ := ret[0].([]*types.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockPrincipalStoreMockRecorder) ListServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockPrincipalStore)(nil).ListServices), arg0)
}

// ListUsers mocks base method.
func (m *MockPrincipalStore) ListUsers(arg0 context.Context, arg1 *types.UserFilter) ([]*types.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0, arg1)
	ret0, _ := ret[0].([]*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockPrincipalStoreMockRecorder) ListUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockPrincipalStore)(nil).ListUsers), arg0, arg1)
}

// UpdateService mocks base method.
func (m *MockPrincipalStore) UpdateService(arg0 context.Context, arg1 *types.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockPrincipalStoreMockRecorder) UpdateService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockPrincipalStore)(nil).UpdateService), arg0, arg1)
}

// UpdateServiceAccount mocks base method.
func (m *MockPrincipalStore) UpdateServiceAccount(arg0 context.Context, arg1 *types.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServiceAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateServiceAccount indicates an expected call of UpdateServiceAccount.
func (mr *MockPrincipalStoreMockRecorder) UpdateServiceAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServiceAccount", reflect.TypeOf((*MockPrincipalStore)(nil).UpdateServiceAccount), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockPrincipalStore) UpdateUser(arg0 context.Context, arg1 *types.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockPrincipalStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockPrincipalStore)(nil).UpdateUser), arg0, arg1)
}

// MockSpaceStore is a mock of SpaceStore interface.
type MockSpaceStore struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceStoreMockRecorder
}

// MockSpaceStoreMockRecorder is the mock recorder for MockSpaceStore.
type MockSpaceStoreMockRecorder struct {
	mock *MockSpaceStore
}

// NewMockSpaceStore creates a new mock instance.
func NewMockSpaceStore(ctrl *gomock.Controller) *MockSpaceStore {
	mock := &MockSpaceStore{ctrl: ctrl}
	mock.recorder = &MockSpaceStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceStore) EXPECT() *MockSpaceStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockSpaceStore) Count(arg0 context.Context, arg1 int64, arg2 *types.SpaceFilter) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockSpaceStoreMockRecorder) Count(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockSpaceStore)(nil).Count), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockSpaceStore) Create(arg0 context.Context, arg1 *types.Space) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSpaceStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSpaceStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockSpaceStore) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSpaceStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSpaceStore)(nil).Delete), arg0, arg1)
}

// Find mocks base method.
func (m *MockSpaceStore) Find(arg0 context.Context, arg1 int64) (*types.Space, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*types.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockSpaceStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockSpaceStore)(nil).Find), arg0, arg1)
}

// FindByRef mocks base method.
func (m *MockSpaceStore) FindByRef(arg0 context.Context, arg1 string) (*types.Space, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByRef", arg0, arg1)
	ret0, _ := ret[0].(*types.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByRef indicates an expected call of FindByRef.
func (mr *MockSpaceStoreMockRecorder) FindByRef(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByRef", reflect.TypeOf((*MockSpaceStore)(nil).FindByRef), arg0, arg1)
}

// List mocks base method.
func (m *MockSpaceStore) List(arg0 context.Context, arg1 int64, arg2 *types.SpaceFilter) ([]*types.Space, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSpaceStoreMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSpaceStore)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockSpaceStore) Update(arg0 context.Context, arg1 *types.Space) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockSpaceStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSpaceStore)(nil).Update), arg0, arg1)
}

// UpdateOptLock mocks base method.
func (m *MockSpaceStore) UpdateOptLock(arg0 context.Context, arg1 *types.Space, arg2 func(*types.Space) error) (*types.Space, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOptLock", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOptLock indicates an expected call of UpdateOptLock.
func (mr *MockSpaceStoreMockRecorder) UpdateOptLock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOptLock", reflect.TypeOf((*MockSpaceStore)(nil).UpdateOptLock), arg0, arg1, arg2)
}

// MockRepoStore is a mock of RepoStore interface.
type MockRepoStore struct {
	ctrl     *gomock.Controller
	recorder *MockRepoStoreMockRecorder
}

// MockRepoStoreMockRecorder is the mock recorder for MockRepoStore.
type MockRepoStoreMockRecorder struct {
	mock *MockRepoStore
}

// NewMockRepoStore creates a new mock instance.
func NewMockRepoStore(ctrl *gomock.Controller) *MockRepoStore {
	mock := &MockRepoStore{ctrl: ctrl}
	mock.recorder = &MockRepoStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepoStore) EXPECT() *MockRepoStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockRepoStore) Count(arg0 context.Context, arg1 int64, arg2 *types.RepoFilter) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockRepoStoreMockRecorder) Count(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockRepoStore)(nil).Count), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockRepoStore) Create(arg0 context.Context, arg1 *types.Repository) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepoStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepoStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockRepoStore) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepoStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepoStore)(nil).Delete), arg0, arg1)
}

// Find mocks base method.
func (m *MockRepoStore) Find(arg0 context.Context, arg1 int64) (*types.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*types.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRepoStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepoStore)(nil).Find), arg0, arg1)
}

// FindByRef mocks base method.
func (m *MockRepoStore) FindByRef(arg0 context.Context, arg1 string) (*types.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByRef", arg0, arg1)
	ret0, _ := ret[0].(*types.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByRef indicates an expected call of FindByRef.
func (mr *MockRepoStoreMockRecorder) FindByRef(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByRef", reflect.TypeOf((*MockRepoStore)(nil).FindByRef), arg0, arg1)
}

// List mocks base method.
func (m *MockRepoStore) List(arg0 context.Context, arg1 int64, arg2 *types.RepoFilter) ([]*types.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRepoStoreMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepoStore)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockRepoStore) Update(arg0 context.Context, arg1 *types.Repository) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepoStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepoStore)(nil).Update), arg0, arg1)
}

// UpdateOptLock mocks base method.
func (m *MockRepoStore) UpdateOptLock(arg0 context.Context, arg1 *types.Repository, arg2 func(*types.Repository) error) (*types.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOptLock", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOptLock indicates an expected call of UpdateOptLock.
func (mr *MockRepoStoreMockRecorder) UpdateOptLock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOptLock", reflect.TypeOf((*MockRepoStore)(nil).UpdateOptLock), arg0, arg1, arg2)
}
