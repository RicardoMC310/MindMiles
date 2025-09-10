package models

import "time"


type CompanyModel struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	RazaoSocial      string    `json:"razao_social"`
	CNPJ             string    `json:"cnpj"`
	NaturezaJuridica string    `json:"natureza_juridica"`
	IE               string    `json:"ie"`
	IM               string    `json:"im"`
	Email            string    `json:"email"`
	Tel              string    `json:"tel"`
	Site             string    `json:"site"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}