package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/handler"
	libConfig "github.com/usagiga/Incipit/back/lib/config"
	"github.com/usagiga/Incipit/back/middleware"
	"github.com/usagiga/Incipit/back/model"
	"log"
)

func main() {
	// Load config
	config := &entity.Config{}
	err := libConfig.Load(config)
	if err != nil {
		log.Fatalln(err)
	}

	// Connect to DB
	db := ConnectToDB(config.MySQLUser, config.MySQLPassword, config.MySQLHost, config.MySQLPort)
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
	linkValidator := model.NewLinkValidator(config.IncipitHost)

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
	installHandler := handler.NewInstallHandler(installerModel)

	// Register to gin
	router := gin.Default()
	normalScenarioGroup := router.Group("/")
	normalScenarioGroup.Use(installInterceptor.HandleNeededInstall)

	adminUserGroup := normalScenarioGroup.Group("/api/admin/")
	adminUserGroup.Use(authInterceptor.Handle)
	adminUserGroup.GET("/", adminUserHandler.HandleGetAdmin)
	adminUserGroup.POST("/", adminUserHandler.HandleCreateAdmin)
	adminUserGroup.PATCH("/", adminUserHandler.HandleUpdateAdmin)
	adminUserGroup.DELETE("/", adminUserHandler.HandleDeleteAdmin)

	loginGroup := normalScenarioGroup.Group("/api/login")
	loginGroup.POST("/", adminAuthHandler.HandleLogin)
	loginGroup.POST("/refresh", adminAuthHandler.HandleRefreshToken)

	linkGroup := normalScenarioGroup.Group("/api/link")
	linkGroup.Use(authInterceptor.Handle)
	linkGroup.GET("/", linkHandler.HandleGetLink)
	linkGroup.POST("/", linkHandler.HandleCreateLink)
	linkGroup.PATCH("/", linkHandler.HandleUpdateLink)
	linkGroup.DELETE("/", linkHandler.HandleDeleteLink)
	normalScenarioGroup.GET("/api/link/shortened", linkHandler.HandleGetLinkByShortURL) // Unnecessary auth

	installerGroup := router.Group("/api/install")
	installerGroup.Use(installInterceptor.HandleRedundantInstall)
	installerGroup.POST("/", installHandler.HandleInstall)

	// Launch
	port := fmt.Sprintf(":%d", config.IncipitPort)
	err = router.Run(port)
	if err != nil {
		log.Fatalln(err)
	}
}
