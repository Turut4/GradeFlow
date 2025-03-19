package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/Turut4/GradeFlow/internal/store"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const userCtx = "user"

func (api *application) authTokenMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return api.unauthorizedResponse(c, fmt.Errorf("authorization header is missing"))
	}

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return api.unauthorizedResponse(c, fmt.Errorf("authorization header is malformed %s ", authHeader))
	}

	token, err := api.authenticator.ValidateToken(parts[1])
	if err != nil {
		return api.unauthorizedResponse(c, err)
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	sub, ok := claims["sub"].(float64)
	if !ok {
		return api.unauthorizedResponse(c, fmt.Errorf("invalid 'sub' claim type: expected float64, got %T", claims["sub"]))
	}

	userID := int64(sub)
	user, err := api.store.Users.GetByID(c.UserContext(), userID)
	if err != nil {
		return api.unauthorizedResponse(c, err)
	}

	c.Locals(userCtx, user)
	return c.Next()
}

func (app *application) checkRolePrecedence(ctx context.Context, user *store.User, roleName string) (bool, error) {
	role, err := app.store.Roles.GetByName(ctx, roleName)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			return false, fmt.Errorf("role don't exists")
		default:
			return false, err
		}
	}

	return user.Role.Level >= role.Level, nil
}
