package interfaces

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ISistemaDesenhar interface {
	Desenhar(cj ICenaJogo, tela *ebiten.Image)
}
