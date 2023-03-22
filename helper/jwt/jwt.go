package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lukamindo/pet-reminder/app/db"
)

var jwtSecretKey = []byte("jwt_secret_key")

// CreateJWT func will used to create the JWT while signing in and signing out
func CreateJWT(email string) (response string, err error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &db.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err == nil {
		return tokenString, nil
	}
	return "", err
}

// VerifyToken func will used to Verify the JWT Token while using APIS
func VerifyToken(tokenString string) (email string, err error) {
	claims := &db.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if token != nil {
		return claims.Email, nil
	}
	return "", err
}
