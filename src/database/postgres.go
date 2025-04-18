package database

import (
	"fmt"
	"snapp/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbClient *gorm.DB

func InitDB(cfg *config.Config) error {
	var err error
	cnn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.Port, cfg.Postgres.Sslmode,
	)
	DbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}
	database, _:= DbClient.DB()
	err = database.Ping()
	if err != nil {
		return err
	}

	database.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	database.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	database.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifeTime * time.Minute)

	return nil
}

func GetDB() *gorm.DB {
	return DbClient
}

func CloseDB() {
	con, _:= DbClient.DB()
	con.Close()
}