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

type MovimentadorVertical struct {
}

func (self *MovimentadorVertical) Mover(entidade ecs.Entidade, sistemaColisao interfaces.ISistemaColisao, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	posY := 0.0

	tomadaDeDecicao := r.Intn(100)

	if tomadaDeDecicao >= 50 {
		posY = objeto.GetY1() + utils.BOT_VELOCIDADE_NORMAL
	} else {
		posY = objeto.GetY1() - utils.BOT_VELOCIDADE_NORMAL
	}

	// 1. Cria o retângulo da PRÓXIMA posição vertical pretendida
	proximoCorpo := geometria.NovoRetangulo(objeto.GetX1(), posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 2. Cria o retângulo da posição ATUAL do bot para usar no filtro de auto-colisão
	corpoAtual := geometria.NovoRetangulo(objeto.GetX1(), objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 3. Validação de segurança para não deixar o cálculo de posY sair dos limites do mundo
	if posY >= mundo.PosYmax(utils.BOT_TAMANHO_MUNDO) {
		posY = mundo.PosYmax(utils.BOT_TAMANHO_MUNDO)
	} else if posY <= mundo.GetY() {
		posY = mundo.GetY()
	}

	// Atualiza o retângulo pretendido caso ele tenha sido ajustado pelas bordas do mundo acima
	proximoCorpo = geometria.NovoRetangulo(objeto.GetX1(), posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 4. Checagem final de impacto seco: Se livre, move. Se houver obstáculo, trava na hora!
	if mundo.EstaDentroDireto(objeto.GetX1(), posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) &&
		!sistemaColisao.VaiColidir("BOT", entidade, corpoAtual, proximoCorpo).Status {
		objeto.SetPosicao(objeto.GetX1(), posY)
	}
}

func (self *MovimentadorVertical) GetTipo() string {
	return "VERTICAL"
}

func (self *MovimentadorVertical) GetCor() color.Color {
	return cores.AMARELO
}
