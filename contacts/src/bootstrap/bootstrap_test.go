package bootstrap

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDotEnv(t *testing.T) {
	loadDotEnv()

	assert.NotNil(t, os.Getenv("http_port"))
	assert.NotNil(t, os.Getenv("http_domain"))
	assert.NotNil(t, os.Getenv("log_filename"))
}
