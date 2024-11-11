package utils

import (
	"bytes"
	"fmt"
	"net/http"
)

func MakeRequestGet(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro ao fazer a solicitação GET: %s\n", err)
		return
	}
}
func MakeRequestPost(url string, requestBody []byte) {
	_, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("Erro ao fazer a solicitação POST: %s\n", err)
		return
	}
}
