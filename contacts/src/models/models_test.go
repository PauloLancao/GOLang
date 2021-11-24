package models

import (
	"testing"

	"github.com/paulolancao/go-contacts/gerrors"
	"github.com/stretchr/testify/assert"
)

func TestContactErrors(t *testing.T) {
	c := Contact{}
	err := c.ErrUnknown()
	assert.Equal(t, "Error unknown", err.Error())

	tc, ok := err.(gerrors.Channelerror)
	if ok {
		assert.Equal(t, "Error unknown", tc.Error())
		assert.Equal(t, 500, tc.StatusCode())
	} else {
		assert.Fail(t, "type cast failed on gerrors_test")
	}

	err = c.ErrDuplicate()
	assert.Equal(t, "Error duplicate", err.Error())

	tc, ok = err.(gerrors.Channelerror)
	if ok {
		assert.Equal(t, "Error duplicate", tc.Error())
		assert.Equal(t, 400, tc.StatusCode())
	} else {
		assert.Fail(t, "type cast failed on gerrors_test")
	}
}
