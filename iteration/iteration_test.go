package iteration

import "testing"

func TestRepeat(t *testing.T){
	repeated := Repeat("a")

	wanted:= "aaaaa"

	if repeated!=wanted{
		t.Errorf("expected %q but got %q", wanted, repeated)
	}
}

func BenchmarkRepeat(b *testing.B){
	for i:=0; i<b.N; i++{
		Repeat("a")
	}
}