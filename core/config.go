package core

import (
	"time"
	"io/ioutil"
)

import (
	"github.com/pkg/errors"
	"github.com/segmentio/objconv/json"
)

type Service struct {
	Status        Status        `json:"status"`
	DeployedAt    time.Time     `json:"deployed_at"`
	ServiceConfig ServiceConfig `json:"service_config"`
}

type ServiceConfig struct {
	Name          string   `json:"name"`
	Namespace     string   `json:"namespace"`
	DockerRunArgs []string `json:"docker_run_args"`
}

type configs map[string]ServiceConfig

func loadConfig(fileName string) (configs, error) {

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	config := make(configs)

	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return config, nil
}
