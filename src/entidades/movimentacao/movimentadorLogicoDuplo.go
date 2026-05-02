package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/interfaces"
	"math/rand"
)

type MovimentadorLogicoDuplo struct {
	ciclos       int
	ciclosMaximo int
	varia        bool
	dx           float64
	dy           float64
}

func (mld *MovimentadorLogicoDuplo) Mover(mundo geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
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
		mld.varia = false
	} else {
		//fmt.Println("\t ciclo :: %d", mld.ciclos)
	}

	posX := objeto.GetX() + mld.dx
	if posX >= mundo.PosXmax(personagens.BOT_TAMANHO) {
		posX = mundo.PosXmax(personagens.BOT_TAMANHO)
		mld.bateu()
	} else if posX <= mundo.GetX() {
		posX = mundo.GetX()
		mld.bateu()
	}

	posY := objeto.GetY() + mld.dy
	if posY >= mundo.PosYmax(personagens.BOT_TAMANHO) {
		posY = mundo.PosYmax(personagens.BOT_TAMANHO)
		mld.bateu()

	} else if posY <= mundo.GetY() {
		posY = mundo.GetY()
		mld.bateu()
	}
	objeto.SetPosicao(posX, posY)
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
		mld.dy = 0.0
		if tomadaDeDecicaoEsqOuDir >= 50 {
			mld.dx = +personagens.BOT_VELOCIDADE_NORMAL
		} else {
			mld.dx = -personagens.BOT_VELOCIDADE_NORMAL
		}
	} else {
		tomadaDeDecicaoSobeOuDesce := r.Intn(100)
		mld.dx = 0.0
		if tomadaDeDecicaoSobeOuDesce >= 50 {
			mld.dy = +personagens.BOT_VELOCIDADE_NORMAL
		} else {
			mld.dy = -personagens.BOT_VELOCIDADE_NORMAL
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
		mld.dx = -personagens.BOT_VELOCIDADE_NORMAL
		mld.dy = -personagens.BOT_VELOCIDADE_NORMAL

	} else if tomadaDeDecicao >= 25 && tomadaDeDecicao < 50 {
		mld.dx = -personagens.BOT_VELOCIDADE_NORMAL
		mld.dy = +personagens.BOT_VELOCIDADE_NORMAL

	} else if tomadaDeDecicao >= 50 && tomadaDeDecicao < 75 {
		mld.dx = +personagens.BOT_VELOCIDADE_NORMAL
		mld.dy = -personagens.BOT_VELOCIDADE_NORMAL

	} else if tomadaDeDecicao >= 75 && tomadaDeDecicao <= 100 {
		mld.dx = +personagens.BOT_VELOCIDADE_NORMAL
		mld.dy = +personagens.BOT_VELOCIDADE_NORMAL
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
