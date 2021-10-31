package main

// Adds all numbers inside an array and returns the sum
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}