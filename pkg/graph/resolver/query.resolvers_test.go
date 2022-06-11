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
