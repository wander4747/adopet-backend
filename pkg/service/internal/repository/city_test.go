package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	"github.com/nleof/goyesql"
	"github.com/wander4747/adopet-backend/pkg/entity"
)

func TestNewCity(t *testing.T) {
	config, _ := NewMockConfig()
	repository := NewCity(*config)
	require.NotNil(t, repository)
}

func Test_city_FindByStateID(t *testing.T) {
	ctx := context.Background()
	queries := goyesql.MustParseBytes(cityQueries)
	config, mock := NewMockConfig()

	repository := city{
		db:      config.DB,
		queries: queries,
		cache:   config.Cache,
	}

	t.Run("Success", func(t *testing.T) {
		expected := []*entity.City{
			{ID: 1, Name: "Belo Horizonte", StateID: 1},
		}
		rows := sqlmock.NewRows([]string{"id", "name", "state_id"}).AddRow(1, "Belo Horizonte", 1)
		mock.ExpectQuery(regexp.QuoteMeta(queries["by-state-id"])).WithArgs(1).WillReturnRows(rows)

		got, err := repository.FindByStateID(ctx, 1)
		require.Equal(t, expected, got)
		require.Nil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		repository.cache.Flush(ctx)
		mock.ExpectQuery(regexp.QuoteMeta(queries["by-state-id"])).WillReturnError(errors.New("fail"))

		_, err := repository.FindByStateID(ctx, 1)
		require.Error(t, err)
	})
}
