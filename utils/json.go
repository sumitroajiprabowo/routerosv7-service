package utils

import "github.com/gofiber/fiber/v2"

func ParseRequestBody(ctx *fiber.Ctx, requestBody interface{}) error {
	err := ctx.BodyParser(requestBody)
	IfError(err)
	return nil
}
