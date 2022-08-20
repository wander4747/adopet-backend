//go:generate mockgen -destination=../../../../test/mock/repository/city.go -package=repository -source=$GOFILE repository

package repository

import (
	"context"
	"fmt"

	"github.com/wander4747/adopet-backend/pkg/infrastructure/cache"

	"github.com/jmoiron/sqlx"

	"github.com/nleof/goyesql"

	// nolint
	_ "embed"

	"github.com/wander4747/adopet-backend/config"
	"github.com/wander4747/adopet-backend/pkg/entity"
)

//go:embed queries/city.sql
var cityQueries []byte

type City interface {
	FindByStateID(ctx context.Context, stateID int) ([]*entity.City, error)
}

type city struct {
	db      *sqlx.DB
	cache   cache.CacheInterface
	queries goyesql.Queries
}

func NewCity(config config.Config) City {
	return &city{
		db:      config.DB,
		cache:   config.Cache,
		queries: goyesql.MustParseBytes(cityQueries),
	}
}

func (c city) FindByStateID(ctx context.Context, stateID int) (cities []*entity.City, err error) {
	err = c.cache.GetSet(ctx, c.ByStateIdCacheKey(stateID), &cities, func() (interface{}, error) {
		err = c.db.Select(&cities, c.queries["by-state-id"], stateID)
		if err != nil {
			return nil, err
		}

		return &cities, nil
	}, cache.OneWeek)

	return cities, err
}

func (c city) ByStateIdCacheKey(stateID int) string {
	return fmt.Sprintf("city-stateid-%d", stateID)
}
