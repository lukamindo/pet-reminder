package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lukamindo/pet-reminder/app/domain"
	"github.com/lukamindo/pet-reminder/helper/conn"
)

func New(e *echo.Echo) {
	usersGroup(e.Group("/users"))
	petGroup(e.Group("/pet"))
	// blockedGroup(e.Group("/test"))
	// socialAuthGroup(e.Group("/auth"))
}

// func socialAuthGroup(g *echo.Group) {
// 	g.GET("/login/google", googleLogin())
// 	g.GET("/google/callback", googleCallback())
// 	g.GET("/login/facebook", facebookLogin())
// 	g.GET("/facebook/callback", facebookCallback())
// }

func usersGroup(g *echo.Group) {
	as := domain.NewUserService(conn.New())
	g.POST("/register", userRegisterHandler(as))
	g.POST("/login", UserLoginHandler(as))
}

func petGroup(g *echo.Group) {
}

// func blockedGroup(g *echo.Group) {
// 	config := echojwt.Config{
// 		NewClaimsFunc: func(c echo.Context) jwt.Claims {
// 			return new(auth.Claims)
// 		},
// 		SigningKey: []byte(os.Getenv(constant.EnvJWT_SECRET_KEY)),
// 	}
// 	g.Use(echojwt.WithConfig(config))
// 	g.GET("", blocked)
// }

// func blocked(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(*auth.Claims)
// 	email := claims.Email
// 	return c.String(http.StatusOK, "Welcome "+email+"!")
// }
