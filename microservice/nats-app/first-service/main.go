package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	if err := nc.Publish("foo", []byte("Hello NATS!!!")); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Message sent to NATS.")

}
