package main

import (
	"fmt"
	"reflect"
)

func main() {
	//array
	a1 := [3]int{1, 2, 3}

	// slice
	s1 := []int{1, 2, 3}

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{0, 1, 2, 3, 4}

	// slice nao é um array. somente define um pedaço de um array
	// internamente
	// podemos imaginar um slice como: tamanho e um ponteiro para um elemento de um array
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	// novo slice, mas internamente ele aponta para o mesmo array (a2)
	s3 := a2[:2]
	fmt.Println(a2, s3)

	// é possivel fazer slices de um slice
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
