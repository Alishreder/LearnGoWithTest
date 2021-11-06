package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCounter := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCounter)
		for i := 0; i < wantedCounter; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCounter)
	})
}

func assertCounter(t testing.TB, counter *Counter, want int) {
	t.Helper()
	got := counter.Value()
	if got != want {
		t.Errorf("got %d, but want %d", got, want)
	}
}
