package core

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func Test_getEnv(t *testing.T) {
	assert.Equal(t, "defaultKey", getEnv("KEY", "defaultKey"))
	os.Setenv("KEY", "newKey")
	assert.Equal(t, "newKey", getEnv("KEY", "defaultKey"))
}
