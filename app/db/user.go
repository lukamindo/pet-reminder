package db

import jwt "github.com/dgrijalva/jwt-go"

// ErrorResponse is struct for sending error message with code.
type ErrorResponse struct {
	Code    int
	Message string
}

// SuccessResponse is struct for sending error message with code.
type SuccessResponse struct {
	Code     int
	Message  string
	Response interface{}
}

// Claims is  a struct that will be encoded to a JWT.
// jwt.StandardClaims is an embedded type to provide expiry time
type Claims struct {
	Email string
	jwt.StandardClaims
}

// RegistationParams is struct to read the request body
type RegistationParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginParams is struct to read the request body
type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SuccessfulLoginResponse is struct to send the request response
type SuccessfulLoginResponse struct {
	Email     string
	AuthToken string
}

// UserDetails is struct used for user details
type UserDetails struct {
	Name     string
	Email    string
	Password string
}
