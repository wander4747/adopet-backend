package model

import (
	"strconv"

	"github.com/wander4747/adopet-backend/pkg/entity"
)

type User struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	CityID      int      `json:"cityID"`
	StateID     int      `json:"stateID"`
	Phone       *string  `json:"phone"`
	Description *string  `json:"description"`
	Password    string   `json:"password"`
	Photo       *string  `json:"photo"`
	Type        TypeUser `json:"type"`
	ShowEmail   bool     `json:"showEmail"`
	ShowPhone   bool     `json:"showPhone"`
	ShowAddress bool     `json:"showAddress"`
	Address     string   `json:"address"`
	Number      *int     `json:"number"`
	ZipCode     string   `json:"zipCode"`
	Complement  *string  `json:"complement"`
	TotalPets   *int     `json:"totalPets"`
}

func NewUserModel(user *entity.User) *User {
	return &User{
		ID:          strconv.FormatInt(int64(user.ID), 10),
		Name:        user.Name,
		Email:       user.Email,
		CityID:      user.CityID,
		StateID:     user.StateID,
		Phone:       user.Phone,
		Description: user.Description,
		Password:    user.Password,
		Photo:       user.Photo,
		Type:        TypeUser(user.Type),
		ShowEmail:   user.ShowEmail,
		ShowPhone:   user.ShowPhone,
		ShowAddress: user.ShowAddress,
		Address:     user.Address,
		Number:      user.Number,
		ZipCode:     user.ZipCode,
		Complement:  user.Complement,
		TotalPets:   user.TotalPets,
	}
}

func NewUserEntity(user NewUser) entity.User {
	return entity.User{
		Name:        user.Name,
		Email:       user.Email,
		CityID:      user.CityID,
		StateID:     user.StateID,
		Phone:       user.Phone,
		Description: user.Description,
		Password:    user.Password,
		Photo:       user.Photo,
		Type:        entity.TypeUser(NewUserType(user.Type)),
		ShowEmail:   user.ShowEmail,
		ShowPhone:   user.ShowPhone,
		ShowAddress: user.ShowAddress,
		Address:     user.Address,
		Number:      user.Number,
		ZipCode:     user.ZipCode,
		Complement:  user.Complement,
		TotalPets:   user.TotalPets,
	}
}

func NewUserType(kind TypeUser) string {
	if kind == TypeUserModerator {
		return "moderator"
	}

	return "normal"
}
