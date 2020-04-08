package arraysiteration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestArraySumOdd(t *testing.T) {
	var sum = naturalOrderSumOdd(arr)

	assert.Equal(t, 25, sum)
}

func TestArraySumEven(t *testing.T) {
	var sum = naturalOrderSumEven(arr)

	assert.Equal(t, 30, sum)
}
