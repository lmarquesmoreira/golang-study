package main

import (
	"fmt"
	"time"
)

// operadores relacionais

func main() {
	fmt.Println("Strings:", "banana" == "banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	// esta nao se preocupa com comparar a referencia de memoria e sim o valor
	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas com .equals():", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	fmt.Println("Pessoas:", p1 == p2)
}
