package middleware

import (
	"github.com/gin-gonic/gin"
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
