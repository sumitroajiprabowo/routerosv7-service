package exception

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/sumitroajiprabowo/routerosv7-service/model"
)

// ErrorBadRequest is a function to return 400 Bad Request
func ErrorBadRequest(ctx *fiber.Ctx, err error) error {
	// Create new ErrorInputResponse instance
	webResponse := model.ErrorInputResponse{
		Code:   fiber.StatusBadRequest,            // Set response code
		Status: "Bad Request",                     // Set response status
		Data:   json.RawMessage(error.Error(err)), // Set response data
	}
	// Return response
	return ctx.Status(fiber.StatusBadRequest).JSON(webResponse)
}

// ErrorNotFound is a function to return 404 Not Found
func ErrorNotFound(ctx *fiber.Ctx, err error) error {
	// Create new WebResponse instance
	webResponse := model.WebResponse{
		Code:    fiber.StatusNotFound, // Set response code
		Status:  "Not Found",          // Set response status
		Message: err.Error(),          // Set response message
	}
	// Return response
	return ctx.Status(fiber.StatusNotFound).JSON(webResponse)
}

// ErrorInternalServerError is a function to return 500 Internal Server Error
func ErrorInternalServerError(ctx *fiber.Ctx, err error) error {
	// Create new WebResponse instance
	webResponse := model.WebResponse{
		Code:    fiber.StatusInternalServerError, // Set response code
		Status:  "Internal Server Error",         // Set response status
		Message: err.Error(),                     // Set response message
	}
	// Return response
	return ctx.Status(fiber.StatusInternalServerError).JSON(webResponse)
}

// ErrorConflict is a function to return 409 Conflict
func ErrorConflict(ctx *fiber.Ctx, err error) error {
	// Create new WebResponse instance
	webResponse := model.WebResponse{
		Code:    fiber.StatusConflict, // Set response code
		Status:  "Conflict",           // Set response status
		Message: err.Error(),          // Set response message
	}
	// Return response
	return ctx.Status(fiber.StatusConflict).JSON(webResponse)
}

// ErrorUnauthorized is a function to return 401 Unauthorized
func ErrorUnauthorized(ctx *fiber.Ctx, err error) error {
	// Create new WebResponse instance
	webResponse := model.WebResponse{
		Code:    fiber.StatusUnauthorized, // Set response code
		Status:  "Unauthorized",           // Set response status
		Message: err.Error(),              // Set response message
	}
	// Return response
	return ctx.Status(fiber.StatusUnauthorized).JSON(webResponse)
}
