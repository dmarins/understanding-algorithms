package binarysearch

func BinarySearch(numbers []int, target int) *int {
	low := 0
	high := len(numbers) - 1

	for low <= high {
		middle := (low + high) / 2
		kick := numbers[middle]

		if kick == target {
			return &middle
		}

		if kick > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return nil
}
