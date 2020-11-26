package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/florianwoelki/insta-clone/internal"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key") // TODO: refactor to env variable

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
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
	claims := &claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// create jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
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
