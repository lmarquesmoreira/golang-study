package main

import (
	"fmt"
	"time"
)

func isOdd(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func odds(n int, c chan int) {
	start := 2

	for i := 0; i < n; i++ {
		for possibleOdd := start; ; possibleOdd++ {
			if isOdd(possibleOdd) {
				c <- possibleOdd
				start = possibleOdd + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)

	go odds(60, c)

	for odd := range c {
		fmt.Printf("%d ", odd)
	}

	fmt.Printf("\n")
}
