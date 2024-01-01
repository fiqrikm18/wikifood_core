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
	"time"
)

const (
	Postgres  = "postgres"
	Mysql     = "mysql"
	Sqlite    = "sqlite"
	SqlServer = "sqlserver"
)

type DBConnConf func(opt *DBConf)

type DBConf struct {
	AppConfig *AppConfig
}

func NewDatabaseConn(opts ...DBConnConf) (*gorm.DB, error) {
	var conf *AppConfig
	var err error
	conf, err = NewAppConfig()
	if len(opts) > 0 {
		dbConf := DBConf{}
		for _, opt := range opts {
			opt(&dbConf)
		}

		conf = dbConf.AppConfig
	}

	dialector, err := getDatabaseDialector(conf)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbConn, _ := db.DB()
	dbConn.SetMaxOpenConns(100)
	dbConn.SetMaxIdleConns(10)
	dbConn.SetConnMaxLifetime(time.Minute * 30)
	dbConn.SetConnMaxIdleTime(time.Minute * 30)

	return db, nil
}

func WithCustomConfig(config *AppConfig) DBConnConf {
	return func(opt *DBConf) {
		opt.AppConfig = config
	}
}

func getDatabaseDialector(conf *AppConfig) (gorm.Dialector, error) {
	dbConf := conf.Database

	switch dbConf.DBDriver {
	case Postgres:
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", dbConf.Host, dbConf.Username, dbConf.Password, dbConf.Name, dbConf.Port, dbConf.TimeZone)
		return postgres.Open(dsn), nil
	case Mysql:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Name)
		return mysql.Open(dsn), nil
	case SqlServer:
		dsn := fmt.Sprintf("sqlserver:%s:%s@%s:%s?database=%s", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Name)
		return sqlserver.Open(dsn), nil
	case Sqlite:
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
