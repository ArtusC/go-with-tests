package mocking

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &DefaultSleeper{}

	Countdown(buffer, sleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q | want %q", got, want)
	}

}
