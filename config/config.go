package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/wander4747/adopet-backend/pkg/infrastructure/database"
)

type Config struct {
	DB *sqlx.DB
}

func NewConfig() *Config {
	db := newMysql()

	return &Config{
		DB: db,
	}
}

func newMysql() *sqlx.DB {
	db, err := database.NewMysql()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
