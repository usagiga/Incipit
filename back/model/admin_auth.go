package model

import "github.com/usagiga/Incipit/back/entity"

type AdminAuthModelImpl struct {

}

func NewAdminAuthModel() AdminAuthModel {
	return &AdminAuthModelImpl{}
}

func (m *AdminAuthModelImpl) Authorize(accToken *entity.AccessToken) (user *entity.AdminUser, err error) {
	panic("implement me")
}

func (m *AdminAuthModelImpl) Login(name, password string) (accToken *entity.AccessToken, refToken *entity.RefreshToken, err error) {
	panic("implement me")
}

func (m *AdminAuthModelImpl) RenewAccessToken(refTokenStr string) (accToken *entity.AccessToken, refToken *entity.RefreshToken, err error) {
	panic("implement me")
}
