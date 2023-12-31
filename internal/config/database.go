package config

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"os"
	"path"
)

const (
	Postgres  = "postgres"
	Mysql     = "mysql"
	Sqlite    = "sqlite"
	SqlServer = "sqlserver"
)

func NewDatabaseConn() (*gorm.DB, error) {
	conf, err := NewAppConfig()
	if err != nil {
		return nil, err
	}

	dialector, err := getDatabaseDialector(conf)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDatabaseDialector(conf *AppConfig) (gorm.Dialector, error) {
	dbConf := conf.Database

	switch dbConf.DBDriver {
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			dbConf.Host, dbConf.Username, dbConf.Password, dbConf.Name, dbConf.Port, dbConf.TimeZone)
		return postgres.Open(dsn), nil
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Name)
		return mysql.Open(dsn), nil
	case "sqlserver":
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Name)
		return sqlserver.Open(dsn), nil
	case "sqlite":
		workDir, _ := os.Getwd()
		confPath := path.Join(workDir, "../../")
		_, err := os.Stat(confPath + "/" + dbConf.Name + ".db")
		if err != nil {
			return nil, err
		}
		return sqlite.Open(confPath + "/" + dbConf.Name + ".db"), nil
	default:
		return nil, errors.New("database is not supported")
	}
}
