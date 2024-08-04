package sum_list

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersList ...[]int) []int {
	var sums []int

	for _, numbers := range numbersList {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers))
		}

	}
	return sums
}
