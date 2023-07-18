package main

import (
	"log"
	"chat-server/configs"
	"chat-server/cmd/http"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	http.Cmd(&config)
}