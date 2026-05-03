package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EntidadeID int

type Entidade interface {
	GetTipo() string
	GetComponente(id string) interface{}
	Atualizar()
	Desenhar(screen *ebiten.Image)
	DesenharMapa(screen *ebiten.Image, mapaX float64, mapaY float64)
}
