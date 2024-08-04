package hello_world

import (
	"fmt"
	"testing"
)

const name = "Soroush"

func TestHello(t *testing.T) {
	t.Run("saying hello + names", func(t *testing.T) {

		want := fmt.Sprintf("Hello, %s", name)

		got := Hello(name, "English")

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello, world when empty name is provided", func(t *testing.T) {
		const emptyName = ""
		want := "Hello, World"

		got := Hello(emptyName, "English")

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello in Persian", func(t *testing.T) {
		want := fmt.Sprintf("Salam, %s", name)

		got := Hello(name, "Persian")

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello in Spanish", func(t *testing.T) {
		want := fmt.Sprintf("Hola, %s", name)

		got := Hello(name, "Spanish")

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello in Persian, with empty name", func(t *testing.T) {
		want := "Hola, World"

		got := Hello("", "Spanish")

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
