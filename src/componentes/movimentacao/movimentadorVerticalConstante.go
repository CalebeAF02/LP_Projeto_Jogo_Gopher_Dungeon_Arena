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

type MovimentadorVerticalConstante struct {
	direcao int
	ciclos  int
}

func (self *MovimentadorVerticalConstante) Mover(entidade ecs.Entidade, sistemaColisao interfaces.ISistemaColisao, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	// 1. Definir a direção inicial se for um novo ciclo
	if self.ciclos == 0 {
		// Garante que não seja 0 para não ficar parado
		self.direcao = r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1
		// Decide aleatoriamente se começa subindo ou descendo
		if r.Float32() < 0.5 {
			self.direcao *= -1
		}
	}

	// 2. Calcular a intenção de movimento
	posY := objeto.GetY1() + float64(self.direcao)
	limiteInferior := mundo.PosYmax(utils.BOT_TAMANHO_MUNDO)
	limiteSuperior := mundo.GetY() // Usando mundo.GetY() para garantir consistência com o topo do mapa

	// 3. Checar colisões com as bordas do mundo
	if posY >= limiteInferior {
		posY = limiteInferior
		self.direcao = -(r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1) // Inverte para subir
		self.ciclos = 0
	} else if posY <= limiteSuperior {
		posY = limiteSuperior
		self.direcao = r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1 // Inverte para descer
		self.ciclos = 0
	}

	// 4. Cria os retângulos para o teste de colisão ECS
	proximoCorpo := geometria.NovoRetangulo(objeto.GetX1(), posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)
	corpoAtual := geometria.NovoRetangulo(objeto.GetX1(), objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 5. Teste de Colisão Seca (Mundo + Outras Entidades)
	if mundo.EstaDentroDireto(objeto.GetX1(), posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) &&
		!sistemaColisao.VaiColidir("BOT", entidade, corpoAtual, proximoCorpo).Status {
		// Caminho livre: Aplica o movimento vertical e incrementa o ciclo
		objeto.SetPosicao(objeto.GetX1(), posY)
		self.ciclos++
	} else {
		// BATEU SECO em outra entidade (Jogador ou Bot): Cancela o movimento do frame
		// COMPORTAMENTO INTELIGENTE: Inverte o sinal da direção para ele começar a andar para o lado oposto no próximo frame
		self.direcao *= -1
		self.ciclos = 0
	}

	// 6. Reset de ciclo por tempo
	if self.ciclos >= utils.BOT_CICLOS_REPETICAO {
		self.ciclos = 0
	}
}

func (self *MovimentadorVerticalConstante) GetTipo() string {
	return "VERTICAL_CONSTANTE"
}

func (self *MovimentadorVerticalConstante) GetCor() color.Color {
	return cores.AMARELO_ESCURO
}
