package utils

import (
	"fmt"
	"regexp"
)

func ValidaTelefone(phone string) bool {
	regex := `^(?:(?:\+|00)?(55)\s?)?(?:\(?([1-9][0-9])\)?\s?)?(?:((?:9\d|[2-9])\d{3})\-?(\d{4}))$`

	matched, err := regexp.MatchString(regex, phone)
	if err != nil {
		// Tratar erro se a expressão regular for inválida
		return false
	}

	return matched
}

func SeparaDDD(telefone string) (ddd, numero string, err error) {
	padrao := `^(?:\+?55\s?)?(?:(\d{2})\s*)?(9?\d{8})$`
	regex := regexp.MustCompile(padrao)
	matches := regex.FindStringSubmatch(telefone)
	if len(matches) >= 3 {
		ddd = matches[1]
		numero = matches[2]
		return ddd, numero, nil
	}
	return "", "", fmt.Errorf("telefone inválido")
}
