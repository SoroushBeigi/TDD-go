package concurrentcounter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing 3 times", func(t *testing.T) {
		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)

	})

	t.Run("incrementing 1000 times concurrently", func(t *testing.T) {
		const incrementCount = 1000
		counter := Counter{}
		var wg sync.WaitGroup
	wg.Add(incrementCount)

		for i:=0 ; i<incrementCount;i++{
			go func ()  {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, incrementCount)

	})

}

func assertCounter(t *testing.T, counter Counter, want int) {
	t.Helper()
	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}
}
