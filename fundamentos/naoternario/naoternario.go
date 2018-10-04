package main

import (
	"fmt"
)

// não existe operador ternário em go

func obterResultado(nota float64) string {
	// nota >= 6 ? "aprovado" : "reprovado"
	if nota >= 6 {
		return "aprovado"
	}
	return "reprovado"
}

func main() {
	fmt.Println(obterResultado(6.3))
}
