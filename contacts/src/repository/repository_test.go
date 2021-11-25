package repository

import (
	"testing"

	"github.com/paulolancao/go-contacts/models"
	"github.com/stretchr/testify/assert"
)

func TestGetContacts(t *testing.T) {
	repo := Init()

	contacts := repo.GetContacts()

	assert.NotNil(t, contacts)
	assert.NotEmpty(t, contacts)
	assert.Equal(t, 3, len(contacts))
}

func TestGetContact(t *testing.T) {
	repo := Init()

	contact, err := repo.GetContact(uint(1))

	assert.Nil(t, err)
	assert.NotNil(t, contact)
}

func TestCreateContact(t *testing.T) {
	repo := Init()

	contact := models.Contact{Age: "35", Email: "test99@test.com",
		Firstname: "firstname1", Fullname: "fullname1", Lastname: "lastname1",
		Middlename: "middlename1", Phone: "11111111"}

	err := repo.CreateContact(contact)

	assert.Nil(t, err)

	contacts := repo.GetContacts()

	assert.NotNil(t, contacts)
	assert.NotEmpty(t, contacts)

	insertedContact := contacts[len(contacts)-1]

	assert.NotNil(t, insertedContact)
	assert.Equal(t, insertedContact.Addresses, []models.Address{})
}

func TestCreateContactDuplicateEmail(t *testing.T) {
	repo := Init()

	contact := models.Contact{Age: "35", Email: "test1@test.com",
		Firstname: "firstname1", Fullname: "fullname1", Lastname: "lastname1",
		Middlename: "middlename1", Phone: "11111111"}

	err := repo.CreateContact(contact)

	assert.NotNil(t, err)
	assert.Equal(t, models.Contact{}.ErrDuplicate().Error(), err.Error())
}

func TestUpdateContactDuplicateEmail(t *testing.T) {
	repo := Init()

	contacts := repo.GetContacts()

	assert.NotNil(t, contacts)
	assert.NotEmpty(t, contacts)

	contact := contacts[0]
	contact.Email = "test2@test.com"

	err := repo.UpdateContact(contact)

	assert.NotNil(t, err)
	assert.Equal(t, models.Contact{}.ErrDuplicate().Error(), err.Error())
}

func TestDeleteContactDuplicateEmail(t *testing.T) {
	repo := Init()

	contacts := repo.GetContacts()

	assert.NotNil(t, contacts)
	assert.NotEmpty(t, contacts)

	contact := contacts[0]

	delContact, err := repo.DeleteContact(contact.ID)

	assert.Nil(t, err)
	assert.Equal(t, delContact.ID, contact.ID)
	assert.Equal(t, len(contacts)-1, len(repo.GetContacts()))
}
