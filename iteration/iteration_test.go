//go:build unit
// +build unit

package iteration_test

import (
	"testing"

	it "github.com/ArtusC/go-with-tests/iteration"
)

func TestRepeat(t *testing.T) {
	repeated := it.Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it.Repeat("a")
	}
}
