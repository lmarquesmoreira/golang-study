package main

import (
	"fmt"
)

func main() {

	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel": 100.00,
			"Guga":    200.00,
		},
		"J": {
			"Jose": 102.00,
		},
		"P": {
			"Pedro": 120.00,
		},
	}

	delete(funcsPorLetra, "P")
	for letra, funcs := range funcsPorLetra {
		fmt.Println("Group:", letra)
		for name, sal := range funcs {
			fmt.Println("User:", name, "Sal:", sal)
		}
	}

}
