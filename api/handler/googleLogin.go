package handler

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// 	"github.com/lukamindo/pet-reminder/pkg/auth"
// 	"github.com/lukamindo/pet-reminder/pkg/server"
// )

// // googleLogin function will initiate the Facebook Login
// func googleLogin() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		OAuth2Config := auth.GoogleGetOAuthConfig()
// 		url := OAuth2Config.AuthCodeURL(auth.GoogleGetRandomOAuthStateString())
// 		return c.Redirect(http.StatusTemporaryRedirect, url)
// 	}
// }

// // googleCallback function will handle the Google Login Callback
// func googleCallback() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		// get params
// 		code := c.QueryParam("code")
// 		state := c.QueryParam("state")

// 		// compare states
// 		if state != auth.GoogleGetRandomOAuthStateString() {
// 			return server.ErrBadRequest(errors.New("state is incorrect"))
// 		}

// 		// config
// 		OAuth2Config := auth.GoogleGetOAuthConfig()

// 		// exchange code for token
// 		token, err := OAuth2Config.Exchange(c.Request().Context(), code)
// 		if err != nil {
// 			return server.ErrBadRequest(err)
// 		}

// 		// use google api to get user info
// 		user, err := auth.GoogleGetUserInfo(OAuth2Config, token)
// 		if err != nil {
// 			return err
// 		}

// 		return server.Success(c, user)
// 	}
// }
