package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"math/rand"
)

type MovimentadorVerticalConstante struct {
	direcao int
	ciclos  int
}

func (mvc *MovimentadorVerticalConstante) Mover(game interfaces.IGame, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	// 1. Definir a direção inicial se for um novo ciclo
	if mvc.ciclos == 0 {
		// Garante que não seja 0 para não ficar parado
		mvc.direcao = r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1
		// Decide aleatoriamente se começa subindo ou descendo
		if r.Float32() < 0.5 {
			mvc.direcao *= -1
		}
	}

	// 2. Calcular a intenção de movimento
	posY := objeto.GetY() + float64(mvc.direcao)
	limiteInferior := mundo.PosYmax(utils.BOT_TAMANHO_MUNDO)

	// 3. Checar colisões com as bordas do mundo
	if posY >= limiteInferior {
		posY = limiteInferior
		mvc.direcao = -(r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1) // Inverte para subir
		mvc.ciclos = 0
	} else if posY <= 0 {
		posY = 0
		mvc.direcao = r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1 // Inverte para descer
		mvc.ciclos = 0
	}

	corpo := geometria.NovoRetangulo(objeto.GetX(), posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 4. Aplicar o movimento
	if mundo.EstaDentroDireto(objeto.GetX(), posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) && !game.ColideComBarreiras(corpo) {
		objeto.SetPosicao(objeto.GetX(), posY)
		mvc.ciclos++

	}

	// 5. Reset de ciclo por tempo
	if mvc.ciclos >= utils.BOT_CICLOS_REPETICAO {
		mvc.ciclos = 0
	}
}

func (mvc *MovimentadorVerticalConstante) GetTipo() string {
	return "VERTICAL_CONSTANTE"
}
