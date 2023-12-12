//go:build unit
// +build unit

package reflection_test

import (
	"testing"

	re "github.com/ArtusC/go-with-tests/reflection/v0"
)

func TestWalk(t *testing.T) {
	expected := "Artus"

	var got []string

	x := struct {
		Name string
	}{expected}

	re.Walk(x, func(input string) {
		got = append(got, input)
	})

	if got[0] != expected {
		t.Errorf("got %q, want %q", got[0], expected)
	}
}
