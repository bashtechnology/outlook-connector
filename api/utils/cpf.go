package utils

import (
	"github.com/Nhanderu/brdoc"
)

func ValidaCPF(cpf string) bool {
	if !brdoc.IsCPF(cpf) {
		// fmt.Println("CPF inválido")
		return false
	}
	return true
}

// ValidaCNPJ verifica se o CNPJ é válido.
func ValidaCNPJ(cnpj string) bool {
	if !brdoc.IsCNPJ(cnpj) {
		// fmt.Println("CNPJ inválido")
		return false
	}
	return true
}
