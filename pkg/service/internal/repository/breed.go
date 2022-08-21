//go:generate mockgen -destination=../../../../test/mock/repository/breed.go -package=repository -source=$GOFILE repository

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

//go:embed queries/breed.sql
var breedQueries []byte

type Breed interface {
	FindByAnimalID(ctx context.Context, stateID int) ([]*entity.Breed, error)
}

type breed struct {
	db      *sqlx.DB
	cache   cache.CacheInterface
	queries goyesql.Queries
}

func NewBreed(config config.Config) Breed {
	return &breed{
		db:      config.DB,
		cache:   config.Cache,
		queries: goyesql.MustParseBytes(breedQueries),
	}
}

func (b breed) FindByAnimalID(ctx context.Context, animalID int) (breeds []*entity.Breed, err error) {
	err = b.cache.GetSet(ctx, b.byAnimalIdCacheKey(animalID), &breeds, func() (interface{}, error) {
		err = b.db.Select(&breeds, b.queries["by-animal-id"], animalID)
		if err != nil {
			return nil, err
		}

		return &breeds, nil
	}, cache.OneWeek)

	return breeds, err
}

func (b breed) byAnimalIdCacheKey(animalId int) string {
	return fmt.Sprintf("breed-animalid-%d", animalId)
}
