//go:build unit
// +build unit

package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	srv "github.com/ArtusC/go-with-tests/application/json-routing-embedding-version/v2/server"
	helper "github.com/ArtusC/go-with-tests/application/json-routing-embedding-version/v2/test_helper"
)

type StubPlayerStore struct {
	scores      map[string]int
	playerCalls []string
	league      []srv.Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) PostPlayerScore(name string) {
	s.playerCalls = append(s.playerCalls, name)
}

func (s *StubPlayerStore) GetLeague() []srv.Player {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Artus":  20,
			"Amanda": 10,
		},
		nil,
		nil,
	}

	server := srv.NewPlayerServer(&store)

	t.Run("returns score of Artus", func(t *testing.T) {
		request := helper.NewGetScoreRequest("Artus")
		response := httptest.NewRecorder()

		server.Handler.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		helper.AssertEqual(t, response.Code, http.StatusOK)
		helper.AssertEqual(t, got, want)
	})

	t.Run("returns score of Amanda", func(t *testing.T) {
		request := helper.NewGetScoreRequest("Amanda")
		response := httptest.NewRecorder()

		server.Handler.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		helper.AssertEqual(t, response.Code, http.StatusOK)
		helper.AssertEqual(t, got, want)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := helper.NewGetScoreRequest("UnregisteredName")
		response := httptest.NewRecorder()

		server.Handler.ServeHTTP(response, request)

		got := response.Result().StatusCode
		want := http.StatusNotFound

		helper.AssertEqual(t, got, want)
	})
}

func TestStatusOnPOSTPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}

	server := srv.NewPlayerServer(&store)

	t.Run("test status code on POST", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/players/Artur", nil)
		response := httptest.NewRecorder()

		server.Handler.ServeHTTP(response, request)

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
		nil,
	}

	server := srv.NewPlayerServer(&store)

	t.Run("test POST", func(t *testing.T) {
		player := "Tiao"

		request := helper.NewPostPlayerRequest(player)
		response := httptest.NewRecorder()

		server.Handler.ServeHTTP(response, request)

		helper.AssertEqual(t, response.Code, http.StatusAccepted)

		if len(store.playerCalls) != 1 {
			t.Errorf("got %d calls to PostPlayerScore want %d", len(store.playerCalls), 1)
		}

		if store.playerCalls[0] != player {
			t.Errorf("didn't store correct player, got %q, want %q", store.playerCalls[0], player)
		}

	})
}

func TestLeague(t *testing.T) {
	t.Run("returns leagueTAble as JSON", func(t *testing.T) {
		wantedLeague := []srv.Player{
			{"Player_1", 1},
			{"Player_2", 2},
			{"Player_3", 4},
		}

		store := StubPlayerStore{
			map[string]int{},
			nil,
			wantedLeague,
		}

		server := srv.NewPlayerServer(&store)

		request, _ := http.NewRequest("GET", "/league", nil)
		respnse := httptest.NewRecorder()

		server.Handler.ServeHTTP(respnse, request)

		var got []srv.Player

		err := json.NewDecoder(respnse.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response %q into Player slice: %v", respnse.Body, err)
		}

		helper.AssertEqual(t, respnse.Code, http.StatusOK)

		if !reflect.DeepEqual(got, wantedLeague) {
			t.Errorf("got %v, wanted %v", got, wantedLeague)
		}
	})
}