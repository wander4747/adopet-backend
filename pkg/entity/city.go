package entity

type City struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	StateID int    `json:"stateId" db:"state_id"`
}
