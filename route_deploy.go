package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"time"
)

func deployHandler(appState *AppState) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload DockerHubWebHookPayload
		err := ctx.BindJSON(&payload)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if service, ok := appState.ServiceStates[payload.Repository.RepoName]; ok {
			log(fmt.Sprintf("Start deploying %v", payload.Repository.RepoName))
			appState.ServiceStates[payload.Repository.RepoName].Status = Deploying

			log(fmt.Sprintf("Pulling %v:latest from Docker Hub", payload.Repository.RepoName))

			imageName := latestImageName(payload.Repository.RepoName)
			err = pullDockerImage(imageName)
			if err != nil {
				ctx.String(http.StatusInternalServerError, "Fail to pull %s\n", imageName)
				return
			}

			log(fmt.Sprintf("Removing exisiting containers for %s", payload.Repository.RepoName))
			err := removeDockerContainers(payload.Repository.RepoName)
			if err != nil {
				ctx.String(http.StatusInternalServerError, "Fail to remove docker containers for %s\n", payload.Repository.RepoName)
				return
			}

			log(fmt.Sprintf("Launching a new container for %s", imageName))
			err = runDockerContainer(imageName, service.ServiceConfig.DockerRunArgs...)
			if err != nil {
				ctx.String(http.StatusInternalServerError, "Fail to run Docker container for %s\n", imageName)
				return
			}

			appState.ServiceStates[payload.Repository.RepoName].DeployedAt = time.Now()
			appState.ServiceStates[payload.Repository.RepoName].Status = Running
			log(fmt.Sprintf("Container for %s is now up and running", imageName))
			ctx.Writer.WriteHeader(http.StatusNoContent)
			return
		}
		log(fmt.Sprintf("No configuration for %v", payload.Repository.RepoName))
		ctx.Writer.WriteHeader(http.StatusBadRequest)
	}
}
