package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"os/exec"
)

type WebHookPayload struct {
	PushData    PushData   `json:"push_data"`
	CallbackUrl string     `json:"callback_url"`
	Repository  Repository `json:"repository"`
}

type PushData struct {
	PushAt int      `json:"pushed_at"`
	Images []string `json:"images"`
	Tag    string   `json:"tag"`
	Pusher string   `json:"pusher"`
}

type Repository struct {
	CommentCount    int     `json:"comment_count"`
	DateCreated     float32 `json:"date_created"`
	Description     string  `json:"description"`
	Dockerfile      string  `json:"dockerfile"`
	FullDescription string  `json:"full_description"`
	IsOfficial      bool    `json:"is_official"`
	IsPrivate       bool    `json:"is_private"`
	IsTrusted       bool    `json:"is_trusted"`
	Name            string  `json:"name"`
	Namespace       string  `json:"namespace"`
	Owner           string  `json:"owner"`
	RepoName        string  `json:"repo_name"`
	RepoUrl         string  `json:"repo_url"`
	StarCount       int     `json:"star_count"`
	Status          string  `json:"status"`
}

type Service struct {
	Name        string
	Status      Status
	RepoName    string
	Namespace   string
	ContainerId string
	RunOptions  string
	CMD         string
	DeployedAt  time.Time
}

type Status int

const (
	Running   Status = iota
	Deploying
	Stopped
)

func main() {
	services := map[string]*Service{
		"teamyapp/teamy_api": {
			"Teamy API",
			Stopped,
			"teamy_api",
			"teamyapp",
			"952a74e8d3fe",
			"-p 8081:8080",
			"",
			time.Now(),
		},
		"teamyapp/teamy_frontend": {
			"Teamy Web Frontend",
			Stopped,
			"teamy_frontend",
			"teamyapp",
			"",
			"",
			"",
			time.Now(),
		},
		"rivalrygame/rivalrygame_frontend": {
			"Rivalry Game Frontend",
			Stopped,
			"rivalrygame_frontend",
			"rivalrygame",
			"",
			"",
			"",
			time.Now(),
		},
	}

	router := gin.Default()
	router.Static("/assets", "./public")
	router.LoadHTMLGlob("views/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "services.html", gin.H{
			"services": services,
		})
	})

	router.POST("/deploy", func(ctx *gin.Context) {
		var payload WebHookPayload
		err := ctx.BindJSON(&payload)
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}

		if service, ok := services[payload.Repository.RepoName]; ok {
			fmt.Printf("[DOCKER DEPLOY] Start deploying %v\n", payload.Repository.RepoName)
			services[payload.Repository.RepoName].Status = Deploying

			fmt.Printf("[DOCKER DEPLOY] Pulling %v:latest from Docker Hub\n", payload.Repository.RepoName)

			imageName := payload.Repository.RepoName + ":latest"

			cmd := exec.Command("docker", "pull", imageName)

			err := cmd.Run()

			if err != nil {
				ctx.Writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			if service.ContainerId != "" {
				fmt.Printf("[DOCKER DEPLOY] Stopping container %v\n", service.ContainerId)
				cmd = exec.Command("docker", "stop", service.ContainerId)
				cmd.Run()

				fmt.Printf("[DOCKER DEPLOY] Removing container %v\n", service.ContainerId)
				cmd = exec.Command("docker", "rm", service.ContainerId)
				cmd.Run()
			}

			fmt.Printf("[DOCKER DEPLOY] Running a new contrainer for %v\n", imageName)
			cmd = exec.Command("docker", "run", "-id", service.ContainerId)
			containerId, _ := cmd.CombinedOutput()
			services[payload.Repository.RepoName].ContainerId = string(containerId)

			if service.CMD != "" {
				fmt.Printf("[DOCKER DEPLOY] Running %v inside container\n", service.ContainerId)
			}

			ctx.Writer.WriteHeader(http.StatusNoContent)
			return
		}

		fmt.Printf("[DOCKER DEPLOY] No configuration for %v\n", payload.Repository.RepoName)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
	})
	router.Run(":8080")
}
