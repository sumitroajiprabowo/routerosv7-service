package repository

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sumitroajiprabowo/routerosv7-restfull-api"
	"github.com/sumitroajiprabowo/routerosv7-service/request"
	"github.com/sumitroajiprabowo/routerosv7-service/utils"
)

// PPPSecretRepository is a contract about something that this repository can do
type PPPSecretRepository interface {
	GetDataByRemoteAddress(ctx *fiber.Ctx, ipAddr, username, password, remoteAddress string) (
		map[string]interface{}, error,
	) // Get data by remote address from router os v7 device
	AddPPPSecret(ctx *fiber.Ctx, request request.CreatePPPoERequest) (
		map[string]interface{}, error,
	) // Add PPP Secret to router os v7 device
	DeletePPPSecret(
		ctx *fiber.Ctx, id string, request request.DeletePPPoERequest,
	) error // Delete PPP Secret from router os v7 device
}

// pppSecretRepository is a struct that represents the PPPSecretRepository contract
type pppSecretRepository struct{}

// NewPPPSecretRepository is a constructor to create a new PPPSecretRepository instance
func NewPPPSecretRepository() PPPSecretRepository {
	return &pppSecretRepository{}
}

func (p *pppSecretRepository) GetDataByRemoteAddress(ctx *fiber.Ctx, ipAddr, username, password, remoteAddress string) (
	map[string]interface{}, error,
) {
	// Command to get ppp secret by remote address
	cmd := fmt.Sprintf("ppp/secret?remote-address=%s", remoteAddress)

	/*
		Get data from router os v7 device using routerosv7_restfull_api package and store it in data variable as
		interface type and err variable as error type if error occurs when getting data from router os v7 device using
		routerosv7_restfull_api package then return error message and if success return data and nil
	*/
	data, err := routerosv7_restfull_api.Print(ctx.Context(), ipAddr, username, password, cmd)
	utils.IfError(err) // If error return 404 Not Found

	// Type asserts the response to []interface{} since that's what's expected
	response := data.([]interface{})

	// If response length is 0 then return error message
	if len(response) == 0 {
		return nil, errors.New("failed to get data") // Return error message
	}

	// If response length is more than 1 then return error message
	return response[0].(map[string]interface{}), nil
}

// AddPPPSecret is a function to add PPP Secret to router os v7 device
func (p *pppSecretRepository) AddPPPSecret(
	ctx *fiber.Ctx, request request.CreatePPPoERequest,
) (map[string]interface{}, error) {

	// Create a new AddRequest using the constructor
	payload := fmt.Sprintf(`{
       "name": "%s",
       "password": "%s",
       "profile": "%s",
       "remote-address": "%s",
       "service": "pppoe"
	}`, request.UsernamePPoE, request.PasswordPPPoE, request.ProfilePPPoE, request.RemoteAddressPPPoE)

	// Command to add ppp secret
	cmd := "ppp/secret"

	/*
		Add ppp secret to a router os v7 device using routerosv7_restfull_api package and store it in data variable as
		map[string]interface{} type and err variable as error type if error occurs when adding ppp secret to a router os
		v7 device using routerosv7_restfull_api package then return error message and if success return data and nil
	*/
	data, err := routerosv7_restfull_api.Add(ctx.Context(), request.RouterIpAddr, request.RouterUsername,
		request.RouterPassword, cmd, []byte(payload))

	utils.IfError(err) // If error return 404 Not Found

	// Type asserts the response to map[string]interface{} since that's what's expected
	response, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response format: %v", data) // Return error message
	}

	// If response length is 0 then return error message
	return response, nil
}

func (p *pppSecretRepository) DeletePPPSecret(ctx *fiber.Ctx, id string, request request.DeletePPPoERequest) error {

	// Create a new DeleteRequest using the constructor and set the id
	cmd := fmt.Sprintf("ppp/secret/%s", id)

	// Create a new DeleteRequest using the constructor
	_, err := routerosv7_restfull_api.Remove(ctx.Context(), request.RouterIpAddr, request.RouterUsername,
		request.RouterPassword, cmd)
	if err == nil {
		return err // Return error message
	}

	// Return nil if success
	return nil
}
