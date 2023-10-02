package model

type UpdateAddressRequest struct {
	ContactId  string `validate:"required" json:"contactId"`
	AddressId  string `validate:"required" json:"addressId"`
	Street     string `validate:"max=200" json:"street"`
	City       string `validate:"max=100" json:"city"`
	Province   string `validate:"max=100" json:"province"`
	Country    string `validate:"required,max=100" json:"country"`
	PostalCode string `validate:"max=10" json:"postalCode"`
}
