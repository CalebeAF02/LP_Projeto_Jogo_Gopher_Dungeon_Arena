package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"math/rand"
)

type MovimentadorHorizontalConstante struct {
	direcao int
	ciclos  int
}

func (mhc *MovimentadorHorizontalConstante) Mover(game interfaces.IGame, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {

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
	posX := objeto.GetX() + float64(mhc.direcao)
	limiteEsquerda := mundo.GetX()
	limiteDireita := mundo.PosXmax(utils.BOT_TAMANHO_MUNDO)

	// 3. Verificação de bordas e inversão
	if posX >= limiteDireita {
		posX = limiteDireita
		// Sorteia nova velocidade e garante que seja negativa (esquerda)
		mhc.direcao = -(r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1)
		mhc.ciclos = 0
	} else if posX <= limiteEsquerda {
		posX = limiteEsquerda
		// Sorteia nova velocidade e garante que seja positiva (direita)
		mhc.direcao = r.Intn(utils.BOT_VELOCIDADE_MAXIMA-1) + 1
		mhc.ciclos = 0
	}

	mhc.ciclos += 1

	corpo := geometria.NovoRetangulo(posX, objeto.GetY(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	if mundo.EstaDentroDireto(posX, objeto.GetY(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) && !game.ColideComBarreiras(corpo) {
		objeto.SetPosicao(posX, objeto.GetY())
	} else {
		//objeto.SetPosicao(posX, objeto.GetY())
	}

}
func (mhc *MovimentadorHorizontalConstante) GetTipo() string {
	return "HORIZONTAL_CONSTANTE"
}
