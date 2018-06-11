package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"io/ioutil"
	"github.com/segmentio/objconv/json"
	"log"
)

func deployHandler(serviceStates serviceStates, errLogger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		b, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			errLogger.Printf("%+v", err)
			ctx.String(http.StatusInternalServerError, "Fail to parse JSON\n")
			return
		}
		log.Println(string(b))

		var payload DockerHubWebHookPayload

		err = json.Unmarshal(b, &payload)

		err = ctx.BindJSON(&payload)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if service, ok := serviceStates[payload.Repository.RepoName]; ok {
			log.Printf("Start deploying %v", payload.Repository.RepoName)
			serviceStates[payload.Repository.RepoName].Status = Deploying

			log.Printf("Pulling %v:latest from Docker Hub", payload.Repository.RepoName)

			imageName := latestImageName(payload.Repository.RepoName)
			log.Printf("Removing exisiting containers for %s", imageName)
			err := removeDockerContainers(imageName)
			if err != nil {
				errLogger.Printf("%+v", err)
				ctx.String(http.StatusInternalServerError, "Fail to remove docker containers for %s\n", imageName)
				return
			}

			err = pullDockerImage(imageName)
			if err != nil {
				errLogger.Printf("%+v", err)
				ctx.String(http.StatusInternalServerError, "Fail to pull %s\n", imageName)
				return
			}

			log.Printf("Launching a new container for %s", imageName)
			err = runDockerContainer(imageName, service.ServiceConfig.DockerRunArgs...)
			if err != nil {
				errLogger.Printf("%+v", err)
				ctx.String(http.StatusInternalServerError, "Fail to run Docker container for %s\n", imageName)
				return
			}

			serviceStates[payload.Repository.RepoName].DeployedAt = time.Now()
			serviceStates[payload.Repository.RepoName].Status = Running
			log.Printf("Container for %s is now up and running", imageName)
			ctx.Writer.WriteHeader(http.StatusNoContent)
			return
		}
		log.Printf("No configuration for %v", payload.Repository.RepoName)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
	}
}
