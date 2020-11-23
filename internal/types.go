package internal

import (
	"database/sql"
	"time"
)

// User is the database struct
type User struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Email       *string      `json:"email"`
	Username    string       `json:"username"`
	Age         uint8        `json:"age"`
	Birthday    *time.Time   `json:"birthday"`
	ActivatedAt sql.NullTime `json:"-"`
	CreatedAt   time.Time    `json:"-"`
	UpdatedAt   time.Time    `json:"-"`
}
