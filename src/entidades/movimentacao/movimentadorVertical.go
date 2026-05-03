package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"math/rand"
)

type MovimentadorVertical struct {
}

func (mb *MovimentadorVertical) Mover(game interfaces.IGame, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	posY := 0.0

	tomadaDeDecicao := r.Intn(100)

	if tomadaDeDecicao >= 50 {
		posY = objeto.GetY() + utils.BOT_VELOCIDADE_NORMAL
	} else {
		posY = objeto.GetY() - utils.BOT_VELOCIDADE_NORMAL
	}

	corpo := geometria.NovoRetangulo(objeto.GetX(), posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	if !mundo.EstaDentroDireto(objeto.GetX(), posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) && !game.ColideComBarreiras(corpo) {
		if posY >= mundo.PosYmax(utils.BOT_TAMANHO_MUNDO) {
			posY = mundo.PosYmax(utils.BOT_TAMANHO_MUNDO)
		} else if posY <= mundo.GetY() {
			posY = mundo.GetY()
		}
	}

	if mundo.EstaDentroDireto(objeto.GetX(), posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) && !game.ColideComBarreiras(corpo) {
		objeto.SetPosicao(objeto.GetX(), posY)
	}
}

func (mb *MovimentadorVertical) GetTipo() string {
	return "VERTICAL"
}
