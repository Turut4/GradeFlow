package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Turut4/GradeFlow/internal/store"
	"github.com/golang-jwt/jwt/v5"
)

type userKey string

const userCtx userKey = "user"

func (app *application) authTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			app.unauthorizedErrorResponse(
				w,
				r,
				fmt.Errorf("authorization header is missing"),
			)
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			app.unauthorizedErrorResponse(
				w,
				r,
				fmt.Errorf("authorization header is malformed %s ", authHeader),
			)
			return
		}

		token, err := app.authenticator.ValidateToken(parts[1])
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		sub, ok := claims["sub"].(float64)
		if !ok {
			app.unauthorizedErrorResponse(
				w,
				r,
				fmt.Errorf(
					"invalid 'sub' claim type: expected float64, got %T",
					claims["sub"],
				),
			)
			return
		}

		userID := uint(sub)
		ctx := r.Context()
		user, err := app.store.Users.GetByID(ctx, userID)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}
		ctx = context.WithValue(ctx, userCtx, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) checkRolePrecedence(
	ctx context.Context,
	user *store.User,
	roleName string,
) (bool, error) {
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

func parseUserFromCtx(r *http.Request) *store.User {
	user, ok := r.Context().Value(userCtx).(*store.User)

	if !ok {
		return nil
	}

	return user
}
