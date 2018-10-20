package main

import (
	"fmt"

	"github.com/lmarquesmoreira/html"
)

func forward(from <-chan string, to chan string) {
	for {
		to <- <-from
	}
}

//multiplexar - mistura as mensagens num canal

func join(entry1, entry2 <-chan string) <-chan string {
	c := make(chan string)
	go forward(entry1, c)
	go forward(entry2, c)
	return c
}

func main() {
	c := join(
		html.GetTitle("https://google.com", "https://bing.com"),
		html.GetTitle("https://amazon.com", "https://youtube.com"),
	)

	fmt.Println("1)", <-c, "|", <-c)
	fmt.Println("2):", <-c, "|", <-c)
}
