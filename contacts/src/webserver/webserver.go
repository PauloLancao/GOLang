package webserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/paulolancao/go-contacts/controllers"
	"github.com/paulolancao/go-contacts/di"
)

const defaultPort = "8000"
const defaultDomain = "localhost"

// Init webserver
func Init(di *di.DI) {

	router := mux.NewRouter()

	// Get port from .env file, we did not specify any port so this should
	// return an empty string when tested locally
	port := os.Getenv("http_port")
	domain := os.Getenv("http_domain")

	if strings.TrimSpace(port) == "" {
		port = defaultPort
	}
	if strings.TrimSpace(domain) == "" {
		domain = defaultDomain
	}

	log.Printf("Webserver port: %s", port)
	log.Printf("Webserver domain: %s", domain)

	// Tradicional approach controller -> service -> repository
	router.HandleFunc("/contacts", controllers.GetContacts(*di.ContactService)).Methods("GET")
	router.HandleFunc("/contact", controllers.CreateContact(*di.ContactService)).Methods("POST")
	router.HandleFunc("/contact/{id:[0-9]+}", controllers.GetContact(*di.ContactService)).Methods("GET")
	router.HandleFunc("/contact/{id:[0-9]+}", controllers.UpdateContact(*di.ContactService)).Methods("PUT")
	router.HandleFunc("/contact/{id:[0-9]+}", controllers.DeleteContact(*di.ContactService)).Methods("DELETE")

	// Channels approach controller -> repository
	router.HandleFunc("/ch/contacts", controllers.GetChannelContacts(*di.ContactRepository)).Methods("GET")
	router.HandleFunc("/ch/contact", controllers.CreateChannelContact(*di.ContactRepository)).Methods("POST")
	router.HandleFunc("/ch/contact/{id:[0-9]+}", controllers.GetChannelContact(*di.ContactRepository)).Methods("GET")
	router.HandleFunc("/ch/contact/{id:[0-9]+}", controllers.UpdateChannelContact(*di.ContactRepository)).Methods("PUT")
	router.HandleFunc("/ch/contact/{id:[0-9]+}", controllers.DeleteChannelContact(*di.ContactRepository)).Methods("DELETE")

	log.Print("Webserver init started successfully")

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", domain, port), router)
	if err != nil {
		log.Print(err)
	}
}
