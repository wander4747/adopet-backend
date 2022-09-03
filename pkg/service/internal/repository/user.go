//go:generate mockgen -destination=../../../../test/mock/repository/user.go -package=repository -source=$GOFILE repository

package repository

import (
	"context"

	"github.com/wander4747/adopet-backend/pkg/infrastructure/cache"

	"github.com/jmoiron/sqlx"

	"github.com/nleof/goyesql"

	// nolint
	_ "embed"

	"github.com/wander4747/adopet-backend/config"
	"github.com/wander4747/adopet-backend/pkg/entity"
)

//go:embed queries/user.sql
var userQueries []byte

type User interface {
	Create(ctx context.Context, user entity.User) (*entity.User, error)
}

type user struct {
	db      *sqlx.DB
	cache   cache.CacheInterface
	queries goyesql.Queries
}

func (u user) Create(_ context.Context, user entity.User) (us *entity.User, err error) {
	us = &user

	passwordHash, err := us.GeneratePassword(us.Password)
	if err != nil {
		return nil, err
	}

	us.Password = string(passwordHash)
	
	row, err := u.db.NamedExec(u.queries["create"], us)
	if err != nil {
		return nil, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return nil, err
	}

	us.ID = int(id)

	return us, nil
}

func NewUser(config config.Config) User {
	return &user{
		db:      config.DB,
		cache:   config.Cache,
		queries: goyesql.MustParseBytes(userQueries),
	}
}
