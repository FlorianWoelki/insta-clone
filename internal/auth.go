package internal

import "github.com/dgrijalva/jwt-go"

// JWTKey is used to create the signature
var JWTKey = []byte("secret_key") // TODO: refactor to env variable

// Claims is a struct that will be encoded in a jwt
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
