//go:generate mockgen -destination=../../../../test/mock/repository/animal.go -package=repository -source=$GOFILE repository

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

//go:embed queries/animal.sql
var animalQueries []byte

type Animal interface {
	All(ctx context.Context) ([]*entity.Animal, error)
}

type animal struct {
	db      *sqlx.DB
	cache   cache.CacheInterface
	queries goyesql.Queries
}

func NewAnimal(config config.Config) Animal {
	return &animal{
		db:      config.DB,
		cache:   config.Cache,
		queries: goyesql.MustParseBytes(animalQueries),
	}
}

func (a animal) All(ctx context.Context) ([]*entity.Animal, error) {
	var animals []*entity.Animal

	err := a.cache.GetSet(ctx, a.AllCacheKey(), &animals, func() (interface{}, error) {
		err := a.db.Select(&animals, a.queries["all"])
		if err != nil {
			return nil, err
		}

		return &animals, nil
	}, cache.OneWeek)

	return animals, err
}

func (a animal) AllCacheKey() string {
	return "animal-all"
}
