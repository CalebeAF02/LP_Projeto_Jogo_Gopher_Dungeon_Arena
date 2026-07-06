package movimentacao

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"image/color"

	"math/rand"
)

type MovimentadorDiagonal struct {
}

func (self *MovimentadorDiagonal) Mover(entidade ecs.Entidade, sistemaColisao interfaces.ISistemaColisao, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	posX := 0.0
	posY := 0.0

	tomadaDeDecicao := r.Intn(100)

	if tomadaDeDecicao >= 0 && tomadaDeDecicao < 25 {
		posX = objeto.GetX1() - utils.BOT_VELOCIDADE_NORMAL
		posY = objeto.GetY1() - utils.BOT_VELOCIDADE_NORMAL
	} else if tomadaDeDecicao >= 25 && tomadaDeDecicao < 50 {
		posX = objeto.GetX1() - utils.BOT_VELOCIDADE_NORMAL
		posY = objeto.GetY1() + utils.BOT_VELOCIDADE_NORMAL
	} else if tomadaDeDecicao >= 50 && tomadaDeDecicao < 75 {
		posX = objeto.GetX1() + utils.BOT_VELOCIDADE_NORMAL
		posY = objeto.GetY1() + utils.BOT_VELOCIDADE_NORMAL
	} else if tomadaDeDecicao >= 75 {
		posX = objeto.GetX1() + utils.BOT_VELOCIDADE_NORMAL
		posY = objeto.GetY1() - utils.BOT_VELOCIDADE_NORMAL
	}

	// 1. Cria o retângulo da PRÓXIMA posição completa (X e Y juntos)
	proximoCorpo := geometria.NovoRetangulo(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 2. Cria o retângulo da posição ATUAL do bot para passar como filtro de auto-colisão
	corpoAtual := geometria.NovoRetangulo(objeto.GetX1(), objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 3. Checa de uma vez só: se a diagonal inteira estiver livre, ele se move.
	// Se encostar em parede, jogador ou bot, a condição falha e ele FICA PARADO na posição atual (BATEU).
	if mundo.EstaDentroDireto(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) && !sistemaColisao.VaiColidir("BOT", entidade, corpoAtual, proximoCorpo).Status {
		objeto.SetPosicao(posX, posY)
	}

}

func (self *MovimentadorDiagonal) GetTipo() string {
	return "DIAGONAL"
}

func (self *MovimentadorDiagonal) GetCor() color.Color {
	return cores.ROSA
}
