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

func TestNewUser(t *testing.T) {
	config, _ := NewMockConfig()
	repository := NewUser(*config)
	require.NotNil(t, repository)
}

func Test_user_Create(t *testing.T) {
	ctx := context.Background()
	queries := goyesql.MustParseBytes(userQueries)
	config, mock := NewMockConfig()

	repository := user{
		db:      config.DB,
		queries: queries,
		cache:   config.Cache,
	}

	entity := entity.User{
		ID:          1,
		Name:        "Wander",
		Email:       "wander.douglas14@gmail.com",
		CityID:      1,
		StateID:     1,
		Phone:       nil,
		Description: nil,
		Password:    "1234",
		Photo:       nil,
		Type:        "normal",
		ShowEmail:   false,
		ShowPhone:   false,
		ShowAddress: false,
		Address:     "Rua 123",
		Number:      nil,
		ZipCode:     "12345-678",
		Complement:  nil,
		TotalPets:   nil,
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO users").WithArgs(entity.CityID, entity.StateID, entity.Name, entity.Email,
			entity.Phone, entity.Description, entity.Password, entity.Photo, entity.Type, entity.ShowEmail, entity.ShowPhone, entity.ShowAddress,
			entity.Address, entity.Number, entity.ZipCode, entity.Complement, entity.TotalPets).WillReturnResult(sqlmock.NewResult(1, 1))

		got, err := repository.Create(ctx, entity)
		require.Equal(t, &entity, got)
		require.Nil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO users").WithArgs(entity.CityID, entity.StateID, entity.Name, entity.Email,
			entity.Phone, entity.Description, entity.Password, entity.Photo, entity.Type, entity.ShowEmail, entity.ShowPhone, entity.ShowAddress,
			entity.Address, entity.Number, entity.ZipCode, entity.Complement, entity.TotalPets).WillReturnError(errors.New("fail"))

		_, err := repository.Create(ctx, entity)
		require.Error(t, err)
	})
}
