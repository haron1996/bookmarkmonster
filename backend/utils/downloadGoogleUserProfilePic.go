package utils

import (
	"crypto/tls"
	"errors"
	"fmt"
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

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}

	client := &http.Client{Transport: tr}

	request, err := http.NewRequest("GET", profilePicLocation, nil)
	if err != nil {
		return nil, fmt.Errorf("could not instantiate new request: %v", err)
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0) Gecko/20100101 Firefox/81.0")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %v", err)
	}

	defer response.Body.Close()

	_, err = io.Copy(f, response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not copy file: %v", err)
	}

	return f, nil
}
