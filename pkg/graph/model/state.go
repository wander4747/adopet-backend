package model

import (
	"strconv"

	"github.com/wander4747/adopet-backend/pkg/entity"
)

type State struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Initials string `json:"initials"`
}

func NewState(state entity.State) *State {
	return &State{
		ID:       strconv.FormatInt(int64(state.ID), 10),
		Name:     state.Name,
		Initials: state.Initials,
	}
}
