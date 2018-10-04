package main

import (
	"fmt"
	"time"
)

// o laco for é a unica estrutura de repeticao em go

func main() {
	i := 1

	// lembra o while
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// inicializacao da varivel
	// condicao
	// incremento
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("par")
		} else {
			fmt.Println("impar")
		}
	}

	// lanco infinito
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 2)
	}
}
