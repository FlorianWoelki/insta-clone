package handlers

import (
	"log"
)

// Accounts handler for getting, creating and updating accounts
type Accounts struct {
	logger *log.Logger
}

// NewAccounts returns a new accounts handler with the given logger
func NewAccounts(logger *log.Logger) *Accounts {
	return &Accounts{logger}
}
