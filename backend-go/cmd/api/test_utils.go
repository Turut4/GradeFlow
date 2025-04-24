package main

import (
	"testing"

	"github.com/Turut4/GradeFlow/internal/auth"
	"github.com/Turut4/GradeFlow/internal/store"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T, cfg config) *application {
	t.Helper()

	store := store.NewMockStore()
	logger := zap.NewNop().Sugar()
	// Uncomment to enable logs
	// logger := zap.Must(zap.NewProduction()).Sugar()
	testAuth := &auth.TestAuthenticator{}
	return &application{
		config:        cfg,
		store:         store,
		logger:        logger,
		authenticator: testAuth,
	}
}
