package services

import (
	"github.com/paulolancao/go-contacts/models"
)

// Service provides access to the contacts repository.
type Service interface {
	GetContacts() []models.Contact
	GetContact(uint) (models.Contact, error)
	CreateContact(models.Contact) error
	UpdateContact(models.Contact) error
	DeleteContact(uint) (models.Contact, error)
}
