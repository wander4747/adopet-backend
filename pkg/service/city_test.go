package service

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/wander4747/adopet-backend/test/mock/repository"

	"testing"

	"github.com/wander4747/adopet-backend/config"
	"github.com/wander4747/adopet-backend/pkg/entity"
)

func TestNewCity(t *testing.T) {
	config, _ := config.NewMockConfig()
	service := NewCity(*config)
	require.NotNil(t, service)
}

func Test_city_FindByStateID(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := repository.NewMockCity(ctrl)
	service := city{
		repository: repositoryMock,
	}

	t.Run("Success", func(t *testing.T) {
		expected := []*entity.City{
			{ID: 1, Name: "Belo Horizonte", StateID: 1},
		}
		repositoryMock.EXPECT().FindByStateID(ctx, 1).Return(expected, nil)

		got, err := service.FindByStateID(ctx, 1)

		require.Equal(t, expected, got)
		require.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		repositoryMock.EXPECT().FindByStateID(ctx, 1).Return(nil, errors.New("fail"))

		got, err := service.FindByStateID(ctx, 1)

		require.Error(t, err)
		require.Nil(t, got)
	})
}
