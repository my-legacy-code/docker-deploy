package core

import (
	"io/ioutil"
	"github.com/pkg/errors"
	"github.com/segmentio/objconv/json"
)

type DockerHubWebHookPayload struct {
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

func loadDockerHubPayload(mockPayloadFilename string) (DockerHubWebHookPayload, error) {
	payload := new(DockerHubWebHookPayload)

	b, err := ioutil.ReadFile(mockPayloadFilename)
	if err != nil {
		return *payload, errors.WithStack(err)
	}

	err = json.Unmarshal(b, payload)

	if err != nil {
		return *payload, errors.WithStack(err)
	}

	return *payload, nil
}
