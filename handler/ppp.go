package handler

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sumitroajiprabowo/routerosv7-service/exception"
	"github.com/sumitroajiprabowo/routerosv7-service/request"
	"github.com/sumitroajiprabowo/routerosv7-service/service"
	"github.com/sumitroajiprabowo/routerosv7-service/utils"
)

// PPPSecretHandler is a contract about something that this handler can do
type PPPSecretHandler interface {
	AddPPPSecret(c *fiber.Ctx) error    // Add PPP Secret
	DeletePPPSecret(c *fiber.Ctx) error // Delete PPP Secret
}

// pppSecretHandler is a struct that represents the PPPSecretHandler contract
type pppSecretHandler struct {
	PPPSecretService service.PPPSecretService // PPPSecretService contract
	Validate         *validator.Validate      // Validator contract
}

// NewPPPSecretHandler is a constructor to create a new PPPSecretHandler instance
func NewPPPSecretHandler(pppSecretService service.PPPSecretService, validate *validator.Validate) PPPSecretHandler {
	return &pppSecretHandler{
		PPPSecretService: pppSecretService, // PPPSecretService contract
		Validate:         validate,         // Validator contract
	}
}

// AddPPPSecret Godoc
// @Summary Add PPP Secret
// @Description Add PPP Secret
// @Tags PPP
// @Accept json
// @Produce json
// @Param pppSecret body request.CreatePPPoERequest true "Add PPP Secret"
// @Success 201 {object} model.WebResponse "Created
// @Failure 400 {object} model.WebResponse "Bad Request"
// @Failure 401 {object} model.WebResponse "Unauthorized"
// @Failure 404 {object} model.WebResponse "Not Found"
// @Failure 409 {object} model.WebResponse "Conflict"
// @Failure 500 {object} model.WebResponse "Internal Server Error"
// @Router /ppp/secret/add [post]
func (p *pppSecretHandler) AddPPPSecret(ctx *fiber.Ctx) error {
	// Get data from body request
	req := new(request.CreatePPPoERequest) // Create new CreatePPPoERequest instance

	// Parse body request to struct
	err := utils.ParseRequestBody(ctx, req) // Parse body request to struct
	utils.IfError(err)                      // If error return 400 Bad Request

	// Validate request data
	msg, err := utils.ValidateRequest(req) // Validate request body

	// If error return 400 Bad Request
	if err != nil {
		return exception.ErrorBadRequest(ctx, errors.New(msg)) // Return error message
	}

	// Add PPP Secret
	err = p.PPPSecretService.AddPPPSecret(ctx, *req) // Add PPP Secret to router os v7 device
	if err != nil {
		return exception.ErrorInternalServerError(ctx, err) // Return 500 Internal Server Error
	}

	// Return 201 Created if success from service
	return nil
}

// DeletePPPSecret Godoc
// @Summary Delete PPP Secret
// @Description Delete PPP Secret
// @Tags PPP
// @Accept json
// @Produce json
// @Param pppSecret body request.DeletePPPoERequest true "Delete PPP Secret"
// @Success 204 {object} model.WebResponse "No Content"
// @Failure 400 {object} model.WebResponse "Bad Request"
// @Failure 401 {object} model.WebResponse "Unauthorized"
// @Failure 404 {object} model.WebResponse "Not Found"
// @Failure 500 {object} model.WebResponse "Internal Server Error"
// @Router /ppp/secret/delete [delete]
func (p *pppSecretHandler) DeletePPPSecret(ctx *fiber.Ctx) error {
	// Get data from body request
	req := new(request.DeletePPPoERequest) // Create new DeletePPPoERequest instance

	// Parse body request to struct
	err := utils.ParseRequestBody(ctx, req) // Parse body request to struct

	// If error return 400 Bad Request
	utils.IfError(err) // If error return 400 Bad Request

	// Validate request data
	msg, err := utils.ValidateRequest(req) // Validate request body

	// If error return 400 Bad Request
	if err != nil {
		return exception.ErrorBadRequest(ctx, errors.New(msg)) // Return error message
	}

	// Delete PPP Secret
	err = p.PPPSecretService.DeletePPPSecret(ctx, *req)

	// If error return 500 Internal Server Error
	if err != nil {
		return exception.ErrorInternalServerError(ctx, err) // Return 500 Internal Server Error
	}

	// Return 204 No Content if success from service
	return nil
}
