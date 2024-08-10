package concurrentcounter

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing 3 times", func(t *testing.T) {
		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)

	})

}

func assertCounter(t *testing.T, counter Counter, want int) {
	t.Helper()
	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}
}
