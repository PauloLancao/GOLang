package services

import (
	"github.com/paulolancao/go-contacts/models"
	"github.com/paulolancao/go-contacts/repository"
)

type service struct {
	contactRepository repository.Repository
}

// Init Service like DI via constructor
func Init(contactRepository *repository.Repository) Service {
	return &service{*contactRepository}
}

// Service layer were some business logic can be added
func (s *service) GetContacts() []models.Contact {
	return s.contactRepository.GetContacts()
}

func (s *service) GetContact(id uint) (models.Contact, error) {
	return s.contactRepository.GetContact(id)
}

func (s *service) CreateContact(c models.Contact) error {
	return s.contactRepository.CreateContact(c)
}

func (s *service) UpdateContact(c models.Contact) error {
	return s.contactRepository.UpdateContact(c)
}

func (s *service) DeleteContact(id uint) (models.Contact, error) {
	return s.contactRepository.DeleteContact(id)
}
