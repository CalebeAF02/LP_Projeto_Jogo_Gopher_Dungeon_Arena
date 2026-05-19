package sistema

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/outros"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	proximo          int
	mundo            *geometria.Retangulo
	entidades        map[ecs.EntidadeID]ecs.Entidade
	camera           *ecs.Camera
	miniMapa         *ecs.MiniMapa
	aleatorio        *rand.Rand
	sistemaAtualizar []ISistemaAtualizar
	sistemaDesenhar  []ISistemaDesenhar
}

func NovoGame() *Game {
	mundo := geometria.NovoRetangulo(0, 0, config.MUNDO_LARGURA, config.MUNDO_ALTURA)
	entidades := make(map[ecs.EntidadeID]ecs.Entidade)
	camera := ecs.NovaCamera(mundo)
	miniMapa := ecs.NovoMiniMapa(mundo, geometria.NovoPonto(10, 10), camera)
	aleatorio := config.GeradorAleatorio()

	g := Game{mundo: mundo, entidades: entidades, aleatorio: aleatorio}

	g.sistemaAtualizar = []ISistemaAtualizar{
		&SistemaInput{},
		&SistemaIA{},
		&SistemaSpawn{},
		&SistemaMovimento{},
		&SistemaEntidades{},
		&SistemaCamera{},
		&SistemaDebug{},
	}

	g.sistemaDesenhar = []ISistemaDesenhar{
		&SistemaDesenhar{},
	}

	g.SetMiniMapa(miniMapa)
	g.SetCamera(camera)

	SpawnJogadores(&g)

	SpawnParedesAoRedor(&g, 20)
	//SpawnParedesEspecificas(&g)
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

func (g *Game) GetMiniMapa() *ecs.MiniMapa {
	return g.miniMapa
}
func (g *Game) GetSistemaAtualizar() []ISistemaAtualizar {
	return g.sistemaAtualizar
}
func (g *Game) GetSistemaDesenhar() []ISistemaDesenhar {
	return g.sistemaDesenhar
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

func (g *Game) Update() error {
	for _, sistema := range g.sistemaAtualizar {
		sistema.Atualizar(g)
	}
	return nil
}

func (g *Game) Draw(tela *ebiten.Image) {
	for _, sistema := range g.sistemaDesenhar {
		sistema.Desenhar(g, tela)
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
	g.proximo = 0
}
