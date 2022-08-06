package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/require"

	"github.com/nleof/goyesql"
	"github.com/wander4747/adopet-backend/pkg/entity"
)

func TestNewState(t *testing.T) {
	config, _ := NewMockConfig()
	repository := NewState(*config)
	require.NotNil(t, repository)
}

func Test_state_All(t *testing.T) {
	ctx := context.Background()
	queries := goyesql.MustParseBytes(stateQueries)
	config, mock := NewMockConfig()

	repository := state{
		db:      config.DB,
		queries: queries,
		cache:   config.Cache,
	}

	t.Run("Success", func(t *testing.T) {
		expected := []*entity.State{
			{ID: 1, Name: "Minas Gerais", Initials: "MG"},
			{ID: 2, Name: "São Paulo", Initials: "SP"},
		}
		rows := sqlmock.NewRows([]string{"id", "name", "initials"}).AddRow(1, "Minas Gerais", "MG").AddRow(2, "São Paulo", "SP")
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
