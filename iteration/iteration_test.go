package iteration

import "testing"

func TestRepeat(t *testing.T){
	repeated := Repeat("a")

	wanted:= "aaaaa"

	if repeated!=wanted{
		t.Errorf("expected %q but got %q", wanted, repeated)
	}
}