package slices_tut




func Sum(numbers []int) int {
	var ans int
	for _, val := range numbers {
		ans += val
	}
	return ans
}

func SumAll(numbersToSum ...[]int) []int{
	length := len(numbersToSum)

	sums:= make([]int, length)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums;
}