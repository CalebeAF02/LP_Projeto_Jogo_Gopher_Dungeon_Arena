package sistema

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ISistemaDesenhar interface {
	Desenhar(g *Game, tela *ebiten.Image)
}
