package assert

import (
	"fmt"
	"net/http"
	"testing"
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

func NewGetScoreRequest(name string) *http.Request {
	player := fmt.Sprintf("/players/%s", name)
	request, _ := http.NewRequest("GET", player, nil)

	return request
}

func NewPostPlayerRequest(name string) *http.Request {
	req, _ := http.NewRequest("POST", fmt.Sprintf("/players/%s", name), nil)
	return req
}
