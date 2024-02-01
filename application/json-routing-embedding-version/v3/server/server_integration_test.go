//go:build integration
// +build integration

package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	srv "github.com/ArtusC/go-with-tests/application/json-routing-embedding-version/v3/server"
	store "github.com/ArtusC/go-with-tests/application/json-routing-embedding-version/v3/store"
	helper "github.com/ArtusC/go-with-tests/application/json-routing-embedding-version/v3/test_helper"
)

func TestRecordPlayerAndRetrieveThem(t *testing.T) {
	store := store.NewInMemoryPlayerStore()
	server := srv.NewPlayerServer(store)
	player := "Artur"

	server.Handler.ServeHTTP(httptest.NewRecorder(), helper.NewPostPlayerRequest(player))
	server.Handler.ServeHTTP(httptest.NewRecorder(), helper.NewPostPlayerRequest(player))
	server.Handler.ServeHTTP(httptest.NewRecorder(), helper.NewPostPlayerRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()

		server.Handler.ServeHTTP(response, helper.NewGetScoreRequest(player))

		helper.AssertEqual(t, response.Code, http.StatusOK)
		helper.AssertEqual(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		t.Helper()
		response := httptest.NewRecorder()

		server.Handler.ServeHTTP(response, helper.NewGetLeagueRequest())
		helper.AssertEqual(t, response.Code, http.StatusOK)

		got := helper.NewGetLeagueFromResponse(t, response.Body)
		want := []srv.Player{
			{"Artur", 3},
		}

		helper.AssertLeague(t, got, want)
	})

}
