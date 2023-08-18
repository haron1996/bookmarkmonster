package main

import (
	"log"
	"net/http"

	"github.com/kwandapchumba/bookmarkmonster/router"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	server := &http.Server{
		Addr:    config.PORT,
		Handler: router.Router(),
	}

	log.Fatal(server.ListenAndServe())
}
