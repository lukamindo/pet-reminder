package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lukamindo/pet-reminder/app/domain"
	"github.com/lukamindo/pet-reminder/app/request"
	"github.com/lukamindo/pet-reminder/pkg/server"
)

func userRegister(s domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var urr request.UserRegister
		err := c.Bind(&urr)
		if err != nil {
			return server.ErrBadRequest(err)
		}
		registerResponse, err := s.Register(c.Request().Context(), urr)
		if err != nil {
			return err
		}
		return server.Success(c, registerResponse)
	}
}

func userLogin(s domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var ulr request.UserLogin
		err := c.Bind(&ulr)
		if err != nil {
			return server.ErrBadRequest(err)
		}
		loginResponse, err := s.Login(c.Request().Context(), ulr)
		if err != nil {
			return err
		}
		return server.Success(c, loginResponse)
	}
}
