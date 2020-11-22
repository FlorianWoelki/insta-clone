package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ListSingle handles get requests and returns single user by id
func (a *Accounts) ListSingle(rw http.ResponseWriter, r *http.Request) {
	// TODO: get single account from database
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		a.logger.Printf("Something went wrong converting id to int: %s\n", err)
		return
	}

	a.logger.Println(id)
}
