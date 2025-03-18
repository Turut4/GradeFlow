package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (api *application) internalError(c *fiber.Ctx, error error) error {
	api.logger.Errorw("internal error", "method", c.Method(), "path", c.Path(), "error", error.Error())

	return writeJSONError(c, http.StatusInternalServerError, "the server encountered a problem")
}

func (api *application) notFoundResponse(c *fiber.Ctx, error error) error {
	api.logger.Warnw("not found", "method", c.Method(), "path", c.Path(), "error", error.Error())

	return writeJSONError(c, http.StatusNotFound, "not found")
}

func (api *application) badResquestResponse(c *fiber.Ctx, error error) error {
	api.logger.Warnw("bad request", "method", c.Method(), "path", c.Path(), "error", error.Error())

	return writeJSONError(c, http.StatusBadRequest, "bad request")
}

func (api *application) forbiddenResponse(c *fiber.Ctx, error error) error {
	api.logger.Warnw("forbidden", "method", c.Method(), "path", c.Path(), "error", error.Error())

	return writeJSONError(c, http.StatusForbidden, "bad request")
}

func (api *application) unauthorizedResponse(c *fiber.Ctx, error error) error {
	api.logger.Warnw("unauthorized", "method", c.Method(), "path", c.Path(), "error", error.Error())

	return writeJSONError(c, http.StatusUnauthorized, "unauthorized")
}
