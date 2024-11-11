package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func PriMaiuscula(input string) string {
	lowercase := strings.ToLower(input)
	return strings.Title(lowercase)
}
func ExtrairNumeroDaString(s string) (int64, error) {
	re := regexp.MustCompile(`(\d+)`)
	match := re.FindStringSubmatch(s)
	if len(match) >= 2 {
		return strconv.ParseInt(match[1], 10, 32)
	}
	return 0, fmt.Errorf("número não encontrado na string")
}

func TrocaObjetoDaString(substituicoes map[string]string, texto string) string {
	for chave, valor := range substituicoes {
		texto = strings.ReplaceAll(texto, chave, valor)
	}
	return texto
}

func IgnoraUltimoCaracterePonteiro(str *string) string {
	if str != nil && len(*str) > 0 {
		return (*str)[:len(*str)-1]
	}
	return ""
}
func RemoverUltimoCaracterePonteiro(str *string) {
	if str != nil && len(*str) > 0 {
		*str = (*str)[:len(*str)-1]
	}
}
func UltimoCaracterePonteiro(str *string) string {
	if str != nil && len(*str) > 0 {
		return (*str)[len(*str)-1:]
	}
	return ""
}

func IgnoraUltimoCaractere(str string) string {
	if len(str) > 0 {
		return str[:len(str)-1]
	}
	return ""
}
func UltimoCaractere(str string) string {
	if len(str) > 0 {
		return str[len(str)-1:]
	}
	return ""
}

func LimitCaracteres(name string, maxLen int) string {
	// Verificar se o nome do arquivo já é curto o suficiente
	if len(name) <= maxLen {
		return name
	}

	// Limitar o nome do arquivo para o comprimento máximo
	return name[:maxLen] + "..."
}
