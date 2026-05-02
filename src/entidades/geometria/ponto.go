package geometria

import (
	"fmt"
)

type Ponto struct {
	posX, posY float64
}

func NovoPonto(x float64, y float64) *Ponto {
	return &Ponto{posX: x, posY: y}
}

func (p *Ponto) GetX() float64 {
	return p.posX
}

func (p *Ponto) GetY() float64 {
	return p.posY
}

func (p *Ponto) SetX(x float64) {
	p.posX = x
}

func (p *Ponto) SetY(y float64) {
	p.posY = y
}

func (p *Ponto) SetPosicao(x float64, y float64) {
	p.posX = x
	p.posY = y
}

func (p *Ponto) ToString() string {
	return fmt.Sprintf("X: %.0f | Y: %.0f", p.posX, p.posY)
}
