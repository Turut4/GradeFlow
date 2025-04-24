package main

import (
	"net/http"
	"net/http/httptest"
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

func executeRequest(
	req *http.Request,
	mux http.Handler,
) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}
