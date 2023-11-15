// utils/graceful.go

package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// GracefulShutdown menerima aplikasi Fiber dan menangani proses graceful shutdown
func GracefulShutdown(app *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	// This blocks the main thread until an interrupt is received
	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()
	fmt.Println("Application is shutting down...")
}
