package middleware

import (
	"github.com/gin-gonic/gin"
)

// InstallInterceptor treats installing if Incipit is needed to initialize
type InstallInterceptor interface {
	// Handle is to handle validation in middleware chain
	Handle(c *gin.Context)
}

// AuthInterceptor treats authentication user
type AuthInterceptor interface {
	// Handle is to handle validation in middleware chain
	Handle(c *gin.Context)
}
