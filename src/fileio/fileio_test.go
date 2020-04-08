package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringArrToByteArrAndBack(t *testing.T) {
	content, err := stringArrToByteArr()

	assert.NoError(t, err)
	assert.NotNil(t, content)

	contentStr := byteArrToStringArr(content)

	assert.NotNil(t, contentStr)

	assert.Equal(t, cities, contentStr)
}
