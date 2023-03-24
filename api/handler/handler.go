package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/lukamindo/pet-reminder/app/db"
	"github.com/lukamindo/pet-reminder/app/domain"
	"github.com/lukamindo/pet-reminder/app/request"
	"github.com/lukamindo/pet-reminder/helper/auth"
	"github.com/lukamindo/pet-reminder/helper/conn"
	"github.com/lukamindo/pet-reminder/helper/server"
	"golang.org/x/oauth2"
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
	g.GET("/login/facebook", InitFacebookLogin())
	g.GET("/facebook/callback", HandleFacebookLogin())
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
		OAuth2Config := auth.GetFacebookOAuthConfig()
		url := OAuth2Config.AuthCodeURL(auth.GetRandomOAuthStateString())
		return c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

// HandleFacebookLogin function will handle the Facebook Login Callback
func HandleFacebookLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		state := c.FormValue("state")
		code := c.FormValue("code")

		if state != auth.GetRandomOAuthStateString() {
			return c.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}

		OAuth2Config := auth.GetFacebookOAuthConfig()

		token, err := OAuth2Config.Exchange(oauth2.NoContext, code)
		if err != nil || token == nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}

		fbUserDetails, fbUserDetailsError := auth.GetUserInfoFromFacebook(token.AccessToken)
		if fbUserDetailsError != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}
		if fbUserDetails == nil {
			//TODO: ესაა დამატებული runtime ერრორირს გამო და ამოსაგდებია და საფიქრია
			return c.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}

		authToken, authTokenError := SignInUser(c.Request().Context(), conn.New(), *fbUserDetails)
		if authTokenError != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}

		cookie := &http.Cookie{Name: "Authorization", Value: "Bearer " + authToken, Path: "/"}
		c.SetCookie(cookie)

		return c.Redirect(http.StatusTemporaryRedirect, "/profile")
	}
}

// SignInUser Used for Signing In the Users
func SignInUser(c context.Context, conn *sqlx.DB, facebookUserDetails auth.FacebookUserDetails) (string, error) {
	if facebookUserDetails == (auth.FacebookUserDetails{}) {
		return "", errors.New("User details Can't be empty")
	}

	if facebookUserDetails.Email == "" {
		return "", errors.New("Last Name can't be empty")
	}

	if facebookUserDetails.Name == "" {
		return "", errors.New("Password can't be empty")
	}

	user, err := db.UserByEmail(c, conn, facebookUserDetails.Email)
	if err != nil {
		return "", errors.New("db problem have to change")
	}
	if user == nil {
		err = db.CreateUser(c, conn, request.RegistationParams{
			Username: facebookUserDetails.Name,
			Email:    facebookUserDetails.Email,
			Password: "",
		})
		if err != nil {
			return "", errors.New("Error occurred registration")
		}
	}

	tokenString, _ := auth.CreateJWT(facebookUserDetails.Email)
	if tokenString == "" {
		return "", errors.New("Unable to generate Auth token")
	}

	return tokenString, nil
}
