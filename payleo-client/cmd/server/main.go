package main

import (
	"fmt"
	"log"
	"net/http"

	"payleo-client/internal/config"
	"payleo-client/internal/handlers"
)

func main() {

	cfg := config.Load()

	http.HandleFunc(
		"/payleo/callback",
		handlers.PayleoCallback,
	)

	port := cfg.ServerPort

	if port == "" {
		port = "8080"
	}

	fmt.Printf(
		"Server listening on :%s\n",
		port,
	)

	log.Fatal(
		http.ListenAndServe(
			":"+port,
			nil,
		),
	)
}