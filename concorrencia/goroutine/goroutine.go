package main

import (
	"fmt"
	"time"
)

func fale(pessoa, text string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao %d)\n", pessoa, text, i+1)
	}
}

func main() {
	// fale("Maria", "Pq vc nao fala", 3)
	// fale("Joao", "so posso falar depois de vc", 1)

	// go fale("Maria", "Ei", 5)
	// go fale("Joao", "Opa", 5)
	// time.Sleep(time.Second * 10)

	go fale("Maria", "entendi", 10)
	fale("Joao", "Parabens", 5)

}
