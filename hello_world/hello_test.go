package hello_world

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Soroush"
	want := fmt.Sprintf("Hello, %s",name)

	got := Hello("Soroush")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
