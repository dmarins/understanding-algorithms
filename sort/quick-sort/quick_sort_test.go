package quicksort_test

import (
	"testing"

	quicksort "github.com/dmarins/understanding-algorithms/sort/quick-sort"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort_ShouldReturnASortedSlice(t *testing.T) {
	var numbers []int = []int{64, 11, 25}

	result := quicksort.QuickSort(numbers)

	assert.EqualValues(t, result, []int{11, 25, 64})
}

func TestQuickSort_WhenSliceAlreadySorted(t *testing.T) {
	var numbers []int = []int{11, 12, 22, 25, 64}

	result := quicksort.QuickSort(numbers)

	assert.EqualValues(t, result, []int{11, 12, 22, 25, 64})
}

func TestQuickSort_WithOneElement(t *testing.T) {
	var numbers []int = []int{11}

	result := quicksort.QuickSort(numbers)

	assert.EqualValues(t, result, []int{11})
}
