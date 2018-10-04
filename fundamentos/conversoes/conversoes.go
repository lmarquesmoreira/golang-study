package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // float64
	y := 2   // int

	// obrigatório fazer a conversao explicita
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado ao concatenar strings
	// string(<inteiro>) ele olha a tabela ASCII
	fmt.Println("Teste " + string(123))
	// use
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("falso")

	}

}
