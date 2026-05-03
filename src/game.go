package src

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/movimentacao"
	"Gopher_Dungeon_Arena/src/entidades/outros"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/enum/componentes"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"fmt"
	"image/color"
	"math/rand"
	"reflect"
	"sort"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	proximo       int
	mundo         *geometria.Retangulo
	entidades     map[ecs.EntidadeID]ecs.Entidade
	aleatorio     *rand.Rand
	miniMapa      *ecs.MiniMapa
	camera        *ecs.Camera
	framesGeracao int
}

func NovoGame() *Game {
	mundo := geometria.NovoRetangulo(0, 0, config.MUNDO_LARGURA, config.MUNDO_ALTURA)
	entidades := make(map[ecs.EntidadeID]ecs.Entidade)
	aleatorio := config.GeradorAleatorio()
	camera := ecs.NovaCamera(mundo)
	miniMapa := ecs.NovoMiniMapa(mundo, geometria.NovoPonto(10, 10), camera)

	g := Game{mundo: mundo, entidades: entidades, aleatorio: aleatorio, framesGeracao: 0}

	g.SetMiniMapa(miniMapa)
	g.SetCamera(camera)

	// Jogadores
	j1 := personagens.NovoJogador(&g, "Jogador 1")
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
	t1 := outros.NovoTime(&g, "Vermelhao - Time_Vermelho", cores.AZUL)
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

	// --- LETRA L ---
	outros.NovaParede(&g, geometria.NovoPonto(200, 400))
	outros.NovaParede(&g, geometria.NovoPonto(200, 430))
	outros.NovaParede(&g, geometria.NovoPonto(200, 460))
	outros.NovaParede(&g, geometria.NovoPonto(200, 490))
	outros.NovaParede(&g, geometria.NovoPonto(230, 490))
	outros.NovaParede(&g, geometria.NovoPonto(260, 490))

	// --- LETRA U ---
	outros.NovaParede(&g, geometria.NovoPonto(320, 400))
	outros.NovaParede(&g, geometria.NovoPonto(320, 430))
	outros.NovaParede(&g, geometria.NovoPonto(320, 460))
	outros.NovaParede(&g, geometria.NovoPonto(320, 490))
	outros.NovaParede(&g, geometria.NovoPonto(350, 490))
	outros.NovaParede(&g, geometria.NovoPonto(380, 490))
	outros.NovaParede(&g, geometria.NovoPonto(380, 460))
	outros.NovaParede(&g, geometria.NovoPonto(380, 430))
	outros.NovaParede(&g, geometria.NovoPonto(380, 400))

	// --- LETRA A ---
	outros.NovaParede(&g, geometria.NovoPonto(440, 490))
	outros.NovaParede(&g, geometria.NovoPonto(440, 460))
	outros.NovaParede(&g, geometria.NovoPonto(440, 430))
	outros.NovaParede(&g, geometria.NovoPonto(470, 400))
	outros.NovaParede(&g, geometria.NovoPonto(500, 400))
	outros.NovaParede(&g, geometria.NovoPonto(530, 430))
	outros.NovaParede(&g, geometria.NovoPonto(530, 460))
	outros.NovaParede(&g, geometria.NovoPonto(530, 490))
	outros.NovaParede(&g, geometria.NovoPonto(470, 450))
	outros.NovaParede(&g, geometria.NovoPonto(500, 450))

	// --- LETRA N ---
	outros.NovaParede(&g, geometria.NovoPonto(590, 490))
	outros.NovaParede(&g, geometria.NovoPonto(590, 460))
	outros.NovaParede(&g, geometria.NovoPonto(590, 430))
	outros.NovaParede(&g, geometria.NovoPonto(590, 400))
	outros.NovaParede(&g, geometria.NovoPonto(620, 430))
	outros.NovaParede(&g, geometria.NovoPonto(650, 460))
	outros.NovaParede(&g, geometria.NovoPonto(680, 400))
	outros.NovaParede(&g, geometria.NovoPonto(680, 430))
	outros.NovaParede(&g, geometria.NovoPonto(680, 460))
	outros.NovaParede(&g, geometria.NovoPonto(680, 490))

	g.labirinto()

	ConstruirParedeAoRedor(&g, mundo, 20)

	//Bot
	for id := 0; id < 1; id++ {
		g.CriarBot(&movimentacao.MovimentadorSimples{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorVertical{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorVerticalConstante{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorHorizontal{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorHorizontalConstante{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorDiagonal{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorLogicoLinha{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorLogicoDiagonal{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorLogicoDuplo{}, g.OrganizaPosicaoAleatoriaBot())
	}

	return &g
}

func (g *Game) Layout(l, a int) (int, int) {
	return config.JANELA_LARGURA, config.JANELA_ALTURA
}

func (g *Game) CriarEntidade() ecs.EntidadeID {
	entidade := ecs.EntidadeID(g.proximo)
	g.proximo++
	return entidade
}

func (g *Game) CriarBot(movendo interfaces.Movimentador, posicao *geometria.Ponto) {
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

func (g *Game) CriarBotAleatorio() {

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
func (g *Game) OrganizaPosicaoAleatoriaBot() *geometria.Ponto {
	larguraBot := float64(utils.BOT_TAMANHO_MUNDO)
	alturaBot := float64(utils.BOT_TAMANHO_MUNDO)

	// Tentamos encontrar uma posição válida (limitamos as tentativas para evitar loop infinito)
	for tentativas := 0; tentativas < 100; tentativas++ {
		x := float64(g.aleatorio.Intn(int(g.mundo.PosXmax(utils.BOT_TAMANHO_MUNDO))))
		y := float64(g.aleatorio.Intn(int(g.mundo.PosYmax(utils.BOT_TAMANHO_MUNDO))))

		if g.PosicaoEstaLivre(x, y, larguraBot, alturaBot) {
			return geometria.NovoPonto(x, y)
		}
	}

	return nil
}

func (g *Game) PosicaoEstaLivre(x, y float64, largura, altura float64) bool {
	// Cria um retângulo temporário para representar o corpo do bot na posição sorteada
	corpoBot := geometria.NovoRetangulo(x, y, largura, altura)

	for _, e := range g.GetEntidades() {
		if e.GetTipo() == entidades.PAREDE.String() {
			if corpoParede := e.GetComponente(componentes.CORPO.String()); corpoParede != nil {
				// Se colidir com qualquer parede, a posição não está livre
				if corpoBot.Colide(corpoParede.(*geometria.Retangulo)) {
					return false
				}
			}
		}
	}
	return true // Não colidiu com nenhuma parede
}

func (g *Game) GetEntidades() map[ecs.EntidadeID]ecs.Entidade {
	return g.entidades
}
func (g *Game) GetTimes() []*outros.Time {
	listaTimes := []*outros.Time{}

	for _, e := range g.GetEntidades() {
		if e.GetTipo() == "TIME" {
			listaTimes = append(listaTimes, e.(*outros.Time))
		}
	}
	//fmt.Printf("Quantidade de Times %d\n", len(listaTimes))

	return listaTimes
}
func (g *Game) GetMundo() *geometria.Retangulo {
	return g.mundo
}
func (g *Game) GetAleatorio() *rand.Rand {
	return g.aleatorio
}
func (g *Game) GetLargura() float64 {
	return g.mundo.GetLargura()
}
func (g *Game) GetAltura() float64 {
	return g.mundo.GetAltura()
}
func (g *Game) GetCamera() *ecs.Camera {
	return g.camera
}

func (g *Game) SetEntidade(nEntidade ecs.EntidadeID, posicao ecs.Entidade) {
	g.entidades[nEntidade] = posicao
}
func (g *Game) SetMiniMapa(miniMapa *ecs.MiniMapa) {
	g.miniMapa = miniMapa
}
func (g *Game) SetCamera(camera *ecs.Camera) {
	g.camera = camera
}

func (g *Game) ColideComBarreiras(eu *geometria.Retangulo) bool {

	for _, e := range g.GetEntidades() {
		if e.GetTipo() == entidades.PAREDE.String() {
			if corpoParede := e.GetComponente(componentes.CORPO.String()); corpoParede != nil {
				if eu.Colide(corpoParede.(*geometria.Retangulo)) {
					return true
				}
			}
		}
	}

	return false
}

func (g *Game) Update() error {
	// Atalho para debugar as entidades no terminal
	if ebiten.IsKeyPressed(ebiten.KeyF1) {
		g.ListarPrincipaisEntidades()
	} else if ebiten.IsKeyPressed(ebiten.KeyF2) {
		g.ListarEntidadesOrdenadas()
	}

	// --- LÓGICA DE TEMPO PARA BOTS ---
	g.framesGeracao++

	// 180 frames = 3 segundos (em 60 FPS)
	if g.framesGeracao >= 180 {
		g.framesGeracao = 0

		// Sorteia uma posição válida (longe de paredes)
		pos := g.OrganizaPosicaoAleatoriaBot()

		// Gera o bot com um movimentador aleatório
		g.GerarBot(pos.GetX(), pos.GetY())
	}

	// --- ATUALIZAÇÃO DA CÂMERA ---
	lTimes := g.GetTimes()
	if len(lTimes) > 0 && len(lTimes[0].GetJogadores()) > 0 {
		jogador := lTimes[0].GetJogador(0)
		g.camera.OrganizarCameraPeloJogador(jogador.GetPosicao())
	}

	// --- ATUALIZAÇÃO DAS ENTIDADES ---
	for _, entidade := range g.entidades {
		entidade.Atualizar()
	}

	return nil
}

func (g *Game) Draw(tela *ebiten.Image) {
	tela.Fill(color.RGBA{20, 20, 20, 255})

	margemMundo := geometria.NovoRetangulo(g.GetCamera().GetX()+g.mundo.GetX(), g.GetCamera().GetY()+g.mundo.GetY(), g.mundo.GetLargura(), g.mundo.GetAltura())
	utils.MargemInterna(tela, margemMundo, utils.JOGADOR_TAMANHO_MUNDO, cores.BRANCO)

	for _, entidade := range g.entidades {
		entidade.Desenhar(tela)
	}

	if config.PROPORCAO_MUNDO > 1 {
		g.miniMapa.Desenhar(tela)

		for _, entidade := range g.entidades {
			entidade.DesenharMapa(tela, g.miniMapa.GetX(), g.miniMapa.GetY())
		}
	}
}

func (g *Game) labirinto() {

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

func ConstruirParedeAoRedor(g *Game, mundo *geometria.Retangulo, passo float64) {
	xMin := mundo.GetX()
	yMin := mundo.GetY()
	xMax := mundo.GetX() + mundo.GetLargura()
	yMax := mundo.GetY() + mundo.GetAltura()

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

func (g *Game) GerarBot(x float64, y float64) {

	tipo := g.aleatorio.Intn(10)
	if tipo == 0 {
		g.CriarBot(&movimentacao.MovimentadorHorizontal{}, geometria.NovoPonto(x, y))
	} else if tipo == 1 {
		g.CriarBot(&movimentacao.MovimentadorHorizontalConstante{}, geometria.NovoPonto(x, y))
	} else if tipo == 2 {
		g.CriarBot(&movimentacao.MovimentadorVertical{}, geometria.NovoPonto(x, y))
	} else if tipo == 3 {
		g.CriarBot(&movimentacao.MovimentadorVerticalConstante{}, geometria.NovoPonto(x, y))
	} else if tipo == 4 {
		g.CriarBot(&movimentacao.MovimentadorDiagonal{}, geometria.NovoPonto(x, y))
	} else if tipo == 5 {
		g.CriarBot(&movimentacao.MovimentadorLogicoDiagonal{}, geometria.NovoPonto(x, y))
	} else {
		g.CriarBot(&movimentacao.MovimentadorDiagonal{}, geometria.NovoPonto(x, y))
	}

}

func (g *Game) ListarEntidadesOrdenadas() {
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("   RELATÓRIO DE ENTIDADES (ORDENADO POR ID)")
	fmt.Println(strings.Repeat("=", 40))

	entidades := g.GetEntidades()
	if len(entidades) == 0 {
		fmt.Println("O mundo está vazio.")
		return
	}

	// 1. Extrair todas as chaves (IDs) do mapa
	ids := make([]int, 0, len(entidades))
	for id := range entidades {
		ids = append(ids, int(id))
	}

	// 2. Ordenar os IDs numericamente
	sort.Ints(ids)

	// 3. Percorrer os IDs já ordenados
	for _, idInt := range ids {
		id := ecs.EntidadeID(idInt)
		entidade := entidades[id]

		v := reflect.ValueOf(entidade)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		t := v.Type()

		fmt.Printf("\n>>> ID: %d | CLASSE: %s\n", id, t.Name())

		// Listar Atributos
		for i := 0; i < v.NumField(); i++ {
			campoNome := t.Field(i).Name
			campoValor := v.Field(i)

			// Pula campos internos ou privados (que começam com letra minúscula) se desejar
			if t.Field(i).PkgPath != "" {
				continue
			}

			fmt.Printf("    %-18s : %v\n", campoNome, campoValor)
		}
	}
	fmt.Println("\n" + strings.Repeat("=", 40))
}

func (g *Game) ListarPrincipaisEntidades() {
	entidades := g.GetEntidades()
	if len(entidades) == 0 {
		fmt.Println("O mundo está vazio.")
		return
	}

	totalParedes := 0
	var idsOutros []int

	for id, e := range entidades {
		// Importante: verifique se seu GetTipo() retorna exatamente "PAREDE"
		// ou use a constante entidades.PAREDE.String()
		if e.GetTipo() == "PAREDE" {
			totalParedes++
		} else {
			idsOutros = append(idsOutros, int(id))
		}
	}

	sort.Ints(idsOutros)

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("   RELATÓRIO COMPLETO DE ATRIBUTOS (Paredes: %d)\n", totalParedes)
	fmt.Println(strings.Repeat("=", 60))

	for _, idInt := range idsOutros {
		id := ecs.EntidadeID(idInt)
		entidade := entidades[id]

		v := reflect.ValueOf(entidade)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		t := v.Type()

		fmt.Printf("\n[ID: %d] CLASSE: %s\n", id, strings.ToUpper(t.Name()))
		fmt.Println(strings.Repeat("-", 30))

		for i := 0; i < v.NumField(); i++ {
			campoNome := t.Field(i).Name
			campoValor := v.Field(i)

			// Esta parte garante que apresente até valores complexos (como Ponto ou Retângulo)
			// de forma legível no terminal
			valorFormatado := fmt.Sprintf("%v", campoValor)

			// Se o valor for uma Interface ou Ponteiro interno, tentamos extrair o conteúdo real
			if campoValor.Kind() == reflect.Ptr && !campoValor.IsNil() {
				valorFormatado = fmt.Sprintf("%v", campoValor.Elem())
			}

			fmt.Printf("   %-20s : %s\n", campoNome, valorFormatado)
		}
	}

	fmt.Println("\n" + strings.Repeat("-", 60))
	fmt.Printf(" -> %d PAREDES FORAM AGRUPADAS PARA LIMPEZA VISUAL\n", totalParedes)
	fmt.Println(strings.Repeat("=", 60))
}
