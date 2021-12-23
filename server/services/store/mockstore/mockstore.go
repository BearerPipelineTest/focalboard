// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/focalboard/server/services/store (interfaces: Store)

// Package mockstore is a generated GoMock package.
package mockstore

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/focalboard/server/model"
	store "github.com/mattermost/focalboard/server/services/store"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddUpdateCategoryBlock mocks base method.
func (m *MockStore) AddUpdateCategoryBlock(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUpdateCategoryBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUpdateCategoryBlock indicates an expected call of AddUpdateCategoryBlock.
func (mr *MockStoreMockRecorder) AddUpdateCategoryBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUpdateCategoryBlock", reflect.TypeOf((*MockStore)(nil).AddUpdateCategoryBlock), arg0, arg1, arg2)
}

// CleanUpSessions mocks base method.
func (m *MockStore) CleanUpSessions(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanUpSessions", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanUpSessions indicates an expected call of CleanUpSessions.
func (mr *MockStoreMockRecorder) CleanUpSessions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanUpSessions", reflect.TypeOf((*MockStore)(nil).CleanUpSessions), arg0)
}

// CreateCategory mocks base method.
func (m *MockStore) CreateCategory(arg0 model.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockStoreMockRecorder) CreateCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockStore)(nil).CreateCategory), arg0)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 *model.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0)
}

// DeleteBlock mocks base method.
func (m *MockStore) DeleteBlock(arg0 store.Container, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBlock indicates an expected call of DeleteBlock.
func (mr *MockStoreMockRecorder) DeleteBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBlock", reflect.TypeOf((*MockStore)(nil).DeleteBlock), arg0, arg1, arg2)
}

// DeleteCategory mocks base method.
func (m *MockStore) DeleteCategory(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockStoreMockRecorder) DeleteCategory(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockStore)(nil).DeleteCategory), arg0, arg1, arg2)
}

// DeleteSession mocks base method.
func (m *MockStore) DeleteSession(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStoreMockRecorder) DeleteSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStore)(nil).DeleteSession), arg0)
}

// GetActiveUserCount mocks base method.
func (m *MockStore) GetActiveUserCount(arg0 int64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveUserCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveUserCount indicates an expected call of GetActiveUserCount.
func (mr *MockStoreMockRecorder) GetActiveUserCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveUserCount", reflect.TypeOf((*MockStore)(nil).GetActiveUserCount), arg0)
}

// GetAllBlocks mocks base method.
func (m *MockStore) GetAllBlocks(arg0 store.Container) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBlocks", arg0)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBlocks indicates an expected call of GetAllBlocks.
func (mr *MockStoreMockRecorder) GetAllBlocks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBlocks", reflect.TypeOf((*MockStore)(nil).GetAllBlocks), arg0)
}

// GetBlock mocks base method.
func (m *MockStore) GetBlock(arg0 store.Container, arg1 string) (*model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", arg0, arg1)
	ret0, _ := ret[0].(*model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockStoreMockRecorder) GetBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockStore)(nil).GetBlock), arg0, arg1)
}

// GetBlockCountsByType mocks base method.
func (m *MockStore) GetBlockCountsByType() (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockCountsByType")
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockCountsByType indicates an expected call of GetBlockCountsByType.
func (mr *MockStoreMockRecorder) GetBlockCountsByType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockCountsByType", reflect.TypeOf((*MockStore)(nil).GetBlockCountsByType))
}

// GetBlocksWithParent mocks base method.
func (m *MockStore) GetBlocksWithParent(arg0 store.Container, arg1 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocksWithParent", arg0, arg1)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocksWithParent indicates an expected call of GetBlocksWithParent.
func (mr *MockStoreMockRecorder) GetBlocksWithParent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocksWithParent", reflect.TypeOf((*MockStore)(nil).GetBlocksWithParent), arg0, arg1)
}

// GetBlocksWithParentAndType mocks base method.
func (m *MockStore) GetBlocksWithParentAndType(arg0 store.Container, arg1, arg2 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocksWithParentAndType", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocksWithParentAndType indicates an expected call of GetBlocksWithParentAndType.
func (mr *MockStoreMockRecorder) GetBlocksWithParentAndType(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocksWithParentAndType", reflect.TypeOf((*MockStore)(nil).GetBlocksWithParentAndType), arg0, arg1, arg2)
}

// GetBlocksWithRootID mocks base method.
func (m *MockStore) GetBlocksWithRootID(arg0 store.Container, arg1 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocksWithRootID", arg0, arg1)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocksWithRootID indicates an expected call of GetBlocksWithRootID.
func (mr *MockStoreMockRecorder) GetBlocksWithRootID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocksWithRootID", reflect.TypeOf((*MockStore)(nil).GetBlocksWithRootID), arg0, arg1)
}

// GetBlocksWithType mocks base method.
func (m *MockStore) GetBlocksWithType(arg0 store.Container, arg1 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocksWithType", arg0, arg1)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocksWithType indicates an expected call of GetBlocksWithType.
func (mr *MockStoreMockRecorder) GetBlocksWithType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocksWithType", reflect.TypeOf((*MockStore)(nil).GetBlocksWithType), arg0, arg1)
}

// GetCategory mocks base method.
func (m *MockStore) GetCategory(arg0 string) (*model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", arg0)
	ret0, _ := ret[0].(*model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockStoreMockRecorder) GetCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockStore)(nil).GetCategory), arg0)
}

// GetParentID mocks base method.
func (m *MockStore) GetParentID(arg0 store.Container, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParentID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParentID indicates an expected call of GetParentID.
func (mr *MockStoreMockRecorder) GetParentID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParentID", reflect.TypeOf((*MockStore)(nil).GetParentID), arg0, arg1)
}

// GetRegisteredUserCount mocks base method.
func (m *MockStore) GetRegisteredUserCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegisteredUserCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegisteredUserCount indicates an expected call of GetRegisteredUserCount.
func (mr *MockStoreMockRecorder) GetRegisteredUserCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegisteredUserCount", reflect.TypeOf((*MockStore)(nil).GetRegisteredUserCount))
}

// GetRootID mocks base method.
func (m *MockStore) GetRootID(arg0 store.Container, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRootID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRootID indicates an expected call of GetRootID.
func (mr *MockStoreMockRecorder) GetRootID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRootID", reflect.TypeOf((*MockStore)(nil).GetRootID), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 string, arg1 int64) (*model.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(*model.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetSharing mocks base method.
func (m *MockStore) GetSharing(arg0 store.Container, arg1 string) (*model.Sharing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSharing", arg0, arg1)
	ret0, _ := ret[0].(*model.Sharing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSharing indicates an expected call of GetSharing.
func (mr *MockStoreMockRecorder) GetSharing(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSharing", reflect.TypeOf((*MockStore)(nil).GetSharing), arg0, arg1)
}

// GetSubTree2 mocks base method.
func (m *MockStore) GetSubTree2(arg0 store.Container, arg1 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubTree2", arg0, arg1)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubTree2 indicates an expected call of GetSubTree2.
func (mr *MockStoreMockRecorder) GetSubTree2(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubTree2", reflect.TypeOf((*MockStore)(nil).GetSubTree2), arg0, arg1)
}

// GetSubTree3 mocks base method.
func (m *MockStore) GetSubTree3(arg0 store.Container, arg1 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubTree3", arg0, arg1)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubTree3 indicates an expected call of GetSubTree3.
func (mr *MockStoreMockRecorder) GetSubTree3(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubTree3", reflect.TypeOf((*MockStore)(nil).GetSubTree3), arg0, arg1)
}

// GetSystemSetting mocks base method.
func (m *MockStore) GetSystemSetting(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSystemSetting", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSystemSetting indicates an expected call of GetSystemSetting.
func (mr *MockStoreMockRecorder) GetSystemSetting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemSetting", reflect.TypeOf((*MockStore)(nil).GetSystemSetting), arg0)
}

// GetSystemSettings mocks base method.
func (m *MockStore) GetSystemSettings() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSystemSettings")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSystemSettings indicates an expected call of GetSystemSettings.
func (mr *MockStoreMockRecorder) GetSystemSettings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemSettings", reflect.TypeOf((*MockStore)(nil).GetSystemSettings))
}

// GetUserByEmail mocks base method.
func (m *MockStore) GetUserByEmail(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockStoreMockRecorder) GetUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockStore)(nil).GetUserByEmail), arg0)
}

// GetUserByID mocks base method.
func (m *MockStore) GetUserByID(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockStoreMockRecorder) GetUserByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockStore)(nil).GetUserByID), arg0)
}

// GetUserByUsername mocks base method.
func (m *MockStore) GetUserByUsername(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStoreMockRecorder) GetUserByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUserByUsername), arg0)
}

// GetUserCategoryBlocks mocks base method.
func (m *MockStore) GetUserCategoryBlocks(arg0, arg1 string) ([]model.CategoryBlocks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserCategoryBlocks", arg0, arg1)
	ret0, _ := ret[0].([]model.CategoryBlocks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserCategoryBlocks indicates an expected call of GetUserCategoryBlocks.
func (mr *MockStoreMockRecorder) GetUserCategoryBlocks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserCategoryBlocks", reflect.TypeOf((*MockStore)(nil).GetUserCategoryBlocks), arg0, arg1)
}

// GetUserWorkspaces mocks base method.
func (m *MockStore) GetUserWorkspaces(arg0 string) ([]model.UserWorkspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserWorkspaces", arg0)
	ret0, _ := ret[0].([]model.UserWorkspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserWorkspaces indicates an expected call of GetUserWorkspaces.
func (mr *MockStoreMockRecorder) GetUserWorkspaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWorkspaces", reflect.TypeOf((*MockStore)(nil).GetUserWorkspaces), arg0)
}

// GetUsersByWorkspace mocks base method.
func (m *MockStore) GetUsersByWorkspace(arg0 string) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByWorkspace", arg0)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByWorkspace indicates an expected call of GetUsersByWorkspace.
func (mr *MockStoreMockRecorder) GetUsersByWorkspace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByWorkspace", reflect.TypeOf((*MockStore)(nil).GetUsersByWorkspace), arg0)
}

// GetWorkspace mocks base method.
func (m *MockStore) GetWorkspace(arg0 string) (*model.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspace", arg0)
	ret0, _ := ret[0].(*model.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspace indicates an expected call of GetWorkspace.
func (mr *MockStoreMockRecorder) GetWorkspace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspace", reflect.TypeOf((*MockStore)(nil).GetWorkspace), arg0)
}

// GetWorkspaceCount mocks base method.
func (m *MockStore) GetWorkspaceCount() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaceCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspaceCount indicates an expected call of GetWorkspaceCount.
func (mr *MockStoreMockRecorder) GetWorkspaceCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaceCount", reflect.TypeOf((*MockStore)(nil).GetWorkspaceCount))
}

// HasWorkspaceAccess mocks base method.
func (m *MockStore) HasWorkspaceAccess(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasWorkspaceAccess", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasWorkspaceAccess indicates an expected call of HasWorkspaceAccess.
func (mr *MockStoreMockRecorder) HasWorkspaceAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasWorkspaceAccess", reflect.TypeOf((*MockStore)(nil).HasWorkspaceAccess), arg0, arg1)
}

// InsertBlock mocks base method.
func (m *MockStore) InsertBlock(arg0 store.Container, arg1 *model.Block, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBlock indicates an expected call of InsertBlock.
func (mr *MockStoreMockRecorder) InsertBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBlock", reflect.TypeOf((*MockStore)(nil).InsertBlock), arg0, arg1, arg2)
}

// PatchBlock mocks base method.
func (m *MockStore) PatchBlock(arg0 store.Container, arg1 string, arg2 *model.BlockPatch, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchBlock", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchBlock indicates an expected call of PatchBlock.
func (mr *MockStoreMockRecorder) PatchBlock(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchBlock", reflect.TypeOf((*MockStore)(nil).PatchBlock), arg0, arg1, arg2, arg3)
}

// RefreshSession mocks base method.
func (m *MockStore) RefreshSession(arg0 *model.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshSession indicates an expected call of RefreshSession.
func (mr *MockStoreMockRecorder) RefreshSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshSession", reflect.TypeOf((*MockStore)(nil).RefreshSession), arg0)
}

// SetSystemSetting mocks base method.
func (m *MockStore) SetSystemSetting(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSystemSetting", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSystemSetting indicates an expected call of SetSystemSetting.
func (mr *MockStoreMockRecorder) SetSystemSetting(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSystemSetting", reflect.TypeOf((*MockStore)(nil).SetSystemSetting), arg0, arg1)
}

// Shutdown mocks base method.
func (m *MockStore) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockStoreMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockStore)(nil).Shutdown))
}

// UpdateCategory mocks base method.
func (m *MockStore) UpdateCategory(arg0 model.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockStoreMockRecorder) UpdateCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockStore)(nil).UpdateCategory), arg0)
}

// UpdateSession mocks base method.
func (m *MockStore) UpdateSession(arg0 *model.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSession indicates an expected call of UpdateSession.
func (mr *MockStoreMockRecorder) UpdateSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSession", reflect.TypeOf((*MockStore)(nil).UpdateSession), arg0)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0)
}

// UpdateUserPassword mocks base method.
func (m *MockStore) UpdateUserPassword(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockStoreMockRecorder) UpdateUserPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockStore)(nil).UpdateUserPassword), arg0, arg1)
}

// UpdateUserPasswordByID mocks base method.
func (m *MockStore) UpdateUserPasswordByID(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPasswordByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPasswordByID indicates an expected call of UpdateUserPasswordByID.
func (mr *MockStoreMockRecorder) UpdateUserPasswordByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPasswordByID", reflect.TypeOf((*MockStore)(nil).UpdateUserPasswordByID), arg0, arg1)
}

// UpsertSharing mocks base method.
func (m *MockStore) UpsertSharing(arg0 store.Container, arg1 model.Sharing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertSharing", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertSharing indicates an expected call of UpsertSharing.
func (mr *MockStoreMockRecorder) UpsertSharing(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertSharing", reflect.TypeOf((*MockStore)(nil).UpsertSharing), arg0, arg1)
}

// UpsertWorkspaceSettings mocks base method.
func (m *MockStore) UpsertWorkspaceSettings(arg0 model.Workspace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertWorkspaceSettings", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWorkspaceSettings indicates an expected call of UpsertWorkspaceSettings.
func (mr *MockStoreMockRecorder) UpsertWorkspaceSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWorkspaceSettings", reflect.TypeOf((*MockStore)(nil).UpsertWorkspaceSettings), arg0)
}

// UpsertWorkspaceSignupToken mocks base method.
func (m *MockStore) UpsertWorkspaceSignupToken(arg0 model.Workspace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertWorkspaceSignupToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWorkspaceSignupToken indicates an expected call of UpsertWorkspaceSignupToken.
func (mr *MockStoreMockRecorder) UpsertWorkspaceSignupToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWorkspaceSignupToken", reflect.TypeOf((*MockStore)(nil).UpsertWorkspaceSignupToken), arg0)
}
