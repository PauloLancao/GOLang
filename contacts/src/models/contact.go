package models

import (
	"net/http"
	"net/url"

	"github.com/paulolancao/go-contacts/gerrors"
)

// Contact struct
type Contact struct {
	ID         uint      `json:"id"`
	Fullname   string    `json:"fullname"`
	Firstname  string    `json:"firstname"`
	Middlename string    `json:"middlename"`
	Lastname   string    `json:"lastname"`
	Email      string    `json:"email"`
	Age        string    `json:"age"`
	Phone      string    `json:"phone"`
	Addresses  []Address `json:"addresses"`
}

// Validate model
func (c Contact) Validate() url.Values {
	errs := url.Values{}

	// Validator code here
	// check if the title empty
	if c.Firstname == "" {
		errs.Add("firstname", "The firstname field is required!")
	}

	if c.Lastname == "" {
		errs.Add("lastname", "The lastname field is required!")
	}

	if c.Age == "" {
		errs.Add("age", "The age field is required!")
	}

	return errs
}

// ErrUnknown is used when a contact isn't found.
func (c Contact) ErrUnknown() error {
	return gerrors.New(http.StatusInternalServerError, gerrors.ErrUnknown)
}

// ErrDuplicate is used when email contact already exists.
func (c Contact) ErrDuplicate() error {
	return gerrors.New(http.StatusBadRequest, gerrors.ErrDuplicate)
}
