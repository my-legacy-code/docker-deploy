package core

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_latestImageName(t *testing.T) {
	assert.Equal(t, "org1/app1:latest", latestImageName("org1/app1"))
}

func Test_pullDockerImage(t *testing.T) {
	imageName := latestImageName("alpine")
	assert.Nil(t, pullDockerImage(imageName))
}

func Test_runDockerContainer(t *testing.T) {
	imageName := latestImageName("alpine")
	assert.Nil(t, runDockerContainer(imageName))
}

func Test_getContainerIds(t *testing.T) {
	imageName := latestImageName("alpine")
	assert.Nil(t, runDockerContainer(imageName))
	assert.Nil(t, runDockerContainer(imageName))

	containerIds, err := getContainerIds(imageName)
	assert.Nil(t, err)
	assert.True(t, len(containerIds) > 2)
}

func Test_removeDockerContainers(t *testing.T) {
	imageName := latestImageName("alpine")

	assert.Nil(t, runDockerContainer(imageName))
	assert.Nil(t, runDockerContainer(imageName))
	containerIds, err := getContainerIds(imageName)
	assert.Nil(t, err)
	assert.True(t, len(containerIds) > 1)

	assert.Nil(t, removeDockerContainers(imageName))
	containerIds, err = getContainerIds(imageName)
	assert.Nil(t, err)
	assert.Nil(t, containerIds)
}

func TestCheckContainerStatus(t *testing.T) {
	imageName := latestImageName("nginx")
	assert.Nil(t, removeDockerContainers(imageName))

	containerIds, err := getContainerIds(imageName)
	assert.True(t, len(containerIds) == 0)

	assert.Nil(t, runDockerContainer(imageName))
	containerIds, err = getContainerIds(imageName)
	assert.Nil(t, err)

	assert.True(t, len(containerIds) == 1)

	isRunning, err := isContainerRunning(containerIds[0])
	assert.Nil(t, err)
	assert.True(t, isRunning)

	assert.Nil(t, stopDockerContainer(containerIds[0]))

	isRunning, err = isContainerRunning(containerIds[0])
	assert.Nil(t, err)
	assert.False(t, isRunning)

	containerIds, err = getContainerIds(imageName)
	assert.Nil(t, err)

	assert.True(t, len(containerIds) == 1)
}
