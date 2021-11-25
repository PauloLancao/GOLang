package router

import (
	"logging"
	"net/http/httptest"
	"storage"
	"strings"
	"testing"
)

func BenchmarkAddRecord(b *testing.B) {
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

	// act
	for i := 0; i < b.N; i++ {
		Handler(rc, storage, StoreExecutor(rc, CmdCreate))
	}
}

func BenchmarkGet(b *testing.B) {
	// arrange
	httpwriter := httptest.NewRecorder()
	httprequest := httptest.NewRequest("GET", "http://localhost:8080/kvs/k1", nil)

	rc := NewContext(httpwriter, httprequest)
	rc.Builder = Builder{ExtractParams: true}

	storage := storage.Start()
	defer storage.Close()

	logging.CreateLogger([]string{})
	defer logging.Close()

	// act
	for i := 0; i < b.N; i++ {
		Handler(rc, storage, StoreExecutor(rc, CmdGet))
	}
}
