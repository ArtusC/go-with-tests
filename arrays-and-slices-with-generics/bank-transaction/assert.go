package arrays_and_slices_with_generics

import "testing"

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want this %+v, but this %+v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, but want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, but want false", got)
	}
}

// func AssertEqual(t *testing.T, got, want interface{}) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("got %+v, want %+v", got, want)
// 	}
// }

// func AssertNotEqual(t *testing.T, got, want interface{}) {
// 	t.Helper()
// 	if got == want {
// 		t.Errorf("didn't want this %d, but this %d", got, want)
// 	}
// }
