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


func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
