package middleware

import (
	"os"

	"github.com/labstack/echo/v4/middleware"
	"github.com/lukamindo/pet-reminder/app/constant"
	"github.com/lukamindo/pet-reminder/pkg/auth"
)

var jwtSecretKey = []byte(os.Getenv(constant.EnvJWT_SECRET_KEY))

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	Claims:     new(auth.Claims),
	SigningKey: jwtSecretKey,
})
