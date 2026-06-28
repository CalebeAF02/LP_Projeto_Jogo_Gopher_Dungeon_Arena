package src

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/cenas"
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/nivel"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	CenaCorrente interfaces.ICena
	CenaJogo     interfaces.ICenaJogo
}

func NovoGame() *Game {
	cenaCorrente := cenas.CenaMenuIniciar{}

	game := Game{CenaCorrente: &cenaCorrente, CenaJogo: nil}
	cenaJogo := cenas.NovoCenaJogo(&game)

	cenaCorrente.SetGame(&game)
	cenaCorrente.SetFonteCache(*assets.FonteCacheCriar())
	game.SetCenaJogo(cenaJogo)

	nivel.Converter()

	return &game
}

func (g *Game) GetCena() interfaces.ICena {
	return g.CenaCorrente
}
func (g *Game) GetCenaJogo() interfaces.ICenaJogo {
	return g.CenaJogo
}

func (g *Game) SetCenaJogo(cj interfaces.ICenaJogo) {
	g.CenaJogo = cj
}
func (g *Game) SetCena(cena interfaces.ICena) {
	g.CenaCorrente = cena
}

func (g *Game) Update() error {
	g.CenaCorrente.Update()

	//fmt.Println(g.CenaCorrente.GetNome())

	return nil
}

func (g *Game) Draw(tela *ebiten.Image) {
	g.CenaCorrente.Draw(tela)
}
func (g *Game) Layout(l, a int) (int, int) {
	return config.JANELA_LARGURA, config.JANELA_ALTURA
}

func (g *Game) IniciarJogo() {
	g.CenaJogo.ReIniciar()
	g.CenaCorrente = g.CenaJogo
}

func (g *Game) ReiniciarMudarTelaMenuIniciar() {
	cenaCorrente := cenas.CenaMenuIniciar{}
	cenaCorrente.SetFonteCache(*assets.FonteCacheCriar())
	cenaCorrente.SetGame(g)
	g.CenaJogo = cenas.NovoCenaJogo(g)
	g.SetCena(&cenaCorrente)
}

func (g *Game) MudarTelaMenuIniciar() {
	cenaCorrente := cenas.CenaMenuIniciar{}
	cenaCorrente.SetFonteCache(*assets.FonteCacheCriar())
	cenaCorrente.SetGame(g)
	g.SetCena(&cenaCorrente)
}

func (g *Game) MudarTelaProgresso() {
	cenaCorrente := cenas.CenaProgresso{}
	cenaCorrente.SetFonteCache(*assets.FonteCacheCriar())
	cenaCorrente.SetGame(g)
	g.SetCena(&cenaCorrente)
}

func (g *Game) Pausar() {
	cenaPause := cenas.CenaMenuPause{}
	cenaPause.SetGame(g)
	g.SetCena(&cenaPause)
}

func (g *Game) Voltar() {
	g.SetCena(g.GetCenaJogo())
}

func (g *Game) Sair() {
	os.Exit(0)
}

func (g *Game) GetNome() string {
	return "GAME"
}
