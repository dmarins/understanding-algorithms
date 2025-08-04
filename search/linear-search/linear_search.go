package linearsearch

func LinearSearch(numbers []int, target int) *int {

	n := len(numbers)
	for i := 0; i < n; i++ {
		if numbers[i] == target {
			return &i
		}
	}

	return nil
}
