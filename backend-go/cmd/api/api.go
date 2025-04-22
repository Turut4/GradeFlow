package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Turut4/GradeFlow/internal/auth"
	"github.com/Turut4/GradeFlow/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type application struct {
	cfg           config
	logger        *zap.SugaredLogger
	store         store.Storage
	authenticator auth.Authenticator
}

type ocrConfig struct {
	addr string
}

type config struct {
	addr string
	db   dbConfig
	auth authConfig
	env  string
	ocr  ocrConfig
}

type authConfig struct {
	basic basicConfig
	token tokenConfig
}

type tokenConfig struct {
	secret string
	exp    time.Duration
	iss    string
}

type basicConfig struct {
	user string
	pass string
}

type dbConfig struct {
	addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Use(app.authTokenMiddleware)
			r.Get("/:userID", app.GetUserHandler)
		})

		r.Route("/authentication", func(r chi.Router) {
			r.Post("/users", app.registerUserHandler)
			r.Post("/token", app.createTokenHandler)
		})

		r.Route("/exams", func(r chi.Router) {
			r.Use(app.authTokenMiddleware)
			r.Post("/", app.createExamHandler)
			r.Get("/:examID", app.GetExamHandler)
			r.Get("/:examID/answer-sheet", app.GetAnswerSheetHandler)
		})

		r.Route("/answer-sheet", func(r chi.Router) {
			r.Post("/process-gabarito", app.processAnswersSheet)
		})
	})

	return r
}

func (app *application) run() {
	mux := app.mount()
	srv := &http.Server{
		Addr:         app.cfg.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app.logger.Infof("Server running %s\n", app.cfg.addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error initializing server: %v", err)
	}
}
