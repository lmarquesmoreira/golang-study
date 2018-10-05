package main

import "fmt"

func main() {
	// usando [...], deixamos a cargo do compilar, definir o tamanho do array
	numeros := [...]int{1, 2, 3, 4, 5}

	// lembra o foreach
	for index, value := range numeros {
		fmt.Printf("%d) %d\n", index, value)
	}

	for _, value := range numeros {
		fmt.Printf("%d ", value)
	}
}
