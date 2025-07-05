package main

import (
	"log"

	"gitub.com/TPautras/ecom/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
