package repository

import (
	"github.com/paulolancao/go-contacts/models"
)

// LoadData into Contacts Storage
func LoadData() []models.Contact {

	memoryStorage := new(memorystorage)
	memoryStorage.contacts = make([]models.Contact, 0)

	contact1 := models.Contact{ID: 1, Age: "35", Email: "test1@test.com",
		Firstname: "firstname1", Fullname: "fullname1", Lastname: "lastname1",
		Middlename: "middlename1", Phone: "11111111", Addresses: []models.Address{}}

	address1 := models.Address{ID: 1, AddressLine1: "line1", AddressLine2: "line2",
		City: "Harrogate", Country: "UK", County: "York", Postcode: "HG1 1AA"}

	contact2 := models.Contact{ID: 2, Age: "25", Email: "test2@test.com",
		Firstname: "firstname2", Fullname: "fullname2", Lastname: "lastname2",
		Middlename: "middlename2", Phone: "22222222", Addresses: []models.Address{address1}}

	address2 := models.Address{ID: 2, AddressLine1: "line2", AddressLine2: "line3",
		City: "Harrogate2", Country: "UK2", County: "York2", Postcode: "HG12 2AA"}

	contact3 := models.Contact{ID: 3, Age: "45", Email: "test3@test.com",
		Firstname: "firstname3", Fullname: "fullname3", Lastname: "lastname3",
		Middlename: "middlename3", Phone: "33333333", Addresses: []models.Address{address1, address2}}

	memoryStorage.contacts = append(memoryStorage.contacts, contact1)
	memoryStorage.contacts = append(memoryStorage.contacts, contact2)
	memoryStorage.contacts = append(memoryStorage.contacts, contact3)

	return memoryStorage.contacts
}
