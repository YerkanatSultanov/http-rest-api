package main

import (
	"github.com/YerkanatSultanov/http-rest-api/internal/app/apiserver"
	"log"
)

func main() {
	api := apiserver.New()

	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
}
