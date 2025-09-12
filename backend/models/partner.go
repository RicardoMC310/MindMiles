package models

type PartnerModel struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Tel string `json:"tel"`
	CPF string `json:"cpf"`
	Age int `json:"age"`
	Password string `json:"password"`
}