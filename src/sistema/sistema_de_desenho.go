package sistema

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type SistemaDesenhar struct{}

func (s *SistemaDesenhar) Desenhar(g *Game, tela *ebiten.Image) {

	tela.Fill(color.RGBA{20, 20, 20, 255})

	margemMundo := geometria.NovoRetangulo(
		g.GetCamera().GetX()+g.mundo.GetX(),
		g.GetCamera().GetY()+g.mundo.GetY(),
		g.mundo.GetLargura(),
		g.mundo.GetAltura(),
	)

	utils.MargemInterna(
		tela,
		margemMundo,
		utils.JOGADOR_TAMANHO_MUNDO,
		cores.BRANCO,
	)

	for _, entidade := range g.entidades {
		entidade.Desenhar(tela)
	}

	if config.PROPORCAO_MUNDO > 1 {

		g.miniMapa.Desenhar(tela)

		for _, entidade := range g.entidades {
			entidade.DesenharMapa(
				tela,
				g.miniMapa.GetX(),
				g.miniMapa.GetY(),
			)
		}
	}
}
