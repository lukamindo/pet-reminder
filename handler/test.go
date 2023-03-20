package handler

import (
	"github.com/labstack/echo"
	"github.com/lukamindo/pet-reminder/helper/server"
)

func test() echo.HandlerFunc {
	return func(c echo.Context) error {
		ret := "hi"
		return server.Success(c, ret)
	}
}
