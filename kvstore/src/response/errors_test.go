package response

import (
	"net/http"
	"testing"
)

func TestNewErrorFunc(t *testing.T) {
	errc := New(ErrNotFound, "TestNewErrorFunc")

	apiErrc, ok := errc.(APIError)
	if !ok {
		t.Error("TestNewErrorFunc error type conversion")
	}

	if apiErrc.statusCode != http.StatusBadRequest {
		t.Errorf("TestNewErrorFunc error apierror expected %d got %d", apiErrc.statusCode, http.StatusBadRequest)
	}

	if apiErrc.ErrorCode != ErrNotFound {
		t.Errorf("TestNewErrorFunc error apierror expected %d got %d", apiErrc.ErrorCode, ErrNotFound)
	}

	if apiErrc.ErrorMessage != ErrorCodeText(ErrNotFound) {
		t.Errorf("TestNewErrorFunc error apierror expected %s got %s", apiErrc.ErrorMessage, ErrorCodeText(ErrNotFound))
	}
}
