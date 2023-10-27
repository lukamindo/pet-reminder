package domain

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lukamindo/pet-reminder/app/db"
	"github.com/lukamindo/pet-reminder/app/request"
	"github.com/lukamindo/pet-reminder/app/response"
	"github.com/lukamindo/pet-reminder/pkg/auth"
	"github.com/lukamindo/pet-reminder/pkg/encrypt"
	"github.com/lukamindo/pet-reminder/pkg/server"
	"github.com/lukamindo/pet-reminder/pkg/validator"
)

type UserService struct {
	connDB *sqlx.DB
}

func NewUserService(connDB *sqlx.DB) UserService {
	return UserService{
		connDB: connDB,
	}
}

// Register method creates user in DB
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
	userID, err := db.UserCreate(c, s.connDB, urrDB)
	if err != nil {
		return nil, server.ErrInternalDB(err)
	}
	ret := response.User{
		ID:       *userID,
		Username: urr.Username,
		Email:    urr.Email,
		//TODO: ეს რამდენად უნდა???????? და თუ კი ბაზიდან დავაბრუნებინო ესეც
		CreatedAt: time.Now().UTC(),
	}
	return &ret, nil
}

// Login is for Signing In the Users
func (s UserService) Login(c context.Context, ulr request.UserLogin) (*response.SuccessfulLoginResponse, error) {
	err := validator.ValidateStruct(ulr)
	if err != nil {
		return nil, server.ErrBadRequest(err)
	}
	// Get User
	user, err := db.UserByEmail(c, s.connDB, ulr.Email)
	if user == nil {
		return nil, server.ErrBadRequest(errors.New("incorrect email or password"))
	}
	if err != nil {
		return nil, server.ErrInternalDB(err)
	}
	// Check password
	isValidPwd := encrypt.CheckPassword(user.Password, ulr.Password)
	if !isValidPwd {
		return nil, server.ErrBadRequest(errors.New("incorrect email or password"))
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
