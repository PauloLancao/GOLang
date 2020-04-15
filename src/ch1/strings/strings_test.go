package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatString(t *testing.T) {
	var formattedStr = formatString()
	assert.Equal(t, "Hi, iam paulo", formattedStr)
}
