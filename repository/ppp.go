package repository

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sumitroajiprabowo/routerosv7-restfull-api"
	"github.com/sumitroajiprabowo/routerosv7-service/request"
	"github.com/sumitroajiprabowo/routerosv7-service/utils"
)

type PPPSecretRepository interface {
	GetDataByRemoteAddress(ctx *fiber.Ctx, ipAddr, username, password, remoteAddress string) (
		map[string]interface{}, error,
	)
	AddPPPSecret(ctx *fiber.Ctx, request request.CreatePPPoERequest) (map[string]interface{}, error)
	DeletePPPSecret(ctx *fiber.Ctx, id string, request request.DeletePPPoERequest) error
}

type pppSecretRepository struct{}

func NewPPPSecretRepository() PPPSecretRepository {
	return &pppSecretRepository{}
}

func (p *pppSecretRepository) GetDataByRemoteAddress(ctx *fiber.Ctx, ipAddr, username, password, remoteAddress string) (
	map[string]interface{}, error,
) {
	cmd := fmt.Sprintf("ppp/secret?remote-address=%s", remoteAddress)
	data, err := routerosv7_restfull_api.Print(ctx.Context(), ipAddr, username, password, cmd)
	utils.IfError(err)

	response := data.([]interface{})

	if len(response) == 0 {
		return nil, errors.New("failed to get data")
	}

	return response[0].(map[string]interface{}), nil
}

func (p *pppSecretRepository) AddPPPSecret(
	ctx *fiber.Ctx, request request.CreatePPPoERequest,
) (map[string]interface{}, error) {
	payload := fmt.Sprintf(`{
       "name": "%s",
       "password": "%s",
       "profile": "%s",
       "remote-address": "%s",
       "service": "pppoe"
	}`, request.UsernamePPoE, request.PasswordPPPoE, request.ProfilePPPoE, request.RemoteAddressPPPoE)

	cmd := "ppp/secret"

	data, err := routerosv7_restfull_api.Add(ctx.Context(), request.RouterIpAddr, request.RouterUsername,
		request.RouterPassword, cmd, []byte(payload))

	utils.IfError(err)

	// Type asserts the response to map[string]interface{} since that's what's expected
	response, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response format: %v", data)
	}

	return response, nil
}

func (p *pppSecretRepository) DeletePPPSecret(ctx *fiber.Ctx, id string, request request.DeletePPPoERequest) error {

	// Create a new DeleteRequest using the constructor and set the id
	cmd := fmt.Sprintf("ppp/secret/%s", id)

	// Create a new DeleteRequest using the constructor
	_, err := routerosv7_restfull_api.Remove(ctx.Context(), request.RouterIpAddr, request.RouterUsername,
		request.RouterPassword, cmd)

	utils.IfError(err)

	return nil
}
