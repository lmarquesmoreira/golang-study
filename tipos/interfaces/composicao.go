package interfaces

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmw7 struct {}


func (b bmw7) ligarTurbo(){
	fmt.Println("turbo")
}

func (b bmw7) fazerBaliza(){
	fmt.Println("baliza")
}

func ExecuteDemoInterfaceComposicao()  {
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}


