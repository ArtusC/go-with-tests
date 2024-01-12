package arrays_and_slices_with_generics

func Reduce[A any](collection []A, accumulator func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}

	return result
}

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
	// sum := 0
	// for _, n := range numbers {
	// 	sum += n
	// }
	// return sum
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
	// var sums []int
	// for _, numbers := range numbersToSum {
	// 	if len(numbers) == 0 {
	// 		sums = append(sums, 0)
	// 	} else {
	// 		tail := numbers[1:]
	// 		sums = append(sums, Sum(tail))
	// 	}
	// }
	// return sums

}
