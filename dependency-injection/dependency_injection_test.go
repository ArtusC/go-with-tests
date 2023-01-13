package dependency_injection

import (
	"bytes"
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q | want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("test correct message", func(t *testing.T) {
		buffer := bytes.Buffer{}

		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})
}
