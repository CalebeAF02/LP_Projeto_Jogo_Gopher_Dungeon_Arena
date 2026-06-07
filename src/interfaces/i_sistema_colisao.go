package interfaces

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
)

type ISistemaColisao interface {
	SetCenaJogo(cj ICenaJogo)
	VaiColidir(origem string, origemEntidade ecs.Entidade, meuCorpoAtual *geometria.Retangulo, proximoCorpo *geometria.Retangulo) *ecs.RespostaColisao
	ColideComTipo(eu *geometria.Retangulo, tipoDesejado string) bool
}
