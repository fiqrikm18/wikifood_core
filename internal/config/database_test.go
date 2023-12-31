package config

import (
	"bytes"
	"github.com/spf13/viper"
	"testing"

	"github.com/stretchr/testify/assert"
)

var confMock = []byte(`
port: 8080
database:
  username: postgres
  password: postgres
  host: 127.0.0.1
  port: 5432
  db_name: wikifood
  db_timezone: Asia/Jakarta
  db_driver: postgres
`)

func TestNewDatabaseConnection(t *testing.T) {
	dbConn, err := NewDatabaseConn()
	assert.Nil(t, err)
	assert.NotNil(t, dbConn)

	appConf := AppConfig{}
	err = viper.ReadConfig(bytes.NewBuffer(confMock))
	assert.Nil(t, err)

	err = viper.Unmarshal(&appConf)
	assert.Nil(t, err)
	dbConn, err = NewDatabaseConn(WithCustomConfig(&appConf))
	assert.Nil(t, err)
	assert.NotNil(t, dbConn)

	appConf.Database.DBDriver = "test"
	dbConn, err = NewDatabaseConn(WithCustomConfig(&appConf))
	assert.NotNil(t, err)
	assert.Nil(t, dbConn)

	appConf.Database.DBDriver = Postgres
	appConf.Database.Name = "test"
	dbConn, err = NewDatabaseConn(WithCustomConfig(&appConf))
	assert.NotNil(t, err)
	assert.Nil(t, dbConn)
}

func TestGetDBDialector(t *testing.T) {
	appConf := AppConfig{}
	err := viper.ReadConfig(bytes.NewBuffer(confMock))
	assert.Nil(t, err)

	err = viper.Unmarshal(&appConf)
	assert.Nil(t, err)

	appConf.Database.DBDriver = Postgres
	dialector, err := getDatabaseDialector(&appConf)
	assert.Nil(t, err)
	assert.NotNil(t, dialector)

	appConf.Database.DBDriver = Mysql
	dialector, err = getDatabaseDialector(&appConf)
	assert.Nil(t, err)
	assert.NotNil(t, dialector)

	appConf.Database.DBDriver = SqlServer
	dialector, err = getDatabaseDialector(&appConf)
	assert.Nil(t, err)
	assert.NotNil(t, dialector)

	appConf.Database.DBDriver = Sqlite
	dialector, err = getDatabaseDialector(&appConf)
	assert.Nil(t, err)
	assert.NotNil(t, dialector)

	appConf.Database.Name = "test_db"
	dialector, err = getDatabaseDialector(&appConf)
	assert.NotNil(t, err)
	assert.Nil(t, dialector)

	appConf.Database.DBDriver = "unsupportedDB"
	dialector, err = getDatabaseDialector(&appConf)
	assert.NotNil(t, err)
	assert.Nil(t, dialector)
}
