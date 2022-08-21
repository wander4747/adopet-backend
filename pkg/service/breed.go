//go:generate go run -mod=mod github.com/golang/mock/mockgen -package=service -source=$GOFILE -destination=../../test/mock/service/breed.go

package service

import (
	"context"

	"github.com/wander4747/adopet-backend/config"
	"github.com/wander4747/adopet-backend/pkg/entity"
	"github.com/wander4747/adopet-backend/pkg/service/internal/repository"
)

type Breed interface {
	FindByAnimalID(ctx context.Context, stateID int) ([]*entity.Breed, error)
}

type breed struct {
	repository repository.Breed
}

func NewBreed(config config.Config) Breed {
	return &breed{
		repository: repository.NewBreed(config),
	}
}

func (b breed) FindByAnimalID(ctx context.Context, animalID int) ([]*entity.Breed, error) {
	return b.repository.FindByAnimalID(ctx, animalID)
}
