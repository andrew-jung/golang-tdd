package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// `make` creates a slice (dynamic array?) with starting capacity of 2nd param, can index and re-assign values.
	// sums := make([]int, lengthOfNumbers)

	var sums []int
	for _, numbers := range numbersToSum {
		// `append` takes a slice and a new value, and returns the appended slice (will reallocate more memory if needed)
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		tail := numbers[1:] // Like in python [low:high], index 1 and beyond.
		sums = append(sums, Sum(tail))
	}
	return sums
}
