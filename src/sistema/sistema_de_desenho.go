package sistema

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type SistemaDesenhar struct{}

func (s *SistemaDesenhar) Desenhar(cj interfaces.ICenaJogo, tela *ebiten.Image) {

	tela.Fill(cores.BRANCO)

	margemMundo := geometria.NovoRetangulo(
		cj.GetCamera().GetX()+cj.GetMundo().GetX(),
		cj.GetCamera().GetY()+cj.GetMundo().GetY(),
		cj.GetMundo().GetLargura(),
		cj.GetMundo().GetAltura(),
	)

	utils.MargemInterna(
		tela,
		margemMundo,
		utils.JOGADOR_TAMANHO_MUNDO,
		cores.PRETO,
	)

	for _, entidade := range cj.GetEntidades() {
		entidade.Desenhar(tela)
	}

	if config.PROPORCAO_MUNDO > 1 {

		cj.GetMiniMapa().Desenhar(tela)

		for _, entidade := range cj.GetEntidades() {
			entidade.DesenharMapa(
				tela,
				cj.GetMiniMapa().GetX(),
				cj.GetMiniMapa().GetY(),
			)
		}
	}
}
