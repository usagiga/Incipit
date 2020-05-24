package model

import (
	"github.com/usagiga/Incipit/back/entity"
)

type InstallerModelImpl struct {

}

func NewInstallerModel() InstallerModel {
	return &InstallerModelImpl{}
}

func (m *InstallerModelImpl) CreateNewAdmin(email, password string) (accToken *entity.AccessToken, refToken *entity.RefreshToken, err error) {
	panic("implement me")
}
