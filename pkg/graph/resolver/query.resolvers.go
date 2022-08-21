package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

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

func (r *queryResolver) Cities(ctx context.Context, stateID string) ([]*model.City, error) {
	ID, err := strconv.Atoi(stateID)
	if err != nil {
		return nil, err
	}

	cities, err := r.Services.CityService.FindByStateID(ctx, ID)
	if err != nil {
		return nil, err
	}

	var citiesCollection []*model.City
	for _, c := range cities {
		citiesCollection = append(citiesCollection, model.NewCity(*c))
	}

	return citiesCollection, nil
}

func (r *queryResolver) Breeds(ctx context.Context, animalID string) ([]*model.Breed, error) {
	ID, err := strconv.Atoi(animalID)
	if err != nil {
		return nil, err
	}

	breeds, err := r.Services.BreedService.FindByAnimalID(ctx, ID)
	if err != nil {
		return nil, err
	}

	var breedsCollection []*model.Breed
	for _, b := range breeds {
		breedsCollection = append(breedsCollection, model.NewBreed(*b))
	}

	return breedsCollection, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
