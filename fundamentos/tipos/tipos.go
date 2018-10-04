package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numericos inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (so positivos)... uint8 (1 byte) uint16 (2 bytes) uint32 (4 bytes) uint64 (8 bytes equivalente ao long)
	var b byte = 255
	fmt.Println("o byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1, ", o tipo de i1 é", reflect.TypeOf(i1))

	// rune ==> representa um mapeamento da tabela Unicode (int32)
	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2), i2)

	var x float32 = 49.99
	fmt.Println("o tipo de x é", reflect.TypeOf(x), ", o tipo do literal 49.99 é", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo), !bo)

	// string
	s1 := "ola meu nome é lucas"
	fmt.Println(s1+"!", ", o tamanho da string é", len(s1))

	// string com multiplas linhas
	s2 := `Olá
	meu 
	nome 
	é
	lucas
	`
	fmt.Println(s2+"!", ", o tamanho da string é", len(s2))

	// char?
	// não existe o tipo char
	// var x char = 'b'
	char := 'a'
	fmt.Println("o tipo do char é", reflect.TypeOf(char), ", valor unicode =>", char)

}
