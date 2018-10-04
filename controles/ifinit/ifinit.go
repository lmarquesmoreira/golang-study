package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// inicia uma variavel dentro do contexto do if; condicao do if
	// é possivel fazer a mesma coisa com switch
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("ganhou")
	} else {
		fmt.Println("perdeu")
	}
}
