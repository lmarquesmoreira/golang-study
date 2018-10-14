package main

import (
	"fmt"
)

func closure() func() {
	x := 10
	var function = func() {
		fmt.Println(x)
	}
	return function
}

func f1() func() {
	fmt.Println("f1")
	return func() {
		fmt.Println("fA")
	}
}

func main() {
	f1()
	f1()()

	x := 20
	fmt.Println(x)
	imprimeX := closure()
	imprimeX()
}
