package movimentacao

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"math/rand"
)

type MovimentadorHorizontalConstante struct {
	direcao int
	ciclos  int
}

func (mhc *MovimentadorHorizontalConstante) Mover(entidade ecs.Entidade, sistemaColisao interfaces.ISistemaColisao, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {

	// 1. Início de ciclo ou inicialização (garante direção aleatória no começo)
	if mhc.ciclos >= utils.BOT_CICLOS_REPETICAO || (mhc.ciclos == 0 && mhc.direcao == 0) {
		mhc.ciclos = 0
		mhc.direcao = r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1

		// Se for o início absoluto, decide aleatoriamente a direção
		if r.Float32() < 0.5 {
			mhc.direcao *= -1
		}
	}

	// 2. Cálculo da nova intenção de posição
	posX := objeto.GetX1() + float64(mhc.direcao)
	limiteEsquerda := mundo.GetX()
	limiteDireita := mundo.PosXmax(utils.BOT_TAMANHO_MUNDO)

	// 3. Verificação de bordas do mundo e inversão automática
	if posX >= limiteDireita {
		posX = limiteDireita
		mhc.direcao = -(r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1)
		mhc.ciclos = 0
	} else if posX <= limiteEsquerda {
		posX = limiteEsquerda
		mhc.direcao = r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1
		mhc.ciclos = 0
	}

	mhc.ciclos += 1

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
		mhc.direcao *= -1
		mhc.ciclos = 0
	}

}
func (mhc *MovimentadorHorizontalConstante) GetTipo() string {
	return "HORIZONTAL_CONSTANTE"
}
