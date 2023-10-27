package handler

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/labstack/echo/v4"

// 	"github.com/lukamindo/pet-reminder/pkg/auth"
// 	"github.com/lukamindo/pet-reminder/pkg/server"
// )

// // facebookLogin function will initiate the Facebook Login
// func facebookLogin() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		OAuth2Config := auth.FacebookGetOAuthConfig()
// 		url := OAuth2Config.AuthCodeURL(auth.FacebookGetRandomOAuthStateString())
// 		return c.Redirect(http.StatusTemporaryRedirect, url)
// 	}
// }

// // facebookCallback function will handle the Facebook Login Callback
// func facebookCallback() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		// get params
// 		code := c.QueryParam("code")
// 		state := c.QueryParam("state")

// 		// compare states
// 		if state != auth.FacebookGetRandomOAuthStateString() {
// 			return server.ErrBadRequest(errors.New("state is incorrect"))
// 		}

// 		// config
// 		OAuth2Config := auth.FacebookGetOAuthConfig()

// 		// exchange code for token
// 		token, err := OAuth2Config.Exchange(c.Request().Context(), code)
// 		if err != nil {
// 			return server.ErrBadRequest(err)
// 		}

// 		// use facebook api to get user info
// 		user, err := auth.FacebookGetUserInfo(OAuth2Config, token)
// 		if err != nil {
// 			return err
// 		}

// 		return server.Success(c, user)
// 	}
// }
