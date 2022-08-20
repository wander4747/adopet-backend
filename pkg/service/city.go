//go:generate go run -mod=mod github.com/golang/mock/mockgen -package=service -source=$GOFILE -destination=../../test/mock/service/city.go

package service

import (
	"context"

	"github.com/wander4747/adopet-backend/config"
	"github.com/wander4747/adopet-backend/pkg/entity"
	"github.com/wander4747/adopet-backend/pkg/service/internal/repository"
)

type City interface {
	FindByStateID(ctx context.Context, stateID int) ([]*entity.City, error)
}

type city struct {
	repository repository.City
}

func NewCity(config config.Config) City {
	return &city{
		repository: repository.NewCity(config),
	}
}

func (c city) FindByStateID(ctx context.Context, stateID int) ([]*entity.City, error) {
	return c.repository.FindByStateID(ctx, stateID)
}
