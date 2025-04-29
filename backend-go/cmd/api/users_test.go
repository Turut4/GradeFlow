package main

import (
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {
	app := newTestApplication(t, config{})
	mux := app.mount()

	testToken, err := app.authenticator.GenerateToken(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("should not allow unauthenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatalf("error creating the request: %v", err)
		}

		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should not allow invalid token", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatalf("error creating the request: %v", err)
		}

		req.AddCookie(&http.Cookie{
			Name:  "jwt",
			Value: "invalid-test-token",
		})

		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusForbidden, rr.Code)
	})

	t.Run("should return OK for authenticated users", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatalf("error creating the request: %v", err)
		}

		req.AddCookie(&http.Cookie{
			Name:  "jwt",
			Value: testToken,
		})

		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusOK, rr.Code)
	})

}
