package model

type AddressResponse struct {
	Id         string `json:"id"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
}
