package main

import (
	"fmt"
)

func media(numeros ...float64) (res float64) {
	for _, num := range numeros {
		res += num
	}
	res = res / float64(len(numeros))
	return
}

func main() {
	fmt.Printf("Media: %.2f\n", media(2.3, 2, 4))
}
