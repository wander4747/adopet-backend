package model

import (
	"strconv"

	"github.com/wander4747/adopet-backend/pkg/entity"
)

type Animal struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewAnimal(animal entity.Animal) *Animal {
	return &Animal{
		ID:   strconv.FormatInt(int64(animal.ID), 10),
		Name: animal.Name,
	}
}
