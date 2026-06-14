package sistema

import (
	"Gopher_Dungeon_Arena/src/cenas/informativos"
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
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

	jogadoresVivos := 0

	for _, entidade := range cj.GetEntidades() {

		entidade.Desenhar(tela)

		if jogador, ok := entidade.(*personagens.Jogador); ok {
			if jogador.GetEntidade() != nil {
				jogadoresVivos++
			}
		}
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

	if jogadoresVivos == 0 {

		informativos.InformativoPerdeu(cj, tela)

	}
}
