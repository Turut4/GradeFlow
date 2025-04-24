package main

import (
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {
	t.Run("should not allow unauthenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatalf("error creating the request: %v", err)
		}
	})
}
