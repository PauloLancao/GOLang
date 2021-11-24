package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRollDices(t *testing.T) {
	diceValue, err := rollDices()

	assert.NoError(t, err)
	assert.True(t, diceValue >= 1 && diceValue <= 6)
}

func TestRollDicesNTimes(t *testing.T) {
	dicesResult := rollDicesNTimes()

	assert.True(t, len(dicesResult) == 50)
}
