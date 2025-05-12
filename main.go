package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create channel to listen for signals
	sigChan := make(chan os.Signal, 1)

	// Register for SIGINT (Ctrl+C) and SIGTERM
	// Note: SIGKILL cannot be caught or handled
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a channel to signal when to exit
	done := make(chan bool, 1)

	// Start goroutine to handle signals
	go func() {
		sig := <-sigChan
		fmt.Printf("Received signal: %v\n", sig)
		fmt.Println("Waiting for 5 seconds before shutting down...")

		// Wait for a while before terminating
		time.Sleep(5 * time.Second)

		fmt.Println("Shutting down now")
		done <- true
	}()

	// Keep the main goroutine running until signaled to exit
	fmt.Println("Application running. Press Ctrl+C to terminate")
	<-done
	fmt.Println("Application terminated")
}