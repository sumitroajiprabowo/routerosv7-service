package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/sumitroajiprabowo/routerosv7-restfull-api"
	"github.com/sumitroajiprabowo/routerosv7-service/exception"
	"github.com/sumitroajiprabowo/routerosv7-service/model"
	"github.com/sumitroajiprabowo/routerosv7-service/repository"
	"github.com/sumitroajiprabowo/routerosv7-service/request"
	"github.com/sumitroajiprabowo/routerosv7-service/utils"
)

// PPPSecretService is a contract about something that this service can do
type PPPSecretService interface {
	GetDataByRemoteAddress(ctx *fiber.Ctx, ipAddr, username, password, remoteAddress string) (
		map[string]interface{}, error,
	)
	AddPPPSecret(ctx *fiber.Ctx, request request.CreatePPPoERequest) error
	DeletePPPSecret(ctx *fiber.Ctx, request request.DeletePPPoERequest) error
}

// pppSecretService is a struct that represents the PPPSecretService contract
type pppSecretService struct {
	PPPSecretRepo repository.PPPSecretRepository
}

// NewPPPSecretService is a constructor to create a new PPPSecretService instance
func NewPPPSecretService(pppSecretRepo repository.PPPSecretRepository) PPPSecretService {
	return &pppSecretService{
		PPPSecretRepo: pppSecretRepo,
	}
}

// GetDataByRemoteAddress is a function to get data by remote address from repository
func (p *pppSecretService) GetDataByRemoteAddress(ctx *fiber.Ctx, ipAddr, username, password, remoteAddress string) (
	map[string]interface{}, error,
) {
	// Get data from repository
	data, err := p.PPPSecretRepo.GetDataByRemoteAddress(ctx, ipAddr, username, password, remoteAddress)

	// If error return 404 Not Found
	utils.IfError(err)

	// If success return data and nil
	return data, nil
}

// AddPPPSecret is a function to add PPP Secret to repository
func (p *pppSecretService) AddPPPSecret(ctx *fiber.Ctx, request request.CreatePPPoERequest) error {

	// Check Authentication to Mikrotik Device
	_, err := routerosv7_restfull_api.Auth(ctx.Context(), routerosv7_restfull_api.AuthConfig{
		Host:     request.RouterIpAddr,   // Get Router IP Address from request body
		Username: request.RouterUsername, // Get Router Username from request body
		Password: request.RouterPassword, // Get Router Password from request body
	})

	// If error return 401 Unauthorized
	if err != nil {
		return exception.ErrorUnauthorized(ctx, errors.New("authentication failed")) // Return 401 Unauthorized
	}

	// Check if PPP Secret already exist
	_, err = p.PPPSecretRepo.GetDataByRemoteAddress(ctx, request.RouterIpAddr, request.RouterUsername,
		request.RouterPassword, request.RemoteAddressPPPoE)

	// If error return 404 Not Found
	if err == nil {
		return exception.ErrorConflict(ctx, errors.New("ppp secret already exist"))
	}

	// Add PPP Secret
	_, err = p.PPPSecretRepo.AddPPPSecret(ctx, request)

	// If error return 500 Internal Server Error
	if err != nil {
		return exception.ErrorInternalServerError(ctx, err)
	}

	// If success return 201 Created
	response := model.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Add PPP Secret",
	}

	// If success return 201 Created
	return ctx.Status(fiber.StatusOK).JSON(response)
}

// DeletePPPSecret is a function to delete PPP Secret from repository by remote address
func (p *pppSecretService) DeletePPPSecret(ctx *fiber.Ctx, request request.DeletePPPoERequest) error {
	// Check Authentication to Mikrotik Device
	_, err := routerosv7_restfull_api.Auth(ctx.Context(), routerosv7_restfull_api.AuthConfig{
		Host:     request.RouterIpAddr,   // Get Router IP Address from request body
		Username: request.RouterUsername, // Get Router Username from request body
		Password: request.RouterPassword, // Get Router Password from request body
	})

	// If error return 401 Unauthorized
	if err != nil {
		return exception.ErrorUnauthorized(ctx, errors.New("authentication failed")) // Return 401 Unauthorized
	}

	// Check if PPP Secret already exist
	data, err := p.PPPSecretRepo.GetDataByRemoteAddress(ctx, request.RouterIpAddr, request.RouterUsername,
		request.RouterPassword, request.RemoteAddressPPPoE)

	// If error return 404 Not Found
	if err != nil {
		return exception.ErrorNotFound(ctx, errors.New("ppp secret not found")) // Return 404 Not Found
	}

	// Get Data .id from map data
	id := data[".id"].(string)

	// Delete PPP Secret
	err = p.PPPSecretRepo.DeletePPPSecret(ctx, id, request)

	// If error return 500 Internal Server Error
	if err != nil {
		return exception.ErrorInternalServerError(ctx, err) // Return 500 Internal Server Error
	}

	// If success return 204 No Content
	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}
