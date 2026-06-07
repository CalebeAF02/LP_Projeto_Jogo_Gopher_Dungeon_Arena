package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"math/rand"
)

type MovimentadorVertical struct {
}

func (mb *MovimentadorVertical) Mover(cenaJogo interfaces.ICenaJogo, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
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
		!cenaJogo.VaiColidir(corpoAtual, proximoCorpo).Status {
		objeto.SetPosicao(objeto.GetX1(), posY)
	}
}

func (mb *MovimentadorVertical) GetTipo() string {
	return "VERTICAL"
}
