package main

import (
	"fmt"
)

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("numero invalido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func fatorialSimples(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorialSimples(n-1)
	}
}

func main() {
	result, _ := fatorial(5)
	fmt.Println(result)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	result1 := fatorialSimples(5)
	fmt.Println(result1)

}
