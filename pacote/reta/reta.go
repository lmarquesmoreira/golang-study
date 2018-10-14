package main

import (
	"math"
)

// iniciando com letra maiúscula o modificador de acesso será PÚBLICO
// com visibilidade dentro e fora do pacote

// iniciando com letra minúscula ou com "_" o modificador de acesso será PRIVADO
// com visibilidade apenas dentro do pacote

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
