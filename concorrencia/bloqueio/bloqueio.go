package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operacao bloqueante
}

func main() {
	ch := make(chan int) // canal sem buffer

	go rotina(ch) // async

	fmt.Println(<-ch) // sync
	fmt.Println("foi lido")
	fmt.Println(<-ch) // deadlock, pois nao tem mais possibilidade de chegar dados no canal

}
