package handler

import (
	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	adminGroup(e.Group("/admin"))
}

func adminGroup(g *echo.Group) {
	g.GET("/test", test())
}
