package service

import (
	"context"
	"errors"
	"testing"

	"github.com/wander4747/adopet-backend/test/mock/repository"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/wander4747/adopet-backend/config"
	"github.com/wander4747/adopet-backend/pkg/entity"
)

func TestNewState(t *testing.T) {
	config, _ := config.NewMockConfig()
	service := NewState(*config)
	require.NotNil(t, service)
}

func Test_state_All(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := repository.NewMockState(ctrl)
	service := state{
		repository: repositoryMock,
	}

	t.Run("Success", func(t *testing.T) {
		expected := []*entity.State{
			{ID: 1, Name: "Minas Gerais", Initials: "MG"},
			{ID: 2, Name: "São Paulo", Initials: "SP"},
		}
		repositoryMock.EXPECT().All(ctx).Return(expected, nil)

		got, err := service.All(ctx)

		require.Equal(t, expected, got)
		require.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		repositoryMock.EXPECT().All(ctx).Return(nil, errors.New("fail"))

		got, err := service.All(ctx)

		require.Error(t, err)
		require.Nil(t, got)
	})
}
