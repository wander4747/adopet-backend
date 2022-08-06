package service

import "github.com/wander4747/adopet-backend/config"

type All struct {
	AnimalService Animal
	StateService  State
}

func NewService() All {
	config := config.NewConfig()

	return All{
		AnimalService: NewAnimal(*config),
		StateService:  NewState(*config),
	}
}
