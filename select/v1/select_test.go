//go:build unit
// +build unit

package racer_test

import (
	"testing"

	ra "github.com/ArtusC/go-with-tests/select/v1"
)

func TestRacer(t *testing.T) {
	fasturl := "http://facebook.com"
	slowUrl := "http://quii.dev"

	want := fasturl
	got := ra.Racer(slowUrl, fasturl)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
