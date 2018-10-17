package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	ch <- 10 // enviando dados para o canal (escrita)
	<-ch     // recebendo dados do canal (leitura use <variavel> <-ch)
	ch <- 2
	fmt.Println(<-ch)
}
