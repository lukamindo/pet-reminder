package auth

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/lukamindo/pet-reminder/app/constant"
)

// Claims is  a struct that will be encoded to a JWT.
// jwt.StandardClaims is an embedded type to provide expiry time
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var jwtSecretKey = []byte(os.Getenv(constant.EnvJWT_SECRET_KEY))

// CreateJWT func will used to create the JWT while signing in and signing out
func CreateJWT(email string) (string, error) {
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecretKey)
	if err == nil {
		return tokenString, nil
	}
	return "", err
}
