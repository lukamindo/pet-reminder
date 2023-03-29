package auth

import (
	"context"
	"encoding/json"
	"os"

	"github.com/lukamindo/pet-reminder/app/constant"
	"github.com/lukamindo/pet-reminder/helper/server"
	"golang.org/x/oauth2"
	facebookOAuth "golang.org/x/oauth2/facebook"
)

// FacebookUserDetails is struct used for user details
type FacebookUserDetails struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// FacebookGetOAuthConfig will return the config to call facebook Login
func FacebookGetOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv(constant.EnvFacebookClientID),
		ClientSecret: os.Getenv(constant.EnvFacebookClientSecret),
		RedirectURL:  os.Getenv(constant.EnvFacebookRedirectURL),
		Endpoint:     facebookOAuth.Endpoint,
		Scopes:       []string{"email"},
	}
}

// FacebookGetRandomOAuthStateString will return random string
func FacebookGetRandomOAuthStateString() string {
	return "SomeRandomStringAlgorithmForMoreSecurityFacebook"
}

// FacebookGetUserInfo will return information of user which is fetched from facebook
func FacebookGetUserInfo(gauth *oauth2.Config, token *oauth2.Token) (*FacebookUserDetails, error) {
	var fud FacebookUserDetails

	// make client from oauth with token
	client := gauth.Client(context.Background(), token)

	// fetch user data
	resp, err := client.Get("https://graph.facebook.com/me?fields=id,name,email")
	if err != nil {
		return nil, server.ErrBadRequest(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&fud)
	if err != nil {
		return nil, server.ErrBadRequest(err)
	}
	return &fud, nil
}
