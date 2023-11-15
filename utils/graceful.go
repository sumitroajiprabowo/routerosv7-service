// utils/graceful.go

package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

// GracefulShutdown menerima aplikasi Fiber dan menangani proses graceful shutdown
func GracefulShutdown(app *fiber.App) {
	// Create a channel to receive OS signals from the kernel
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// This blocks the main thread until an interrupt is received
	<-c // blocks here until ctrl-c is pressed
	fmt.Println("Gracefully shutting down...")
	// Shutdown application
	_ = app.Shutdown()
	// Close the channel
	fmt.Println("Application is shutting down...")
}
