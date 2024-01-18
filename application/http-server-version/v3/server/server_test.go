//go:build unit
// +build unit

package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	srv "github.com/ArtusC/go-with-tests/application/http-server-version/v3/server"
	assert "github.com/ArtusC/go-with-tests/application/http-server-version/v3/test_helper"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Artus":  20,
			"Amanda": 10,
		},
	}

	server := &srv.PlayerServer{&store}

	t.Run("returns score of Artus", func(t *testing.T) {
		request := newGetScoreRequest("Artus")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assert.AssertEqual(t, response.Code, http.StatusOK)
		assert.AssertEqual(t, got, want)
	})

	t.Run("returns score of Amanda", func(t *testing.T) {
		request := newGetScoreRequest("Amanda")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assert.AssertEqual(t, response.Code, http.StatusOK)
		assert.AssertEqual(t, got, want)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("UnregisteredName")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Result().StatusCode
		want := http.StatusNotFound

		assert.AssertEqual(t, got, want)
	})
}

func TestPOSTPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
	}

	server := &srv.PlayerServer{&store}

	t.Run("test status code on POST", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/players/Artur", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.AssertEqual(t, response.Code, http.StatusAccepted)
	})
}

func newGetScoreRequest(name string) *http.Request {
	player := fmt.Sprintf("/players/%s", name)
	request, _ := http.NewRequest("GET", player, nil)

	return request
}
