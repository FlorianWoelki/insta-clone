package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/florianwoelki/insta-clone/internal"
	"golang.org/x/crypto/bcrypt"
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login lets the user login with a valid json structure defined in the credentials struct
func (a *Accounts) Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	var creds credentials
	if err := internal.FromJSON(&creds, r.Body); err != nil {
		a.logger.Printf("[ERROR] Deserializing login body %v", err)
		http.Error(rw, "Deserializing login body", http.StatusBadRequest)
		return
	}

	// check wether email exists in database
	var account internal.Account
	if err := a.db.Where("email = ?", creds.Email).First(&account).Error; err != nil {
		a.logger.Printf("[ERROR] Couldn't find any email %s, error: %v", creds.Email, err)
		http.Error(rw, "Couldn't login the user, wrong email", http.StatusNotFound)
		return
	}

	// compare and hash the password with bcrypt
	passErr := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(creds.Password))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		a.logger.Printf("[ERROR] Wrong password, error: %v", passErr)
		http.Error(rw, "Entered wrong password", http.StatusForbidden)
		return
	}

	// define jwt information
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &internal.Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// create jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(internal.JWTKey)
	if err != nil {
		a.logger.Printf("[ERROR] Token creation has failed, error: %v", err)
		http.Error(rw, "Something went wrong internally", http.StatusInternalServerError)
		return
	}

	// set client cookie with name "token" as the generated jwt
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

// Refresh allows the user to get a new refreshed token
func (a *Accounts) Refresh(rw http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			a.logger.Printf("[ERROR] Couldn't find any token cookie, error: %v", err)
			http.Error(rw, "Not authorized", http.StatusUnauthorized)
			return
		}

		a.logger.Printf("[ERROR] Something went wrong while getting the token cookie, error: %v", err)
		http.Error(rw, "Something went wrong internally", http.StatusBadRequest)
		return
	}

	tokenStr := c.Value
	claims := &internal.Claims{}

	// parse jwt string and store the result in claims
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

	// check if time in token has elapsed
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		a.logger.Println("[ERROR] Time for token has elapsed")
		http.Error(rw, "Time for token has elapsed", http.StatusBadRequest)
		return
	}

	// create a new token for the current use
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(internal.JWTKey)
	if err != nil {
		a.logger.Printf("[ERROR] Token creation has failed, error: %v", err)
		http.Error(rw, "Something went wrong internally", http.StatusInternalServerError)
		return
	}

	// set the new token as the users token cookie
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   tokenStr,
		Expires: expirationTime,
	})
}
