package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_initServiceState(t *testing.T) {
	testConfigs := configs{
		"org1/app1": ServiceConfig{
			Name:          "app1",
			Namespace:     "org1",
			DockerRunArgs: []string{"arg1", "arg2"},
		},
		"org2/app2": ServiceConfig{
			Name:          "app2",
			Namespace:     "org2",
			DockerRunArgs: []string{"arg3", "arg4"},
		},
	}

	testServiceStatus := initServiceState(testConfigs)
	assert.Equal(t,
		ServiceStates{
			"org1/app1": &Service{
				Status: Initialized,
				ServiceConfig: ServiceConfig{
					Name:          "app1",
					Namespace:     "org1",
					DockerRunArgs: []string{"arg1", "arg2"},
				},
			},
			"org2/app2": &Service{
				Status: Initialized,
				ServiceConfig: ServiceConfig{
					Name:          "app2",
					Namespace:     "org2",
					DockerRunArgs: []string{"arg3", "arg4"},
				},
			},
		},
		testServiceStatus)
}
