package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/usagiga/Incipit/back/entity"
)

type AdminUserValidatorImpl struct {

}

func NewAdminUserValidator() AdminUserValidator {
	return &AdminUserValidatorImpl{}
}

func (m *AdminUserValidatorImpl) Handle(c *gin.Context) {
	panic("implement me")
}

func (m *AdminUserValidatorImpl) ValidateAll(user *entity.AdminUser) (err error) {
	panic("implement me")
}

func (m *AdminUserValidatorImpl) ValidateID(id uint) (err error) {
	panic("implement me")
}

func (m *AdminUserValidatorImpl) ValidateName(name string) (err error) {
	panic("implement me")
}

func (m *AdminUserValidatorImpl) ValidateScreenName(scName string) (err error) {
	panic("implement me")
}

func (m *AdminUserValidatorImpl) ValidatePassword(password string) (err error) {
	panic("implement me")
}
