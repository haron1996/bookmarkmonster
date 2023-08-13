package utils

import (
	"errors"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GenerateGoogleAuthUrl() (string, error) {
	b, err := os.ReadFile("google-auth-credentials.json")
	if err != nil {
		return "", errors.New("could not read google-auth-credentials.json")
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email")
	if err != nil {
		return "", errors.New("could not construct config")
	}

	return config.AuthCodeURL("state-token", oauth2.AccessTypeOffline), nil
}
