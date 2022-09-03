//go:generate go run -mod=mod github.com/golang/mock/mockgen -package=service -source=$GOFILE -destination=../../test/mock/service/user.go

package service

import (
	"context"
	"github.com/wander4747/adopet-backend/config"

	"github.com/wander4747/adopet-backend/pkg/entity"
	"github.com/wander4747/adopet-backend/pkg/service/internal/repository"
)

type User interface {
	Create(ctx context.Context, user entity.User) (*entity.User, error)
}

type user struct {
	repository repository.User
}

func NewUser(config config.Config) User {
	return &user{
		repository: repository.NewUser(config),
	}
}

func (u user) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	return u.repository.Create(ctx, user)
}
