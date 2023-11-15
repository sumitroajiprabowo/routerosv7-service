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

type PPPSecretService interface {
	GetDataByRemoteAddress(ctx *fiber.Ctx, ipAddr, username, password, remoteAddress string) (
		map[string]interface{}, error,
	)
	AddPPPSecret(ctx *fiber.Ctx, request request.CreatePPPoERequest) error
	DeletePPPSecret(ctx *fiber.Ctx, request request.DeletePPPoERequest) error
}

type pppSecretService struct {
	PPPSecretRepo repository.PPPSecretRepository
}

func NewPPPSecretService(pppSecretRepo repository.PPPSecretRepository) PPPSecretService {
	return &pppSecretService{
		PPPSecretRepo: pppSecretRepo,
	}
}

func (p *pppSecretService) GetDataByRemoteAddress(ctx *fiber.Ctx, ipAddr, username, password, remoteAddress string) (
	map[string]interface{}, error,
) {
	data, err := p.PPPSecretRepo.GetDataByRemoteAddress(ctx, ipAddr, username, password, remoteAddress)
	utils.IfError(err)
	return data, nil
}

func (p *pppSecretService) AddPPPSecret(ctx *fiber.Ctx, request request.CreatePPPoERequest) error {

	// Check Authentication to Mikrotik Device
	_, err := routerosv7_restfull_api.Auth(ctx.Context(), routerosv7_restfull_api.AuthConfig{
		Host:     request.RouterIpAddr,
		Username: request.RouterUsername,
		Password: request.RouterPassword,
	})

	if err != nil {
		return exception.ErrorUnauthorized(ctx, errors.New("authentication failed"))
	}

	// Check if PPP Secret already exist
	_, err = p.PPPSecretRepo.GetDataByRemoteAddress(ctx, request.RouterIpAddr, request.RouterUsername,
		request.RouterPassword, request.RemoteAddressPPPoE)
	if err == nil {
		return exception.ErrorConflict(ctx, errors.New("ppp secret already exist"))
	}

	_, err = p.PPPSecretRepo.AddPPPSecret(ctx, request)
	if err != nil {
		return exception.ErrorInternalServerError(ctx, err)
	}
	response := model.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Add PPP Secret",
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (p *pppSecretService) DeletePPPSecret(ctx *fiber.Ctx, request request.DeletePPPoERequest) error {
	// Check Authentication to Mikrotik Device
	_, err := routerosv7_restfull_api.Auth(ctx.Context(), routerosv7_restfull_api.AuthConfig{
		Host:     request.RouterIpAddr,
		Username: request.RouterUsername,
		Password: request.RouterPassword,
	})

	// If error return 401 Unauthorized
	if err != nil {
		return exception.ErrorUnauthorized(ctx, errors.New("authentication failed"))
	}

	// Check if PPP Secret already exist
	data, err := p.PPPSecretRepo.GetDataByRemoteAddress(ctx, request.RouterIpAddr, request.RouterUsername,
		request.RouterPassword, request.RemoteAddressPPPoE)

	// If error return 404 Not Found
	if err != nil {
		return exception.ErrorNotFound(ctx, errors.New("ppp secret not found"))
	}

	// Get Data .id from map data
	id := data[".id"].(string)

	// Delete PPP Secret
	err = p.PPPSecretRepo.DeletePPPSecret(ctx, id, request)

	// If error return 500 Internal Server Error
	if err != nil {
		return exception.ErrorInternalServerError(ctx, err)
	}

	// If success return 204 No Content
	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}
