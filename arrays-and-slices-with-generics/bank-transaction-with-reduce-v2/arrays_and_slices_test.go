//go:build unit
// +build unit

package arrays_and_slices_with_generics_test

import (
	"reflect"
	"testing"

	arr "github.com/ArtusC/go-with-tests/arrays-and-slices-with-generics/bank-transaction-with-reduce-v2"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := arr.Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {

	got := arr.SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := arr.SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := arr.SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		checkSums(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		arr.AssertEqual(t, arr.Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}
		arr.AssertEqual(t, arr.Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestBankTransaction(t *testing.T) {
	var (
		artus  = arr.Account{Name: "Artus", Balance: 100}
		amanda = arr.Account{Name: "Amanda", Balance: 175}
		paulo  = arr.Account{Name: "Paulo", Balance: 200}

		transactions = []arr.BankTransactions{
			arr.NewBankTransaction(artus, amanda, 100),
			arr.NewBankTransaction(paulo, artus, 25),
		}
	)

	newBalanceFor := func(account arr.Account) float64 {
		return arr.NewBalanceFor(account, transactions).Balance
	}

	arr.AssertEqual(t, newBalanceFor(artus), 25)
	arr.AssertEqual(t, newBalanceFor(amanda), 275)
	arr.AssertEqual(t, newBalanceFor(paulo), 175)
}
