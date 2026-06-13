package interfaces

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"image/color"
	"math/rand"
)

type HabilidadeMovimentacao interface {
	GetX1() float64
	GetY1() float64
	SetPosicao(x, y float64)
}

type Movimentador interface {
	Mover(entidade ecs.Entidade, sistemaColisao ISistemaColisao, mundo *geometria.Retangulo, bot HabilidadeMovimentacao, r *rand.Rand)
	GetTipo() string
	GetCor() color.Color
}
