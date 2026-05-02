package movimentacao

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/interfaces"
	"math/rand"
)

type MovimentadorLogicoLinha struct {
	ciclos       int
	ciclosMaximo int
	varia        bool
	dx           float64
	dy           float64
}

func (mll *MovimentadorLogicoLinha) Mover(mundo geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
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
	posX := objeto.GetX() + mll.dx
	if posX >= mundo.PosXmax(personagens.BOT_TAMANHO) {
		posX = mundo.PosXmax(personagens.BOT_TAMANHO)
		mll.bateu()
		alterar = false
	} else if posX <= mundo.GetX() {
		posX = mundo.GetX()
		mll.bateu()
		alterar = false

	}

	posY := objeto.GetY() + mll.dy
	if posY >= mundo.PosYmax(personagens.BOT_TAMANHO) {
		posY = mundo.PosYmax(personagens.BOT_TAMANHO)
		mll.bateu()
		alterar = false

	} else if posY <= mundo.GetY() {
		posY = mundo.GetY()
		mll.bateu()
		alterar = false

	}

	if alterar {
		objeto.SetPosicao(posX, posY)
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
		mll.dy = 0.0
		if tomadaDeDecicaoEsqOuDir >= 50 {
			mll.dx = +personagens.BOT_VELOCIDADE_NORMAL
		} else {
			mll.dx = -personagens.BOT_VELOCIDADE_NORMAL
		}
	} else {
		tomadaDeDecicaoSobeOuDesce := r.Intn(100)
		mll.dx = 0.0
		if tomadaDeDecicaoSobeOuDesce >= 50 {
			mll.dy = +personagens.BOT_VELOCIDADE_NORMAL
		} else {
			mll.dy = -personagens.BOT_VELOCIDADE_NORMAL
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
