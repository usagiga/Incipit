package model

import (
	"github.com/usagiga/Incipit/back/entity"
	"golang.org/x/xerrors"
)

type InstallerModelImpl struct {
	adminModel     AdminModel
	adminAuthModel AdminAuthModel

	// Cache for `IsNeededInstall()`
	isNeeded bool
}

func NewInstallerModel(adminModel AdminModel, adminAuthModel AdminAuthModel) InstallerModel {
	return &InstallerModelImpl{adminModel: adminModel, adminAuthModel: adminAuthModel, isNeeded: true}
}

func (m *InstallerModelImpl) CreateNewAdmin(name, screenName, password string) (accToken *entity.AccessToken, refToken *entity.RefreshToken, err error) {
	newUser := &entity.AdminUser{
		Name:       name,
		ScreenName: screenName,
		Password:   password,
	}

	// Add AdminUser
	_, err = m.adminModel.Add(newUser)
	if err != nil {
		return nil, nil, xerrors.Errorf("Can't add first admin: %w", err)
	}

	// Login
	accToken, refToken, err = m.adminAuthModel.Login(name, password)
	if err != nil {
		return nil, nil, xerrors.Errorf("Can't login first admin: %w", err)
	}

	return accToken, refToken, nil
}

func (m *InstallerModelImpl) IsNeededInstall() (isNeeded bool, err error) {
	// Use cached result
	// Once isNeeded turn false, Incipit return cache value
	if !m.isNeeded {
		return false, nil
	}

	// Find all admin
	// TODO : `Find()` should be `First()` with no condition or `Count()`
	users, err := m.adminModel.Find()
	if err != nil {
		return false, xerrors.Errorf("Can't find admin users: %w", err)
	}

	// Judge
	isNeeded = len(users) <= 0

	// Store result into cache
	m.isNeeded = isNeeded

	return isNeeded, nil
}
