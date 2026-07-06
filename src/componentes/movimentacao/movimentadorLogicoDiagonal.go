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

type MovimentadorLogicoDiagonal struct {
	ciclos       int
	ciclosMaximo int
	varia        bool
	direcaoX     float64
	direcaoY     float64
}

func (self *MovimentadorLogicoDiagonal) Mover(entidade ecs.Entidade, sistemaColisao interfaces.ISistemaColisao, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {

	self.ciclos += 1
	if self.ciclos >= self.ciclosMaximo {
		self.varia = true
		//fmt.Printf("Cheguei ao Maximo :: %d\n", mld.ciclosMaximo)
	}

	if self.varia {
		self.MovimentoDiagonal(r)
		self.ciclos = 0
		self.varia = false
	} else {
		//fmt.Println("\t ciclo :: %d", mld.ciclos)

	}

	// 1. Cálculo da nova intenção de posição
	posX := objeto.GetX1() + self.direcaoX
	posY := objeto.GetY1() + self.direcaoY

	// 2. Verificação de bordas e colisão com os limites do mundo
	if posX >= mundo.PosXmax(utils.BOT_TAMANHO_MUNDO) {
		posX = mundo.PosXmax(utils.BOT_TAMANHO_MUNDO)
		self.bateu()
	} else if posX <= mundo.GetX() {
		posX = mundo.GetX()
		self.bateu()
	}

	if posY >= mundo.PosYmax(utils.BOT_TAMANHO_MUNDO) {
		posY = mundo.PosYmax(utils.BOT_TAMANHO_MUNDO)
		self.bateu()
	} else if posY <= mundo.GetY() {
		posY = mundo.GetY()
		self.bateu()
	}

	// 3. Cria os retângulos para o teste de colisão ECS
	proximoCorpo := geometria.NovoRetangulo(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)
	corpoAtual := geometria.NovoRetangulo(objeto.GetX1(), objeto.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)

	// 4. Teste de Colisão Seca (Mundo + Outras Entidades)
	if mundo.EstaDentroDireto(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO) &&
		!sistemaColisao.VaiColidir("BOT", entidade, corpoAtual, proximoCorpo).Status {
		// Caminho livre: Aplica a nova posição na diagonal
		objeto.SetPosicao(posX, posY)
	} else {
		// BATEU SECO: Cancela o movimento deste frame (não chama SetPosicao)
		// COMPORTAMENTO INTELIGENTE: Aciona a sua função para recalcular uma nova direção no próximo frame
		self.bateu()
	}

}

func (self *MovimentadorLogicoDiagonal) MovimentoDiagonal(r *rand.Rand) {
	tomadaDeDecicaoCiclo := r.Intn(100)
	//fmt.Printf("Tomada de Decisao :: %d\n", tomadaDeDecicaoCiclo)

	if tomadaDeDecicaoCiclo < 50 {
		self.ciclosMaximo = 30

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

func (self *MovimentadorLogicoDiagonal) bateu() {
	self.ciclos = 0
	self.varia = true
	//fmt.Printf("Bati! :: %d\n", mld.ciclos)

}

func (self *MovimentadorLogicoDiagonal) GetTipo() string {
	return "LOGICO_DIAGONAL"
}

func (self *MovimentadorLogicoDiagonal) GetCor() color.Color {
	return cores.ROSA_ESCURO
}
