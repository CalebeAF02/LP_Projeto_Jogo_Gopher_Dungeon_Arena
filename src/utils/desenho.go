package utils

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func MargemInterna(tela *ebiten.Image, desenho *geometria.Retangulo, margem float64, cor color.Color) {
	MargemLinha(tela, desenho, margem*(-1), cor)
	MargemQuina(tela, desenho, margem*(-1), cor)
}

func MargemExterna(tela *ebiten.Image, desenho *geometria.Retangulo, margem float64, cor color.Color) {
	MargemLinha(tela, desenho, margem, cor)
	MargemQuina(tela, desenho, margem, cor)
}

func MargemLinha(tela *ebiten.Image, desenho *geometria.Retangulo, margem float64, cor color.Color) {
	// linha da esquerda
	ebitenutil.DrawRect(tela, desenho.GetX()-margem, desenho.GetY(), margem, desenho.GetAltura(), cor)

	// linha da direita
	ebitenutil.DrawRect(tela, desenho.GetX()+desenho.GetLargura(), desenho.GetY(), margem, desenho.GetAltura(), cor)

	// linha de cima
	ebitenutil.DrawRect(tela, desenho.GetX(), desenho.GetY()-margem, desenho.GetLargura(), margem, cor)

	// linha de baixo
	ebitenutil.DrawRect(tela, desenho.GetX(), desenho.GetY()+desenho.GetAltura(), desenho.GetLargura(), margem, cor)
}

func MargemQuina(tela *ebiten.Image, desenho *geometria.Retangulo, margem float64, cor color.Color) {
	// linha da esquerda de cima
	ebitenutil.DrawRect(tela, desenho.GetX()-margem, desenho.GetY()-margem, margem, margem, cor)

	// linha da direita de cima
	ebitenutil.DrawRect(tela, desenho.GetX()+desenho.GetLargura(), desenho.GetY()-margem, margem, margem, cor)

	// linha da esquerda de baixo
	ebitenutil.DrawRect(tela, desenho.GetX()-margem, desenho.GetY()+desenho.GetAltura(), margem, margem, cor)

	// linhada direita de baixo
	ebitenutil.DrawRect(tela, desenho.GetX()+desenho.GetLargura(), desenho.GetY()+desenho.GetAltura(), margem, margem, cor)
}
