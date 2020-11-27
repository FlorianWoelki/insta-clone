package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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

		// try to validate the input account
		errs := a.validator.Validate(account)
		if len(errs) != 0 {
			a.logger.Printf("[ERROR] Validating account: %v", errs)
			http.Error(rw, fmt.Sprintf("Error validating account %v", errs), http.StatusUnprocessableEntity)
			return
		}

		// add account to context
		ctx := context.WithValue(r.Context(), KeyAccount{}, account)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}

// MiddlewareValidateToken validates if the jwt in the cookie is valid
func (a *Accounts) MiddlewareValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				a.logger.Printf("[ERROR] Cookie is not set, error %v", err)
				http.Error(rw, "Not authorized", http.StatusUnauthorized)
				return
			}
			a.logger.Printf("[ERROR] Something went wrong while fetching the token cookie, error %v", err)
			http.Error(rw, "Something internally went wrong", http.StatusBadRequest)
			return
		}

		tokenStr := c.Value

		claims := &internal.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return internal.JWTKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				a.logger.Printf("[ERROR] Invalid token, error: %v", err)
				http.Error(rw, "Invalid token", http.StatusUnauthorized)
				return
			}

			a.logger.Printf("[ERROR] Something went wrong while validating the token cookie, error: %v", err)
			http.Error(rw, "Something went wrong internally", http.StatusBadRequest)
			return
		}

		// check if token is invalid
		if !token.Valid {
			a.logger.Println("[ERROR] Invalid token")
			http.Error(rw, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(rw, r)
	})
}
