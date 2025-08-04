package selectionsort

func SelectionSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		indexOfTheSmallestNumber := i

		for j := i + 1; j < n; j++ {
			if numbers[j] < numbers[indexOfTheSmallestNumber] {
				indexOfTheSmallestNumber = j
			}

		}

		if indexOfTheSmallestNumber != i {
			numbers[i], numbers[indexOfTheSmallestNumber] = numbers[indexOfTheSmallestNumber], numbers[i]

			// Same as here
			// temp := numbers[i]
			// numbers[i] = numbers[indexOfTheSmallestNumber]
			// numbers[indexOfTheSmallestNumber] = temp
		}
	}
}
