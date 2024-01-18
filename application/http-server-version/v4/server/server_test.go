//go:build unit
// +build unit

package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	srv "github.com/ArtusC/go-with-tests/application/http-server-version/v4/server"
	helper "github.com/ArtusC/go-with-tests/application/http-server-version/v4/test_helper"
)

type StubPlayerStore struct {
	scores      map[string]int
	playerCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) PostPlayerScore(name string) {
	s.playerCalls = append(s.playerCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Artus":  20,
			"Amanda": 10,
		},
		nil,
	}

	server := &srv.PlayerServer{&store}

	t.Run("returns score of Artus", func(t *testing.T) {
		request := helper.NewGetScoreRequest("Artus")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		helper.AssertEqual(t, response.Code, http.StatusOK)
		helper.AssertEqual(t, got, want)
	})

	t.Run("returns score of Amanda", func(t *testing.T) {
		request := helper.NewGetScoreRequest("Amanda")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		helper.AssertEqual(t, response.Code, http.StatusOK)
		helper.AssertEqual(t, got, want)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := helper.NewGetScoreRequest("UnregisteredName")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Result().StatusCode
		want := http.StatusNotFound

		helper.AssertEqual(t, got, want)
	})
}

func TestStatusOnPOSTPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}

	server := &srv.PlayerServer{&store}

	t.Run("test status code on POST", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/players/Artur", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		helper.AssertEqual(t, response.Code, http.StatusAccepted)

		if len(store.playerCalls) != 1 {
			t.Errorf("got %d calls to PostPlayerScore want %d", len(store.playerCalls), 1)
		}
	})
}

func TestPOSTPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}

	server := &srv.PlayerServer{&store}

	t.Run("test POST", func(t *testing.T) {
		player := "Tiao"

		request := helper.NewPostPlayerRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		helper.AssertEqual(t, response.Code, http.StatusAccepted)

		if len(store.playerCalls) != 1 {
			t.Errorf("got %d calls to PostPlayerScore want %d", len(store.playerCalls), 1)
		}

		if store.playerCalls[0] != player {
			t.Errorf("didn't store correct player, got %q, want %q", store.playerCalls[0], player)
		}

	})
}
