package main

import (
	"fmt"
)

// a linguagem go não permite a aritmética de ponteiros
// ponteiro nada mais é do q uma referencia de memoria

func main() {
	i := 1
	var p *int
	p = &i // pega o endereco da variavel i e coloca em p

	// endereco de p => &p
	// valor em p => p
	// valor que p referencia => *p
	// valor de i => i
	// endereco de i => &i
	fmt.Println(&p, p, *p, i, &i)

	*p++
	fmt.Println(&p, p, *p, i, &i)
}
