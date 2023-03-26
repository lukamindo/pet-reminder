package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lukamindo/pet-reminder/app/domain"
	"github.com/lukamindo/pet-reminder/app/request"
	"github.com/lukamindo/pet-reminder/helper/server"
)

func userRegisterHandler(s domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {

		var urr request.UserRegister

		err := c.Bind(&urr)
		if err != nil {
			return echo.ErrBadRequest
		}

		user, err := s.Register(c.Request().Context(), urr)
		if user == nil {
			return err
		}
		if err != nil {
			return err
		}
		return server.Success(c, user)
	}
}

func UserLoginHandler(s domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {

		var ulr request.UserLogin

		err := c.Bind(&ulr)
		if err != nil {
			return err
		}

		loginResponse, err := s.Login(c.Request().Context(), ulr)
		if loginResponse == nil {
			return err
		}
		if err != nil {
			return err
		}
		return server.Success(c, loginResponse)
	}
}
