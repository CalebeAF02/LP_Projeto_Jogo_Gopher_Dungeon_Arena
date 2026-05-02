package interfaces

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"math/rand"
)

type HabilidadeMovimentacao interface {
	GetX() float64
	GetY() float64
	SetPosicao(x, y float64)
}

type Movimentador interface {
	Mover(mundo geometria.Retangulo, bot HabilidadeMovimentacao, r *rand.Rand)
	GetTipo() string
}
