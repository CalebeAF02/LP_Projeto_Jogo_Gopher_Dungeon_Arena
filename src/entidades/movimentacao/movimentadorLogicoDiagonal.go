package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"math/rand"
)

type MovimentadorLogicoDiagonal struct {
	ciclos       int
	ciclosMaximo int
	varia        bool
	direcaoX     float64
	direcaoY     float64
}

func (mld *MovimentadorLogicoDiagonal) Mover(game interfaces.IGame, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {

	mld.ciclos += 1
	if mld.ciclos >= mld.ciclosMaximo {
		mld.varia = true
		//fmt.Printf("Cheguei ao Maximo :: %d\n", mld.ciclosMaximo)
	}

	if mld.varia {
		mld.MovimentoDiagonal(r)
		mld.ciclos = 0
		mld.varia = false
	} else {
		//fmt.Println("\t ciclo :: %d", mld.ciclos)

	}

	// 1. Cálculo da nova intenção de posição
	posX := objeto.GetX1() + mld.direcaoX
	posY := objeto.GetY1() + mld.direcaoY

	// 2. Verificação de bordas e colisão com os limites do mundo
	if posX >= mundo.PosXmax(utils.BOT_TAMANHO_MUNDO) {
		posX = mundo.PosXmax(utils.BOT_TAMANHO_MUNDO)
		mld.bateu()
	} else if posX <= mundo.GetX() {
		posX = mundo.GetX()
		mld.bateu()
	}

	if posY >= mundo.PosYmax(utils.BOT_TAMANHO_MUNDO) {
		posY = mundo.PosYmax(utils.BOT_TAMANHO_MUNDO)
		mld.bateu()
	} else if posY <= mundo.GetY() {
		posY = mundo.GetY()
		mld.bateu()
	}

	// 3. Cria os retângulos para o teste de colisão ECS
	proximoCorpo := geometria.NovoRetangulo(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)
	corpoAtual := geometria.NovoRetangulo(objeto.GetX1(), objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 4. Teste de Colisão Seca (Mundo + Outras Entidades)
	if mundo.EstaDentroDireto(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) &&
		!game.VaiColidir(corpoAtual, proximoCorpo) {
		// Caminho livre: Aplica a nova posição na diagonal
		objeto.SetPosicao(posX, posY)
	} else {
		// BATEU SECO: Cancela o movimento deste frame (não chama SetPosicao)
		// COMPORTAMENTO INTELIGENTE: Aciona a sua função para recalcular uma nova direção no próximo frame
		mld.bateu()
	}

}

func (mld *MovimentadorLogicoDiagonal) MovimentoDiagonal(r *rand.Rand) {
	tomadaDeDecicaoCiclo := r.Intn(100)
	//fmt.Printf("Tomada de Decisao :: %d\n", tomadaDeDecicaoCiclo)

	if tomadaDeDecicaoCiclo < 50 {
		mld.ciclosMaximo = 30

	} else {
		mld.ciclosMaximo = tomadaDeDecicaoCiclo
	}

	//fmt.Printf("Mudei de ideia heheh :: %d\n", mld.ciclos)
	//fmt.Printf("Ciclo Maximo :: %d\n", mld.ciclosMaximo)

	tomadaDeDecicao := r.Intn(100)

	if tomadaDeDecicao >= 0 && tomadaDeDecicao < 25 {
		mld.direcaoX = -utils.BOT_VELOCIDADE_NORMAL
		mld.direcaoY = -utils.BOT_VELOCIDADE_NORMAL

	} else if tomadaDeDecicao >= 25 && tomadaDeDecicao < 50 {
		mld.direcaoX = -utils.BOT_VELOCIDADE_NORMAL
		mld.direcaoY = +utils.BOT_VELOCIDADE_NORMAL

	} else if tomadaDeDecicao >= 50 && tomadaDeDecicao < 75 {
		mld.direcaoX = +utils.BOT_VELOCIDADE_NORMAL
		mld.direcaoY = -utils.BOT_VELOCIDADE_NORMAL

	} else if tomadaDeDecicao >= 75 && tomadaDeDecicao <= 100 {
		mld.direcaoX = +utils.BOT_VELOCIDADE_NORMAL
		mld.direcaoY = +utils.BOT_VELOCIDADE_NORMAL
	}
}

func (mld *MovimentadorLogicoDiagonal) bateu() {
	mld.ciclos = 0
	mld.varia = true
	//fmt.Printf("Bati! :: %d\n", mld.ciclos)

}

func (mld *MovimentadorLogicoDiagonal) GetTipo() string {
	return "LOGICO_DIAGONAL"
}
