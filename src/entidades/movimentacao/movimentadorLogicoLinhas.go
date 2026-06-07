package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"math/rand"
)

type MovimentadorLogicoLinha struct {
	ciclos       int
	ciclosMaximo int
	varia        bool
	direcaoX     float64
	direcaoY     float64
}

func (mll *MovimentadorLogicoLinha) Mover(game interfaces.IGame, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	mll.ciclos += 1
	if mll.ciclos >= mll.ciclosMaximo {
		mll.varia = true
		//fmt.Printf("Cheguei ao Maximo :: %d\n", mll.ciclosMaximo)
	}

	if mll.varia {
		mll.MovimentoLinear(r)
		mll.ciclos = 0
		mll.varia = false
	} else {
		//fmt.Println("\t ciclo :: %d", mll.ciclos)
	}

	alterar := true
	posX := objeto.GetX1() + mll.direcaoX
	if posX >= mundo.PosXmax(utils.BOT_TAMANHO_MUNDO) {
		posX = mundo.PosXmax(utils.BOT_TAMANHO_MUNDO)
		mll.bateu()
		alterar = false
	} else if posX <= mundo.GetX() {
		posX = mundo.GetX()
		mll.bateu()
		alterar = false

	}

	posY := objeto.GetY1() + mll.direcaoY
	if posY >= mundo.PosYmax(utils.BOT_TAMANHO_MUNDO) {
		posY = mundo.PosYmax(utils.BOT_TAMANHO_MUNDO)
		mll.bateu()
		alterar = false

	} else if posY <= mundo.GetY() {
		posY = mundo.GetY()
		mll.bateu()
		alterar = false

	}

	// Se não bateu nos limites do mundo, verifica colisão com entidades
	if alterar {
		// 1. Cria o retângulo da PRÓXIMA posição pretendida
		proximoCorpo := geometria.NovoRetangulo(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

		// 2. Cria o retângulo da posição ATUAL para ignorar a auto-colisão no ECS
		corpoAtual := geometria.NovoRetangulo(objeto.GetX1(), objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

		// 3. Teste de Colisão Rígida
		if mundo.EstaDentroDireto(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) &&
			!game.VaiColidir(corpoAtual, proximoCorpo).Status {
			// Caminho livre: atualiza a posição
			objeto.SetPosicao(posX, posY)
		} else {
			// BATEU SECO em outra entidade: Cancela o movimento do frame
			// E aciona o comportamento para mudar de direção em linha reta
			mll.bateu()
		}
	}
}

func (mll *MovimentadorLogicoLinha) MovimentoLinear(r *rand.Rand) {
	tomadaDeDecicaoCiclo := r.Intn(100)
	//fmt.Printf("Tomada de Decisao :: %d\n", tomadaDeDecicaoCiclo)

	if tomadaDeDecicaoCiclo < 50 {
		mll.ciclosMaximo = 30

	} else {
		mll.ciclosMaximo = tomadaDeDecicaoCiclo
	}

	//fmt.Printf("Mudei de ideia heheh :: %d\n", mll.ciclos)
	//fmt.Printf("Ciclo Maximo :: %d\n", mll.ciclosMaximo)

	tomadaDeDecicaoXouY := r.Intn(100)

	if tomadaDeDecicaoXouY > 50 {
		tomadaDeDecicaoEsqOuDir := r.Intn(100)
		mll.direcaoY = 0.0
		if tomadaDeDecicaoEsqOuDir >= 50 {
			mll.direcaoX = +utils.BOT_VELOCIDADE_NORMAL
		} else {
			mll.direcaoX = -utils.BOT_VELOCIDADE_NORMAL
		}
	} else {
		tomadaDeDecicaoSobeOuDesce := r.Intn(100)
		mll.direcaoX = 0.0
		if tomadaDeDecicaoSobeOuDesce >= 50 {
			mll.direcaoY = +utils.BOT_VELOCIDADE_NORMAL
		} else {
			mll.direcaoY = -utils.BOT_VELOCIDADE_NORMAL
		}
	}
}

func (mll *MovimentadorLogicoLinha) bateu() {
	mll.ciclos = 0
	mll.varia = true
	//fmt.Printf("Bati! :: %d\n", mll.ciclos)

}

func (mll *MovimentadorLogicoLinha) GetTipo() string {
	return "LOGICO_LINHA"
}
