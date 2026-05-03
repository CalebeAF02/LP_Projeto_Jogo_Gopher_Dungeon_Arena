package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"

	"math/rand"
)

type MovimentadorDiagonal struct {
}

func (md *MovimentadorDiagonal) Mover(game interfaces.IGame, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	posX := 0.0
	posY := 0.0

	tomadaDeDecicao := r.Intn(100)

	if tomadaDeDecicao >= 0 && tomadaDeDecicao < 25 {
		posX = objeto.GetX() - utils.BOT_VELOCIDADE_NORMAL
		posY = objeto.GetY() - utils.BOT_VELOCIDADE_NORMAL
	} else if tomadaDeDecicao >= 25 && tomadaDeDecicao < 50 {
		posX = objeto.GetX() - utils.BOT_VELOCIDADE_NORMAL
		posY = objeto.GetY() + utils.BOT_VELOCIDADE_NORMAL
	} else if tomadaDeDecicao >= 50 && tomadaDeDecicao < 75 {
		posX = objeto.GetX() + utils.BOT_VELOCIDADE_NORMAL
		posY = objeto.GetY() + utils.BOT_VELOCIDADE_NORMAL
	} else if tomadaDeDecicao >= 75 {
		posX = objeto.GetX() + utils.BOT_VELOCIDADE_NORMAL
		posY = objeto.GetY() - utils.BOT_VELOCIDADE_NORMAL
	}

	corpo := geometria.NovoRetangulo(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	if mundo.EstaDentroDireto(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) && !game.ColideComBarreiras(corpo) {
		objeto.SetPosicao(posX, posY)
	}

}

func (md *MovimentadorDiagonal) GetTipo() string {
	return "DIAGONAL"
}
