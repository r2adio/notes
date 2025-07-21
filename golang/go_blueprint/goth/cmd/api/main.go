package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"goth/internal/server"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	//signal.NotifyContext: modern go idiom that creates a ctx that cancels when os signals arrive
	//SIGINT: ctrl+c from terminal, SIGTERM: std termination signal from process managers(docker kubernetes systemd)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	//blocks until signal is received
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop() // Allow Ctrl+C to force shutdown

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling. thus server shutdown with timeout
	//w/o timeout, server might wait forever for long-running requests
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {

	server := server.NewServer() //create server instance

	// Create a done channel to signal when the shutdown is complete
	//its a buffered channel that signals when shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	//runs in a different goroutine, thus it doesnt block main execution
	//the signal handler must be registered before the server starts to catch interrrupts
	go gracefulShutdown(server, done)

	//blocking call: main goroutine waits here while server handles requests
	//returns when: server shuts down(gracefully/ forced)
	err := server.ListenAndServe()

	//http.ErrServerClosed is expected during graceful shutdown
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	//main go routine waits for shutdown completion before exiting
	<-done
	log.Println("Graceful shutdown complete.")
}
