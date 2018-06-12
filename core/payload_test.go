package core

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadDockerHubPayload(t *testing.T) {
	payload, err := loadDockerHubPayload("testdata/dockerhub_payload.json")
	assert.Nil(t, err)

	expectedPayload := DockerHubWebHookPayload{
		PushData: PushData{
			PushAt: 1528486026,
			Images: []string{},
			Tag:    "latest",
			Pusher: "teamyapp",
		},
		CallbackUrl: "https://registry.hub.docker.com/u/teamyapp/teamy_api/hook/22g0ag1d5if044a3edbi10c2if35adi0a/",
		Repository: Repository{
			Status:          "Active",
			Description:     "Backend for Teamy, an effective task manager for start-ups and small companies.",
			IsTrusted:       false,
			FullDescription: "",
			RepoUrl:         "https://hub.docker.com/r/teamyapp/teamy_api",
			Owner:           "teamyapp",
			IsOfficial:      false,
			IsPrivate:       false,
			Name:            "teamy_api",
			Namespace:       "teamyapp",
			StarCount:       0,
			CommentCount:    0,
			DateCreated:     1528357169,
			RepoName:        "teamyapp/teamy_api",
		},
	}

	assert.Equal(t, expectedPayload, payload)
}
