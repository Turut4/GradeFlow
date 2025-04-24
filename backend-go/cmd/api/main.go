package main

import (
	"time"

	"github.com/Turut4/GradeFlow/internal/auth"
	"github.com/Turut4/GradeFlow/internal/db"
	"github.com/Turut4/GradeFlow/internal/env"
	"github.com/Turut4/GradeFlow/internal/store"
	"go.uber.org/zap"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString(
				"DB_ADDR",
				"postgres://user:password@localhost:5433/gradeflow?sslmode=disable",
			),
		},
		auth: authConfig{
			basic: basicConfig{
				user: env.GetString("AUTH_BASIC_USER", "admin"),
				pass: env.GetString("AUTH_BASIC_PASS", "admin"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "example"),
				exp:    time.Hour * 24 * 3, // 3 days
				iss:    "gradeflow",
			},
		},
		env: env.GetString("ENVIROMENT", "development"),
		ocr: ocrConfig{
			addr: env.GetString(
				"OCR_ADDR",
				"http://localhost:8000/process-gabarito",
			),
		},
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	db, err := db.NewDB(cfg.db.addr)
	if err != nil {
		logger.Fatalw("error connecting to db", "error", err)
	}

	logger.Info("DB connection stablished")
	store := store.NewStorage(db)
	auth := auth.NewJWTAuthenticator(
		cfg.auth.token.secret,
		cfg.auth.token.iss,
		cfg.auth.token.iss,
	)
	app := &application{
		config:        cfg,
		logger:        logger,
		store:         store,
		authenticator: auth,
	}

	app.run()
}
