package main

import (
	"net/http"
	"time"

	"github.com/Turut4/GradeFlow/internal/store"
	"github.com/golang-jwt/jwt/v5"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"min=3,max=72,required"`
	Password string `json:"password" validate:"required,min=3,max=72"`
	Email    string `json:"email"    validate:"required,email"`
}

func (app *application) registerUserHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var payload RegisterUserPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		Role: store.Role{
			Name: "user",
		},
	}

	if err := app.store.Users.Create(r.Context(), user); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type CreateUserTokenPayload struct {
	Email    string `json:"email"    validate:"required,max=255,email"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

func (app *application) createTokenHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var payload CreateUserTokenPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user, err := app.store.Users.GetByEmail(r.Context(), payload.Email)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.invalidCredentialsResponse(w, r)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := user.ComparePassword(payload.Password); err != nil {
		app.forbiddenErrorResponse(w, r)
		return
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(app.config.auth.token.exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": app.config.auth.token.iss,
		"aud": app.config.auth.token.iss,
	}

	token, err := app.authenticator.GenerateToken(claims)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Expires:  time.Now().Add(app.config.auth.token.exp),
	})

	if err := app.jsonResponse(w, http.StatusOK, "login realizado com sucesso"); err != nil {
		app.internalServerError(w, r, err)
	}
}
