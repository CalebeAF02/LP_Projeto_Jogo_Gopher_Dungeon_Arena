package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/interfaces"
	"math/rand"
)

type MovimentadorSimples struct {
}

func (ms *MovimentadorSimples) Mover(mundo geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {

	antesX := objeto.GetX()
	antesY := objeto.GetY()

	tomadaDeDecicao := r.Intn(100)

	if tomadaDeDecicao > 50 {
		dX := r.Intn(personagens.BOT_VELOCIDADE_MAXIMA)
		posX := 0.0
		if dX >= 5 {
			posX = objeto.GetX() + personagens.BOT_VELOCIDADE_NORMAL
		} else {
			posX = objeto.GetX() - personagens.BOT_VELOCIDADE_NORMAL
		}
		if posX >= mundo.PosXmax(personagens.BOT_TAMANHO) {
			posX = mundo.PosXmax(personagens.BOT_TAMANHO)
		} else if posX <= mundo.GetX() {
			posX = mundo.GetX()
		}
		objeto.SetPosicao(posX, antesY)
	} else {
		dY := r.Intn(personagens.BOT_VELOCIDADE_MAXIMA)
		posY := 0.0
		if dY >= 5 {
			posY = objeto.GetY() + personagens.BOT_VELOCIDADE_NORMAL
		} else {
			posY = objeto.GetY() - personagens.BOT_VELOCIDADE_NORMAL
		}
		if posY <= mundo.GetY() {
			posY = mundo.GetY()
		} else if posY >= mundo.PosYmax(personagens.BOT_TAMANHO) {
			posY = mundo.PosYmax(personagens.BOT_TAMANHO)
		}
		objeto.SetPosicao(antesX, posY)
	}

}

func (ms *MovimentadorSimples) GetTipo() string {
	return "SIMPLES"
}
