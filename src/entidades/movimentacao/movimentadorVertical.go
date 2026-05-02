package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/interfaces"
	"math/rand"
)

type MovimentadorVertical struct {
}

func (mb *MovimentadorVertical) Mover(mundo geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	posY := 0.0

	tomadaDeDecicao := r.Intn(100)

	if tomadaDeDecicao >= 50 {
		posY = objeto.GetY() + personagens.BOT_VELOCIDADE_NORMAL
	} else {
		posY = objeto.GetY() - personagens.BOT_VELOCIDADE_NORMAL
	}

	if !mundo.EstaDentro(objeto.GetX(), posY, personagens.BOT_TAMANHO, personagens.BOT_TAMANHO) {
		if posY >= mundo.PosYmax(personagens.BOT_TAMANHO) {
			posY = mundo.PosYmax(personagens.BOT_TAMANHO)
		} else if posY <= mundo.GetY() {
			posY = mundo.GetY()
		}
	}
	objeto.SetPosicao(objeto.GetX(), posY)

}

func (mb *MovimentadorVertical) GetTipo() string {
	return "VERTICAL"
}
