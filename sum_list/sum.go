package sum_list

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersList ...[]int) []int {
	slicesCount := len(numbersList)
	sums := make([]int, slicesCount)

	for _, numbers := range numbersList {
		sums = append(sums,Sum(numbers))
	}
	return sums
}
