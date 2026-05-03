package outros

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/componentes"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Parede struct {
	game        interfaces.IGame
	entidade    ecs.EntidadeID
	corpo       *geometria.Retangulo
	Componentes map[string]interface{}
}

func NovaParede(game interfaces.IGame, posicao *geometria.Ponto) *Parede {

	nEntidade := game.CriarEntidade()
	nParede := Parede{game: game, entidade: nEntidade, corpo: geometria.NovoRetangulo(posicao.GetX(), posicao.GetY(), 30, 30)}

	game.SetEntidade(nEntidade, &nParede)

	nParede.AdicionarComponente(componentes.CORPO.String(), nParede.corpo)

	return &nParede
}

func (b *Parede) GetX() float64 {
	return b.corpo.GetX()
}

func (b *Parede) GetY() float64 {
	return b.corpo.GetY()
}

func (b *Parede) GetTipo() string {
	return "PAREDE"
}

func (b *Parede) Atualizar() {

}

func (e *Parede) GetComponente(id string) interface{} {
	return e.Componentes[id]
}

func (e *Parede) AdicionarComponente(id string, comp interface{}) {
	if e.Componentes == nil {
		e.Componentes = make(map[string]interface{})
	}
	e.Componentes[id] = comp
}

func (b *Parede) Desenhar(tela *ebiten.Image) {
	ebitenutil.DrawRect(tela, b.game.GetCamera().GetX()+b.GetX(), b.game.GetCamera().GetY()+b.GetY(), 30, 30, cores.CINZA_ESCURO)
}

func (b *Parede) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawRect(tela, mapaX+(b.GetX()/config.PROPORCAO_MAPA), mapaY+(b.GetY()/config.PROPORCAO_MAPA), utils.BOT_TAMANHO_MAPA, utils.BOT_TAMANHO_MAPA, cores.VERMELHO)
}
