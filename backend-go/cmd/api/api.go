package main

import (
	"log"
	"time"

	"github.com/Turut4/GradeFlow/internal/auth"
	"github.com/Turut4/GradeFlow/internal/store"
	"github.com/gofiber/fiber/v2"
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

func (api *application) mount() *fiber.App {
	app := fiber.New()

	app.Route("/api/v1", func(router fiber.Router) {
		router.Route("/users", func(router fiber.Router) {
			router.Use(api.authTokenMiddleware)
			router.Get("/:userID", api.GetUserHandler)
		})

		router.Route("/authentication", func(router fiber.Router) {
			router.Post("/users", api.registerUserHandler)
			router.Post("/token", api.createTokenHandler)
		})

		router.Route("/exams", func(router fiber.Router) {
			router.Use(api.authTokenMiddleware)
			router.Post("/", api.createExamHandler)
			router.Get("/:examID", api.GetExamPDFHandler)
		})

		router.Route("/answer-sheet", func(router fiber.Router) {
			router.Post("/process-gabarito", api.processAnswersSheet)
		})
	})

	return app
}

func (api *application) run() {
	app := api.mount()

	api.logger.Infof("Server running %s\n", api.cfg.addr)
	if err := app.Listen(api.cfg.addr); err != nil {
		log.Fatalf("Error initializing server: %v", err)
	}
}
