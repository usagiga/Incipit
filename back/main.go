package main

import (
	"github.com/gin-gonic/gin"
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/handler"
	"github.com/usagiga/Incipit/back/lib/config"
	"github.com/usagiga/Incipit/back/middleware"
	"github.com/usagiga/Incipit/back/model"
	"log"
	"strconv"
)

func main() {
	// Load config
	c := &entity.Config{}
	err := config.Load(c)
	if err != nil {
		log.Fatalln(err)
	}

	// Connect to DB
	db := ConnectToDB(c.MySQLUser, c.MySQLPassword, c.MySQLHost, c.MySQLPort)
	defer db.Close()

	// Auto migrate
	err = Migrate(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Build modules
	adminUserValidator, err := model.NewAdminUserValidator()
	if err != nil {
		log.Fatalln(err)
	}
	linkValidator := model.NewLinkValidator(c.IncipitHost)

	hashModel := model.NewHashModel()
	adminModel := model.NewAdminModel(db, hashModel, adminUserValidator)
	adminAuthModel := model.NewAdminAuthModel(db, adminModel, hashModel)
	linkModel := model.NewLinkModel(db, linkValidator)
	installerModel := model.NewInstallerModel(adminModel, adminAuthModel)

	authInterceptor := middleware.NewAuthInterceptor(adminAuthModel)
	installInterceptor := middleware.NewInstallInterceptor(installerModel)

	adminUserHandler := handler.NewAdminUserHandler(adminModel)
	adminAuthHandler := handler.NewAdminAuthHandler(adminAuthModel)
	linkHandler := handler.NewLinkHandler(linkModel)

	// Register to gin
	router := gin.Default()
	router.Use(installInterceptor.Handle)

	adminUserGroup := router.Group("/api/admin/")
	adminUserGroup.Use(authInterceptor.Handle)
	adminUserGroup.GET("/", adminUserHandler.HandleGetAdmin)
	adminUserGroup.POST("/", adminUserHandler.HandleCreateAdmin)
	adminUserGroup.PATCH("/", adminUserHandler.HandleUpdateAdmin)
	adminUserGroup.DELETE("/", adminUserHandler.HandleDeleteAdmin)

	loginGroup := router.Group("/api/login")
	loginGroup.POST("/", adminAuthHandler.HandleLogin)
	loginGroup.POST("/refresh", adminAuthHandler.HandleRefreshToken)

	linkGroup := router.Group("/api/link")
	linkGroup.Use(authInterceptor.Handle)
	linkGroup.GET("/", linkHandler.HandleGetLink)
	linkGroup.POST("/", linkHandler.HandleCreateLink)
	linkGroup.PATCH("/", linkHandler.HandleUpdateLink)
	linkGroup.DELETE("/", linkHandler.HandleDeleteLink)
	router.GET("/api/link/shortened", linkHandler.HandleGetLinkByShortURL) // Unnecessary auth

	// Launch
	port := strconv.Itoa(c.IncipitPort)
	err = router.Run(port)
	if err != nil {
		log.Fatalln(err)
	}
}
