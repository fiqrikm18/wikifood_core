package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabaseConnection(t *testing.T) {
	_, err := NewDatabaseConn()
	assert.Nil(t, err)
}
