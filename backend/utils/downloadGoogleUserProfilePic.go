package utils

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadGoogleUserProfilePic(profilePicLocation string) (*os.File, error) {
	f, err := os.Create("googleUserProfilePic.png")
	if err != nil {
		log.Printf("could not create file name: %v", err)
		return nil, errors.New("could not create file name")
	}

	defer f.Close()

	url := profilePicLocation

	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New("could not get url")
	}

	defer res.Body.Close()

	_, err = io.Copy(f, res.Body)
	if err != nil {
		return nil, errors.New("could not copy image")
	}

	return f, nil
}
