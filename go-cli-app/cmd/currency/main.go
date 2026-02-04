package main

import (
	"go-cli-app/internal/command"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("CURRENCY_API_KEY")
	if apiKey == "" {
		log.Fatalln("set konid env var CURRENCY_API_KEY!!")
	}
	if err := command.Execute(apiKey); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
