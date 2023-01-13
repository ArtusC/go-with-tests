package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknow word", func(t *testing.T) {
		_, got := dictionary.Search("unknow")

		assertError(t, got, ErrorNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "just a test"

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExists)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "just a test"

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExists)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestDelete(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrorNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "just a test"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrorNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatalf("should find added word: %s | err: %v", word, err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

/*
	* IMPORTANT NOTE TO REMEMBER:
		A  gotcha with maps is that they can be a nil value. A nil map behaves like an empty map
		 when reading, but attempts to write to a nil map will cause a runtime panic.
		Therefore, you should never initialize an empty map variable:
			var m map[string]string
		Instead, you can initialize an empty map like we were doing above, or use the make keyword to create
		 a map for you:
			var dictionary = map[string]string{}
			// OR
			var dictionary = make(map[string]string)
		Both approaches create an empty hash map and point dictionary at it. Which ensures that you
		 will never get a runtime panic.
*/
