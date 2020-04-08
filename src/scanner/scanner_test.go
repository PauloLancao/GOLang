package scanner

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveIndexAlt(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("paulolancao\n"))

	result, err := stdin.ReadString('\n')
	assertResult := strings.Replace(result, "\n", "", -1)

	assert.NoError(t, err)
	assert.Equal(t, "paulolancao", assertResult)
}
