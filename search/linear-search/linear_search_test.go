package linearsearch_test

import (
	"testing"

	linearsearch "github.com/dmarins/understanding-algorithms/search/linear-search"
	"github.com/stretchr/testify/assert"
)

var numbers []int = []int{1, 3, 5, 7, 9}

func TestLinearSearch_ShouldReturnTheLastPosition(t *testing.T) {
	expected := linearsearch.LinearSearch(numbers, 9)

	assert.Equal(t, *expected, 4)
}

func TestLinearSearch_ShouldReturnNil(t *testing.T) {
	expected := linearsearch.LinearSearch(numbers, -1)

	assert.Nil(t, expected)
}
