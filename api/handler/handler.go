package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lukamindo/pet-reminder/app/domain"
	"github.com/lukamindo/pet-reminder/app/request"
	"github.com/lukamindo/pet-reminder/helper/auth"
	"github.com/lukamindo/pet-reminder/helper/conn"
	"github.com/lukamindo/pet-reminder/helper/server"
)

func New(e *echo.Echo) {
	authenticationGroup(e.Group("/auth"))
	adminGroup(e.Group("/admin"))
}

func adminGroup(g *echo.Group) {
	// g.GET("/test", test())
}

func authenticationGroup(g *echo.Group) {
	as := domain.NewAuthService(conn.New())
	g.POST("/signin", signInUser(as))
	g.POST("/signup", signUpUser(as))
	g.POST("/login/facebook", InitFacebookLogin())
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

func signUpUser(s domain.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {

		var urr request.RegistationParams

		err := c.Bind(&urr)
		if err != nil {
			return err
		}

		loginResponse, err := s.SignUp(c.Request().Context(), urr)
		if loginResponse == nil {
			return err
		}
		if err != nil {
			return err
		}
		return server.Success(c, loginResponse)
	}
}

// InitFacebookLogin function will initiate the Facebook Login
func InitFacebookLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var OAuth2Config = auth.GetFacebookOAuthConfig()
		url := OAuth2Config.AuthCodeURL(auth.GetRandomOAuthStateString())
		return c.Redirect(http.StatusTemporaryRedirect, url)
	}
}
