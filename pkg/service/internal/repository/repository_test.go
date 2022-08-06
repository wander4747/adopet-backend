package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/wander4747/adopet-backend/config"
	"github.com/wander4747/adopet-backend/pkg/infrastructure/cache"
)

func NewMockConfig() (*config.Config, sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	cache := cache.NewLocal()

	c := config.Config{
		DB:    sqlxDB,
		Cache: cache,
	}

	return &c, mock
}
