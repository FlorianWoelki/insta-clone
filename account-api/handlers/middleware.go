package handlers

import (
	"context"
	"net/http"

	"github.com/florianwoelki/insta-clone/internal"
)

// MiddlewareValidateAccount checks wether the input is a valid json formatted
// account. When it is it will pass the next http handler
func (a *Accounts) MiddlewareValidateAccount(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		account := &internal.Account{}

		err := internal.FromJSON(account, r.Body)
		if err != nil {
			a.logger.Println("[ERROR] Deserializing account", err)
			http.Error(rw, "Error reading account", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), &KeyAccount{}, account)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
