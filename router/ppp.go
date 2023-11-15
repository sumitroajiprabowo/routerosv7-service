package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sumitroajiprabowo/routerosv7-service/handler"
)

type PPPRouter struct {
	pppSecretHandler handler.PPPSecretHandler
}

func NewPPPRouter(pppSecretHandler handler.PPPSecretHandler) *PPPRouter {
	return &PPPRouter{
		pppSecretHandler: pppSecretHandler,
	}
}

func (r *PPPRouter) SetupPPPRoutes(app *fiber.App) {
	pppGroup := app.Group("/api/v1/ppp")
	pppGroup.Post("/secret/add", r.pppSecretHandler.AddPPPSecret)
	pppGroup.Delete("/secret/delete", r.pppSecretHandler.DeletePPPSecret)
}
