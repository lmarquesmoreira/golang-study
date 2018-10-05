package main

import (
	"fmt"
)

func main() {
	funcESalarios := map[string]float64{
		"Jose":     1000.00,
		"Gabriela": 2400.00,
		"Pedro":    2000.00,
	}

	funcESalarios["Rafael"] = 1350.0

	delete(funcESalarios, "Inexistente")
	for name, sal := range funcESalarios {
		fmt.Println(name, sal)
	}
}
