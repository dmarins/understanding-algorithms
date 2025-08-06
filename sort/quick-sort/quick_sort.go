package quicksort

func QuickSort(numbers []int) []int {
	n := len(numbers)
	if n < 2 {
		return numbers
	}

	pivot := numbers[0]
	var smaller []int
	var larger []int

	for _, item := range numbers[1:] {
		if item <= pivot {
			smaller = append(smaller, item)
		} else {
			larger = append(larger, item)
		}
	}

	orderedLarger := QuickSort(smaller)
	result := append(orderedLarger, pivot)
	orderedSmaller := QuickSort(larger)
	result = append(result, orderedSmaller...)

	return result
}
