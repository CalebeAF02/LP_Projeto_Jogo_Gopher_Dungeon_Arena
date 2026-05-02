package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/interfaces"
	"math/rand"
)

type MovimentadorHorizontal struct {
}

func (mh *MovimentadorHorizontal) Mover(mundo geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	posX := 0.0

	tomadaDeDecicao := r.Intn(100)

	if tomadaDeDecicao >= 50 {
		posX = objeto.GetX() + personagens.BOT_VELOCIDADE_NORMAL
	} else {
		posX = objeto.GetX() - personagens.BOT_VELOCIDADE_NORMAL
	}

	if !mundo.EstaDentro(posX, objeto.GetY(), personagens.BOT_TAMANHO, personagens.BOT_TAMANHO) {
		if posX >= mundo.PosXmax(personagens.BOT_TAMANHO) {
			posX = mundo.PosXmax(personagens.BOT_TAMANHO)
		} else if posX <= mundo.GetX() {
			posX = mundo.GetX()
		}
	}
	objeto.SetPosicao(posX, objeto.GetY())
}
func (mh *MovimentadorHorizontal) GetTipo() string {
	return "HORIZONTAL"
}
