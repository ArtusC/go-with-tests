//go:build integration
// +build integration

package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	srv "github.com/ArtusC/go-with-tests/application/http-server-version/v4/server"
	store "github.com/ArtusC/go-with-tests/application/http-server-version/v4/store"
	helper "github.com/ArtusC/go-with-tests/application/http-server-version/v4/test_helper"
)

func TestRecordPlayerAndRetrieveThem(t *testing.T) {
	store := store.NewInMemoryPlayerStore()
	server := srv.PlayerServer{store}
	player := "Artur"

	server.ServeHTTP(httptest.NewRecorder(), helper.NewPostPlayerRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), helper.NewPostPlayerRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), helper.NewPostPlayerRequest(player))

	response := httptest.NewRecorder()

	server.ServeHTTP(response, helper.NewGetScoreRequest(player))

	helper.AssertEqual(t, response.Code, http.StatusOK)
	helper.AssertEqual(t, response.Body.String(), "3")

}
