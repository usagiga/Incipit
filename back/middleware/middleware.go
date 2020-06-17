package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/usagiga/Incipit/back/entity"
)

// InstallInterceptor treats installing if Incipit is needed to initialize
type InstallInterceptor interface {
	// Handle is to handle validation in middleware chain
	Handle(c *gin.Context)
	IsNeededInstall() (isNeeded bool, err error)
}

// AuthInterceptor treats authentication user
type AuthInterceptor interface {
	// Handle is to handle validation in middleware chain
	Handle(c *gin.Context)
	IsNeededLogin(token string) (user *entity.AdminUser, err error)
}
