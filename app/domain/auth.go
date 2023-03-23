package domain

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lukamindo/pet-reminder/app/response"
)

type AuthService struct {
	connDB *sqlx.DB
}

// NewAuthService constructs new service
func NewAuthService(connDB *sqlx.DB) AuthService {
	return AuthService{
		connDB: connDB,
	}
}

// Logins returns a struct
func (s AuthService) Login(c context.Context) (*response.SuccessfulLoginResponse, error) {
	return nil, nil
}
