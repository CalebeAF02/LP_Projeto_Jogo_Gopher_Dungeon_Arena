package sistema

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/outros"
	"Gopher_Dungeon_Arena/src/enum/componentes"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/utils"
	"image/color"
	"math/rand"
	"os"

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

	SpawnJogadores(&g)

	SpawnParedesAoRedor(&g, 20)
	SpawnParedesEspecificas(&g)
	SpawnLabirinto(&g)

	SpawnBots(&g)

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
		ListarPrincipaisEntidades(g)
	} else if ebiten.IsKeyPressed(ebiten.KeyF2) {
		ListarEntidadesOrdenadas(g)
	} else if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.Sair()
		os.Exit(0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyB) {
		CriarBotAleatorio(g)
	}

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

func (g *Game) Sair() {
	// Limpando entidades
	g.entidades = nil

	// Limpando estruturas principais
	g.mundo = nil
	g.camera = nil
	g.miniMapa = nil
	g.aleatorio = nil

	// Zerando contadores
	g.framesGeracao = 0
	g.proximo = 0
}
