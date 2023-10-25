package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lukamindo/pet-reminder/app/domain"
	"github.com/lukamindo/pet-reminder/helper/conn"
)

func New(e *echo.Echo) {
	usersGroup(e.Group("/users"))
	petGroup(e.Group("/pets"))
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
	us := domain.NewUserService(conn.New())
	g.POST("/register", userRegister(us))
	g.POST("/login", userLogin(us))
}

func petGroup(g *echo.Group) {
	ps := domain.NewPetService(conn.New())
	g.POST("/create", petCreate(ps))
	g.GET("", petList(ps))
	g.GET("/:id", petByID(ps))
	g.PUT("/:id", petUpdateByID(ps))
	g.DELETE("/:id", petDeleteByID(ps))

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
