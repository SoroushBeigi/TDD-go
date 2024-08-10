package concurrentcounter

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing 3 times", func(t *testing.T) {
		counter := Counter{}
		
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("got %d, want %d", counter.Value(), 3)
		}

	})
}
