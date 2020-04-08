package numberbetween

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdIoInput(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("222\n"))

	result, err := stdin.ReadString('\n')
	assertResult, err1 := strconv.Atoi(strings.Replace(result, "\n", "", -1))

	assert.NoError(t, err)
	assert.NoError(t, err1)
	assert.Equal(t, 222, assertResult)
}
