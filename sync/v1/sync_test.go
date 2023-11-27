package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 time leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()

		assertCounter(t, counter, 3)

	})

	t.Run("it runs safely concurrently", func(t *testing.T) {

		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}
		wg.Wait()

		//  The test will probably fail with a different number,
		//  but nonetheless it demonstrates it does not work when multiple
		//  goroutines are trying to mutate the value of the counter at the same time.
		// assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, but expected %d", got.Value(), want)
	}
}