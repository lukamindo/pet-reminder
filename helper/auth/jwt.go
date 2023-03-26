package auth

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// Claims is  a struct that will be encoded to a JWT.
// jwt.StandardClaims is an embedded type to provide expiry time
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var jwtSecretKey = []byte("jwt_secret_key")

// CreateJWT func will used to create the JWT while signing in and signing out
func CreateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
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
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if token != nil {
		return claims.Email, nil
	}
	return "", err
}
