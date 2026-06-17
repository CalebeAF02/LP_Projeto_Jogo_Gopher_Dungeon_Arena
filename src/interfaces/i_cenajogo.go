package interfaces

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type ICenaJogo interface {
	GetNome() string
	CriarEntidade() ecs.EntidadeID
	GetSistemaColisao() ISistemaColisao
	GetContadorMortos() string
	GetGame() IGame
	SetGame(game IGame)
	GetEntidades() map[ecs.EntidadeID]ecs.Entidade
	SetEntidade(nEntidade ecs.EntidadeID, e ecs.Entidade)
	GetAleatorio() *rand.Rand
	GetMundo() *geometria.Retangulo
	GetCamera() *ecs.Camera
	GetMiniMapa() *ecs.MiniMapa
	OrganizaPosicaoAleatoriaBot() *geometria.Ponto
	Update() error
	Draw(tela *ebiten.Image)
	SetFonteCache(cache assets.FonteCache)
	CapturouTudo() bool
	ColetadoTudo(status bool)
	MiniMapaEstaVisivel() bool
	ObterFonteCache() *assets.FonteCache
	ObterPontuacaoFaltante() int
	SetFaltaPontuacao(pontos int)
	Concluiu() bool
	EntreiNaSaida()
	EntrouNaSaida() bool
}
