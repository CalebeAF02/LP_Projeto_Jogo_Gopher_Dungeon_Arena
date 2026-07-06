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

type MovimentadorHorizontal struct {
}

func (self *MovimentadorHorizontal) Mover(entidade ecs.Entidade, sistemaColisao interfaces.ISistemaColisao, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	posX := 0.0

	tomadaDeDecicao := r.Intn(100)

	if tomadaDeDecicao >= 50 {
		posX = objeto.GetX1() + utils.BOT_VELOCIDADE_NORMAL
	} else {
		posX = objeto.GetX1() - utils.BOT_VELOCIDADE_NORMAL
	}

	// 1. Cria o retângulo da PRÓXIMA posição horizontal pretendida
	proximoCorpo := geometria.NovoRetangulo(posX, objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 2. Cria o retângulo da posição ATUAL do bot para usar no filtro de auto-colisão
	corpoAtual := geometria.NovoRetangulo(objeto.GetX1(), objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 3. Validação de segurança para não deixar o cálculo de posX sair dos limites do mundo
	if posX >= mundo.PosXmax(utils.BOT_TAMANHO_MUNDO) {
		posX = mundo.PosXmax(utils.BOT_TAMANHO_MUNDO)
	} else if posX <= mundo.GetX() {
		posX = mundo.GetX()
	}

	// Atualiza o retângulo pretendido caso ele tenha sido ajustado pelas bordas do mundo acima
	proximoCorpo = geometria.NovoRetangulo(posX, objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 4. Checagem final: Se a nova posição estiver dentro do mundo E não colidir com ninguém, ele anda.
	// Se bater em parede, jogador ou outro bot, a condição falha e ele para seco no lugar!
	if mundo.EstaDentroDireto(posX, objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) &&
		!sistemaColisao.VaiColidir("BOT", entidade, corpoAtual, proximoCorpo).Status {
		objeto.SetPosicao(posX, objeto.GetY1())
	}

}
func (self *MovimentadorHorizontal) GetTipo() string {
	return "HORIZONTAL"
}

func (self *MovimentadorHorizontal) GetCor() color.Color {
	return cores.MARROM
}
