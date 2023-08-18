package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
)

type GoogleUser struct {
	Id            string
	Email         string
	EmailVerified bool
	Name          string
	GivenName     string
	FamilyName    string
	Picture       string
	Locale        string
}

func ExchangeGoogleAuthCodeForToken(authCode string, ctx context.Context) (*GoogleUser, error) {
	b, err := os.ReadFile("google-auth-credentials.json")
	if err != nil {
		return nil, errors.New("could not read google-auth-credentials.json")
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email")
	if err != nil {
		return nil, errors.New("could not construct config")
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, errors.New("could not retrieve token from web")
	}

	googleUserInfoEndpoint := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", tok.AccessToken)

	response, err := http.Get(googleUserInfoEndpoint)
	if err != nil {
		return nil, errors.New("could not get url")
	}

	defer response.Body.Close()

	var user *GoogleUser

	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		log.Printf("could not decode response body: %v", err)
	}

	return user, nil
}
