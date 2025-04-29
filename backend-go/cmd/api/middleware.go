package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Turut4/GradeFlow/internal/store"
	"github.com/golang-jwt/jwt/v5"
)

type userKey string

const userCtx userKey = "user"

func (app *application) authTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			app.unauthorizedErrorResponse(
				w,
				r,
				fmt.Errorf("cookie de autenticação ausente"),
			)
			return
		}
		token, err := app.authenticator.ValidateToken(cookie.Value)
		if err != nil {
			app.forbiddenErrorResponse(w, r)
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
		user, err := app.store.Users.GetByID(r.Context(), userID)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}
		ctx := context.WithValue(r.Context(), userCtx, user)

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
