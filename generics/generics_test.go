//go:build unit
// +build unit

package generics_test

import (
	"testing"

	gn "github.com/ArtusC/go-with-tests/generics"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("assert on integers", func(t *testing.T) {
		gn.AssertEqual(t, 1, 1)
		gn.AssertNotEqual(t, 1, 2)
	})

	t.Run("assert on strings", func(t *testing.T) {
		gn.AssertEqual(t, "hello", "hello")
		gn.AssertNotEqual(t, "hello", "bye")
	})

	// gn.AssertEqual(t, 1, "1") // uncomment to see the compilation error
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInt := new(gn.Stack[int])

		// check stack is empty
		gn.AssertTrue(t, myStackOfInt.IsEmpty())

		// add a value, and chek is not empty anymore
		myStackOfInt.Push(123)
		gn.AssertFalse(t, myStackOfInt.IsEmpty())

		// add another value, and pop it back again
		myStackOfInt.Push(456)
		value, _ := myStackOfInt.Pop()
		gn.AssertEqual(t, value, 456)

		value, _ = myStackOfInt.Pop()
		gn.AssertEqual(t, value, 123)

		gn.AssertTrue(t, myStackOfInt.IsEmpty())

		// can get the numbers we put as a numbers, not untyped inteface{}
		myStackOfInt.Push(1)
		myStackOfInt.Push(2)

		firstNum, _ := myStackOfInt.Pop()
		secondNum, _ := myStackOfInt.Pop()

		gn.AssertEqual(t, firstNum+secondNum, 3)
	})
}
