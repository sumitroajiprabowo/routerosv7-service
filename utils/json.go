package utils

import "github.com/gofiber/fiber/v2"

// ParseRequestBody function to parse request body
func ParseRequestBody(ctx *fiber.Ctx, requestBody interface{}) error {
	// Parse body request to struct
	err := ctx.BodyParser(requestBody)
	// If error return error
	IfError(err)
	// If no error return nil
	return nil
}
