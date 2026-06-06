package interfaces

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"math/rand"
)

type HabilidadeMovimentacao interface {
	GetX1() float64
	GetY1() float64
	SetPosicao(x, y float64)
}

type Movimentador interface {
	Mover(game IGame, mundo *geometria.Retangulo, bot HabilidadeMovimentacao, r *rand.Rand)
	GetTipo() string
}
