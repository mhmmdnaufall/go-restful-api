package model

type CreateContactRequest struct {
	FirstName string `validate:"required,max=100" json:"firstName"`
	LastName  string `validate:"max=100" json:"lastName"`
	Email     string `validate:"omitempty,email,max=100" json:"email"`
	Phone     string `validate:"omitempty,e164,max=100" json:"phone"`
}
