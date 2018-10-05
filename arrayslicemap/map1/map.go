package main

import (
	"fmt"
)

func main() {
	var aprovados map[int]string
	aprovados = make(map[int]string)

	aprovados[1] = "Maria"
	aprovados[2] = "Pedro"
	aprovados[3] = "Ana"

	fmt.Println(aprovados)

	for id, name := range aprovados {
		fmt.Println(id, name)
	}

	fmt.Println(aprovados[2])

	delete(aprovados, 1)
	fmt.Println(aprovados)

}
