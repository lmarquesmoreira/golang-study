package main

import (
	"fmt"
)

// defer adia a execução de uma função até o final do retorno da função

func main() {
	defer fmt.Println("World")
	defer fmt.Println("Lucas")
	fmt.Println("hello")
}
