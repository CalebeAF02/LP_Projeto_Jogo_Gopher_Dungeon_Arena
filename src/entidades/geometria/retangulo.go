package geometria

type Retangulo struct {
	x       float64
	y       float64
	largura float64
	altura  float64
}

func NovoRetangulo(x float64, y float64, largura float64, altura float64) *Retangulo {
	return &Retangulo{x: x, y: y, largura: largura, altura: altura}
}

func (self *Retangulo) GetX() float64 {
	return self.x
}
func (self *Retangulo) GetY() float64 {
	return self.y

}
func (self *Retangulo) GetLargura() float64 {
	return self.largura

}
func (self *Retangulo) GetAltura() float64 {
	return self.altura

}

func (self *Retangulo) SetX(x float64) {
	self.x = x
}
func (self *Retangulo) SetY(y float64) {
	self.y = y

}

func (self *Retangulo) SetPosicao(x float64, y float64) {
	self.x = x
	self.y = y
}

func (self *Retangulo) SetLargura(l float64) {
	self.largura = l

}
func (self *Retangulo) SetAltura(a float64) {
	self.altura = a

}

func (self *Retangulo) PosXmax(largura_obj int) float64 {
	return float64(int(self.GetLargura()) - largura_obj)
}

func (self *Retangulo) PosYmax(altura_obj int) float64 {
	return float64(int(self.GetAltura()) - altura_obj)
}

func (self *Retangulo) EstaDentro(r2 *Retangulo) bool {
	// Extremidades do Retângulo (Mundo)
	mundoX1 := self.GetX()
	mundoY1 := self.GetY()
	mundoX2 := self.GetX() + self.GetLargura()
	mundoY2 := self.GetY() + self.GetAltura()

	// Extremidades do Objeto
	outroX1 := r2.GetX()
	outroY1 := r2.GetY()
	outroX2 := outroX1 + r2.GetLargura()
	outroY2 := outroY1 + r2.GetAltura()

	return outroX1 >= mundoX1 && outroX2 <= mundoX2 && outroY1 >= mundoY1 && outroY2 <= mundoY2
}

func (self *Retangulo) EstaDentroDireto(outroX1 float64, outroY1 float64, outroLargura float64, outroAltura float64) bool {
	// Extremidades do Retângulo (Mundo)
	mundoX1 := self.GetX()
	mundoY1 := self.GetY()
	mundoX2 := self.GetX() + self.GetLargura()
	mundoY2 := self.GetY() + self.GetAltura()

	// Extremidades do Objeto

	outroX2 := outroX1 + outroLargura
	outroY2 := outroY1 + outroAltura

	return outroX1 >= mundoX1 && outroX2 <= mundoX2 && outroY1 >= mundoY1 && outroY2 <= mundoY2
}

// MargemInterna: Encolhe o limite de colisão para dentro
func (self *Retangulo) EstaNaMargemInterna(r2 *Retangulo, margem float64) bool {
	// Extremidades do Retângulo (Mundo)
	mundoX1 := self.GetX()
	mundoY1 := self.GetY()
	mundoX2 := mundoX1 + self.GetLargura()
	mundoY2 := mundoY1 + self.GetAltura()

	// Extremidades do Objeto
	outroX1 := r2.GetX()
	outroY1 := r2.GetY()
	outroX2 := outroX1 + r2.GetLargura()
	outroY2 := outroY1 + r2.GetAltura()

	return (outroX1) >= (mundoX1+margem) && (outroX2) <= (mundoX2-margem) && (outroY1) >= (mundoY1+margem) && (outroY2) <= (mundoY2-margem)
}

// MargemExterna: Expande o limite de colisão para fora
func (self *Retangulo) EstaNaMargemExterna(r2 *Retangulo, margem float64) bool {
	// Extremidades do Retângulo (Mundo)
	mundoX1 := self.GetX()
	mundoY1 := self.GetY()
	mundoX2 := mundoX1 + self.GetLargura()
	mundoY2 := mundoY1 + self.GetAltura()

	// Extremidades do Objeto
	outroX1 := r2.GetX()
	outroY1 := r2.GetY()
	outroX2 := outroX1 + r2.GetLargura()
	outroY2 := outroY1 + r2.GetAltura()

	return outroX1 >= mundoX1-margem && outroX2 <= mundoX2+margem && outroY1 >= outroY1-margem && outroY2 <= mundoY2+margem
}

func (self *Retangulo) Colide(r2 *Retangulo) bool {
	// Posições do retangulo 1
	posX1 := self.x
	posY1 := self.y
	posX2 := posX1 + self.largura
	posY2 := posY1 + self.altura

	// Posições do retangulo 2
	outroX1 := r2.x
	outroY1 := r2.y
	outroX2 := outroX1 + r2.largura
	outroY2 := outroY1 + r2.altura

	return posX1 < outroX2 && posX2 > outroX1 && posY1 < outroY2 && posY2 > outroY1
}
