package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "test value"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "test value"

		assertStrings(t, got, want)

	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("sthelse")
		want := errNotFound

		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, want)

	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertError(t, err, errWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)

	})

	t.Run("unknown word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, errWordDoesNotExistUpdate)

	})

}

func TestDelete(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		delErr := dictionary.Delete(word)
		_, searchErr := dictionary.Search(word)

		assertError(t, searchErr, errNotFound)
		assertError(t, delErr, nil)

	})

	t.Run("unknown word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)

		assertError(t, err, errWordDoesNotExistDelete)
	})

}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
