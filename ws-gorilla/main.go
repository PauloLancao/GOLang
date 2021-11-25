package main

import (
	"log"
	"net/http"

	"github.com/PauloLancao/GOLang/ws-gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.HomeLink)
	router.HandleFunc("/event", handlers.CreateEvent).Methods("POST")
	router.HandleFunc("/events", handlers.GetAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", handlers.GetOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", handlers.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", handlers.DeleteEvent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
