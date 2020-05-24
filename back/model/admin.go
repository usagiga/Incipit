package model

import (
	"github.com/usagiga/Incipit/back/entity"
)

type AdminModelImpl struct {

}

func NewAdminModel() AdminModel {
	return &AdminModelImpl{}
}

func (m *AdminModelImpl) Add(user *entity.AdminUser) (added *entity.AdminUser, err error) {
	panic("implement me")
}

func (m *AdminModelImpl) FindOne(id uint) (user *entity.AdminUser, err error) {
	panic("implement me")
}

func (m *AdminModelImpl) Find() (users []entity.AdminUser, err error) {
	panic("implement me")
}

func (m *AdminModelImpl) Update(updating *entity.AdminUser) (updated *entity.AdminUser, err error) {
	panic("implement me")
}

func (m *AdminModelImpl) Delete(id uint) (err error) {
	panic("implement me")
}
