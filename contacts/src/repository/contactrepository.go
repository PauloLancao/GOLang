package repository

import (
	"log"

	"github.com/paulolancao/go-contacts/models"
)

// GetContacts repo
func (m *memorystorage) GetContacts() []models.Contact {
	return m.contacts
}

// GetContact repo
func (m *memorystorage) GetContact(id uint) (models.Contact, error) {
	var contact = models.Contact{}

	for _, c := range m.contacts {
		if c.ID == id {
			return c, nil
		}
	}

	log.Printf("GetContact %d failed", id)

	return contact, contact.ErrUnknown()
}

// CreateContact repo
func (m *memorystorage) CreateContact(c models.Contact) error {
	// Check if email already exists
	err := m.isEmailDuplicate(c, true)

	if err != nil {
		return err
	}

	if c.Addresses == nil {
		c.Addresses = []models.Address{}
	}

	// Validate addresses passed without ID or 0
	if c.Addresses != nil && len(c.Addresses) > 0 {
		for i := range c.Addresses {
			c.Addresses[i].ID = uint(i + 1)
		}
	}

	c.ID = uint(len(m.contacts) + 1)
	m.contacts = append(m.contacts, c)

	return nil
}

// UpdateContact repo
func (m *memorystorage) UpdateContact(c models.Contact) error {
	// Check if email already exists
	err := m.isEmailDuplicate(c, false)

	if err != nil {
		return err
	}

	if c.Addresses == nil {
		c.Addresses = []models.Address{}
	}

	for i, e := range m.contacts {
		if e.ID == c.ID {
			m.contacts[i] = c
			return nil
		}
	}

	log.Printf("UpdateContact %d failed", c.ID)

	return models.Contact{}.ErrUnknown()
}

// DeleteContact repo
func (m *memorystorage) DeleteContact(id uint) (models.Contact, error) {
	var contact = models.Contact{}

	for i, c := range m.contacts {
		if c.ID == id {
			// Swap elem to delete with the one at the end, return n - 1 elems
			m.contacts[len(m.contacts)-1], m.contacts[i] = m.contacts[i], m.contacts[len(m.contacts)-1]
			m.contacts = m.contacts[:len(m.contacts)-1]

			return c, nil
		}
	}

	log.Printf("DeleteContact %d failed", id)

	return contact, contact.ErrUnknown()
}

func (m *memorystorage) isEmailDuplicate(c models.Contact, create bool) error {
	for _, e := range m.contacts {
		if (create && c.Email == e.Email) || (!create && c.ID != e.ID && c.Email == e.Email) {
			log.Printf("CreateContact %d duplicate with email %s", c.ID, c.Email)

			return c.ErrDuplicate()
		}
	}

	return nil
}
