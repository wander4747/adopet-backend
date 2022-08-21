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

func TestNewBreed(t *testing.T) {
	config, _ := config.NewMockConfig()
	service := NewBreed(*config)
	require.NotNil(t, service)
}

func Test_breed_FindByAnimalID(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := repository.NewMockBreed(ctrl)
	service := breed{
		repository: repositoryMock,
	}

	t.Run("Success", func(t *testing.T) {
		expected := []*entity.Breed{
			{ID: 1, Name: "Dalmata"},
		}
		repositoryMock.EXPECT().FindByAnimalID(ctx, 1).Return(expected, nil)

		got, err := service.FindByAnimalID(ctx, 1)

		require.Equal(t, expected, got)
		require.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		repositoryMock.EXPECT().FindByAnimalID(ctx, 1).Return(nil, errors.New("fail"))

		got, err := service.FindByAnimalID(ctx, 1)

		require.Error(t, err)
		require.Nil(t, got)
	})
}
