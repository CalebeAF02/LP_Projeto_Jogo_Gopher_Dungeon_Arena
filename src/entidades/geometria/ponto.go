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

func (self *Ponto) GetX() float64 {
	return self.posX
}

func (self *Ponto) GetY() float64 {
	return self.posY
}

func (self *Ponto) SetX(x float64) {
	self.posX = x
}

func (self *Ponto) SetY(y float64) {
	self.posY = y
}

func (self *Ponto) SetPosicao(x float64, y float64) {
	self.posX = x
	self.posY = y
}

func (self *Ponto) ToString() string {
	return fmt.Sprintf("X: %.0f | Y: %.0f", self.posX, self.posY)
}
