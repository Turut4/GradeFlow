package main

import (
	"net/http"
	"time"

	"github.com/Turut4/GradeFlow/internal/store"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"min=3,max=72,required"`
	Password string `json:"password" validate:"required,min=3,max=72"`
	Email    string `json:"email" validate:"required,email"`
}

func (api *application) registerUserHandler(c *fiber.Ctx) error {
	var payload RegisterUserPayload

	if err := c.BodyParser(&payload); err != nil {
		return api.badRequestResponse(c, err)
	}

	if err := Validate.Struct(payload); err != nil {
		return api.badRequestResponse(c, err)
	}

	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		Role: store.Role{
			Name: "user",
		},
	}

	if err := api.store.Users.Create(c.Context(), user); err != nil {
		return api.internalError(c, err)
	}

	return api.jsonResponse(c, http.StatusCreated, nil)
}

type CreateUserTokenPayload struct {
	Email    string `json:"email" validate:"required,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

func (api *application) createTokenHandler(c *fiber.Ctx) error {
	var payload CreateUserTokenPayload
	if err := c.BodyParser(&payload); err != nil {
		return api.badRequestResponse(c, err)
	}

	if err := Validate.Struct(payload); err != nil {
		return api.badRequestResponse(c, err)
	}

	user, err := api.store.Users.GetByEmail(c.Context(), payload.Email)
	if err != nil {
		return api.badRequestResponse(c, err)
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(api.cfg.auth.token.exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": api.cfg.auth.token.iss,
		"aud": api.cfg.auth.token.iss,
	}

	token, err := api.authenticator.GenerateToken(claims)
	if err != nil {
		return api.internalError(c, err)
	}

	return api.jsonResponse(c, http.StatusCreated, token)
}
