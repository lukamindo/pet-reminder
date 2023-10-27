package auth

import (
	"context"
	"encoding/json"
	"os"

	"github.com/lukamindo/pet-reminder/app/constant"
	"github.com/lukamindo/pet-reminder/pkg/server"
	"golang.org/x/oauth2"
	googleOAuth "golang.org/x/oauth2/google"
)

// GoogleUserDetail struct for google auth user
type GoogleUserDetail struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

// GoogleGetOAuthConfig will return the config to call google Login
func GoogleGetOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv(constant.EnvGoogleClientID),
		ClientSecret: os.Getenv(constant.EnvGoogleClientSecret),
		RedirectURL:  os.Getenv(constant.EnvGoogleRedirectURL),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: googleOAuth.Endpoint,
	}
}

// GoogleGetRandomOAuthStateString will return random string
func GoogleGetRandomOAuthStateString() string {
	return "SomeRandomStringAlgorithmForMoreSecurityGoogle"
}

// GoogleGetUserInfo will return information of user which is fetched from google
func GoogleGetUserInfo(gauth *oauth2.Config, token *oauth2.Token) (*GoogleUserDetail, error) {
	var gud GoogleUserDetail

	// make client from oauth with token
	client := gauth.Client(context.Background(), token)

	// fetch user data
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, server.ErrBadRequest(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&gud)
	if err != nil {
		return nil, server.ErrBadRequest(err)
	}
	return &gud, nil
}
