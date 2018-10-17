package main

import (
	"fmt"
	"time"
)

// Channel é a forma de comunicação entre goroutines
// chan

func mult(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	ch := make(chan int)
	go mult(2, ch) // async
	fmt.Println("executando...")

	a, b := <-ch, <-ch // (sync) recebendo dados do canal
	fmt.Println(a, b)
	fmt.Println(<-ch)
}
