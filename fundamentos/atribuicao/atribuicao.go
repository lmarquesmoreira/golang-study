package main

import (
	"fmt"
)

// operadores de operacao

func main() {
	var b byte = 255
	fmt.Println(b)

	i := 255 // criacao da variavel + atribucao do valor e é feita inferencia de tipo
	i += 3   // i = i + 3
	i -= 3   // i = i - 3
	i /= 2   // i = i / 2
	i *= 2   // i = i * 2
	i %= 2   // i = i % 2
	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}
