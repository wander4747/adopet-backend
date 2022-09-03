package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/wander4747/adopet-backend/pkg/graph/generated"
	"github.com/wander4747/adopet-backend/pkg/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := r.Services.UserService.Create(ctx, model.NewUserEntity(input))
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	return model.NewUserModel(user), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
