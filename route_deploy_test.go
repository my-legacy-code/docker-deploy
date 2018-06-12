package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"bytes"
	"net/http"
)

func setupDeployRoute(t *testing.T) (*gin.Engine, *httptest.ResponseRecorder, DockerHubWebHookPayload) {

	configFilename := "testdata/config_test.json"
	serviceConfig, err := loadConfig(configFilename)
	assert.Nil(t, err)

	setupLogger()
	errLogger := makeErrLogger()
	appState := initAppState(serviceConfig, errLogger)
	router := setupRouter(appState, errLogger)

	recorder := httptest.NewRecorder()

	payload, err := loadDockerHubPayload("testdata/dockerhub_payload.json")
	assert.Nil(t, err)

	return router, recorder, payload
}

func TestDeployService(t *testing.T) {
	router, recorder, payload := setupDeployRoute(t)

	b, err := json.Marshal(payload)
	assert.Nil(t, err)

	req, _ := http.NewRequest("POST", "/api/deploy", bytes.NewReader(b))
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNoContent, recorder.Code)
	assert.Equal(t, "", recorder.Body.String())
}
