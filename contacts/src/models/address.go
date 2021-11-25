package models

// Address struct
type Address struct {
	ID           uint   `json:"id"`
	AddressLine1 string `json:"addressline1"`
	AddressLine2 string `json:"addressline2"`
	Postcode     string `json:"postcode"`
	City         string `json:"city"`
	County       string `json:"county"`
	Country      string `json:"country"`
}
