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

func (r *Retangulo) GetX() float64 {
	return r.x
}
func (r *Retangulo) GetY() float64 {
	return r.y

}
func (r *Retangulo) GetLargura() float64 {
	return r.largura

}
func (r *Retangulo) GetAltura() float64 {
	return r.altura

}

func (r *Retangulo) SetX(x float64) {
	r.x = x
}
func (r *Retangulo) SetY(y float64) {
	r.y = y

}
func (r *Retangulo) SetLargura(l float64) {
	r.largura = l

}
func (r *Retangulo) SetAltura(a float64) {
	r.altura = a

}

func (r *Retangulo) PosXmax(largura_obj int) float64 {
	return float64(int(r.GetLargura()) - largura_obj)
}

func (r *Retangulo) PosYmax(altura_obj int) float64 {
	return float64(int(r.GetAltura()) - altura_obj)
}

func (r *Retangulo) EstaDentro(r2 *Retangulo) bool {
	// Extremidades do Retângulo (Mundo)
	mundoX1 := r.GetX()
	mundoY1 := r.GetY()
	mundoX2 := r.GetX() + r.GetLargura()
	mundoY2 := r.GetY() + r.GetAltura()

	// Extremidades do Objeto
	outroX1 := r2.GetX()
	outroY1 := r2.GetY()
	outroX2 := outroX1 + r2.GetLargura()
	outroY2 := outroY1 + r2.GetAltura()

	return outroX1 >= mundoX1 && outroX2 <= mundoX2 && outroY1 >= mundoY1 && outroY2 <= mundoY2
}

func (r *Retangulo) EstaDentroDireto(outroX1 float64, outroY1 float64, outroLargura float64, outroAltura float64) bool {
	// Extremidades do Retângulo (Mundo)
	mundoX1 := r.GetX()
	mundoY1 := r.GetY()
	mundoX2 := r.GetX() + r.GetLargura()
	mundoY2 := r.GetY() + r.GetAltura()

	// Extremidades do Objeto

	outroX2 := outroX1 + outroLargura
	outroY2 := outroY1 + outroAltura

	return outroX1 >= mundoX1 && outroX2 <= mundoX2 && outroY1 >= mundoY1 && outroY2 <= mundoY2
}

// MargemInterna: Encolhe o limite de colisão para dentro
func (r *Retangulo) EstaNaMargemInterna(r2 *Retangulo, margem float64) bool {
	// Extremidades do Retângulo (Mundo)
	mundoX1 := r.GetX()
	mundoY1 := r.GetY()
	mundoX2 := mundoX1 + r.GetLargura()
	mundoY2 := mundoY1 + r.GetAltura()

	// Extremidades do Objeto
	outroX1 := r2.GetX()
	outroY1 := r2.GetY()
	outroX2 := outroX1 + r2.GetLargura()
	outroY2 := outroY1 + r2.GetAltura()

	return (outroX1) >= (mundoX1+margem) && (outroX2) <= (mundoX2-margem) && (outroY1) >= (mundoY1+margem) && (outroY2) <= (mundoY2-margem)
}

// MargemExterna: Expande o limite de colisão para fora
func (r *Retangulo) EstaNaMargemExterna(r2 *Retangulo, margem float64) bool {
	// Extremidades do Retângulo (Mundo)
	mundoX1 := r.GetX()
	mundoY1 := r.GetY()
	mundoX2 := mundoX1 + r.GetLargura()
	mundoY2 := mundoY1 + r.GetAltura()

	// Extremidades do Objeto
	outroX1 := r2.GetX()
	outroY1 := r2.GetY()
	outroX2 := outroX1 + r2.GetLargura()
	outroY2 := outroY1 + r2.GetAltura()

	return outroX1 >= mundoX1-margem && outroX2 <= mundoX2+margem && outroY1 >= outroY1-margem && outroY2 <= mundoY2+margem
}

func (r *Retangulo) Colide(r2 *Retangulo) bool {
	return r.x < r2.x+r2.largura &&
		r.x+r.largura > r2.x &&
		r.y < r2.y+r2.altura &&
		r.y+r.altura > r2.y
}
