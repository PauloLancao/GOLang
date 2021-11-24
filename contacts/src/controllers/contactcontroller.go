package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/paulolancao/go-contacts/models"
	"github.com/paulolancao/go-contacts/services"
	"github.com/paulolancao/go-contacts/utils"

	"github.com/gorilla/mux"
)

// GetContacts restapi endpoint
func GetContacts(s services.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, http.StatusOK, s.GetContacts())
	}
}

// GetContact restapi endpoint
// /contact/{id:[0-9]+}
func GetContact(s services.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := getContactIDFromRequest(w, r)

		c, err := s.GetContact(uint(id))
		if err != nil {
			log.Printf("Error calling service layer GetContact ID:%d", id)

			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, c)
	}
}

// CreateContact restapi endpoint
func CreateContact(s services.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var c models.Contact
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&c); err != nil {
			log.Print("Error decoding payload")

			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		defer r.Body.Close()

		err := s.CreateContact(c)
		if err != nil {
			log.Print("Error calling service layer CreateContact")

			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, struct{}{})
	}
}

// UpdateContact restapi endpoint
// /contact/{id:[0-9]+}
func UpdateContact(s services.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := getContactIDFromRequest(w, r)

		var c models.Contact
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&c); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		defer r.Body.Close()

		// Update contact with ID
		c.ID = uint(id)
		err := s.UpdateContact(c)

		if err != nil {
			log.Printf("Error calling service layer UpdateContact ID:%d", id)

			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, c)
	}
}

// DeleteContact restapi endpoint
// /contact/{id:[0-9]+}
func DeleteContact(s services.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := getContactIDFromRequest(w, r)

		c, err := s.DeleteContact(uint(id))

		if err != nil {
			log.Printf("Error calling service layer DeleteContact ID:%d", id)

			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, c)
	}
}

func getContactIDFromRequest(w http.ResponseWriter, r *http.Request) (int, error) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Print("Error calling GetContact Atoi")

		utils.RespondWithError(w, http.StatusBadRequest, err)
		return id, err
	}

	return id, nil
}
