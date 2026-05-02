package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/interfaces"
	"math/rand"
)

type MovimentadorHorizontalConstante struct {
	dX     int
	ciclos int
}

func (mhc *MovimentadorHorizontalConstante) Mover(mundo geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {

	if mhc.dX == 0 {
		mhc.dX = r.Intn(personagens.BOT_VELOCIDADE_MAXIMA)
		mhc.ciclos = 0
	}

	posX := objeto.GetX() + float64(mhc.dX)

	if posX >= mundo.PosXmax(personagens.BOT_TAMANHO) {
		posX = mundo.PosXmax(personagens.BOT_TAMANHO)

		mhc.dX = r.Intn(10)
		mhc.dX = mhc.dX * (-1)

	} else if posX <= mundo.GetX() {
		posX = mundo.GetX()

		mhc.dX = r.Intn(10)
		mhc.dX = mhc.dX * (-1)
	}

	mhc.ciclos += 1

	if mhc.ciclos == 10 {
		mhc.dX = r.Intn(10)
	}

	objeto.SetPosicao(posX, objeto.GetY())
}
func (mhc *MovimentadorHorizontalConstante) GetTipo() string {
	return "HORIZONTAL_CONSTANTE"
}
