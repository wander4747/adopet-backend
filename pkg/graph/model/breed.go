package model

import (
	"strconv"

	"github.com/wander4747/adopet-backend/pkg/entity"
)

type Breed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewBreed(breed entity.Breed) *Breed {
	return &Breed{
		ID:   strconv.FormatInt(int64(breed.ID), 10),
		Name: breed.Name,
	}
}
