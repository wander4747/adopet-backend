package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nleof/goyesql"
	"github.com/wander4747/adopet-backend/pkg/entity"

	"github.com/stretchr/testify/require"
)

func TestNewBreed(t *testing.T) {
	config, _ := NewMockConfig()
	repository := NewBreed(*config)
	require.NotNil(t, repository)
}

func Test_breed_FindByAnimalID(t *testing.T) {
	ctx := context.Background()
	queries := goyesql.MustParseBytes(breedQueries)
	config, mock := NewMockConfig()

	repository := breed{
		db:      config.DB,
		queries: queries,
		cache:   config.Cache,
	}

	t.Run("Success", func(t *testing.T) {
		expected := []*entity.Breed{
			{ID: 1, Name: "Dalmata"},
		}
		rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Dalmata")
		mock.ExpectQuery(regexp.QuoteMeta(queries["by-animal-id"])).WithArgs(1).WillReturnRows(rows)

		got, err := repository.FindByAnimalID(ctx, 1)
		require.Equal(t, expected, got)
		require.Nil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		repository.cache.Flush(ctx)
		mock.ExpectQuery(regexp.QuoteMeta(queries["by-animal-id"])).WillReturnError(errors.New("fail"))

		_, err := repository.FindByAnimalID(ctx, 2)
		require.Error(t, err)
	})
}
