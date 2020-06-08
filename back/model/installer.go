package model

import (
	"github.com/usagiga/Incipit/back/entity"
	"golang.org/x/xerrors"
)

type InstallerModelImpl struct {
	adminModel     AdminModel
	adminAuthModel AdminAuthModel
}

func NewInstallerModel(adminModel AdminModel, adminAuthModel AdminAuthModel) InstallerModel {
	return &InstallerModelImpl{adminModel: adminModel, adminAuthModel: adminAuthModel}
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
