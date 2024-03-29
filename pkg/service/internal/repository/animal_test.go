package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nleof/goyesql"
	"github.com/stretchr/testify/require"
	"github.com/wander4747/adopet-backend/pkg/entity"
)

func TestNewAnimal(t *testing.T) {
	config, _ := NewMockConfig()
	repository := NewAnimal(*config)
	require.NotNil(t, repository)
}

func Test_animal_All(t *testing.T) {
	ctx := context.Background()
	queries := goyesql.MustParseBytes(animalQueries)
	config, mock := NewMockConfig()

	repository := animal{
		db:      config.DB,
		queries: queries,
		cache:   config.Cache,
	}

	t.Run("Success", func(t *testing.T) {
		expected := []*entity.Animal{
			{ID: 1, Name: "Cão"},
			{ID: 2, Name: "Gato"},
		}
		rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Cão").AddRow(2, "Gato")
		mock.ExpectQuery(queries["all"]).WillReturnRows(rows)

		got, err := repository.All(ctx)
		require.Equal(t, expected, got)
		require.Nil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		repository.cache.Flush(ctx)
		mock.ExpectQuery(queries["all"]).WillReturnError(errors.New("fail"))

		_, err := repository.All(ctx)
		require.Error(t, err)
	})
}
