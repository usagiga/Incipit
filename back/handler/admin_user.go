package handler

import "github.com/gin-gonic/gin"

type AdminUserHandlerImpl struct {

}

func NewAdminUserHandler() AdminUserHandler {
	return 	&AdminUserHandlerImpl{}
}

func (h *AdminUserHandlerImpl) HandleGetAdmin(c *gin.Context) {
	panic("implement me")
}

func (h *AdminUserHandlerImpl) HandleCreateAdmin(c *gin.Context) {
	panic("implement me")
}

func (h *AdminUserHandlerImpl) HandleUpdateAdmin(c *gin.Context) {
	panic("implement me")
}

func (h *AdminUserHandlerImpl) HandleDeleteAdmin(c *gin.Context) {
	panic("implement me")
}
