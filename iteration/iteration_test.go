package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeating 'a' 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)

		wanted := "aaaaa"

		if repeated != wanted {
			t.Errorf("expected %q but got %q", wanted, repeated)
		}
	})
	t.Run("repeating 'a' 995 times", func(t *testing.T) {
		repeated := Repeat("a", 995)

		wanted := strings.Repeat("a", 995)

		if repeated != wanted {
			t.Errorf("expected %q but got %q", wanted, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
