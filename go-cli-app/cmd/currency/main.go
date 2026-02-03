package main

import (
	"go-cli-app/internal/command"
	"log"
)

func main() {
	if err := command.Execute(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
