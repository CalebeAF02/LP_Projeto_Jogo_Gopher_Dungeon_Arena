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

type MovimentadorSimples struct {
}

func (self *MovimentadorSimples) Mover(entidade ecs.Entidade, sistemaColisao interfaces.ISistemaColisao, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {

	antesX := objeto.GetX1()
	antesY := objeto.GetY1()

	// Cria o retângulo da posição ATUAL estável do bot para o filtro do ECS
	corpoAtual := geometria.NovoRetangulo(antesX, antesY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	tomadaDeDecicao := r.Intn(100)

	if tomadaDeDecicao > 50 {
		dX := r.Intn(100)
		posX := 0.0
		if dX >= 50 {
			posX = objeto.GetX1() + utils.BOT_VELOCIDADE_NORMAL
		} else {
			posX = objeto.GetX1() - utils.BOT_VELOCIDADE_NORMAL
		}
		if posX >= mundo.PosXmax(utils.BOT_TAMANHO_MUNDO) {
			posX = mundo.PosXmax(utils.BOT_TAMANHO_MUNDO)
		} else if posX <= mundo.GetX() {
			posX = mundo.GetX()
		}

		// Cria o retângulo da PRÓXIMA posição pretendida em X
		proximoCorpo := geometria.NovoRetangulo(posX, antesY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

		// Teste de Colisão Seca em X
		if mundo.EstaDentroDireto(posX, antesY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) &&
			!sistemaColisao.VaiColidir("BOT", entidade, corpoAtual, proximoCorpo).Status {
			objeto.SetPosicao(posX, antesY)
		}
	} else {
		dY := r.Intn(100)
		posY := 0.0
		if dY >= 50 {
			posY = objeto.GetY1() + utils.BOT_VELOCIDADE_NORMAL
		} else {
			posY = objeto.GetY1() - utils.BOT_VELOCIDADE_NORMAL
		}
		if posY <= mundo.GetY() {
			posY = mundo.GetY()
		} else if posY >= mundo.PosYmax(utils.BOT_TAMANHO_MUNDO) {
			posY = mundo.PosYmax(utils.BOT_TAMANHO_MUNDO)
		}

		// Cria o retângulo da PRÓXIMA posição pretendida em Y
		proximoCorpo := geometria.NovoRetangulo(antesX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

		// Teste de Colisão Seca em Y
		if mundo.EstaDentroDireto(antesX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) &&
			!sistemaColisao.VaiColidir("BOT", entidade, corpoAtual, proximoCorpo).Status {
			objeto.SetPosicao(antesX, posY)
		}
	}

}

func (self *MovimentadorSimples) GetTipo() string {
	return "SIMPLES"
}

func (self *MovimentadorSimples) GetCor() color.Color {
	return cores.VERDE_LIMAO
}
