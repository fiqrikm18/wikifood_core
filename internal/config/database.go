package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConn() (*gorm.DB, error) {
	conf, err := NewAppConfig()
	if err != nil {
		return nil, err
	}

	dbConf := conf.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbConf.Host, dbConf.Username, dbConf.Password, dbConf.Name, dbConf.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
