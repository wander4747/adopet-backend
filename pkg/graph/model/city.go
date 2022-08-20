package model

import (
	"strconv"

	"github.com/wander4747/adopet-backend/pkg/entity"
)

type City struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	StateID string `json:"stateId"`
}

func NewCity(city entity.City) *City {
	return &City{
		ID:      strconv.FormatInt(int64(city.ID), 10),
		Name:    city.Name,
		StateID: strconv.FormatInt(int64(city.StateID), 10),
	}
}
