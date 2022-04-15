package Sum_all

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers{
		sum += number 
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum) // len is length of numbertosum slice
	sums := make([]int, lengthOfNumbers) // Make will create a new slice starting with capacity of the len

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}