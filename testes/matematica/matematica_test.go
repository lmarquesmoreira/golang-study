package matematica

import "testing"

/*
	use:
		go test ./... 	para executar todos os testes, independente de diretorio
		go test 		para executar os testes do diretorio corrente
*/

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v"

func TestMedia(t *testing.T) {
	t.Parallel()
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)
	if valor != valorEsperado {
		t.Errorf(erroPadrao, valor, valorEsperado)
	}
}
