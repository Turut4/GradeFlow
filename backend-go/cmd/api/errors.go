package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (api *application) processError(c *fiber.Ctx, code int, err error) error {
	api.logger.Errorw("process error", "method", c.Method(), "path", c.Path(), "err", err.Error())
	return writeJSONError(c, code, "the server encountered a problem")
}
func (api *application) internalError(c *fiber.Ctx, err error) error {
	api.logger.Errorw("internal err", "method", c.Method(), "path", c.Path(), "err", err.Error())

	return writeJSONError(c, http.StatusInternalServerError, "the server encountered a problem")
}

func (api *application) notFoundResponse(c *fiber.Ctx, err error) error {
	api.logger.Warnw("not found", "method", c.Method(), "path", c.Path(), "err", err.Error())

	return writeJSONError(c, http.StatusNotFound, "not found")
}

func (api *application) badResquestResponse(c *fiber.Ctx, err error) error {
	api.logger.Warnw("bad request", "method", c.Method(), "path", c.Path(), "err", err.Error())

	return writeJSONError(c, http.StatusBadRequest, err.Error())
}

func (api *application) forbiddenResponse(c *fiber.Ctx, err error) error {
	api.logger.Warnw("forbidden", "method", c.Method(), "path", c.Path(), "err", err.Error())

	return writeJSONError(c, http.StatusForbidden, "bad request")
}

func (api *application) unauthorizedResponse(c *fiber.Ctx, err error) error {
	api.logger.Warnw("unauthorized", "method", c.Method(), "path", c.Path(), "err", err.Error())

	return writeJSONError(c, http.StatusUnauthorized, "unauthorized")
}
