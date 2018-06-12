package main

import (
	"os"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
)

func setupRouter(appState *AppState, errLogger *log.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	apiRoutes := router.Group("/api")
	apiRoutes.GET("/connect", newConnectionHandler(appState, errLogger))
	apiRoutes.POST("/deploy", deployHandler(appState, errLogger))
	return router
}

func monitorServiceStates(appState *AppState, errLogger *log.Logger)  {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for range ticker.C {
			appState.ServiceStates = updateContainerStatus(appState.ServiceStates, errLogger)
			for userId := range appState.Clients {
				pushServiceStates(userId, appState)
			}
		}
	}()
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

	appState := initAppState(serviceConfig, errLogger)
	router := setupRouter(appState, errLogger)
	monitorServiceStates(appState, errLogger)

	router.Run(":" + getEnv("PORT", "3000"))

}
