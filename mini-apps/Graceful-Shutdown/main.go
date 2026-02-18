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

	// Run server in to goroutine
	go func() {
		fmt.Printf("Starting server on 0.0.0.0%s\n", server.Addr)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
		close(serverErr)
	}()

	// Signal for close the server bay Ctrl+C (terminal) and close the port
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Handle in check the serverErr and received stop server from chan and Done to parent with ctx
	select {
	case err := <-serverErr:
		return err
	case <-stop:
		log.Println("🛑 Shutdown signal received...")
	case <-ctx.Done():
		log.Println("Context cancelled.")
	case <-time.After(shutdownTimeout):
		log.Println("Timeout shutdown server.")
	}

	// Handle the realy shutdown the server
	// shutdownCtx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		if closeErr := server.Close(); closeErr != nil {
			return errors.Join(err, closeErr)
		}
		return err
	}

	return nil
}

func main() {
	server := createServer()

	if err := runServer(context.Background(), server, 2*time.Second); err != nil {
		log.Fatalf("RunServerError: %v", err)
	}
}
