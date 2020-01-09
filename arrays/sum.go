package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}

func SumAllTails(numbersToTails ...[]int) (tails []int) {
	for _, numbers := range numbersToTails {
		if len(numbers) == 0 {
			tails = append(tails, 0)
		} else {
			tail := numbers[1:]
			tails = append(tails, Sum(tail))
		}
	}
	return
}
