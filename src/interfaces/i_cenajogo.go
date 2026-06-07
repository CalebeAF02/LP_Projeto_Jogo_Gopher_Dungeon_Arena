package interfaces

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"math/rand"
)

type ICenaJogo interface {
	CriarEntidade() ecs.EntidadeID
	GetEntidades() map[ecs.EntidadeID]ecs.Entidade
	SetEntidade(nEntidade ecs.EntidadeID, e ecs.Entidade)
	GetAleatorio() *rand.Rand
	GetMundo() *geometria.Retangulo
	GetCamera() *ecs.Camera
	GetMiniMapa() *ecs.MiniMapa
	VaiColidir(meuCorpoAtual *geometria.Retangulo, proximoCorpo *geometria.Retangulo) *ecs.RespostaColisao
	OrganizaPosicaoAleatoriaBot() *geometria.Ponto
}
