package entity

import "golang.org/x/crypto/bcrypt"

type TypeUser string

type User struct {
	ID          int      `json:"id"`
	Name        string   `json:"name" db:"name"`
	Email       string   `json:"email" db:"email"`
	CityID      int      `json:"cityID" db:"city_id"`
	StateID     int      `json:"stateID" db:"state_id"`
	Phone       *string  `json:"phone" db:"phone"`
	Description *string  `json:"description" db:"description"`
	Password    string   `json:"password" db:"password"`
	Photo       *string  `json:"photo" db:"photo"`
	Type        TypeUser `json:"type" db:"type"`
	ShowEmail   bool     `json:"showEmail" db:"show_email"`
	ShowPhone   bool     `json:"showPhone" db:"show_phone"`
	ShowAddress bool     `json:"showAddress" db:"show_address"`
	Address     string   `json:"address" db:"address"`
	Number      *int     `json:"number" db:"number"`
	ZipCode     string   `json:"zipCode" db:"zip_code"`
	Complement  *string  `json:"complement" db:"complement"`
	TotalPets   *int     `json:"totalPets" db:"total_pets"`
}

func (user *User) GeneratePassword(password string) {
	user.Password, _ = user.generatePassword(password)
}
func (user *User) generatePassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pass), err
}

func (user *User) CheckPassword(passwordHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}
