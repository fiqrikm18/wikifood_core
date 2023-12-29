package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAppConfig(t *testing.T) {
	conf, err := NewAppConfig()
	assert.Nil(t, err)
	assert.NotNil(t, conf)
	assert.NotEqual(t, conf.Port, "")

	conf, err = NewAppConfig(WithFilePath("path"))
	assert.NotNil(t, err)
	assert.Nil(t, conf)
}
