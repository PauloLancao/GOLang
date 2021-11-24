package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDob(t *testing.T) {
	var dob = calcDob()
	assert.Equal(t, 44, dob)
}
