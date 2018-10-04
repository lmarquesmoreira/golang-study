package main

import (
	"fmt"
)

// operadores unários

func main() {
	x := 1
	y := 2
	b := false
	fmt.Println(x, y, b)

	// apenas postfix
	x++     // x += 1 ou x = x + 1
	y--     // y -= 1 ou y = y - 1
	c := !b // c = negacao de b

	fmt.Println(x, y, c)
}
