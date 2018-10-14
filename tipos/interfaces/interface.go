package interfaces

import "fmt"

type imprimivel interface {
	toString() string
}

type Pessoa struct {
	nome string
	sobrenome string
}

type Product struct {
	nome string
	preco float64
}


func (p Pessoa) toString() string{
	return p.nome + " " + p.sobrenome
}

func (p Product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(i imprimivel) {
	fmt.Println(i.toString())
}

func ExecuteDemoInterface (){
	var coisa imprimivel = Pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = Product{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := Product{"Calça Jeans", 179.90}
	imprimir(p2)
}