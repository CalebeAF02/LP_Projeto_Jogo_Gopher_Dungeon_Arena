package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EntidadeID int

type Entidade interface {
	GetID() EntidadeID
	GetTipo() string
	GetComponente(id string) interface{}
	ExisteComponente(id string) bool
	Atualizar()
	Desenhar(screen *ebiten.Image)
	DesenharMapa(screen *ebiten.Image, mapaX float64, mapaY float64)
}
