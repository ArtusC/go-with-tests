//go:build unit
// +build unit

package integers_test

import (
	"fmt"
	"testing"

	in "github.com/ArtusC/go-with-tests/integers"
)

func TestAdder(t *testing.T) {
	sum := in.Add(2, 2)

	expected := 4
	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := in.Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
