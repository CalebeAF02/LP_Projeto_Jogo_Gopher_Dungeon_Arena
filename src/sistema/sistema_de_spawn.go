package sistema

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/movimentacao"
	"Gopher_Dungeon_Arena/src/entidades/outros"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/interfaces"
)

type SistemaSpawn struct{}

func (s *SistemaSpawn) Atualizar(g *Game) {
	// --- LÓGICA DE TEMPO PARA BOTS ---
	g.framesGeracao++

	// 180 frames = 3 segundos (em 60 FPS)
	if g.framesGeracao >= 180 {
		g.framesGeracao = 0

		// Sorteia uma posição válida (longe de paredes)
		pos := OrganizaPosicaoAleatoriaBot(g)

		// Gera o bot com um movimentador aleatório
		GerarBot(g, pos.GetX(), pos.GetY())
	}
}

func SpawnJogadores(g *Game) {
	// Jogadores
	j1 := personagens.NovoJogador(g, "Jogador 1")
	//j2 := personagens.NovoJogador(&g, "Jogador 2")
	//j3 := personagens.NovoJogador(&g, "Jogador 3")

	//j4 := personagens.NovoJogador(&g, "Jogador 4")
	//j5 := personagens.NovoJogador(&g, "Jogador 5")
	//j6 := personagens.NovoJogador(&g, "Jogador 6")

	j1.SetPosicao(100, 100)
	//j2.SetPosicao(200, 300)
	//j3.SetPosicao(300, 500)

	//j4.SetPosicao(500, 100)
	//j5.SetPosicao(300, 200)
	//j6.SetPosicao(500, 300)

	// Times
	t1 := outros.NovoTime(g, "Vermelhao - Time_Vermelho", cores.AZUL)
	//t2 := outros.NovoTime(&g, "Azulzinhos - Time_Azul", cores.AZUL)

	// Gerenciando
	t1.Adicionnar(j1)
	//t1.Adicionnar(j2)
	//t1.Adicionnar(j3)
	//t1.Posicoes()

	//t2.Adicionnar(j4)
	//t2.Adicionnar(j5)
	//t2.Adicionnar(j6)
	//t2.Posicoes()
}

func SpawnBots(g *Game) {
	//Bot
	for id := 0; id < 1; id++ {
		CriarBot(g, &movimentacao.MovimentadorSimples{}, OrganizaPosicaoAleatoriaBot(g))
		CriarBot(g, &movimentacao.MovimentadorVertical{}, OrganizaPosicaoAleatoriaBot(g))
		CriarBot(g, &movimentacao.MovimentadorVerticalConstante{}, OrganizaPosicaoAleatoriaBot(g))
		CriarBot(g, &movimentacao.MovimentadorHorizontal{}, OrganizaPosicaoAleatoriaBot(g))
		CriarBot(g, &movimentacao.MovimentadorHorizontalConstante{}, OrganizaPosicaoAleatoriaBot(g))
		CriarBot(g, &movimentacao.MovimentadorDiagonal{}, OrganizaPosicaoAleatoriaBot(g))
		CriarBot(g, &movimentacao.MovimentadorLogicoLinha{}, OrganizaPosicaoAleatoriaBot(g))
		CriarBot(g, &movimentacao.MovimentadorLogicoDiagonal{}, OrganizaPosicaoAleatoriaBot(g))
		CriarBot(g, &movimentacao.MovimentadorLogicoDuplo{}, OrganizaPosicaoAleatoriaBot(g))
	}
}

func SpawnParedesEspecificas(g *Game) {
	// --- LETRA L ---
	outros.NovaParede(g, geometria.NovoPonto(200, 400))
	outros.NovaParede(g, geometria.NovoPonto(200, 430))
	outros.NovaParede(g, geometria.NovoPonto(200, 460))
	outros.NovaParede(g, geometria.NovoPonto(200, 490))
	outros.NovaParede(g, geometria.NovoPonto(230, 490))
	outros.NovaParede(g, geometria.NovoPonto(260, 490))

	// --- LETRA U ---
	outros.NovaParede(g, geometria.NovoPonto(320, 400))
	outros.NovaParede(g, geometria.NovoPonto(320, 430))
	outros.NovaParede(g, geometria.NovoPonto(320, 460))
	outros.NovaParede(g, geometria.NovoPonto(320, 490))
	outros.NovaParede(g, geometria.NovoPonto(350, 490))
	outros.NovaParede(g, geometria.NovoPonto(380, 490))
	outros.NovaParede(g, geometria.NovoPonto(380, 460))
	outros.NovaParede(g, geometria.NovoPonto(380, 430))
	outros.NovaParede(g, geometria.NovoPonto(380, 400))

	// --- LETRA A ---
	outros.NovaParede(g, geometria.NovoPonto(440, 490))
	outros.NovaParede(g, geometria.NovoPonto(440, 460))
	outros.NovaParede(g, geometria.NovoPonto(440, 430))
	outros.NovaParede(g, geometria.NovoPonto(470, 400))
	outros.NovaParede(g, geometria.NovoPonto(500, 400))
	outros.NovaParede(g, geometria.NovoPonto(530, 430))
	outros.NovaParede(g, geometria.NovoPonto(530, 460))
	outros.NovaParede(g, geometria.NovoPonto(530, 490))
	outros.NovaParede(g, geometria.NovoPonto(470, 450))
	outros.NovaParede(g, geometria.NovoPonto(500, 450))

	// --- LETRA N ---
	outros.NovaParede(g, geometria.NovoPonto(590, 490))
	outros.NovaParede(g, geometria.NovoPonto(590, 460))
	outros.NovaParede(g, geometria.NovoPonto(590, 430))
	outros.NovaParede(g, geometria.NovoPonto(590, 400))
	outros.NovaParede(g, geometria.NovoPonto(620, 430))
	outros.NovaParede(g, geometria.NovoPonto(650, 460))
	outros.NovaParede(g, geometria.NovoPonto(680, 400))
	outros.NovaParede(g, geometria.NovoPonto(680, 430))
	outros.NovaParede(g, geometria.NovoPonto(680, 460))
	outros.NovaParede(g, geometria.NovoPonto(680, 490))
}

func CriarBot(g *Game, movendo interfaces.Movimentador, posicao *geometria.Ponto) {
	b := personagens.NovoBot(g, 0)
	b.SetPosicao(posicao.GetX(), posicao.GetY())
	b.SetMovimentacao(movendo)

	if movendo.GetTipo() == "LOGICO_LINHA" {
		b.SetCor(cores.AMARELO)
	} else if movendo.GetTipo() == "LOGICO_DIAGONAL" {
		b.SetCor(cores.VERDE)
	} else if movendo.GetTipo() == "LOGICO_DUPLO" {
		b.SetCor(cores.AZUL)
	} else if movendo.GetTipo() == "SIMPLES" {
		b.SetCor(cores.VERDE_LIMAO)
	} else if movendo.GetTipo() == "VERTICAL" {
		b.SetCor(cores.AMARELO_CLARO)
	} else if movendo.GetTipo() == "VERTICAL_CONSTANTE" {
		b.SetCor(cores.AMARELO_ESCURO)
	} else if movendo.GetTipo() == "HORIZONTAL" {
		b.SetCor(cores.MARROM)
	} else if movendo.GetTipo() == "HORIZONTAL_CONSTANTE" {
		b.SetCor(cores.MARROM_ESCURO)
	} else if movendo.GetTipo() == "DIAGONAL" {
		b.SetCor(cores.ROSA)
	}

	b.SetMovimentacao(movendo)
	//fmt.Printf("BOT <%s> | X: %f | Y: %f\n", b.GetMovendoTipo(), b.GetX(), b.GetY())
}

func CriarBotAleatorio(g *Game) {

	for id := 0; id < 10; id++ {
		b := personagens.NovoBot(g, int64(id))
		b.SetPosicao(config.XAleatorio(g.aleatorio), config.YAleatorio(g.aleatorio))

		movimentacaoAleatoria := g.aleatorio.Intn(100)
		if movimentacaoAleatoria >= 0 && movimentacaoAleatoria < 15 {
			b.SetMovimentacao(&movimentacao.MovimentadorSimples{})
			b.SetCor(cores.BRANCO)
		} else if movimentacaoAleatoria >= 15 && movimentacaoAleatoria < 40 {
			b.SetMovimentacao(&movimentacao.MovimentadorVertical{})
			b.SetCor(cores.VERDE)
		} else if movimentacaoAleatoria >= 40 && movimentacaoAleatoria < 60 {
			b.SetMovimentacao(&movimentacao.MovimentadorHorizontalConstante{})
			b.SetCor(cores.LARANJA)
		} else if movimentacaoAleatoria >= 60 && movimentacaoAleatoria < 80 {
			b.SetMovimentacao(&movimentacao.MovimentadorVerticalConstante{})
			b.SetCor(cores.VERDE)
		} else {
			b.SetMovimentacao(&movimentacao.MovimentadorDiagonal{})
			b.SetCor(cores.CIANO)
		}

		if g.aleatorio.Intn(100) < 30 {
		} else {
			valor := g.aleatorio.Intn(100)
			if valor < 30 {

			} else if valor > 30 && valor < 50 {
				v2 := g.aleatorio.Intn(100)
				if v2 > 50 {
				}
			} else if valor > 50 && valor < 70 {

			} else {
				b.SetMovimentacao(&movimentacao.MovimentadorHorizontal{})
				b.SetCor(cores.LARANJA)
			}
		}
	}
}

func SpawnLabirinto(g *Game) {

	inicioX, inicioY := 800.0, 800.0
	tamanho := 600.0
	passo := 30.0

	for i := 0.0; i <= tamanho; i += passo {
		// Topo e Base
		outros.NovaParede(g, geometria.NovoPonto(inicioX+i, inicioY))
		outros.NovaParede(g, geometria.NovoPonto(inicioX+i, inicioY+tamanho))

		// Lateral Direita
		outros.NovaParede(g, geometria.NovoPonto(inicioX+tamanho, inicioY+i))

		// Lateral Esquerda (com abertura entre 950 e 1010 para entrar)
		posAtualY := inicioY + i
		if posAtualY < 950 || posAtualY > 1010 {
			outros.NovaParede(g, geometria.NovoPonto(inicioX, posAtualY))
		}
	}

	// --- Grande Parede Vertical Central (Divide o labirinto em dois) ---
	for y := 860.0; y <= 1340; y += passo {
		if y < 1070 || y > 1130 { // Buraco no meio da parede central para passar
			outros.NovaParede(g, geometria.NovoPonto(1100, y))
		}
	}

	// --- Ala Esquerda (Setor de Entrada) ---
	// Uma barreira horizontal superior
	for x := 860.0; x <= 1040; x += passo {
		outros.NovaParede(g, geometria.NovoPonto(x, 920))
	}

	// Uma barreira horizontal inferior
	for x := 860.0; x <= 1040; x += passo {
		outros.NovaParede(g, geometria.NovoPonto(x, 1280))
	}

	// --- Ala Direita (Setor de Desafio) ---
	// Obstáculo em forma de 'T'
	for x := 1200.0; x <= 1350; x += passo {
		outros.NovaParede(g, geometria.NovoPonto(x, 1000)) // Parte de cima do T
	}
	outros.NovaParede(g, geometria.NovoPonto(1275, 1030))
	outros.NovaParede(g, geometria.NovoPonto(1275, 1060))

	// Paredes tipo "Dentes de Pente" na lateral direita
	for y := 1150.0; y <= 1300; y += passo {
		outros.NovaParede(g, geometria.NovoPonto(1340, y))
		outros.NovaParede(g, geometria.NovoPonto(1250, y))
	}

}

func SpawnParedesAoRedor(g *Game, passo float64) {
	xMin := g.mundo.GetX()
	yMin := g.mundo.GetY()
	xMax := g.mundo.GetX() + g.mundo.GetLargura()
	yMax := g.mundo.GetY() + g.mundo.GetAltura()

	// 1. Paredes Horizontais (Topo e Base)
	for x := xMin; x < xMax; x += passo {
		outros.NovaParede(g, geometria.NovoPonto(x, yMin))       // Linha de cima
		outros.NovaParede(g, geometria.NovoPonto(x, yMax-passo)) // Linha de baixo
	}

	// 2. Paredes Verticais (Esquerda e Direita)
	// Começamos em yMin + passo para não sobrepor os cantos já criados
	for y := yMin + passo; y < yMax-passo; y += passo {
		outros.NovaParede(g, geometria.NovoPonto(xMin, y))       // Lateral esquerda
		outros.NovaParede(g, geometria.NovoPonto(xMax-passo, y)) // Lateral direita
	}
}

func GerarBot(g *Game, x float64, y float64) {

	tipo := g.aleatorio.Intn(10)
	if tipo == 0 {
		CriarBot(g, &movimentacao.MovimentadorHorizontal{}, geometria.NovoPonto(x, y))
	} else if tipo == 1 {
		CriarBot(g, &movimentacao.MovimentadorHorizontalConstante{}, geometria.NovoPonto(x, y))
	} else if tipo == 2 {
		CriarBot(g, &movimentacao.MovimentadorVertical{}, geometria.NovoPonto(x, y))
	} else if tipo == 3 {
		CriarBot(g, &movimentacao.MovimentadorVerticalConstante{}, geometria.NovoPonto(x, y))
	} else if tipo == 4 {
		CriarBot(g, &movimentacao.MovimentadorDiagonal{}, geometria.NovoPonto(x, y))
	} else if tipo == 5 {
		CriarBot(g, &movimentacao.MovimentadorLogicoDiagonal{}, geometria.NovoPonto(x, y))
	} else {
		CriarBot(g, &movimentacao.MovimentadorDiagonal{}, geometria.NovoPonto(x, y))
	}

}
