package core

import (
	"os"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	
	"sync"
)

func setupRouter(appState *AppState, errLogger *log.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("public", true)))
	apiRoutes := router.Group("/api")
	apiRoutes.GET("/connect", newConnectionHandler(appState, errLogger))
	apiRoutes.POST("/deploy", deployHandler(appState, errLogger))
	return router
}

func monitorServiceStates(appState *AppState, errLogger *log.Logger) {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for range ticker.C {
			appState.ServiceStates = updateContainerStatus(appState.ServiceStates, errLogger)
			pushServiceStates(appState)
		}
	}()
}

func LaunchServer(configFilename string, port string) {
	errLogger := makeErrLogger()
	serviceConfig, err := loadConfig(configFilename)

	setupLogger()

	if err != nil {
		fmt.Printf("Cannot read config file %s\n", configFilename)
		os.Exit(1)
	}

	log.Println("Read configurations from " + configFilename)

	appState := initAppState(serviceConfig, errLogger)
	router := setupRouter(appState, errLogger)
	monitorServiceStates(appState, errLogger)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		router.Run(":" + port)
	}()

	log.Println("Docker Deploy is now listening on port " + port)
	log.Println(fmt.Sprintf("You can now visit the dashboard at http://localhost:%s in your broswer", port))
	wg.Wait()
}
