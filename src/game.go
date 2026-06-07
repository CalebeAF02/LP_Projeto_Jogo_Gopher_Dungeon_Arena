package src

import (
	"Gopher_Dungeon_Arena/src/cenas"
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/interfaces"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	CenaCorrente interfaces.ICena
}

func NovoGame() *Game {
	cenaCorrente := cenas.CenaMenuIniciar{}

	ng := Game{CenaCorrente: &cenaCorrente}
	return &ng
}

func (g *Game) Update() error {
	g.CenaCorrente.Update()

	return nil
}

func (g *Game) Draw(tela *ebiten.Image) {
	g.CenaCorrente.Draw()
}
func (g *Game) Layout(l, a int) (int, int) {
	return config.JANELA_LARGURA, config.JANELA_ALTURA
}
