package domain

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lukamindo/pet-reminder/app/db"
	"github.com/lukamindo/pet-reminder/app/request"
	"github.com/lukamindo/pet-reminder/app/response"
	"github.com/lukamindo/pet-reminder/helper/auth"
	"github.com/lukamindo/pet-reminder/helper/encrypt"
	"github.com/lukamindo/pet-reminder/helper/server"
	"github.com/lukamindo/pet-reminder/helper/validator"
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

	err := validator.ValidateStruct(urr)
	if err != nil {
		return nil, server.ErrBadRequest(err)
	}

	hashedPwd, err := encrypt.Password(urr.Password)
	if err != nil {
		return nil, server.ErrInternalDomain(err)
	}

	urrDB := urr.DB(hashedPwd)
	_, err = db.UserCreate(c, s.connDB, urrDB)
	if err != nil {
		return nil, server.ErrInternalDB(err)
	}

	token, err := auth.CreateJWT(urr.Email)
	if token == "" {
		return nil, server.ErrInternalDomain(err)
	}

	ret := response.Player{
		Email: urr.Email,
	}

	return &ret, nil
}

// los Used for Signing In the Users
func (s UserService) Login(c context.Context, ulr request.UserLogin) (*response.SuccessfulLoginResponse, error) {

	err := validator.ValidateStruct(ulr)
	if err != nil {
		return nil, server.ErrBadRequest(err)
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
		return nil, server.ErrInternalDomain(err)
	}

	ret := response.SuccessfulLoginResponse{
		Email:     ulr.Email,
		AuthToken: token,
	}
	return &ret, nil
}
