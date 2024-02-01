//go:build unit
// +build unit

package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	srv "github.com/ArtusC/go-with-tests/application/http-server-version/v2/server"
	assert "github.com/ArtusC/go-with-tests/application/http-server-version/v2/test_helper"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns score of a player", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/players/Artus", nil)
		response := httptest.NewRecorder()

		srv.PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		assert.AssertEqual(t, got, want)
	})
}
