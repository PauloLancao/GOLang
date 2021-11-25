package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondWithError func, message and code are used to track on client side
func RespondWithError(w http.ResponseWriter, statusCode int, payload interface{}) {
	RespondWithJSON(w, statusCode, payload)
}

// RespondWithJSON func
func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Printf("RespondWithJson err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
