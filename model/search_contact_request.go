package model

type SearchContactRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Page  int    `validate:"required" json:"page"`
	Size  int    `validate:"required" json:"size"`
}
