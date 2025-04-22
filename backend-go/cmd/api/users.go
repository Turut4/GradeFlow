package main

import (
	"net/http"
	"strconv"

	"github.com/Turut4/GradeFlow/internal/store"
	"github.com/go-chi/chi/v5"
)

func (app *application) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, string(userCtx))
	userID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user, err := app.store.Users.GetByID(r.Context(), userID)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
		}
	}

	if err := app.jsonResponse(w, http.StatusCreated, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
