package model

type UpdateContactRequest struct {
	Id        string `validate:"required" json:"id"`
	FirstName string `validate:"required,max=100" json:"firstName"`
	LastName  string `validate:"max=100" json:"lastName"`
	Email     string `validate:"email,max=100" json:"email"`
	Phone     string `validate:"e164,max=100" json:"phone"`
}