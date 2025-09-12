package dtos

type PartnerDto struct {
	Name string `json:"name" validate:"required,min=2,max=255"`
	Email string `json:"email" validate:"required,email"`
	CPF string `json:"cpf" validate:"required,cpf"`
	Tel string `json:"tel" validate:"required,min=10,max=20"`
	Age int `json:"age" validate:"required,number"`
	Password string `json:"password" validate:"required,min=8"`
}