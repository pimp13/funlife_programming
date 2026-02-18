package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func createServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Slow request is called")
		time.Sleep(8 * time.Second)
		fmt.Fprintf(w, "Slow request completed at %v\n", time.Now())
	})

	return &http.Server{
		Addr:    ":4224",
		Handler: mux,
	}
}

func runServer(
	ctx context.Context,
	server *http.Server,
	shutdownTimeout time.Duration,
) error {
	serverErr := make(chan error, 1)

	go func() {
		fmt.Printf("Starting server on 0.0.0.0%s\n", server.Addr)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
		close(serverErr)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)


	return <-serverErr
}

func main() {
	server := createServer()

	if err := runServer(context.Background(), server, 2*time.Second); err != nil {
		log.Fatalf("RunServerError: %v", err)
	}
}
