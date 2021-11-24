package repository

import "github.com/paulolancao/go-contacts/models"

// Repository provides access to the contacts storage.
type Repository interface {
	GetContacts() []models.Contact
	GetContact(uint) (models.Contact, error)
	CreateContact(models.Contact) error
	UpdateContact(models.Contact) error
	DeleteContact(uint) (models.Contact, error)
}

// MemoryStorage struct
type memorystorage struct {
	contacts []models.Contact
}

// Init MemoryStorage and array contacts
func Init() Repository {
	memoryContacts := LoadData()
	return &memorystorage{memoryContacts}
}
