package racer

import "testing"

func TestRacer(t *testing.T) {
	slowUrl := "http://facebook.com"
	fasturl := "http://quii.dev"

	want := fasturl
	got := Racer(slowUrl, fasturl)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
