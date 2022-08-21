package resolver

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/wander4747/adopet-backend/pkg/entity"
	"github.com/wander4747/adopet-backend/pkg/graph/model"
	"github.com/wander4747/adopet-backend/pkg/service"
	serviceMock "github.com/wander4747/adopet-backend/test/mock/service"
)

func Test_queryResolver_Animals(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	serviceMock := serviceMock.NewMockAnimal(ctrl)

	resolver := &Resolver{
		Services: service.All{
			AnimalService: serviceMock,
		},
	}

	t.Run("Success without data", func(t *testing.T) {
		var expected []*model.Animal
		serviceMock.EXPECT().All(ctx).
			Return([]*entity.Animal{}, nil)

		got, err := resolver.Query().Animals(ctx)
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})

	t.Run("Success with data", func(t *testing.T) {
		expected := []*model.Animal{
			{
				ID:   "1",
				Name: "Cão",
			},
		}

		expected2 := []*entity.Animal{
			{
				ID:   1,
				Name: "Cão",
			},
		}
		serviceMock.EXPECT().All(ctx).
			Return(expected2, nil)

		got, err := resolver.Query().Animals(ctx)
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})

	t.Run("Error", func(t *testing.T) {
		var expected []*model.Animal
		serviceMock.EXPECT().All(ctx).
			Return(nil, errors.New("fail"))

		got, err := resolver.Query().Animals(ctx)
		require.Error(t, err)
		require.Equal(t, expected, got)
	})
}

func Test_queryResolver_States(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	serviceMock := serviceMock.NewMockState(ctrl)

	resolver := &Resolver{
		Services: service.All{
			StateService: serviceMock,
		},
	}

	t.Run("Success without data", func(t *testing.T) {
		var expected []*model.State
		serviceMock.EXPECT().All(ctx).
			Return([]*entity.State{}, nil)

		got, err := resolver.Query().States(ctx)
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})

	t.Run("Success with data", func(t *testing.T) {
		expected := []*model.State{
			{
				ID:       "1",
				Name:     "Minas Gerais",
				Initials: "MG",
			},
		}

		expectedEntity := []*entity.State{
			{
				ID:       1,
				Name:     "Minas Gerais",
				Initials: "MG",
			},
		}

		serviceMock.EXPECT().All(ctx).
			Return(expectedEntity, nil)

		got, err := resolver.Query().States(ctx)
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})

	t.Run("Error", func(t *testing.T) {
		var expected []*model.State
		serviceMock.EXPECT().All(ctx).
			Return(nil, errors.New("fail"))

		got, err := resolver.Query().States(ctx)
		require.Error(t, err)
		require.Equal(t, expected, got)
	})
}

func Test_queryResolver_Cities(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	serviceMock := serviceMock.NewMockCity(ctrl)

	resolver := &Resolver{
		Services: service.All{
			CityService: serviceMock,
		},
	}

	t.Run("Success without data", func(t *testing.T) {
		var expected []*model.City
		serviceMock.EXPECT().FindByStateID(ctx, 1).
			Return([]*entity.City{}, nil)

		got, err := resolver.Query().Cities(ctx, "1")
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})

	t.Run("Success with data", func(t *testing.T) {
		expected := []*model.City{
			{ID: "1", Name: "Belo Horizonte", StateID: "1"},
		}

		expectedEntity := []*entity.City{
			{ID: 1, Name: "Belo Horizonte", StateID: 1},
		}

		serviceMock.EXPECT().FindByStateID(ctx, 1).
			Return(expectedEntity, nil)

		got, err := resolver.Query().Cities(ctx, "1")
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})

	t.Run("Error", func(t *testing.T) {
		var expected []*model.City
		serviceMock.EXPECT().FindByStateID(ctx, 1).
			Return(nil, errors.New("fail"))

		got, err := resolver.Query().Cities(ctx, "1")
		require.Error(t, err)
		require.Equal(t, expected, got)
	})

	t.Run("Error converter string to int", func(t *testing.T) {
		var expected []*model.City

		got, err := resolver.Query().Cities(ctx, "asdf")
		require.Error(t, err)
		require.Equal(t, expected, got)
	})
}

func Test_queryResolver_Breeds(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	serviceMock := serviceMock.NewMockBreed(ctrl)

	resolver := &Resolver{
		Services: service.All{
			BreedService: serviceMock,
		},
	}

	t.Run("Success without data", func(t *testing.T) {
		var expected []*model.Breed
		serviceMock.EXPECT().FindByAnimalID(ctx, 1).
			Return([]*entity.Breed{}, nil)

		got, err := resolver.Query().Breeds(ctx, "1")
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})

	t.Run("Success with data", func(t *testing.T) {
		expected := []*model.Breed{
			{ID: "1", Name: "Dalmata"},
		}

		expectedEntity := []*entity.Breed{
			{ID: 1, Name: "Dalmata"},
		}

		serviceMock.EXPECT().FindByAnimalID(ctx, 1).
			Return(expectedEntity, nil)

		got, err := resolver.Query().Breeds(ctx, "1")
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})

	t.Run("Error", func(t *testing.T) {
		var expected []*model.Breed
		serviceMock.EXPECT().FindByAnimalID(ctx, 1).
			Return(nil, errors.New("fail"))

		got, err := resolver.Query().Breeds(ctx, "1")
		require.Error(t, err)
		require.Equal(t, expected, got)
	})

	t.Run("Error converter string to int", func(t *testing.T) {
		var expected []*model.Breed

		got, err := resolver.Query().Breeds(ctx, "asdf")
		require.Error(t, err)
		require.Equal(t, expected, got)
	})
}
