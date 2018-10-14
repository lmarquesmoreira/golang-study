package main

import (
	"fmt"
)

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim")
	if numero >= 5000 {
		fmt.Println("valor alto")
		return 5000
	}
	fmt.Println("valor baixo")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(6000))
}
