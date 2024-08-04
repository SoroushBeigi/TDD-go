package hello_world

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello + names",func (t *testing.T)  {
	const name = "Soroush"
	want := fmt.Sprintf("Hello, %s",name)

	got := Hello(name)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	})

	t.Run("saying Hello, world when empty name is provided",func (t *testing.T)  {
		const name = ""
		want := "Hello, World"
	
		got := Hello(name)
	
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		})
	
}
