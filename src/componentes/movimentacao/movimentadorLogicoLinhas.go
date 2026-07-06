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

type MovimentadorLogicoLinha struct {
	ciclos       int
	ciclosMaximo int
	varia        bool
	direcaoX     float64
	direcaoY     float64
}

func (self *MovimentadorLogicoLinha) Mover(entidade ecs.Entidade, sistemaColisao interfaces.ISistemaColisao, mundo *geometria.Retangulo, objeto interfaces.HabilidadeMovimentacao, r *rand.Rand) {
	self.ciclos += 1
	if self.ciclos >= self.ciclosMaximo {
		self.varia = true
		//fmt.Printf("Cheguei ao Maximo :: %d\n", mll.ciclosMaximo)
	}

	if self.varia {
		self.MovimentoLinear(r)
		self.ciclos = 0
		self.varia = false
	} else {
		//fmt.Println("\t ciclo :: %d", mll.ciclos)
	}

	alterar := true
	posX := objeto.GetX1() + self.direcaoX
	if posX >= mundo.PosXmax(utils.BOT_TAMANHO_MUNDO) {
		posX = mundo.PosXmax(utils.BOT_TAMANHO_MUNDO)
		self.bateu()
		alterar = false
	} else if posX <= mundo.GetX() {
		posX = mundo.GetX()
		self.bateu()
		alterar = false

	}

	posY := objeto.GetY1() + self.direcaoY
	if posY >= mundo.PosYmax(utils.BOT_TAMANHO_MUNDO) {
		posY = mundo.PosYmax(utils.BOT_TAMANHO_MUNDO)
		self.bateu()
		alterar = false

	} else if posY <= mundo.GetY() {
		posY = mundo.GetY()
		self.bateu()
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
			!sistemaColisao.VaiColidir("BOT", entidade, corpoAtual, proximoCorpo).Status {
			// Caminho livre: atualiza a posição
			objeto.SetPosicao(posX, posY)
		} else {
			// BATEU SECO em outra entidade: Cancela o movimento do frame
			// E aciona o comportamento para mudar de direção em linha reta
			self.bateu()
		}
	}
}

func (self *MovimentadorLogicoLinha) MovimentoLinear(r *rand.Rand) {
	tomadaDeDecicaoCiclo := r.Intn(100)
	//fmt.Printf("Tomada de Decisao :: %d\n", tomadaDeDecicaoCiclo)

	if tomadaDeDecicaoCiclo < 50 {
		self.ciclosMaximo = 30

	} else {
		self.ciclosMaximo = tomadaDeDecicaoCiclo
	}

	//fmt.Printf("Mudei de ideia heheh :: %d\n", mll.ciclos)
	//fmt.Printf("Ciclo Maximo :: %d\n", mll.ciclosMaximo)

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

func (self *MovimentadorLogicoLinha) bateu() {
	self.ciclos = 0
	self.varia = true
	//fmt.Printf("Bati! :: %d\n", mll.ciclos)

}

func (self *MovimentadorLogicoLinha) GetTipo() string {
	return "LOGICO_LINHA"
}

func (self *MovimentadorLogicoLinha) GetCor() color.Color {
	return cores.LARANJA
}
