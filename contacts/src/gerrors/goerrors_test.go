package gerrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	err := New(500, ErrArgumentException)

	assert.Equal(t, "Error argument exception", err.Error())

	tc, ok := err.(Channelerror)
	if ok {
		assert.Equal(t, "Error argument exception", tc.Error())
		assert.Equal(t, 500, tc.StatusCode())
	} else {
		assert.Fail(t, "type cast failed on gerrors_test")
	}
}

func TestStructError(t *testing.T) {
	ce := channelerror{errorMessage: "struct test", statusCode: 222}

	assert.Equal(t, "struct test", ce.Error())
	assert.Equal(t, 222, ce.StatusCode())
}
