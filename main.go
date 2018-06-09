package main

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter(serviceStates serviceStates) *gin.Engine{
	router := gin.Default()
	router.Static("/assets", "./public")
	router.LoadHTMLGlob("views/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "services.html", gin.H{
			"services": serviceStates,
		})
	})
	router.POST("/deploy", deployHandler(serviceStates))
	return router
}

func main() {

	configFilename := os.Args[1]
	serviceConfig, err := loadConfig(configFilename)

	if err != nil {
		fmt.Printf("Cannot read config file %s\n", configFilename)
		os.Exit(1)
	}

	serviceStates := initServiceState(serviceConfig)

	router := setupRouter(serviceStates)
	router.Run(":" + getEnv("PORT", "3000"))
}
