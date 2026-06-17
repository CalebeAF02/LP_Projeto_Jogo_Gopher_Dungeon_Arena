package sistema

import (
	"Gopher_Dungeon_Arena/src/componentes/movimentacao"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
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

	j1.SetPosicao(300, 300)

	j1.SetNivel(1)

	// Times
	t1 := outros.NovoTime(cj, "Vermelhao - Time_Azul", cores.AZUL)

	// Gerenciando
	t1.Adicionnar(j1)

	j1.CarregarPontuacao()

}

func (s *SistemaSpawn) SpawnarBot(cj interfaces.ICenaJogo, movendo interfaces.Movimentador, posicao *geometria.Ponto) {
	b := personagens.NovoBot(cj, 0)
	b.SetNivelAleatorio()
	b.SetPosicao(posicao.GetX(), posicao.GetY())
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

func (s *SistemaSpawn) SpawnarPortais(cj interfaces.ICenaJogo) {
	// Par 1
	bEntrada1 := objeto.NovoPortalEntrada(cj, 0)
	bEntrada1.SetPosicao(100, 100)

	bSaida1 := objeto.NovoPortalSaida(cj, 0)
	bSaida1.SetPosicao(1255, 100)

	bEntrada1.ConectarSaida(bSaida1)

	// Par 2
	bEntrada2 := objeto.NovoPortalEntrada(cj, 0)
	bEntrada2.SetPosicao(2410, 100)

	bSaida2 := objeto.NovoPortalSaida(cj, 0)
	bSaida2.SetPosicao(1255, 1290)

	bEntrada2.ConectarSaida(bSaida2)

	// Par 3
	bEntrada3 := objeto.NovoPortalEntrada(cj, 0)
	bEntrada3.SetPosicao(100, 1290)

	bSaida3 := objeto.NovoPortalSaida(cj, 0)
	bSaida3.SetPosicao(100, 695)

	bEntrada3.ConectarSaida(bSaida3)

	// Par 4
	bEntrada4 := objeto.NovoPortalEntrada(cj, 0)
	bEntrada4.SetPosicao(2410, 1290)

	bSaida4 := objeto.NovoPortalSaida(cj, 0)
	bSaida4.SetPosicao(2410, 695)

	bEntrada4.ConectarSaida(bSaida4)

	// Par 5
	bEntrada5 := objeto.NovoPortalEntrada(cj, 0)
	bEntrada5.SetPosicao(1255, 695)

	bSaida5 := objeto.NovoPortalSaida(cj, 0)
	bSaida5.SetPosicao(1800, 930)

	bEntrada5.ConectarSaida(bSaida5)
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

	s.spawnContorno(cj)

	s.spawnSetorNorte(cj)

	s.spawnSetorSul(cj)

	s.spawnSetorLeste(cj)

	s.spawnSetorOeste(cj)

	s.spawnCorredores(cj)
}

func (s *SistemaSpawn) criarParede(
	cj interfaces.ICenaJogo,
	x float64,
	y float64,
) {
	objeto.NovaParede(
		cj,
		geometria.NovoPonto(x, y),
	)
}

func (s *SistemaSpawn) spawnSetorNorte(cj interfaces.ICenaJogo) {

	passo := 30.0

	for x := 300.0; x <= 2200; x += passo {

		if x > 900 && x < 1200 {
			continue
		}

		s.criarParede(cj, x, 250)
	}

	for x := 500.0; x <= 1800; x += passo {

		if x > 1300 && x < 1600 {
			continue
		}

		s.criarParede(cj, x, 400)
	}
}

func (s *SistemaSpawn) spawnSetorSul(cj interfaces.ICenaJogo) {

	passo := 30.0

	for x := 400.0; x <= 2100; x += passo {

		if x > 1000 && x < 1300 {
			continue
		}

		s.criarParede(cj, x, 1050)
	}

	for x := 600.0; x <= 2000; x += passo {

		if x > 1600 && x < 1800 {
			continue
		}

		s.criarParede(cj, x, 1200)
	}
}

func (s *SistemaSpawn) spawnSetorOeste(cj interfaces.ICenaJogo) {

	passo := 30.0

	for y := 350.0; y <= 1000; y += passo {

		if y > 600 && y < 750 {
			continue
		}

		s.criarParede(cj, 500, y)
	}

	for y := 450.0; y <= 1100; y += passo {

		if y > 850 && y < 950 {
			continue
		}

		s.criarParede(cj, 700, y)
	}
}

func (s *SistemaSpawn) spawnSetorLeste(cj interfaces.ICenaJogo) {

	passo := 30.0

	for y := 300.0; y <= 1100; y += passo {

		if y > 500 && y < 700 {
			continue
		}

		s.criarParede(cj, 1900, y)
	}

	for y := 400.0; y <= 1000; y += passo {

		if y > 800 && y < 950 {
			continue
		}

		s.criarParede(cj, 2150, y)
	}
}

func (s *SistemaSpawn) spawnCorredores(cj interfaces.ICenaJogo) {

	passo := 30.0

	for x := 850.0; x <= 1700; x += passo {

		if x > 1180 && x < 1380 {
			continue
		}

		s.criarParede(cj, x, 550)
	}

	for x := 850.0; x <= 1700; x += passo {

		if x > 1180 && x < 1380 {
			continue
		}

		s.criarParede(cj, x, 900)
	}
}

func (s *SistemaSpawn) spawnPracaCentral(
	cj interfaces.ICenaJogo,
) {
	// propositalmente vazia
}

func (s *SistemaSpawn) spawnContorno(cj interfaces.ICenaJogo) {

	passo := 30.0

	largura := cj.GetMundo().GetLargura()
	altura := cj.GetMundo().GetAltura()

	for x := 0.0; x <= largura; x += passo {

		s.criarParede(cj, x, 0)

		s.criarParede(cj, x, altura)
	}

	for y := 0.0; y <= altura; y += passo {

		s.criarParede(cj, 0, y)

		s.criarParede(cj, largura, y)
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
