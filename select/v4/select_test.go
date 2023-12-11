//go:build unit
// +build unit

package racer_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	ra "github.com/ArtusC/go-with-tests/select/v4"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(9 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fasturl := fastServer.URL

		want := fasturl
		got, err := ra.Racer(slowUrl, fasturl)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		timeOutServer := 20 * time.Millisecond
		duratioTimeServer := 25 * time.Millisecond
		server := makeDelayedServer(duratioTimeServer)

		defer server.Close()

		_, err := ra.ConfigurableRacer(server.URL, server.URL, timeOutServer)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
