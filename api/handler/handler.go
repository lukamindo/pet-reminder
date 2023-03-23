package handler

import (
	"github.com/labstack/echo"
	"github.com/lukamindo/pet-reminder/app/domain"
	"github.com/lukamindo/pet-reminder/app/request"
	"github.com/lukamindo/pet-reminder/helper/conn"
	"github.com/lukamindo/pet-reminder/helper/server"
)

func New(e *echo.Echo) {
	authenticationGroup(e.Group("/auth"))
	adminGroup(e.Group("/admin"))
}

func adminGroup(g *echo.Group) {
	g.GET("/test", test())
}

func authenticationGroup(g *echo.Group) {
	as := domain.NewAuthService(conn.New())
	g.POST("/signin", signInUser(as))
	// g.POST("/signup", test())
	// g.GET("/userDetails", test())
}

func test() echo.HandlerFunc {
	return func(c echo.Context) error {
		ret := "hi"
		return server.Success(c, ret)
	}
}

func signInUser(s domain.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {

		var ulr request.LoginParams

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
