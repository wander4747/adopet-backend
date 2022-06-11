package service

import (
	"context"
	"errors"
	"testing"

	"github.com/wander4747/adopet-backend/pkg/entity"

	"github.com/golang/mock/gomock"
	"github.com/wander4747/adopet-backend/test/mock/repository"

	"github.com/stretchr/testify/require"
	"github.com/wander4747/adopet-backend/config"
)

func TestNewAnimal(t *testing.T) {
	config, _ := config.NewMockConfig()
	service := NewAnimal(*config)
	require.NotNil(t, service)
}

func Test_animal_All(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := repository.NewMockAnimal(ctrl)
	service := animal{
		repository: repositoryMock,
	}

	t.Run("Success", func(t *testing.T) {
		expected := []*entity.Animal{
			{ID: 1, Name: "Cão"},
			{ID: 2, Name: "Gato"},
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
