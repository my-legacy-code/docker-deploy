package main

import (
	"os/exec"
	"strings"
	"github.com/pkg/errors"
	"fmt"
)

func downloadImage(imageName string) error {
	cmd := exec.Command("docker", "pull", imageName)
	err := cmd.Run()
	return errors.Wrapf(err, "downloadImage(%s) failed", imageName)
}

func run(imageName string) error {
	err := removeContainers(imageName)
	if err != nil {
		return errors.Wrapf(err, "run(%s) failed", imageName)
	}

	cmd := exec.Command("docker", "run", "-d", imageName)
	return errors.Wrapf(cmd.Run(), "run(%s) failed", imageName)
}

// removeContainers removes containers identified by the imageName
func removeContainers(imageName string) error {
	const acceptableError = "exit status 1"

	argStr := fmt.Sprintf("rm $(docker ps -a -q --filter ancestor=%s) --force", imageName)
	args := strings.Split(argStr, " ")

	// todo: collect system return code=1
	cmd := exec.Command("docker", args...)
	err := cmd.Run()
	if err != nil {
		switch err.Error() {
		case acceptableError:
			return nil
		default:
			return errors.Wrapf(cmd.Run(), "removeContainers(%s) failed", imageName)
		}
	}
	return nil
}
