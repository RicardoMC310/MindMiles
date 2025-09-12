package dtos

type ResponseLoginDto struct {
	Token string `json:"token"`
	ExpiresIn string `json:"expiresin"`
}

type RequestLoginDto struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}