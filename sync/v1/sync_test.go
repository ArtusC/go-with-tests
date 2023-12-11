//go:build unit
// +build unit

package sync_test

import (
	"sync"
	"testing"

	sy "github.com/ArtusC/go-with-tests/sync/v1"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 time leaves it at 3", func(t *testing.T) {
		counter := sy.Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()

		assertCounter(t, counter, 3)

	})

	t.Run("it runs safely concurrently", func(t *testing.T) {

		wantedCount := 1000
		counter := sy.Counter{}

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

func assertCounter(t testing.TB, got sy.Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, but expected %d", got.Value(), want)
	}
}
