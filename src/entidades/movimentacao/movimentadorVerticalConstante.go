package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/interfaces"
	"math/rand"
)

type MovimentadorVerticalConstante struct {
	dY     int
	ciclos int
}

func (mvc *MovimentadorVerticalConstante) Mover(mundo geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {

	if mvc.dY == 0 {
		mvc.dY = r.Intn(personagens.BOT_VELOCIDADE_MAXIMA)
		mvc.ciclos = 0
	}

	posY := objeto.GetY() + float64(mvc.dY)

	if posY >= mundo.PosYmax(personagens.BOT_TAMANHO) {
		posY = mundo.PosYmax(personagens.BOT_TAMANHO)

		mvc.dY = r.Intn(personagens.BOT_VELOCIDADE_MAXIMA)
		mvc.dY = mvc.dY * (-1)

	} else if posY <= mundo.GetY() {
		posY = mundo.GetY()

		mvc.dY = r.Intn(personagens.BOT_VELOCIDADE_MAXIMA)
		mvc.dY = mvc.dY * (-1)
	}

	mvc.ciclos += 1

	if mvc.ciclos == personagens.BOT_CICLOS_REPETICAO {
		mvc.dY = r.Intn(personagens.BOT_VELOCIDADE_MAXIMA)
	}

	objeto.SetPosicao(objeto.GetX(), posY)
}

func (mvc *MovimentadorVerticalConstante) GetTipo() string {
	return "VERTICAL_CONSTANTE"
}
