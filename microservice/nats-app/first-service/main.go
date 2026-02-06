package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

/*
sudo docker run -d \
  --name nats \
  -p 4222:4222 \
  -p 8222:8222 \
  nats:latest -js

save to volume
sudo docker run -d \
  --name nats \
  -p 4222:4222 \
  -p 8222:8222 \
  -v nats-data:/data \
  nats:latest -js

*/

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

	// اینجا stream به صورت خودکار ساخته میشه اگر وجود نداشته باشه
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "FOO",
		Subjects: []string{"foo"},
		Storage:  nats.MemoryStorage, // برای dev
	})
	if err != nil {
		log.Println("Stream already exists or error:", err)
	}

	_, err = js.Publish("foo", []byte("Hello JetStream Send Messsage From First Service 🚀"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Message sent!")
}
