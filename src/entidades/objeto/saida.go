package objeto

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/componentes"
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Saida struct {
	cenaJogo    interfaces.ICenaJogo
	entidadeID  ecs.EntidadeID
	Componentes map[string]interface{}
	ciclos      int
}

func NovaSaida(cenaJogo interfaces.ICenaJogo, posicao *geometria.Ponto) *Saida {

	nEntidade := cenaJogo.CriarEntidade()
	nSaida := Saida{cenaJogo: cenaJogo, entidadeID: nEntidade}

	cenaJogo.SetEntidade(nEntidade, &nSaida)

	corpo := geometria.NovoRetangulo(posicao.GetX(), posicao.GetY(), utils.SAIDA_TAMANHO_MUNDO, utils.SAIDA_TAMANHO_MUNDO)

	nSaida.AdicionarComponente(componentes.CORPO.String(), corpo)

	return &nSaida
}

func (self *Saida) GetID() ecs.EntidadeID {
	return self.entidadeID
}

func (self *Saida) ObterCorpo() *geometria.Retangulo {
	if corpocomp := self.GetComponente(componentes.CORPO.String()); corpocomp != nil {
		return corpocomp.(*geometria.Retangulo)
	}
	return nil
}

func (self *Saida) GetX() float64 {
	return self.ObterCorpo().GetX()
}

func (self *Saida) GetY() float64 {
	return self.ObterCorpo().GetY()
}

func (self *Saida) GetLargura() float64 {
	return self.ObterCorpo().GetLargura()
}

func (self *Saida) GetAltura() float64 {
	return self.ObterCorpo().GetAltura()
}

func (self *Saida) GetTipo() string {
	return "SAIDA"
}

func (self *Saida) Atualizar() {

	if self.cenaJogo.CapturouTudo() {

		self.ObterCorpo().SetLargura(60)
		self.ObterCorpo().SetAltura(60)

	}

	self.ciclos += 1

	if self.ciclos > 200 {
		self.ciclos = 0
	}

}

func (self *Saida) GetComponente(id string) interface{} {
	return self.Componentes[id]
}

func (self *Saida) AdicionarComponente(id string, comp interface{}) {
	if self.Componentes == nil {
		self.Componentes = make(map[string]interface{})
	}
	self.Componentes[id] = comp
}

func (self *Saida) Desenhar(tela *ebiten.Image) {

	var raio float64 = utils.SAIDA_TAMANHO_MUNDO / 2
	var raioGrande float64 = (raio)
	var raioPequeno float64 = (raio - 5)
	var raioInterno float64 = (raio - 8)
	var raioMicro float64 = (raio - 10)
	var raioMicro2 float64 = (raio - 25)

	cor := cores.CINZA_ESCURO

	if self.cenaJogo.CapturouTudo() {
		cor = cores.ROSA
	}

	ebitenutil.DrawCircle(tela, self.cenaJogo.GetCamera().GetX()+self.GetX()+raio, self.cenaJogo.GetCamera().GetY()+self.GetY()+raio, raioGrande, cor)
	ebitenutil.DrawCircle(tela, self.cenaJogo.GetCamera().GetX()+self.GetX()+raio, self.cenaJogo.GetCamera().GetY()+self.GetY()+raio, raioPequeno, cores.BRANCO)
	ebitenutil.DrawCircle(tela, self.cenaJogo.GetCamera().GetX()+self.GetX()+raio, self.cenaJogo.GetCamera().GetY()+self.GetY()+raio, raioInterno, cor)
	ebitenutil.DrawCircle(tela, self.cenaJogo.GetCamera().GetX()+self.GetX()+raio, self.cenaJogo.GetCamera().GetY()+self.GetY()+raio, raioMicro, cores.BRANCO)

	ebitenutil.DrawCircle(tela, self.cenaJogo.GetCamera().GetX()+self.GetX()+raio, self.cenaJogo.GetCamera().GetY()+self.GetY()+raio, raioMicro2, cor)

	//ebitenutil.DrawRect(tela, self.cenaJogo.GetCamera().GetX()+self.ObterCorpo().GetX(), self.cenaJogo.GetCamera().GetY()+self.ObterCorpo().GetY(), self.ObterCorpo().GetLargura(), self.ObterCorpo().GetAltura(), cores.VERDE)

	if self.ciclos > 100 {

		ebitenutil.DrawCircle(tela, (self.cenaJogo.GetCamera().GetX()+self.GetX())+(raio)-32, (self.cenaJogo.GetCamera().GetY()+self.GetY())+(raio), 5, cores.ROSA_ESCURO)
		ebitenutil.DrawCircle(tela, (self.cenaJogo.GetCamera().GetX()+self.GetX())+(raio)+32, (self.cenaJogo.GetCamera().GetY()+self.GetY())+(raio), 5, cores.ROSA_ESCURO)

		ebitenutil.DrawCircle(tela, (self.cenaJogo.GetCamera().GetX()+self.GetX())+(raio), (self.cenaJogo.GetCamera().GetY()+self.GetY())+(raio)-32, 5, cores.ROSA_ESCURO)
		ebitenutil.DrawCircle(tela, (self.cenaJogo.GetCamera().GetX()+self.GetX())+(raio), (self.cenaJogo.GetCamera().GetY()+self.GetY())+(raio)+32, 5, cores.ROSA_ESCURO)

	}

	if self.cenaJogo.ObterPontuacaoFaltante() > 0 {
		assets.EscreverNumeroLocal(tela, self.cenaJogo.ObterFonteCache().Normal, self.cenaJogo.GetCamera().GetX()+self.GetX()+(raio)-10, self.cenaJogo.GetCamera().GetY()+self.GetY()+(raio)-15, self.cenaJogo.ObterPontuacaoFaltante())
	}

}

func (self *Saida) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawCircle(tela, mapaX+(self.GetX()/config.PROPORCAO_MAPA), mapaY+(self.GetY()/config.PROPORCAO_MAPA), utils.SAIDA_TAMANHO_MAPA, cores.ROSA)
}

func (self *Saida) ExisteComponente(id string) bool {
	_, existe := self.Componentes[id]
	return existe
}
