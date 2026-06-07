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
	posY := objeto.GetY1() + float64(mvc.direcao)
	limiteInferior := mundo.PosYmax(utils.BOT_TAMANHO_MUNDO)
	limiteSuperior := mundo.GetY() // Usando mundo.GetY() para garantir consistência com o topo do mapa

	// 3. Checar colisões com as bordas do mundo
	if posY >= limiteInferior {
		posY = limiteInferior
		mvc.direcao = -(r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1) // Inverte para subir
		mvc.ciclos = 0
	} else if posY <= limiteSuperior {
		posY = limiteSuperior
		mvc.direcao = r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1 // Inverte para descer
		mvc.ciclos = 0
	}

	// 4. Cria os retângulos para o teste de colisão ECS
	proximoCorpo := geometria.NovoRetangulo(objeto.GetX1(), posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)
	corpoAtual := geometria.NovoRetangulo(objeto.GetX1(), objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 5. Teste de Colisão Seca (Mundo + Outras Entidades)
	if mundo.EstaDentroDireto(objeto.GetX1(), posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) &&
		!game.VaiColidir(corpoAtual, proximoCorpo).Status {
		// Caminho livre: Aplica o movimento vertical e incrementa o ciclo
		objeto.SetPosicao(objeto.GetX1(), posY)
		mvc.ciclos++
	} else {
		// BATEU SECO em outra entidade (Jogador ou Bot): Cancela o movimento do frame
		// COMPORTAMENTO INTELIGENTE: Inverte o sinal da direção para ele começar a andar para o lado oposto no próximo frame
		mvc.direcao *= -1
		mvc.ciclos = 0
	}

	// 6. Reset de ciclo por tempo
	if mvc.ciclos >= utils.BOT_CICLOS_REPETICAO {
		mvc.ciclos = 0
	}
}

func (mvc *MovimentadorVerticalConstante) GetTipo() string {
	return "VERTICAL_CONSTANTE"
}
