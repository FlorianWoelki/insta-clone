package internal

import (
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
)

// User is the database struct
type User struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name" validate:"required"`
	Email       *string      `json:"email" validate:"required"`
	Username    string       `json:"username" validate:"required"`
	Age         uint8        `json:"age" validate:"required"`
	Birthday    *time.Time   `json:"birthday" validate:"required"`
	ActivatedAt sql.NullTime `json:"-"`
	CreatedAt   time.Time    `json:"-"`
	UpdatedAt   time.Time    `json:"-"`
}

// Validate checks the validation of the user struct
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
