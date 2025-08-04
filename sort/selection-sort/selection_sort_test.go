package selectionsort_test

import (
	"testing"

	selectionsort "github.com/dmarins/understanding-algorithms/sort/selection-sort"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSort_ShouldReturnASortedSlice(t *testing.T) {
	var numbers []int = []int{64, 25, 12, 22, 11}

	selectionsort.SelectionSort(numbers)

	assert.EqualValues(t, numbers, []int{11, 12, 22, 25, 64})
}

func TestSelectionSort_WhenSliceAlreadySorted(t *testing.T) {
	var numbers []int = []int{11, 12, 22, 25, 64}

	selectionsort.SelectionSort(numbers)

	assert.EqualValues(t, numbers, []int{11, 12, 22, 25, 64})
}
