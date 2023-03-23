package domain

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lukamindo/pet-reminder/app/db"
	"github.com/lukamindo/pet-reminder/app/request"
	"github.com/lukamindo/pet-reminder/app/response"
	"github.com/lukamindo/pet-reminder/helper/jwt"
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

// SignInUser Used for Signing In the Users
func (s AuthService) Login(c context.Context, ulr request.LoginParams) (*response.SuccessfulLoginResponse, error) {
	if ulr.Email == "" || ulr.Password == "" {
		return nil, fmt.Errorf("bad request")
	}

	//TODO: aq unda daematos  password gacheqva

	user, err := db.UserByEmail(c, s.connDB, ulr.Email)
	if user == nil {
		return nil, fmt.Errorf("internal db")
	}
	if err != nil {
		return nil, err
	}
	fmt.Println(user)

	token, _ := jwt.CreateJWT(ulr.Email)
	if token == "" {
		return nil, fmt.Errorf("internal server error")
	}
	fmt.Println(token)

	ret := response.SuccessfulLoginResponse{
		Email:     ulr.Email,
		AuthToken: token,
	}
	return &ret, nil
}