package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sumitroajiprabowo/routerosv7-service/handler"
	"github.com/sumitroajiprabowo/routerosv7-service/repository"
	"github.com/sumitroajiprabowo/routerosv7-service/router"
	"github.com/sumitroajiprabowo/routerosv7-service/service"
	"github.com/sumitroajiprabowo/routerosv7-service/utils"
)

func SetupFiberApp() *fiber.App {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
	}))

	app.Use(cors.New())

	return app
}

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

	// Use graceful shutdown from utils
	utils.GracefulShutdown(app)
}
