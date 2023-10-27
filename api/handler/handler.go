package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lukamindo/pet-reminder/api/middleware"
	"github.com/lukamindo/pet-reminder/app/domain"
	"github.com/lukamindo/pet-reminder/pkg/conn"
)

func New(e *echo.Echo) {
	usersGroup(e.Group("/users"))
	petGroup(e.Group("/pets"))
	// socialAuthGroup(e.Group("/auth"))
}

// func socialAuthGroup(g *echo.Group) {
// 	g.GET("/login/google", googleLogin())
// 	g.GET("/google/callback", googleCallback())
// 	g.GET("/login/facebook", facebookLogin())
// 	g.GET("/facebook/callback", facebookCallback())
// }

func usersGroup(g *echo.Group) {
	us := domain.NewUserService(conn.New())
	g.POST("/register", userRegister(us))
	g.POST("/login", userLogin(us))
}

func petGroup(g *echo.Group) {
	ps := domain.NewPetService(conn.New())
	g.POST("/create", petCreate(ps), middleware.IsLoggedIn)
	g.GET("", petList(ps))
	g.GET("/:id", petByID(ps))
	g.PUT("/:id", petUpdateByID(ps))
	g.DELETE("/:id", petDeleteByID(ps))

}
