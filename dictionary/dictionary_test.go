package dictionary

import "testing"

func TestSearch(t *testing.T){
	dictionary := map[string]string{"test":"test value"}
	got := Search(dictionary, "test")
	want := "test value"

	if got!=want{
		t.Errorf("got %q want %q given %q",got,want,"test")
	}
}