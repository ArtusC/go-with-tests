//go:build unit
// +build unit

package sync_test

import (
	"sync"
	"testing"

	sy "github.com/ArtusC/go-with-tests/sync/v2"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 time leaves it at 3", func(t *testing.T) {
		counter := sy.NewCounter()
		counter.Increment()
		counter.Increment()
		counter.Increment()

		assertCounter(t, counter, 3)

	})

	t.Run("it runs safely concurrently", func(t *testing.T) {

		wantedCount := 1000
		counter := sy.NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}
		wg.Wait()

		// Use channels when passing ownership of data
		// Use mutexes for managing state
		// go vet
		// 	Remember to use go vet in your build scripts as it can alert you to some
		// 	subtle bugs in your code before they hit your poor users.

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *sy.Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, but expected %d", got.Value(), want)
	}
}
