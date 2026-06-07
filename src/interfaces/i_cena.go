package interfaces

import "github.com/hajimehoshi/ebiten/v2"

type ICena interface {
	GetNome() string
	GetGame() IGame
	SetGame(game IGame)
	Update() error
	Draw(tela *ebiten.Image)
}
