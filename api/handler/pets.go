package handler

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lukamindo/pet-reminder/app/domain"
	"github.com/lukamindo/pet-reminder/app/request"
	"github.com/lukamindo/pet-reminder/helper/server"
)

func petCreate(s domain.PetService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var pcr request.PetCreate
		err := c.Bind(&pcr)
		if err != nil {
			return server.ErrBadRequest(err)
		}
		createResponse, err := s.Create(c.Request().Context(), pcr)
		if err != nil {
			return err
		}
		return server.Success(c, createResponse)
	}
}

func petList(s domain.PetService) echo.HandlerFunc {
	return func(c echo.Context) error {
		petList, err := s.List(c.Request().Context())
		if err != nil {
			return err
		}
		return server.Success(c, petList)
	}
}

func petByID(s domain.PetService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return server.ErrBadRequest(fmt.Errorf("bad pet id"))
		}
		pet, err := s.ByID(c.Request().Context(), id)
		if err != nil {
			return err
		}
		return server.Success(c, pet)
	}
}

func petUpdateByID(s domain.PetService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return server.ErrBadRequest(fmt.Errorf("bad pet id"))
		}
		var pur request.PetUpdate
		err = c.Bind(&pur)
		if err != nil {
			return server.ErrBadRequest(err)
		}
		createResponse, err := s.Update(c.Request().Context(), id, pur)
		if err != nil {
			return err
		}
		return server.Success(c, createResponse)
	}
}

func petDeleteByID(s domain.PetService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return server.ErrBadRequest(fmt.Errorf("bad pet id"))
		}
		err = s.DeleteByID(c.Request().Context(), id)
		if err != nil {
			return err
		}
		return server.Success(c, "OK")
	}
}
