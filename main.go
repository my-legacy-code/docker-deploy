package main

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
)

func setupRouter(appState *AppState) *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	apiRoutes := router.Group("/api")
	apiRoutes.GET("/connections/:user_id", newConnectionHandler(appState))
	router.POST("/deploy", deployHandler(appState))
	return router
}

func main() {

	configFilename := os.Args[1]
	serviceConfig, err := loadConfig(configFilename)

	if err != nil {
		fmt.Printf("Cannot read config file %s\n", configFilename)
		os.Exit(1)
	}

	appState := initAppState(serviceConfig)
	router := setupRouter(appState)
	router.Run(":" + getEnv("PORT", "3000"))
}
