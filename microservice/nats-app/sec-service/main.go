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

	subj, err := nc.SubscribeSync("foo")
	if err != nil {
		log.Fatal(err)
	}

	msg, err := subj.NextMsg(time.Duration(5 * time.Second))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got message from first service:\n %s\n", msg.Data)
}
