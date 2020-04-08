package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestRemoveIndexAlt(t *testing.T) {
	var retArr = RemoveIndexAlt(arr, 4)

	assert.Equal(t, []int{0, 1, 2, 3, 5, 6, 7, 8, 9}, retArr)
}

func TestRemoveIndex(t *testing.T) {
	var retArr = RemoveIndex(arr, 7)

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 8, 9}, retArr)
}
