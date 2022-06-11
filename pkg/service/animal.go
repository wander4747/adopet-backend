//go:generate go run -mod=mod github.com/golang/mock/mockgen -package=service -source=$GOFILE -destination=../../test/mock/service/animal.go

package service

import (
	"context"

	"github.com/wander4747/adopet-backend/config"
	"github.com/wander4747/adopet-backend/pkg/entity"
	"github.com/wander4747/adopet-backend/pkg/service/internal/repository"
)

type Animal interface {
	All(ctx context.Context) ([]*entity.Animal, error)
}

type animal struct {
	repository repository.Animal
}

func NewAnimal(config config.Config) Animal {
	return &animal{
		repository: repository.NewAnimal(config),
	}
}

func (a animal) All(ctx context.Context) ([]*entity.Animal, error) {
	return a.repository.All(ctx)
}
