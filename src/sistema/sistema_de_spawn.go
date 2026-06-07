package sistema

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/movimentacao"
	"Gopher_Dungeon_Arena/src/entidades/objeto"
	"Gopher_Dungeon_Arena/src/entidades/outros"

	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/interfaces"

	"github.com/hajimehoshi/ebiten/v2"
)

type SistemaSpawn struct {
	framesGereacao int
}

func (s *SistemaSpawn) Atualizar(cj interfaces.ICenaJogo) {
	// --- LÓGICA DE TEMPO PARA BOTS ---
	s.framesGereacao++

	// 1860 = 60*30 frames = 30 segundos (em 60 FPS)
	if s.framesGereacao >= 1860 {
		s.framesGereacao = 0

		if pos := cj.OrganizaPosicaoAleatoriaBot(); pos != nil {
			s.SpawnarBotAleatorio(cj, pos.GetX(), pos.GetY())
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyB) {
		s.SpawnarBotsAleatroiamenteNoMundo(cj)
	}
}

func (s *SistemaSpawn) SpawnJogadores(cj interfaces.ICenaJogo) {
	// Jogadores
	j1 := personagens.NovoJogador(cj, "Jogador 1")
	//j2 := personagens.NovoJogador(&cj, "Jogador 2")
	//j3 := personagens.NovoJogador(&cj, "Jogador 3")

	//j4 := personagens.NovoJogador(&cj, "Jogador 4")
	//j5 := personagens.NovoJogador(&cj, "Jogador 5")
	//j6 := personagens.NovoJogador(&cj, "Jogador 6")

	j1.SetPosicao(100, 100)
	//j2.SetPosicao(200, 300)
	//j3.SetPosicao(300, 500)

	//j4.SetPosicao(500, 100)
	//j5.SetPosicao(300, 200)
	//j6.SetPosicao(500, 300)

	j1.SetNivel(1)

	// Times
	t1 := outros.NovoTime(cj, "Vermelhao - Time_Azul", cores.AZUL)
	//t2 := outros.NovoTime(&cj, "Azulzinhos - Time_Vermelho", cores.VERMELHO)

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

func (s *SistemaSpawn) SpawnarBot(cj interfaces.ICenaJogo, movendo interfaces.Movimentador, posicao *geometria.Ponto) {
	b := personagens.NovoBot(cj, 0)
	b.SetNivelAleatorio()
	b.SetPosicao(posicao.GetX(), posicao.GetY())

	// Define as cores com base estrita na assinatura de tipo do movimento injetado
	switch movendo.GetTipo() {
	case "LOGICO_LINHA":
		b.SetCor(cores.AMARELO)
	case "LOGICO_DIAGONAL":
		b.SetCor(cores.VERDE)
	case "LOGICO_DUPLO":
		b.SetCor(cores.AZUL)
	case "SIMPLES":
		b.SetCor(cores.VERDE_LIMAO)
	case "VERTICAL":
		b.SetCor(cores.AMARELO_CLARO)
	case "VERTICAL_CONSTANTE":
		b.SetCor(cores.AMARELO_ESCURO)
	case "HORIZONTAL":
		b.SetCor(cores.MARROM)
	case "HORIZONTAL_CONSTANTE":
		b.SetCor(cores.MARROM_ESCURO)
	case "DIAGONAL":
		b.SetCor(cores.ROSA)
	}

	b.SetMovimentacao(movendo)
	//fmt.Printf("BOT <%s> | X: %f | Y: %f\n", b.GetMovendoTipo(), b.GetX(), b.GetY())

}

func (s *SistemaSpawn) SpawnBotDeCadaTipo(cj interfaces.ICenaJogo) {
	// Cria uma lista dos movimentadores desejados para iterar de forma limpa e segura
	movimentadores := []interfaces.Movimentador{
		&movimentacao.MovimentadorSimples{},
		&movimentacao.MovimentadorVertical{},
		&movimentacao.MovimentadorVerticalConstante{},
		&movimentacao.MovimentadorHorizontal{},
		&movimentacao.MovimentadorHorizontalConstante{},
		&movimentacao.MovimentadorDiagonal{},
		&movimentacao.MovimentadorLogicoLinha{},
		&movimentacao.MovimentadorLogicoDiagonal{},
		&movimentacao.MovimentadorLogicoDuplo{},
	}

	// Varre a lista garantindo que cada bot receba uma posição válida individual
	for _, mov := range movimentadores {
		if pos := cj.OrganizaPosicaoAleatoriaBot(); pos != nil {
			s.SpawnarBot(cj, mov, pos)
		}
	}
}

func (s *SistemaSpawn) SpawnarBotAleatorio(cj interfaces.ICenaJogo, x float64, y float64) {
	// Sorteia dinamicamente entre os 9 tipos de movimentadores que você possui
	tipo := cj.GetAleatorio().Intn(9)
	posicao := geometria.NovoPonto(x, y)

	switch tipo {
	case 0:
		s.SpawnarBot(cj, &movimentacao.MovimentadorSimples{}, posicao)
	case 1:
		s.SpawnarBot(cj, &movimentacao.MovimentadorVertical{}, posicao)
	case 2:
		s.SpawnarBot(cj, &movimentacao.MovimentadorVerticalConstante{}, posicao)
	case 3:
		s.SpawnarBot(cj, &movimentacao.MovimentadorHorizontal{}, posicao)
	case 4:
		s.SpawnarBot(cj, &movimentacao.MovimentadorHorizontalConstante{}, posicao)
	case 5:
		s.SpawnarBot(cj, &movimentacao.MovimentadorDiagonal{}, posicao)
	case 6:
		s.SpawnarBot(cj, &movimentacao.MovimentadorLogicoLinha{}, posicao)
	case 7:
		s.SpawnarBot(cj, &movimentacao.MovimentadorLogicoDiagonal{}, posicao)
	case 8:
		s.SpawnarBot(cj, &movimentacao.MovimentadorLogicoDuplo{}, posicao)
	}
}

func (s *SistemaSpawn) SpawnarBotsAleatroiamenteNoMundo(cj interfaces.ICenaJogo) {
	for id := 0; id < 3; id++ {
		if pos := cj.OrganizaPosicaoAleatoriaBot(); pos != nil {
			s.SpawnarBotAleatorio(cj, pos.GetX(), pos.GetY())
		}
	}
}

func (s *SistemaSpawn) SpawnParedesAoRedor(cj interfaces.ICenaJogo, passo float64) {
	xMin := cj.GetMundo().GetX()
	yMin := cj.GetMundo().GetY()
	xMax := cj.GetMundo().GetX() + cj.GetMundo().GetLargura()
	yMax := cj.GetMundo().GetY() + cj.GetMundo().GetAltura()

	// 1. Paredes Horizontais (Topo e Base)
	for x := xMin; x < xMax; x += passo {
		objeto.NovaParede(cj, geometria.NovoPonto(x, yMin))       // Linha de cima
		objeto.NovaParede(cj, geometria.NovoPonto(x, yMax-passo)) // Linha de baixo
	}

	// 2. Paredes Verticais (Esquerda e Direita)
	// Começamos em yMin + passo para não sobrepor os cantos já criados
	for y := yMin + passo; y < yMax-passo; y += passo {
		objeto.NovaParede(cj, geometria.NovoPonto(xMin, y))       // Lateral esquerda
		objeto.NovaParede(cj, geometria.NovoPonto(xMax-passo, y)) // Lateral direita
	}
}

func (s *SistemaSpawn) SpawnLabirinto(cj interfaces.ICenaJogo) {
	// 1. Coleta os dados reais e dinâmicos do tamanho do mundo
	mundoX := cj.GetMundo().GetX()
	mundoY := cj.GetMundo().GetY()
	mundoLargura := cj.GetMundo().GetLargura()
	mundoAltura := cj.GetMundo().GetAltura()

	// 2. Define o tamanho do Labirinto de forma proporcional ao mundo
	// Aqui fazemos com que ele ocupe 80% da menor dimensão do mapa, por exemplo
	tamanho := mundoLargura * 0.8
	if mundoAltura < mundoLargura {
		tamanho = mundoAltura * 0.8
	}

	passo := 30.0

	// 3. Centraliza o Labirinto dinamicamente no meio do mundo
	inicioX := mundoX + (mundoLargura-tamanho)/2
	inicioY := mundoY + (mundoAltura-tamanho)/2

	// --- Contorno do Labirinto (Gerado com base no novo início dinâmico) ---
	// Calculamos onde ficam as aberturas de forma proporcional ao tamanho
	aberturaEntradaMin := inicioY + (tamanho * 0.4) // Abertura centralizada na lateral
	aberturaEntradaMax := aberturaEntradaMin + (passo * 2)

	for i := 0.0; i <= tamanho; i += passo {
		// Topo e Base
		objeto.NovaParede(cj, geometria.NovoPonto(inicioX+i, inicioY))
		objeto.NovaParede(cj, geometria.NovoPonto(inicioX+i, inicioY+tamanho))

		// Lateral Direita
		objeto.NovaParede(cj, geometria.NovoPonto(inicioX+tamanho, inicioY+i))

		// Lateral Esquerda com abertura dinâmica para o jogador entrar
		posAtualY := inicioY + i
		if posAtualY < aberturaEntradaMin || posAtualY > aberturaEntradaMax {
			objeto.NovaParede(cj, geometria.NovoPonto(inicioX, posAtualY))
		}
	}

	// --- Grande Parede Vertical Central (Divide o labirinto exatamente ao meio) ---
	centroX := inicioX + (tamanho / 2)
	paredeCentralYMin := inicioY + (tamanho * 0.15)
	paredeCentralYMax := inicioY + (tamanho * 0.85)
	aberturaCentralMin := inicioY + (tamanho * 0.45)
	aberturaCentralMax := aberturaCentralMin + (passo * 2)

	for y := paredeCentralYMin; y <= paredeCentralYMax; y += passo {
		if y < aberturaCentralMin || y > aberturaCentralMax { // Buraco dinâmico no meio para passar
			objeto.NovaParede(cj, geometria.NovoPonto(centroX, y))
		}
	}

	// --- Ala Esquerda (Setor de Entrada - Barreiras Horizontais Proporcionais) ---
	alaEsquerdaXMin := inicioX + (tamanho * 0.1)
	alaEsquerdaXMax := centroX - (tamanho * 0.1)

	barreiraSuperiorY := inicioY + (tamanho * 0.25)
	barreiraInferiorY := inicioY + (tamanho * 0.75)

	// Uma barreira horizontal superior
	for x := alaEsquerdaXMin; x <= alaEsquerdaXMax; x += passo {
		objeto.NovaParede(cj, geometria.NovoPonto(x, barreiraSuperiorY))
	}

	// Uma barreira horizontal inferior
	for x := alaEsquerdaXMin; x <= alaEsquerdaXMax; x += passo {
		objeto.NovaParede(cj, geometria.NovoPonto(x, barreiraInferiorY))
	}

	// --- Ala Direita (Setor de Desafio) ---
	// Obstáculo em forma de 'T' proporcional
	tXMin := centroX + (tamanho * 0.15)
	tXMax := inicioX + (tamanho * 0.85)
	tY := inicioY + (tamanho * 0.35)

	for x := tXMin; x <= tXMax; x += passo {
		objeto.NovaParede(cj, geometria.NovoPonto(x, tY)) // Parte de cima do T
	}
	// Haste do T
	centroT_X := tXMin + (tXMax-tXMin)/2
	objeto.NovaParede(cj, geometria.NovoPonto(centroT_X, tY+passo))
	objeto.NovaParede(cj, geometria.NovoPonto(centroT_X, tY+(passo*2)))

	// Paredes tipo "Dentes de Pente" na lateral inferior direita
	dentesYMin := inicioY + (tamanho * 0.6)
	dentesYMax := inicioY + (tamanho * 0.85)
	raiaDireitaX := inicioX + (tamanho * 0.85)
	raiaEsquerdaX := centroX + (tamanho * 0.3)

	for y := dentesYMin; y <= dentesYMax; y += passo {
		objeto.NovaParede(cj, geometria.NovoPonto(raiaDireitaX, y))
		objeto.NovaParede(cj, geometria.NovoPonto(raiaEsquerdaX, y))
	}
}

func (s *SistemaSpawn) SpawnParedesEspecificas(cj interfaces.ICenaJogo) {
	// --- LETRA L ---
	objeto.NovaParede(cj, geometria.NovoPonto(200, 400))
	objeto.NovaParede(cj, geometria.NovoPonto(200, 430))
	objeto.NovaParede(cj, geometria.NovoPonto(200, 460))
	objeto.NovaParede(cj, geometria.NovoPonto(200, 490))
	objeto.NovaParede(cj, geometria.NovoPonto(230, 490))
	objeto.NovaParede(cj, geometria.NovoPonto(260, 490))

	// --- LETRA U ---
	objeto.NovaParede(cj, geometria.NovoPonto(320, 400))
	objeto.NovaParede(cj, geometria.NovoPonto(320, 430))
	objeto.NovaParede(cj, geometria.NovoPonto(320, 460))
	objeto.NovaParede(cj, geometria.NovoPonto(320, 490))
	objeto.NovaParede(cj, geometria.NovoPonto(350, 490))
	objeto.NovaParede(cj, geometria.NovoPonto(380, 490))
	objeto.NovaParede(cj, geometria.NovoPonto(380, 460))
	objeto.NovaParede(cj, geometria.NovoPonto(380, 430))
	objeto.NovaParede(cj, geometria.NovoPonto(380, 400))

	// --- LETRA A ---
	objeto.NovaParede(cj, geometria.NovoPonto(440, 490))
	objeto.NovaParede(cj, geometria.NovoPonto(440, 460))
	objeto.NovaParede(cj, geometria.NovoPonto(440, 430))
	objeto.NovaParede(cj, geometria.NovoPonto(470, 400))
	objeto.NovaParede(cj, geometria.NovoPonto(500, 400))
	objeto.NovaParede(cj, geometria.NovoPonto(530, 430))
	objeto.NovaParede(cj, geometria.NovoPonto(530, 460))
	objeto.NovaParede(cj, geometria.NovoPonto(530, 490))
	objeto.NovaParede(cj, geometria.NovoPonto(470, 450))
	objeto.NovaParede(cj, geometria.NovoPonto(500, 450))

	// --- LETRA N ---
	objeto.NovaParede(cj, geometria.NovoPonto(590, 490))
	objeto.NovaParede(cj, geometria.NovoPonto(590, 460))
	objeto.NovaParede(cj, geometria.NovoPonto(590, 430))
	objeto.NovaParede(cj, geometria.NovoPonto(590, 400))
	objeto.NovaParede(cj, geometria.NovoPonto(620, 430))
	objeto.NovaParede(cj, geometria.NovoPonto(650, 460))
	objeto.NovaParede(cj, geometria.NovoPonto(680, 400))
	objeto.NovaParede(cj, geometria.NovoPonto(680, 430))
	objeto.NovaParede(cj, geometria.NovoPonto(680, 460))
	objeto.NovaParede(cj, geometria.NovoPonto(680, 490))
}
