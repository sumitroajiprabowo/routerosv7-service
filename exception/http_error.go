package exception

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/sumitroajiprabowo/routerosv7-service/model"
)

func ErrorBadRequest(ctx *fiber.Ctx, err error) error {
	webResponse := model.ErrorInputResponse{
		Code:   fiber.StatusBadRequest,
		Status: "Bad Request",
		Data:   json.RawMessage(error.Error(err)),
	}
	return ctx.Status(fiber.StatusBadRequest).JSON(webResponse)
}

func ErrorNotFound(ctx *fiber.Ctx, err error) error {
	webResponse := model.WebResponse{
		Code:    fiber.StatusNotFound,
		Status:  "Not Found",
		Message: err.Error(),
	}
	return ctx.Status(fiber.StatusNotFound).JSON(webResponse)
}

func ErrorInternalServerError(ctx *fiber.Ctx, err error) error {
	webResponse := model.WebResponse{
		Code:    fiber.StatusInternalServerError,
		Status:  "Internal Server Error",
		Message: err.Error(),
	}
	return ctx.Status(fiber.StatusInternalServerError).JSON(webResponse)
}

func ErrorConflict(ctx *fiber.Ctx, err error) error {
	webResponse := model.WebResponse{
		Code:    fiber.StatusConflict,
		Status:  "Conflict",
		Message: err.Error(),
	}
	return ctx.Status(fiber.StatusConflict).JSON(webResponse)
}

func ErrorUnauthorized(ctx *fiber.Ctx, err error) error {
	webResponse := model.WebResponse{
		Code:    fiber.StatusUnauthorized,
		Status:  "Unauthorized",
		Message: err.Error(),
	}
	return ctx.Status(fiber.StatusUnauthorized).JSON(webResponse)
}

func ErrorServiceUnavailable(ctx *fiber.Ctx, err error) error {
	webResponse := model.WebResponse{
		Code:    fiber.StatusServiceUnavailable,
		Status:  "Service Unavailable",
		Message: err.Error(),
	}
	return ctx.Status(fiber.StatusServiceUnavailable).JSON(webResponse)
}
