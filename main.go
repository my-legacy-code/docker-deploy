package main

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter(appState *AppState) *gin.Engine {
	router := gin.Default()
	router.Static("/assets", "./public")
	router.LoadHTMLGlob("views/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "services.html", gin.H{
			"services": appState.ServiceStates,
		})
	})
	router.GET("/api/real-time-channel", realTimeChannelHandler(appState))
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
