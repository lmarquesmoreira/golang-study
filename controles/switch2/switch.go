package main

import (
	"fmt"
	"time"
)

// use para compilar
// go run *go

func main() {
	t := time.Now()

	// switch true
	// busca o primeiro case que seja verdadeiro
	switch {
	case t.Hour() < 12:
		fmt.Println("bom dia")
	case t.Hour() < 18:
		fmt.Println("boa tarde")
	default:
		fmt.Println("boa noite")
	}

	fmt.Println("Desafio...")
	fmt.Println("Nota", notaParaConceito(9.8))
}
