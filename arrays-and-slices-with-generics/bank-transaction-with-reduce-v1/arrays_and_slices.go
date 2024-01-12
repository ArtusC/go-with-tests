package arrays_and_slices_with_generics

/*
Now we're also reducing a different type than the type of the collection.
This sounds scary, but actually just requires us to adjust the type signature of Reduce to make it work.
We won't have to change the function body, and we won't have to change any of our existing callers.
*/
func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}

	return result
}

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums

}

func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbers, sumTail, []int{})
}

type BankTransactions struct {
	From, To string
	Sum      float64
}

func BalanceFor(transactions []BankTransactions, name string) float64 {
	adjustBalance := func(currentBalance float64, t BankTransactions) float64 {
		if t.From == name {
			return currentBalance - t.Sum
		}
		if t.To == name {
			return currentBalance + t.Sum
		}
		return currentBalance
	}
	return Reduce(transactions, adjustBalance, 0.0)
}
