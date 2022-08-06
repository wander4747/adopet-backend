package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/wander4747/adopet-backend/pkg/graph/generated"
	"github.com/wander4747/adopet-backend/pkg/graph/model"
)

// Animals is the resolver for the animals field.
func (r *queryResolver) Animals(ctx context.Context) ([]*model.Animal, error) {
	animals, err := r.Services.AnimalService.All(ctx)
	if err != nil {
		return nil, err
	}

	var animalCollection []*model.Animal
	for _, a := range animals {
		animalCollection = append(animalCollection, model.NewAnimal(*a))
	}

	return animalCollection, nil
}

// States is the resolver for the states field.
func (r *queryResolver) States(ctx context.Context) ([]*model.State, error) {
	states, err := r.Services.StateService.All(ctx)
	if err != nil {
		return nil, err
	}

	var stateCollection []*model.State
	for _, s := range states {
		stateCollection = append(stateCollection, model.NewState(*s))
	}

	return stateCollection, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
