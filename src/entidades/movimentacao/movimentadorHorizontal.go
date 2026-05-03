package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"math/rand"
)

type MovimentadorHorizontal struct {
}

func (mh *MovimentadorHorizontal) Mover(game interfaces.IGame, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	posX := 0.0

	tomadaDeDecicao := r.Intn(100)

	if tomadaDeDecicao >= 50 {
		posX = objeto.GetX() + utils.BOT_VELOCIDADE_NORMAL
	} else {
		posX = objeto.GetX() - utils.BOT_VELOCIDADE_NORMAL
	}

	corpo := geometria.NovoRetangulo(posX, objeto.GetY(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	if !mundo.EstaDentroDireto(posX, objeto.GetY(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) && !game.ColideComBarreiras(corpo) {
		if posX >= mundo.PosXmax(utils.BOT_TAMANHO_MUNDO) {
			posX = mundo.PosXmax(utils.BOT_TAMANHO_MUNDO)
		} else if posX <= mundo.GetX() {
			posX = mundo.GetX()
		}
	}

	if mundo.EstaDentroDireto(posX, objeto.GetY(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) && !game.ColideComBarreiras(corpo) {
		objeto.SetPosicao(posX, objeto.GetY())
	}

}
func (mh *MovimentadorHorizontal) GetTipo() string {
	return "HORIZONTAL"
}
