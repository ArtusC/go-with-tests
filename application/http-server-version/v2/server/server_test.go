//go:build unit
// +build unit

package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	srv "github.com/ArtusC/go-with-tests/application/http-server-version/v2/server"
	assert "github.com/ArtusC/go-with-tests/application/http-server-version/v2/test_helper"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns score of Artus", func(t *testing.T) {
		request := newGetScoreRequest("Artus")
		response := httptest.NewRecorder()

		srv.PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		assert.AssertEqual(t, got, want)
	})

	t.Run("returns score of Amanda", func(t *testing.T) {
		request := newGetScoreRequest("Amanda")
		response := httptest.NewRecorder()

		srv.PlayerServer(response, request)

		got := response.Body.String()
		want := "10"

		assert.AssertEqual(t, got, want)
	})
}

func newGetScoreRequest(name string) *http.Request {
	player := fmt.Sprintf("/players/%s", name)
	request, _ := http.NewRequest("GET", player, nil)

	return request
}
