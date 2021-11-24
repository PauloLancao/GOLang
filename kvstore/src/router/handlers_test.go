package router

import (
	"logging"
	"net/http"
	"net/http/httptest"
	"storage"
	"strings"
	"testing"
)

func TestErrorResponse(t *testing.T) {

	// arrange
	logging.CreateLogger([]string{})
	defer logging.Close()

	httpwriter := httptest.NewRecorder()
	rc := NewContext(httpwriter, nil)

	// act
	ErrorResponse(rc, http.StatusOK, map[string]interface{}{"msg": "test"})

	header := httpwriter.Header()
	bodyResponse := httpwriter.Body.String()
	ct := header.Get("Content-Type")

	// assert
	if httpwriter.Code != http.StatusOK {
		t.Errorf("ResponseWriter expected status code 200 but got %d", httpwriter.Code)
	}

	if ct != "application/json" {
		t.Errorf("ResponseWriter header expected content type 'application/json' but got %s", ct)
	}

	if bodyResponse == "" {
		t.Errorf("ResponseWriter body expected non nil but got %s", "{\"msg\":\"test\"}")
	}
}

func TestAddRecord(t *testing.T) {

	// arrange
	httpwriter := httptest.NewRecorder()
	bodyReader := strings.NewReader(`{"test":"BodyTest"}`)
	httprequest := httptest.NewRequest("POST", "/kvs/k1", bodyReader)

	rc := NewContext(httpwriter, httprequest)
	rc.Builder = Builder{ExtractParams: true, ExtractBody: true}

	storage := storage.Start()
	defer storage.Close()

	logging.CreateLogger([]string{})
	defer logging.Close()

	Handler(rc, storage, StoreExecutor(rc, CmdCreate))
}
