package main

import (
	"fmt"
	"time"
)

func talk(people string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Millisecond * 100)
			c <- fmt.Sprintf("%s talking: %d", people, i)
		}
	}()
	return c
}

func join(entry1, entry2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-entry1:
				c <- s
			case s := <-entry2:
				c <- s
			}
		}
	}()

	return c
}

func main() {

	c := join(talk("Joao"), talk("Maria"))

	for i := 0; i < 6; i++ {
		fmt.Println(<-c)
	}
}
