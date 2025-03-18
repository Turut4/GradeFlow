package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func (api *application) jsonResponse(c *fiber.Ctx, code int, data any) error {
	type envelop struct {
		Data any `json:"data"`
	}
	return c.Status(code).JSON(&envelop{Data: data})
}

func writeJSONError(c *fiber.Ctx, code int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}

	return c.Status(code).JSON(&envelope{Error: message})
}
