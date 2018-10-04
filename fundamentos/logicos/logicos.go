package main

import (
	"fmt"
)

// operadores lógicos
// operador unário !<boolean> (negacao logica)

func main() {
	tv50, tv32, sorvete := comprar(false, false)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}

func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2    // E
	comprarTv32 := trab1 != trab2    // use o diferente para o caso do ou exclusivo
	comprarSorvete := trab1 || trab2 // OU
	return comprarTv50, comprarTv32, comprarSorvete
}
