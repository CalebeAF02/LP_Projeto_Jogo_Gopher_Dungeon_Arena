package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/interfaces"

	"math/rand"
)

type MovimentadorDiagonal struct {
}

func (md *MovimentadorDiagonal) Mover(mundo geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	posX := 0.0
	posY := 0.0

	tomadaDeDecicao := r.Intn(100)

	if tomadaDeDecicao >= 0 && tomadaDeDecicao < 25 {
		posX = objeto.GetX() - personagens.BOT_VELOCIDADE_NORMAL
		posY = objeto.GetY() - personagens.BOT_VELOCIDADE_NORMAL
	} else if tomadaDeDecicao >= 25 && tomadaDeDecicao < 50 {
		posX = objeto.GetX() - personagens.BOT_VELOCIDADE_NORMAL
		posY = objeto.GetY() + personagens.BOT_VELOCIDADE_NORMAL
	} else if tomadaDeDecicao >= 50 && tomadaDeDecicao < 75 {
		posX = objeto.GetX() + personagens.BOT_VELOCIDADE_NORMAL
		posY = objeto.GetY() + personagens.BOT_VELOCIDADE_NORMAL
	} else if tomadaDeDecicao >= 75 {
		posX = objeto.GetX() + personagens.BOT_VELOCIDADE_NORMAL
		posY = objeto.GetY() - personagens.BOT_VELOCIDADE_NORMAL
	}

	if mundo.EstaDentro(posX, objeto.GetY(), personagens.BOT_TAMANHO, personagens.BOT_TAMANHO) {
		objeto.SetPosicao(posX, posY)
	}

}

func (md *MovimentadorDiagonal) GetTipo() string {
	return "DIAGONAL"
}
