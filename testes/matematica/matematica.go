package matematica

import (
	"fmt"
	"strconv"
)

// Somar é responsável por somar dois numeros
func Somar(a, b int) int {
	return a + b
}

// Media é responsável por calcular a média
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArrendonda, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArrendonda
}
