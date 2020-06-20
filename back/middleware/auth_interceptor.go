package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/usagiga/Incipit/back/entity/messages"
	"github.com/usagiga/Incipit/back/model"
	"net/http"
	"strings"
)

type AuthInterceptorImpl struct {
	adminAuthModel model.AdminAuthModel
}

func NewAuthInterceptor(adminAuthModel model.AdminAuthModel) AuthInterceptor {
	return &AuthInterceptorImpl{adminAuthModel: adminAuthModel}
}

func (i *AuthInterceptorImpl) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	splitHeader := strings.Split(authHeader, "Bearer ")
	if len(splitHeader) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "{}")
		return
	}
	accToken := splitHeader[1]

	// Authorize
	user, err := i.adminAuthModel.Authorize(accToken)
	if err != nil {
		resp := messages.NewErrorResponse(err)
		sCode := resp.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, resp)
		return
	}

	// Set user into ctx
	c.Set("user", user)
}
