package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10)

	s[9] = 12
	fmt.Println(s)

	// cria um slice com 10 elementos com um array interno com 20 elementos
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// quando chega o array interno esta cheio ele dobra a capacidade do array interno
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

}
