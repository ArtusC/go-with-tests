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

type Account struct {
	Name    string
	Balance float64
}

func NewBankTransaction(from, to Account, sum float64) BankTransactions {
	return BankTransactions{From: from.Name, To: to.Name, Sum: sum}
}

func NewBalanceFor(account Account, transaction []BankTransactions) Account {
	return Reduce(
		transaction,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, t BankTransactions) Account {
	if t.From == a.Name {
		a.Balance -= t.Sum
	}
	if t.To == a.Name {
		a.Balance += t.Sum
	}
	return a
}
