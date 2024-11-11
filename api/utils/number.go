package utils

func ResetFromMinimum(resultado float64, minimo float64) float64 {
	if resultado < minimo {
		return minimo
	}
	return resultado
}
