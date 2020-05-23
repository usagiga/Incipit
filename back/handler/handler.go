package handler

import "github.com/gin-gonic/gin"

type AdminUserHandler interface {
	// GET /api/admin
	HandleGetAdmin(c *gin.Context)
	// POST /api/admin
	HandleCreateAdmin(c *gin.Context)
	// PATCH /api/admin
	HandleUpdateAdmin(c *gin.Context)
	// DELETE /api/admin
	HandleDeleteAdmin(c *gin.Context)
}

type LinkHandler interface {
	// GET /api/link
	HandleGetLink(c *gin.Context)
	// POST /api/link
	HandleCreateLink(c *gin.Context)
	// PATCH /api/link
	HandleUpdateLink(c *gin.Context)
	// DELETE /api/link
	HandleDeleteLink(c *gin.Context)
}
