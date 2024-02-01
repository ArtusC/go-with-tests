package assert

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"

	srv "github.com/ArtusC/go-with-tests/application/json-routing-embedding-version/v3/server"
)

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want this %+v, but this %+v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, but want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, but want false", got)
	}
}

func AssertLeague(t *testing.T, got, want []srv.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func NewGetScoreRequest(name string) *http.Request {
	player := fmt.Sprintf("/players/%s", name)
	request, _ := http.NewRequest("GET", player, nil)

	return request
}

func NewPostPlayerRequest(name string) *http.Request {
	req, err := http.NewRequest("POST", fmt.Sprintf("/players/%s", name), nil)
	if err != nil {
		panic("Failed to create the POST request to path /players, error: " + err.Error())
	}
	return req
}

func NewGetLeagueRequest() *http.Request {
	req, _ := http.NewRequest("GET", "/league", nil)
	return req
}

func NewGetLeagueFromResponse(t testing.TB, body io.Reader) (league []srv.Player) {
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into Player slice, error: %v", body, err)
	}

	return
}
