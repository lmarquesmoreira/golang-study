package main

import (
	"golang-study/tipos/convertingJson"
	"golang-study/tipos/interfaces"
	"golang-study/tipos/metodos"
	"golang-study/tipos/meutipo"
	"golang-study/tipos/pseudoheranca"
	"golang-study/tipos/struct"
)

func main() {
	tipos.ExecuteDemo()
	tipos.ExecuteDemoAlinhada()
	metodos.ExecuteDemoMetodos()
	pseudoheranca.ExecuteDemoPseudoHeranca()
	meutipo.ExecuteMeuTipo()
	interfaces.ExecuteDemoInterface()
	interfaces.ExecuteDemoInterfaceComposicao()
	interfaces.ExecutaDemoInterfaceTipo()

	convertingJson.ExecuteDemoJsonConvert()
}
