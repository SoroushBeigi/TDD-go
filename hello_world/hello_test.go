package hello_world

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Soroush")

	want := "Hello, Soroush"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
