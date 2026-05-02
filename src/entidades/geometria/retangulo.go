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

func (r *Retangulo) PosXmax(largura_obj int) float64 {
	return float64(int(r.GetLargura()) - largura_obj)
}

func (r *Retangulo) PosYmax(altura_obj int) float64 {
	return float64(int(r.GetAltura()) - altura_obj)
}

func (r *Retangulo) EstaDentro(outro_x float64, outro_y float64, outro_largura float64, outro_altura float64) bool {
	x1 := outro_x
	y1 := outro_y

	if x1 >= r.GetX() && x1 < (r.GetX()+r.GetLargura()) {
		if y1 >= r.GetY() && y1 < (r.GetY()+r.GetAltura()) {
			x2 := outro_x + outro_largura
			y2 := outro_y + outro_altura

			if x2 >= r.GetX() && x2 < (r.GetX()+r.GetLargura()) {
				if y2 >= r.GetY() && y2 < (r.GetY()+r.GetAltura()) {
					return true
				}
			}
		}
	}

	return false
}
