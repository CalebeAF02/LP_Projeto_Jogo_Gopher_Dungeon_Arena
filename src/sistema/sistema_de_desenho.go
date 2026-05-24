package sistema

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type SistemaDesenhar struct{}

func (s *SistemaDesenhar) Desenhar(g *Game, tela *ebiten.Image) {

	tela.Fill(cores.BRANCO)

	margemMundo := geometria.NovoRetangulo(
		g.GetCamera().GetX()+g.GetMundo().GetX(),
		g.GetCamera().GetY()+g.GetMundo().GetY(),
		g.GetMundo().GetLargura(),
		g.GetMundo().GetAltura(),
	)

	utils.MargemInterna(
		tela,
		margemMundo,
		utils.JOGADOR_TAMANHO_MUNDO,
		cores.PRETO,
	)

	for _, entidade := range g.GetEntidades() {
		entidade.Desenhar(tela)
	}

	if config.PROPORCAO_MUNDO > 1 {

		g.GetMiniMapa().Desenhar(tela)

		for _, entidade := range g.GetEntidades() {
			entidade.DesenharMapa(
				tela,
				g.GetMiniMapa().GetX(),
				g.GetMiniMapa().GetY(),
			)
		}
	}
}
