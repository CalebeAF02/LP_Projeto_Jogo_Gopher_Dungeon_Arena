package interfaces

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"math/rand"
)

type IGame interface {
	CriarEntidade() ecs.EntidadeID
	GetEntidades() map[ecs.EntidadeID]ecs.Entidade
	SetEntidade(nEntidade ecs.EntidadeID, bot ecs.Entidade)
	GetAleatorio() *rand.Rand
	GetMundo() geometria.Retangulo
	GetCameraX() float64
	GetCameraY() float64
}
