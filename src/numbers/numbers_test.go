package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveIndexAlt(t *testing.T) {
	var totalSum = sumAll()

	assert.Equal(t, 1629, totalSum)
}
