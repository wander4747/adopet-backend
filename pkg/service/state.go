//go:generate go run -mod=mod github.com/golang/mock/mockgen -package=service -source=$GOFILE -destination=../../test/mock/service/state.go

package service

import (
	"context"

	"github.com/wander4747/adopet-backend/config"
	"github.com/wander4747/adopet-backend/pkg/entity"
	"github.com/wander4747/adopet-backend/pkg/service/internal/repository"
)

type State interface {
	All(ctx context.Context) ([]*entity.State, error)
}

type state struct {
	repository repository.State
}

func NewState(config config.Config) State {
	return &state{
		repository: repository.NewState(config),
	}
}

func (s state) All(ctx context.Context) ([]*entity.State, error) {
	return s.repository.All(ctx)
}
