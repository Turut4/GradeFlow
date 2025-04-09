package main

import (
	"net/http"
	"strconv"

	"github.com/Turut4/GradeFlow/internal/store"
	"github.com/gofiber/fiber/v2"
)

func (api *application) GetUserHandler(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Params("userID"), 10, 64)

	if err != nil {
		return api.badRequestResponse(c, err)
	}

	user, err := api.store.Users.GetByID(c.Context(), userID)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			return api.notFoundResponse(c, err)
		}
	}

	return api.jsonResponse(c, http.StatusOK, &user)
}
