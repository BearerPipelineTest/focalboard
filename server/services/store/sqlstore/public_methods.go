// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

// Code generated by "make generate" from the Store interface
// DO NOT EDIT

// To add a public method, create an entry in the Store interface,
// prefix it with a @withTransaction comment if you need it to be
// transactional and then add a private method in the store itself
// with db sq.BaseRunner as the first parameter before running `make
// generate`

package sqlstore

import (
	"context"

	"github.com/mattermost/focalboard/server/model"
	"github.com/mattermost/focalboard/server/services/store"

	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

func (s *SQLStore) AddUpdateCategoryBlock(userID string, categoryID string, blockID string) error {
	return s.addUpdateCategoryBlock(s.db, userID, categoryID, blockID)

}

func (s *SQLStore) CleanUpSessions(expireTime int64) error {
	return s.cleanUpSessions(s.db, expireTime)

}

func (s *SQLStore) CreateCategory(category model.Category) error {
	return s.createCategory(s.db, category)

}

func (s *SQLStore) CreateSession(session *model.Session) error {
	return s.createSession(s.db, session)

}

func (s *SQLStore) CreateUser(user *model.User) error {
	return s.createUser(s.db, user)

}

func (s *SQLStore) DeleteBlock(c store.Container, blockID string, modifiedBy string) error {
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.deleteBlock(tx, c, blockID, modifiedBy)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "DeleteBlock"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) DeleteCategory(categoryID string, userID string, teamID string) error {
	return s.deleteCategory(s.db, categoryID, userID, teamID)

}

func (s *SQLStore) DeleteSession(sessionID string) error {
	return s.deleteSession(s.db, sessionID)

}

func (s *SQLStore) GetActiveUserCount(updatedSecondsAgo int64) (int, error) {
	return s.getActiveUserCount(s.db, updatedSecondsAgo)

}

func (s *SQLStore) GetAllBlocks(c store.Container) ([]model.Block, error) {
	return s.getAllBlocks(s.db, c)

}

func (s *SQLStore) GetBlock(c store.Container, blockID string) (*model.Block, error) {
	return s.getBlock(s.db, c, blockID)

}

func (s *SQLStore) GetBlockCountsByType() (map[string]int64, error) {
	return s.getBlockCountsByType(s.db)

}

func (s *SQLStore) GetBlocksWithParent(c store.Container, parentID string) ([]model.Block, error) {
	return s.getBlocksWithParent(s.db, c, parentID)

}

func (s *SQLStore) GetBlocksWithParentAndType(c store.Container, parentID string, blockType string) ([]model.Block, error) {
	return s.getBlocksWithParentAndType(s.db, c, parentID, blockType)

}

func (s *SQLStore) GetBlocksWithRootID(c store.Container, rootID string) ([]model.Block, error) {
	return s.getBlocksWithRootID(s.db, c, rootID)

}

func (s *SQLStore) GetBlocksWithType(c store.Container, blockType string) ([]model.Block, error) {
	return s.getBlocksWithType(s.db, c, blockType)

}

func (s *SQLStore) GetCategory(id string) (*model.Category, error) {
	return s.getCategory(s.db, id)

}

func (s *SQLStore) GetParentID(c store.Container, blockID string) (string, error) {
	return s.getParentID(s.db, c, blockID)

}

func (s *SQLStore) GetRegisteredUserCount() (int, error) {
	return s.getRegisteredUserCount(s.db)

}

func (s *SQLStore) GetRootID(c store.Container, blockID string) (string, error) {
	return s.getRootID(s.db, c, blockID)

}

func (s *SQLStore) GetSession(token string, expireTime int64) (*model.Session, error) {
	return s.getSession(s.db, token, expireTime)

}

func (s *SQLStore) GetSharing(c store.Container, rootID string) (*model.Sharing, error) {
	return s.getSharing(s.db, c, rootID)

}

func (s *SQLStore) GetSubTree2(c store.Container, blockID string) ([]model.Block, error) {
	return s.getSubTree2(s.db, c, blockID)

}

func (s *SQLStore) GetSubTree3(c store.Container, blockID string) ([]model.Block, error) {
	return s.getSubTree3(s.db, c, blockID)

}

func (s *SQLStore) GetSystemSetting(key string) (string, error) {
	return s.getSystemSetting(s.db, key)

}

func (s *SQLStore) GetSystemSettings() (map[string]string, error) {
	return s.getSystemSettings(s.db)

}

func (s *SQLStore) GetUserByEmail(email string) (*model.User, error) {
	return s.getUserByEmail(s.db, email)

}

func (s *SQLStore) GetUserByID(userID string) (*model.User, error) {
	return s.getUserByID(s.db, userID)

}

func (s *SQLStore) GetUserByUsername(username string) (*model.User, error) {
	return s.getUserByUsername(s.db, username)

}

func (s *SQLStore) GetUserCategoryBlocks(userID string, teamID string) ([]model.CategoryBlocks, error) {
	return s.getUserCategoryBlocks(s.db, userID, teamID)

}

func (s *SQLStore) GetUserWorkspaces(userID string) ([]model.UserWorkspace, error) {
	return s.getUserWorkspaces(s.db, userID)

}

func (s *SQLStore) GetUsersByWorkspace(workspaceID string) ([]*model.User, error) {
	return s.getUsersByWorkspace(s.db, workspaceID)

}

func (s *SQLStore) GetWorkspace(ID string) (*model.Workspace, error) {
	return s.getWorkspace(s.db, ID)

}

func (s *SQLStore) GetWorkspaceCount() (int64, error) {
	return s.getWorkspaceCount(s.db)

}

func (s *SQLStore) HasWorkspaceAccess(userID string, workspaceID string) (bool, error) {
	return s.hasWorkspaceAccess(s.db, userID, workspaceID)

}

func (s *SQLStore) InsertBlock(c store.Container, block *model.Block, userID string) error {
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.insertBlock(tx, c, block, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "InsertBlock"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) PatchBlock(c store.Container, blockID string, blockPatch *model.BlockPatch, userID string) error {
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.patchBlock(tx, c, blockID, blockPatch, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "PatchBlock"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) RefreshSession(session *model.Session) error {
	return s.refreshSession(s.db, session)

}

func (s *SQLStore) SetSystemSetting(key string, value string) error {
	return s.setSystemSetting(s.db, key, value)

}

func (s *SQLStore) UpdateCategory(category model.Category) error {
	return s.updateCategory(s.db, category)

}

func (s *SQLStore) UpdateSession(session *model.Session) error {
	return s.updateSession(s.db, session)

}

func (s *SQLStore) UpdateUser(user *model.User) error {
	return s.updateUser(s.db, user)

}

func (s *SQLStore) UpdateUserPassword(username string, password string) error {
	return s.updateUserPassword(s.db, username, password)

}

func (s *SQLStore) UpdateUserPasswordByID(userID string, password string) error {
	return s.updateUserPasswordByID(s.db, userID, password)

}

func (s *SQLStore) UpsertSharing(c store.Container, sharing model.Sharing) error {
	return s.upsertSharing(s.db, c, sharing)

}

func (s *SQLStore) UpsertWorkspaceSettings(workspace model.Workspace) error {
	return s.upsertWorkspaceSettings(s.db, workspace)

}

func (s *SQLStore) UpsertWorkspaceSignupToken(workspace model.Workspace) error {
	return s.upsertWorkspaceSignupToken(s.db, workspace)

}
