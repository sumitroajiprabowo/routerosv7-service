package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sumitroajiprabowo/routerosv7-service/handler"
	"github.com/sumitroajiprabowo/routerosv7-service/repository"
	"github.com/sumitroajiprabowo/routerosv7-service/router"
	"github.com/sumitroajiprabowo/routerosv7-service/service"
	"github.com/sumitroajiprabowo/routerosv7-service/utils"
	"time"
)

func SetupFiberApp() *fiber.App {

	// Create a new fiber app with config timeout 10 seconds
	app := fiber.New(fiber.Config{
		ReadTimeout: 10 * time.Second,
	})

	app.Use(recover.New()) // Enable recover

	// Setup logger
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n", // format log
		TimeFormat: "02-Jan-2006 15:04:05",                                 // time format
		TimeZone:   "Asia/Jakarta",                                         // time zone
	}))

	// Setup Swagger
	app.Use(swagger.New(swagger.Config{
		BasePath: "/",                   // Base path of docs endpoint
		FilePath: "./docs/swagger.json", // Filepath of docs endpoint
		Path:     "swagger",             // Path of docs endpoint
	}))

	app.Use(cors.New()) // Enable CORS

	return app // Return fiber app
}

// InitServerHTTP is a function to initialize server
func InitServerHTTP() {
	app := SetupFiberApp()

	validate := validator.New()

	// Initialize Repository
	pppRepo := repository.NewPPPSecretRepository()

	// Initialize Service
	pppService := service.NewPPPSecretService(pppRepo)

	// Initialize Handler
	pppHandler := handler.NewPPPSecretHandler(pppService, validate)

	// Initialize Router
	pppRouter := router.NewPPPRouter(pppHandler)

	// Setup Routes
	pppRouter.SetupPPPRoutes(app)

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	// Use graceful shutdown from utils
	utils.GracefulShutdown(app)
}
