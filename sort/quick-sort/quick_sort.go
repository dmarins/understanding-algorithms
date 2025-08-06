package quicksort

func QuickSort(numbers []int) []int {
	n := len(numbers)
	if n < 2 {
		return numbers
	}

	pivo := numbers[0]
	var minor []int
	var major []int

	for _, item := range numbers[1:] {
		if item <= pivo {
			minor = append(minor, item)
		} else {
			major = append(major, item)
		}
	}

	sortedMenores := QuickSort(minor)
	result := append(sortedMenores, pivo)
	sortedMaiores := QuickSort(major)
	result = append(result, sortedMaiores...)

	return result
}
