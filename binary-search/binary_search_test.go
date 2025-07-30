package binarysearch_test

import (
	"testing"

	binarysearch "github.com/dmarins/understanding-algorithms/binary-search"
	"github.com/stretchr/testify/assert"
)

var numbers []int = []int{1, 3, 5, 7, 9}

func TestBinarySearch_ShouldReturnToPositionOne(t *testing.T) {
	expected := binarysearch.BinarySearch(numbers, 3)

	assert.Equal(t, *expected, 1)
}

func TestBinarySearch_ShouldReturnNil(t *testing.T) {
	expected := binarysearch.BinarySearch(numbers, -1)

	assert.Nil(t, expected)
}
