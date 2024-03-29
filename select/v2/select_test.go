//go:build unit
// +build unit

package racer_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	ra "github.com/ArtusC/go-with-tests/select/v2"
)

func TestRacer(t *testing.T) {

	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fasturl := fastServer.URL

	want := fasturl
	got := ra.Racer(slowUrl, fasturl)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
