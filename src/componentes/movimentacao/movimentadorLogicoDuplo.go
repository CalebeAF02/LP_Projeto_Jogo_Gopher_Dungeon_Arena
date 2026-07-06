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

type MovimentadorLogicoDuplo struct {
	ciclos       int
	ciclosMaximo int
	varia        bool
	direcaoX     float64
	direcaoY     float64
}

func (self *MovimentadorLogicoDuplo) Mover(entidade ecs.Entidade, sistemaColisao interfaces.ISistemaColisao, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	self.ciclos += 1
	if self.ciclos >= self.ciclosMaximo {
		self.varia = true
		//fmt.Printf("Cheguei ao Maximo :: %d\n", mld.ciclosMaximo)
	}

	if self.varia {
		tomadaDeDecicaoTipo := r.Intn(100)
		if tomadaDeDecicaoTipo < 50 {
			self.MovimentoLinear(r)
		} else {
			self.MovimentoDiagonal(r)
		}
		self.ciclos = 0 // Importante: Garante o reset do contador ao mudar de comportamento
		self.varia = false
	} else {
		//fmt.Println("\t ciclo :: %d", mld.ciclos)
	}

	// 1. Cálculo da nova intenção de posição
	posX := objeto.GetX1() + self.direcaoX
	if posX >= mundo.PosXmax(utils.BOT_TAMANHO_MUNDO) {
		posX = mundo.PosXmax(utils.BOT_TAMANHO_MUNDO)
		self.bateu()
	} else if posX <= mundo.GetX() {
		posX = mundo.GetX()
		self.bateu()
	}

	posY := objeto.GetY1() + self.direcaoY
	if posY >= mundo.PosYmax(utils.BOT_TAMANHO_MUNDO) {
		posY = mundo.PosYmax(utils.BOT_TAMANHO_MUNDO)
		self.bateu()
	} else if posY <= mundo.GetY() {
		posY = mundo.GetY()
		self.bateu()
	}

	// 2. Cria os retângulos para o teste de colisão ECS
	proximoCorpo := geometria.NovoRetangulo(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)
	corpoAtual := geometria.NovoRetangulo(objeto.GetX1(), objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 3. Teste de Colisão Seca (Mundo + Outras Entidades)
	if mundo.EstaDentroDireto(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) &&
		!sistemaColisao.VaiColidir("BOT", entidade, corpoAtual, proximoCorpo).Status {
		// Caminho inteiramente livre: Atualiza a posição do agente
		objeto.SetPosicao(posX, posY)
	} else {
		// BATEU SECO: O movimento deste frame é cancelado (não atualiza o SetPosicao)
		// COMPORTAMENTO INTELIGENTE: Aciona o reset para sortear uma nova direção/modo no próximo frame
		self.bateu()
	}

}

func (self *MovimentadorLogicoDuplo) MovimentoLinear(r *rand.Rand) {
	tomadaDeDecicaoCiclo := r.Intn(100)
	//fmt.Printf("Tomada de Decisao :: %d\n", tomadaDeDecicaoCiclo)

	if tomadaDeDecicaoCiclo < 50 {
		self.ciclosMaximo = 100

	} else {
		self.ciclosMaximo = tomadaDeDecicaoCiclo
	}

	//fmt.Printf("Mudei de ideia heheh :: %d\n", mld.ciclos)
	//fmt.Printf("Ciclo Maximo :: %d\n", mld.ciclosMaximo)

	tomadaDeDecicaoXouY := r.Intn(100)

	if tomadaDeDecicaoXouY > 50 {
		tomadaDeDecicaoEsqOuDir := r.Intn(100)
		self.direcaoY = 0.0
		if tomadaDeDecicaoEsqOuDir >= 50 {
			self.direcaoX = +utils.BOT_VELOCIDADE_NORMAL
		} else {
			self.direcaoX = -utils.BOT_VELOCIDADE_NORMAL
		}
	} else {
		tomadaDeDecicaoSobeOuDesce := r.Intn(100)
		self.direcaoX = 0.0
		if tomadaDeDecicaoSobeOuDesce >= 50 {
			self.direcaoY = +utils.BOT_VELOCIDADE_NORMAL
		} else {
			self.direcaoY = -utils.BOT_VELOCIDADE_NORMAL
		}
	}
}

func (self *MovimentadorLogicoDuplo) MovimentoDiagonal(r *rand.Rand) {
	tomadaDeDecicaoCiclo := r.Intn(100)
	//fmt.Printf("Tomada de Decisao :: %d\n", tomadaDeDecicaoCiclo)

	if tomadaDeDecicaoCiclo < 50 {
		self.ciclosMaximo = 100

	} else {
		self.ciclosMaximo = tomadaDeDecicaoCiclo
	}

	//fmt.Printf("Mudei de ideia heheh :: %d\n", mld.ciclos)
	//fmt.Printf("Ciclo Maximo :: %d\n", mld.ciclosMaximo)

	tomadaDeDecicao := r.Intn(100)

	if tomadaDeDecicao >= 0 && tomadaDeDecicao < 25 {
		self.direcaoX = -utils.BOT_VELOCIDADE_NORMAL
		self.direcaoY = -utils.BOT_VELOCIDADE_NORMAL

	} else if tomadaDeDecicao >= 25 && tomadaDeDecicao < 50 {
		self.direcaoX = -utils.BOT_VELOCIDADE_NORMAL
		self.direcaoY = +utils.BOT_VELOCIDADE_NORMAL

	} else if tomadaDeDecicao >= 50 && tomadaDeDecicao < 75 {
		self.direcaoX = +utils.BOT_VELOCIDADE_NORMAL
		self.direcaoY = -utils.BOT_VELOCIDADE_NORMAL

	} else if tomadaDeDecicao >= 75 && tomadaDeDecicao <= 100 {
		self.direcaoX = +utils.BOT_VELOCIDADE_NORMAL
		self.direcaoY = +utils.BOT_VELOCIDADE_NORMAL
	}
}

func (self *MovimentadorLogicoDuplo) bateu() {
	self.ciclos = 0
	self.varia = true
	//fmt.Printf("Bati! :: %d\n", mld.ciclos)

}

func (self *MovimentadorLogicoDuplo) GetTipo() string {
	return "LOGICO_DUPLO"
}

func (self *MovimentadorLogicoDuplo) GetCor() color.Color {
	return cores.VERMELHO
}
