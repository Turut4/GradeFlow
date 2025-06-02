package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Turut4/GradeFlow/internal/auth"
	"github.com/Turut4/GradeFlow/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
)

type application struct {
	config        config
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

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Set-Cookie"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Use(app.authTokenMiddleware)
			r.Get("/{userID}", app.GetUserHandler)
			r.Get("/me", app.getMe)
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/users", app.registerUserHandler)
			r.Post("/token", app.createTokenHandler)
			r.Post("/destroy-token", app.removeTokenHandler)
			r.Post("/register", app.registerUserHandler)
		})

		r.Route("/exams", func(r chi.Router) {
			r.Use(app.authTokenMiddleware)
			r.Post("/", app.createExamHandler)
			r.Get("/{examID}", app.GetExamHandler)
			r.Get("/{examID}/answer-sheet", app.GetAnswerSheetHandler)
		})

		r.Route("/answer-sheet", func(r chi.Router) {
			r.Post("/process-gabarito", app.processAnswersSheet)
		})
	})

	return r
}

func (app *application) run() error {
	mux := app.mount()
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)
	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		app.logger.Infow("signal caught", "signal", s.String())

		shutdown <- srv.Shutdown(ctx)
	}()
	app.logger.Infow("Server running", "addr", app.config.addr, "env", app.config.env)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown
	if err != nil {
		return err
	}

	app.logger.Infow("server has stoped", "addr", app.config.addr, "env", app.config.env)
	return nil

}
