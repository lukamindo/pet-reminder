package domain

import (
	"context"
	"errors"

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
func (s UserService) Register(c context.Context, urr request.UserRegister) (*response.User, error) {

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

	ret := response.User{
		Email: urr.Email,
	}

	return &ret, nil
}

// los Used for Signing In the Users
func (s UserService) Login(c context.Context, ulr request.UserLogin) (*response.SuccessfulLoginResponse, error) {

	// Validate req Struct
	err := validator.ValidateStruct(ulr)
	if err != nil {
		return nil, server.ErrBadRequest(err)
	}

	// Get User
	user, err := db.UserByEmail(c, s.connDB, ulr.Email)
	if user == nil {
		return nil, server.ErrBadRequest(errors.New("Incorrect email or password"))
	}
	if err != nil {
		return nil, server.ErrInternalDB(err)
	}

	// Check password
	isValidPwd := encrypt.CheckPassword(user.Password, ulr.Password)
	if !isValidPwd {
		return nil, server.ErrBadRequest(errors.New("Incorrect email or password"))
	}

	// Create JWT token
	token, err := auth.CreateJWT(ulr.Email)
	if err != nil {
		return nil, server.ErrInternalDomain(err)
	}

	// Create and send SuccessfulLoginResponse struct with token
	ret := response.SuccessfulLoginResponse{
		User:      user.Response(),
		AuthToken: token,
	}
	return &ret, nil
}
