package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	sub, err := js.SubscribeSync("foo")
	if err != nil {
		log.Fatal(err)
	}

	msg, err := sub.NextMsg(5 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Got message:", string(msg.Data))
}
