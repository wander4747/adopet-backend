package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/wander4747/adopet-backend/config"
	"github.com/wander4747/adopet-backend/pkg/entity"
	"github.com/wander4747/adopet-backend/test/mock/repository"
)

func TestNewUser(t *testing.T) {
	config, _ := config.NewMockConfig()
	service := NewUser(*config)
	require.NotNil(t, service)
}

func Test_user_Create(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := repository.NewMockUser(ctrl)
	service := user{
		repository: repositoryMock,
	}

	entity := entity.User{
		ID:          0,
		Name:        "",
		Email:       "",
		CityID:      0,
		StateID:     0,
		Phone:       nil,
		Description: nil,
		Password:    "",
		Photo:       nil,
		Type:        "",
		ShowEmail:   false,
		ShowPhone:   false,
		ShowAddress: false,
		Address:     "",
		Number:      nil,
		ZipCode:     "",
		Complement:  nil,
		TotalPets:   nil,
	}

	t.Run("Success", func(t *testing.T) {
		repositoryMock.EXPECT().Create(ctx, entity).Return(&entity, nil)

		got, err := service.Create(ctx, entity)

		require.Equal(t, &entity, got)
		require.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		repositoryMock.EXPECT().Create(ctx, entity).Return(nil, errors.New("fail"))

		got, err := service.Create(ctx, entity)

		require.Error(t, err)
		require.Nil(t, got)
	})
}
