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

type PPPSecretHandler interface {
	AddPPPSecret(c *fiber.Ctx) error
	DeletePPPSecret(c *fiber.Ctx) error
}

type pppSecretHandler struct {
	PPPSecretService service.PPPSecretService
	Validate         *validator.Validate
}

func NewPPPSecretHandler(pppSecretService service.PPPSecretService, validate *validator.Validate) PPPSecretHandler {
	return &pppSecretHandler{
		PPPSecretService: pppSecretService,
		Validate:         validate,
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
	req := new(request.CreatePPPoERequest)

	// Parse body request to struct
	err := utils.ParseRequestBody(ctx, req)
	utils.IfError(err)

	// Validate request data
	msg, err := utils.ValidateRequest(req) // Validate request body
	if err != nil {
		return exception.ErrorBadRequest(ctx, errors.New(msg))
	}

	// Add PPP Secret
	err = p.PPPSecretService.AddPPPSecret(ctx, *req)
	if err != nil {
		return exception.ErrorInternalServerError(ctx, err)
	}

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
	req := new(request.DeletePPPoERequest)

	// Parse body request to struct
	err := utils.ParseRequestBody(ctx, req)
	utils.IfError(err)

	// Validate request data
	msg, err := utils.ValidateRequest(req) // Validate request body
	if err != nil {
		return exception.ErrorBadRequest(ctx, errors.New(msg))
	}

	// Delete PPP Secret
	err = p.PPPSecretService.DeletePPPSecret(ctx, *req)
	if err != nil {
		return exception.ErrorInternalServerError(ctx, err)
	}

	return nil
}
