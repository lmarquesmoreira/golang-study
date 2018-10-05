package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 10, 20)

	// o s2 aponta para o s1
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, cap(s1), s2, cap(s2))

	s1[0] = 7
	fmt.Println(s1, cap(s1), s2, cap(s2))
}
