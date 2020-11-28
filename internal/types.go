package internal

import (
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
)

// Account is the database struct
// swagger:model
type Account struct {
	// the id for this account
	// required: true
	// min: 1
	ID uint `json:"id"`

	// the name of this account
	// required: true
	Name string `json:"name" validate:"required"`

	// the email of this account
	// required: true
	Email *string `json:"email" validate:"required"`

	// the password for this account
	// required: true
	Password string `json:"-" validate:"required"`

	// the username of this account
	// required: true
	Username string `json:"username" validate:"required"`

	// the age of this account
	Age uint8 `json:"age"`

	// the birthday of this account
	// required: true
	Birthday    *time.Time   `json:"birthday" validate:"required"`
	ActivatedAt sql.NullTime `json:"-"`
	CreatedAt   time.Time    `json:"-"`
	UpdatedAt   time.Time    `json:"-"`
}

// Validate checks the validation of the user struct
func (a *Account) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
