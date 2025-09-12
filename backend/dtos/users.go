package dtos

import "github.com/go-playground/validator/v10"

var Validate = validator.New()

type UsersDto struct {
	Name 	 string  `json:"name" validate:"required,min=2,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Rules    string `json:"rules" validate:"required,oneof=student teacher admin"`
}