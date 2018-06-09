package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_loadConfig(t *testing.T) {
	testConfigs, err := loadConfig("testdata/config_test.json")
	assert.Nil(t, err)
	assert.IsType(t, configs{}, testConfigs)
}
