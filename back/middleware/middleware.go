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

// AdminUserValidator treats validation admin user in request arguments
type AdminUserValidator interface {
	// Handle is to handle validation in middleware chain
	Handle(c *gin.Context)
	ValidateAll(user *entity.AdminUser) (err error)
	ValidateID(id uint) (err error)
	ValidateName(name string) (err error)
	ValidateScreenName(scName string) (err error)
	ValidatePassword(password string) (err error)
}

// LinkValidator treats validation link in request arguments
type LinkValidator interface {
	// Handle is to handle validation in middleware chain
	Handle(c *gin.Context)
	ValidateAll(link *entity.Link) (err error)
	ValidateID(id uint) (err error)
	ValidateURL(url string) (err error)
}

// AuthInterceptor treats authentication user
type AuthInterceptor interface {
	// Handle is to handle validation in middleware chain
	Handle(c *gin.Context)
	IsNeededLogin(token string) (user *entity.AdminUser, err error)
}
