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

type MovimentadorHorizontalConstante struct {
	direcao int
	ciclos  int
}

func (self *MovimentadorHorizontalConstante) Mover(entidade ecs.Entidade, sistemaColisao interfaces.ISistemaColisao, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {

	// 1. Início de ciclo ou inicialização (garante direção aleatória no começo)
	if self.ciclos >= utils.BOT_CICLOS_REPETICAO || (self.ciclos == 0 && self.direcao == 0) {
		self.ciclos = 0
		self.direcao = r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1

		// Se for o início absoluto, decide aleatoriamente a direção
		if r.Float32() < 0.5 {
			self.direcao *= -1
		}
	}

	// 2. Cálculo da nova intenção de posição
	posX := objeto.GetX1() + float64(self.direcao)
	limiteEsquerda := mundo.GetX()
	limiteDireita := mundo.PosXmax(utils.BOT_TAMANHO_MUNDO)

	// 3. Verificação de bordas do mundo e inversão automática
	if posX >= limiteDireita {
		posX = limiteDireita
		self.direcao = -(r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1)
		self.ciclos = 0
	} else if posX <= limiteEsquerda {
		posX = limiteEsquerda
		self.direcao = r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1
		self.ciclos = 0
	}

	self.ciclos += 1

	// 4. Cria os retângulos para o teste de colisão ECS
	proximoCorpo := geometria.NovoRetangulo(posX, objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)
	corpoAtual := geometria.NovoRetangulo(objeto.GetX1(), objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 5. Teste de Colisão Seca (Mundo + Outras Entidades)
	if mundo.EstaDentroDireto(posX, objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) &&
		!sistemaColisao.VaiColidir("BOT", entidade, corpoAtual, proximoCorpo).Status {
		// Caminho livre: Aplica a nova posição
		objeto.SetPosicao(posX, objeto.GetY1())
	} else {
		// BATEU SECO: O movimento deste frame é cancelado (não chama SetPosicao)
		// COMPORTAMENTO INTELIGENTE: Força a inversão de direção para o bot não ficar preso correndo contra o obstáculo
		self.direcao *= -1
		self.ciclos = 0
	}

}
func (self *MovimentadorHorizontalConstante) GetTipo() string {
	return "HORIZONTAL_CONSTANTE"
}

func (self *MovimentadorHorizontalConstante) GetCor() color.Color {
	return cores.MARROM_ESCURO
}
