package main

import (
	"fmt"
)

func imprimirResultado(nota float64) {
	// por convencao não se usa parenteses no if,
	// 		exceto para o uso de precendencias na expressao logica
	// nao e permitido o nao uso das chaves para identificar o bloco
	if nota >= 7 {
		fmt.Println("Aprovado", nota)
	} else {
		fmt.Println("reprovado", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
