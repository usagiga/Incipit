package model

import (
	"github.com/usagiga/Incipit/back/entity"
)

type InstallerModelImpl struct {

}

func NewInstallerModel(adminModel AdminModel, adminAuthModel AdminAuthModel) InstallerModel {
	return &InstallerModelImpl{adminModel: adminModel, adminAuthModel: adminAuthModel}
}

func (m *InstallerModelImpl) CreateNewAdmin(name, screenName, password string) (accToken *entity.AccessToken, refToken *entity.RefreshToken, err error) {
	panic("implement me")
}
