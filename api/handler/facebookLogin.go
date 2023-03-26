package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/lukamindo/pet-reminder/app/db"
	"github.com/lukamindo/pet-reminder/helper/auth"
	"github.com/lukamindo/pet-reminder/helper/conn"
	"golang.org/x/oauth2"
)

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
		code := c.QueryParam("code")
		state := c.QueryParam("state")

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

		return c.Redirect(http.StatusTemporaryRedirect, "https://youtube.com")
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
		_, err = db.UserCreate(c, conn, db.User{
			Username: facebookUserDetails.Name,
			Password: "",
			Email:    facebookUserDetails.Email,
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
