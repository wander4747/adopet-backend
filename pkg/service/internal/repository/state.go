//go:generate mockgen -destination=../../../../test/mock/repository/state.go -package=repository -source=$GOFILE repository

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

//go:embed queries/state.sql
var stateQueries []byte

type State interface {
	All(ctx context.Context) ([]*entity.State, error)
}

type state struct {
	db      *sqlx.DB
	cache   cache.CacheInterface
	queries goyesql.Queries
}

func NewState(config config.Config) State {
	return &state{
		db:      config.DB,
		cache:   config.Cache,
		queries: goyesql.MustParseBytes(stateQueries),
	}
}

func (s state) All(ctx context.Context) ([]*entity.State, error) {
	var states []*entity.State

	err := s.cache.GetSet(ctx, s.AllCacheKey(), &states, func() (interface{}, error) {
		err := s.db.Select(&states, s.queries["all"])
		if err != nil {
			return nil, err
		}

		return &states, nil
	}, cache.OneWeek)

	return states, err
}

func (s state) AllCacheKey() string {
	return "state-all"
}
