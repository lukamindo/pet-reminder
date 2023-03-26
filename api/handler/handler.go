package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/lukamindo/pet-reminder/app/domain"
	"github.com/lukamindo/pet-reminder/helper/conn"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func New(e *echo.Echo) {
	usersGroup(e.Group("/users"))
	testGroup(e.Group("/test"))
}

func usersGroup(g *echo.Group) {
	as := domain.NewUserService(conn.New())
	g.POST("/register", userRegisterHandler(as))
	// g.POST("/login", UserLoginHandler(as))
	// g.GET("/login/facebook", InitFacebookLogin())
	// g.GET("/facebook/callback", HandleFacebookLogin())
}
func testGroup(g *echo.Group) {
	//TODO: უდნა გავჩექო რატო არ მუშაობს იმაზე და მიდლვეარი როგორია
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("jwt_secret_key"),
	}
	g.Use(echojwt.WithConfig(config))
	g.GET("", blocked)
}

func blocked(c echo.Context) error {
	return c.String(http.StatusOK, "blocked")
}
