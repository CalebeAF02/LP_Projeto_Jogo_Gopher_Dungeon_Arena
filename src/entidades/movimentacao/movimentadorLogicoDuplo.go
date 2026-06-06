package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"math/rand"
)

type MovimentadorLogicoDuplo struct {
	ciclos       int
	ciclosMaximo int
	varia        bool
	direcaoX     float64
	direcaoY     float64
}

func (mld *MovimentadorLogicoDuplo) Mover(game interfaces.IGame, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	mld.ciclos += 1
	if mld.ciclos >= mld.ciclosMaximo {
		mld.varia = true
		//fmt.Printf("Cheguei ao Maximo :: %d\n", mld.ciclosMaximo)
	}

	if mld.varia {
		tomadaDeDecicaoTipo := r.Intn(100)
		if tomadaDeDecicaoTipo < 50 {
			mld.MovimentoLinear(r)
		} else {
			mld.MovimentoDiagonal(r)
		}
		mld.ciclos = 0 // Importante: Garante o reset do contador ao mudar de comportamento
		mld.varia = false
	} else {
		//fmt.Println("\t ciclo :: %d", mld.ciclos)
	}

	// 1. Cálculo da nova intenção de posição
	posX := objeto.GetX1() + mld.direcaoX
	if posX >= mundo.PosXmax(utils.BOT_TAMANHO_MUNDO) {
		posX = mundo.PosXmax(utils.BOT_TAMANHO_MUNDO)
		mld.bateu()
	} else if posX <= mundo.GetX() {
		posX = mundo.GetX()
		mld.bateu()
	}

	posY := objeto.GetY1() + mld.direcaoY
	if posY >= mundo.PosYmax(utils.BOT_TAMANHO_MUNDO) {
		posY = mundo.PosYmax(utils.BOT_TAMANHO_MUNDO)
		mld.bateu()
	} else if posY <= mundo.GetY() {
		posY = mundo.GetY()
		mld.bateu()
	}

	// 2. Cria os retângulos para o teste de colisão ECS
	proximoCorpo := geometria.NovoRetangulo(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)
	corpoAtual := geometria.NovoRetangulo(objeto.GetX1(), objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 3. Teste de Colisão Seca (Mundo + Outras Entidades)
	if mundo.EstaDentroDireto(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) &&
		!game.VaiColidir(corpoAtual, proximoCorpo) {
		// Caminho inteiramente livre: Atualiza a posição do agente
		objeto.SetPosicao(posX, posY)
	} else {
		// BATEU SECO: O movimento deste frame é cancelado (não atualiza o SetPosicao)
		// COMPORTAMENTO INTELIGENTE: Aciona o reset para sortear uma nova direção/modo no próximo frame
		mld.bateu()
	}

}

func (mld *MovimentadorLogicoDuplo) MovimentoLinear(r *rand.Rand) {
	tomadaDeDecicaoCiclo := r.Intn(100)
	//fmt.Printf("Tomada de Decisao :: %d\n", tomadaDeDecicaoCiclo)

	if tomadaDeDecicaoCiclo < 50 {
		mld.ciclosMaximo = 30

	} else {
		mld.ciclosMaximo = tomadaDeDecicaoCiclo
	}

	//fmt.Printf("Mudei de ideia heheh :: %d\n", mld.ciclos)
	//fmt.Printf("Ciclo Maximo :: %d\n", mld.ciclosMaximo)

	tomadaDeDecicaoXouY := r.Intn(100)

	if tomadaDeDecicaoXouY > 50 {
		tomadaDeDecicaoEsqOuDir := r.Intn(100)
		mld.direcaoY = 0.0
		if tomadaDeDecicaoEsqOuDir >= 50 {
			mld.direcaoX = +utils.BOT_VELOCIDADE_NORMAL
		} else {
			mld.direcaoX = -utils.BOT_VELOCIDADE_NORMAL
		}
	} else {
		tomadaDeDecicaoSobeOuDesce := r.Intn(100)
		mld.direcaoX = 0.0
		if tomadaDeDecicaoSobeOuDesce >= 50 {
			mld.direcaoY = +utils.BOT_VELOCIDADE_NORMAL
		} else {
			mld.direcaoY = -utils.BOT_VELOCIDADE_NORMAL
		}
	}
}

func (mld *MovimentadorLogicoDuplo) MovimentoDiagonal(r *rand.Rand) {
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

func (mld *MovimentadorLogicoDuplo) bateu() {
	mld.ciclos = 0
	mld.varia = true
	//fmt.Printf("Bati! :: %d\n", mld.ciclos)

}

func (mld *MovimentadorLogicoDuplo) GetTipo() string {
	return "LOGICO_DUPLO"
}
