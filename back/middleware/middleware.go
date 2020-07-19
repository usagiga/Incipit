package middleware

import (
	"github.com/gin-gonic/gin"
)

// InstallInterceptor treats installing if Incipit is needed to initialize
type InstallInterceptor interface {
	// HandleNeededInstall is to handle validation in middleware chain.
	// It will abort requests if this app needs to install
	HandleNeededInstall(c *gin.Context)

	// HandleRedundantInstall is to handle validation in middleware chain.
	// It will abort requests if this app doesn't need to install
	HandleRedundantInstall(c *gin.Context)
}

// AuthInterceptor treats authentication user
type AuthInterceptor interface {
	// Handle is to handle validation in middleware chain
	Handle(c *gin.Context)
}
