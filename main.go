package main

import (
	"os"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
)

func setupRouter(appState *AppState, errLogger *log.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	apiRoutes := router.Group("/api")
	apiRoutes.GET("/connections/:user_id", newConnectionHandler(appState, errLogger))
	apiRoutes.POST("/deploy", deployHandler(appState, errLogger))
	return router
}

func main() {
	setupLogger()
	errLogger := makeErrLogger()

	configFilename := os.Args[1]
	serviceConfig, err := loadConfig(configFilename)

	if err != nil {
		fmt.Printf("Cannot read config file %s\n", configFilename)
		os.Exit(1)
	}

	appState := initAppState(serviceConfig)
	router := setupRouter(appState, errLogger)
	router.Run(":" + getEnv("PORT", "3000"))
}
