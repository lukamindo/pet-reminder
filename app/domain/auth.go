package domain

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lukamindo/pet-reminder/app/db"
	"github.com/lukamindo/pet-reminder/app/request"
	"github.com/lukamindo/pet-reminder/app/response"
	"github.com/lukamindo/pet-reminder/helper/auth"
	"github.com/lukamindo/pet-reminder/helper/encrypt"
)

type UserService struct {
	connDB *sqlx.DB
}

// NewAuthService constructs new service
func NewUserService(connDB *sqlx.DB) UserService {
	return UserService{
		connDB: connDB,
	}
}

// Register method for users to register
func (s UserService) Register(c context.Context, urr request.UserRegister) (*response.Player, error) {
	if urr.Username == "" || urr.Email == "" || urr.Password == "" {
		return nil, fmt.Errorf("bad request")
	}
	hashedPwd, err := encrypt.Password(urr.Password)
	if err != nil {
		return nil, fmt.Errorf("internal server error, while encrypting password")
	}

	urrDB := urr.DB(hashedPwd)
	_, err = db.UserCreate(c, s.connDB, urrDB)
	if err != nil {
		return nil, err
	}

	// token, err := auth.CreateJWT(urr.Email)
	// if token == "" {
	// 	return nil, fmt.Errorf("internal server error")
	// }

	ret := response.Player{
		Email: urr.Email,
	}

	return &ret, nil
}

// los Used for Signing In the Users
func (s UserService) Login(c context.Context, ulr request.UserLogin) (*response.SuccessfulLoginResponse, error) {
	if ulr.Email == "" || ulr.Password == "" {
		return nil, fmt.Errorf("bad request")
	}

	//TODO: aq unda daematos  password gacheqva

	// user, err := db.UserByEmail(c, s.connDB, ulr.Email)
	// if user == nil {
	// 	return nil, fmt.Errorf("internal db")
	// }
	// if err != nil {
	// 	return nil, err
	// }

	token, _ := auth.CreateJWT(ulr.Email)
	if token == "" {
		return nil, fmt.Errorf("internal server error")
	}

	ret := response.SuccessfulLoginResponse{
		Email:     ulr.Email,
		AuthToken: token,
	}
	return &ret, nil
}
