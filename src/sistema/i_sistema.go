package sistema

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ISistemaAtualizar interface {
	Atualizar(g *Game)
}

type ISistemaDesenhar interface {
	Desenhar(g *Game, tela *ebiten.Image)
}
