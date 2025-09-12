package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
	cpfLib "github.com/mvrilo/go-cpf"
)

var Validate = validator.New()

func clearCpf(cpf string) string {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	return cpf
}

func Init() {
	Validate.RegisterValidation("cpf", func (fl validator.FieldLevel) bool {
		ret, _ := cpfLib.Valid(clearCpf(fl.Field().String())) 
		return ret
	})
}