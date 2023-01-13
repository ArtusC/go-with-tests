package hello

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q | want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people in English", func(t *testing.T) {
		got := Hello("Artus", "English")
		want := "Hello, super world and Artus!"
		assertCorrectMessage(t, got, want)

	})
	t.Run("saying hello to people in Portugues", func(t *testing.T) {
		got := Hello("Artus", "Portugues")
		want := "Olá, super mundo e Artus!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world and anything!' when an empty string is supplied in English", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, super world and anything!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world and anything!' when an empty string is supplied in Portugues", func(t *testing.T) {
		got := Hello("", "Portugues")
		want := "Hello, super world and anything!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world and anything!' when name and language are supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, super world and anything!"
		assertCorrectMessage(t, got, want)
	})

}
