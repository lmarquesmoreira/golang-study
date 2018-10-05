package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func myWalk(t *tree.Tree, ch chan int) {
	if t != nil {
		myWalk(t.Left, ch)
		ch <- t.Value
		myWalk(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	myWalk(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {

	channel1 := make(chan int)
	channel2 := make(chan int)

	go Walk(t1, channel1)
	go Walk(t2, channel2)

	for i := range channel1 {
		if i != <-channel2 {
			return false
		}
	}

	return true
}

func main() {
	channel := make(chan int)
	go Walk(tree.New(1), channel)
	for i := range channel {
		fmt.Printf("%d   ", i)
	}
	fmt.Print("\n")

	fmt.Println(Same(tree.New(1), tree.New(1)), "", "")
	fmt.Println(Same(tree.New(1), tree.New(2)), "", "")
}
