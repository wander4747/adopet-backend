package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/wander4747/adopet-backend/pkg/graph/generated"
	"github.com/wander4747/adopet-backend/pkg/graph/model"
)

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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
