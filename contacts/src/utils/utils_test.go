package utils

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRespondWithJSON(t *testing.T) {

	response := httptest.NewRecorder()

	RespondWithJSON(response, 200, map[string]interface{}{"msg": "test"})

	header := response.Header()
	bodyResponse := response.Body.String()
	ct := header.Get("Content-Type")

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "application/json", ct)
	assert.Equal(t, "{\"msg\":\"test\"}", bodyResponse)
}
