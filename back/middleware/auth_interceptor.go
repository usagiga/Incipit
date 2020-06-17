package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/usagiga/Incipit/back/entity"
)

// AuthInterceptor represents ...
type AuthInterceptorImpl struct {

}

func NewAuthInterceptor() AuthInterceptor {
	return &AuthInterceptorImpl{}
}

func (i *AuthInterceptorImpl) Handle(c *gin.Context) {
	panic("implement me")
}

func (i *AuthInterceptorImpl) Authorize(token string) (user *entity.AdminUser, err error) {
	panic("implement me")
}
