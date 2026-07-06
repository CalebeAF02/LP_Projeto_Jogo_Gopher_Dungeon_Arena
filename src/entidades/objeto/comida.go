package objeto

import (
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

type Comida struct {
	cenaJogo    interfaces.ICenaJogo
	entidadeID  ecs.EntidadeID
	estrutura   *geometria.Retangulo
	Componentes map[string]interface{}
	ciclos      int
}

func NovaComida(cenaJogo interfaces.ICenaJogo, posicao *geometria.Ponto) *Comida {

	nEntidade := cenaJogo.CriarEntidade()
	nComida := Comida{cenaJogo: cenaJogo, entidadeID: nEntidade, ciclos: 0, estrutura: geometria.NovoRetangulo(posicao.GetX(), posicao.GetY(), utils.COMIDA_TAMANHO_MUNDO, utils.COMIDA_TAMANHO_MUNDO)}

	cenaJogo.SetEntidade(nEntidade, &nComida)

	nComida.AdicionarComponente(componentes.CORPO.String(), nComida.estrutura)
	nComida.AdicionarComponente(componentes.ENERGIA.String(), &componentes.Energia{Valor: 100, Status: true})

	return &nComida
}

func (self *Comida) ObterEnergia() *componentes.Energia {
	if energia_comp := self.GetComponente(componentes.ENERGIA.String()); energia_comp != nil {
		return energia_comp.(*componentes.Energia)
	}
	return nil
}

func (self *Comida) GetID() ecs.EntidadeID {
	return self.entidadeID
}

func (self *Comida) GetX() float64 {
	return self.estrutura.GetX()
}

func (self *Comida) GetY() float64 {
	return self.estrutura.GetY()
}

func (self *Comida) GetLargura() float64 {
	return self.estrutura.GetLargura()
}

func (self *Comida) GetAltura() float64 {
	return self.estrutura.GetAltura()
}

func (self *Comida) GetTipo() string {
	return "COMIDA"
}

func (self *Comida) Atualizar() {
	self.ciclos += 1

	if self.ciclos > 200 {
		self.ciclos = 0
	}
}

func (self *Comida) GetComponente(id string) interface{} {
	return self.Componentes[id]
}

func (self *Comida) AdicionarComponente(id string, comp interface{}) {
	if self.Componentes == nil {
		self.Componentes = make(map[string]interface{})
	}
	self.Componentes[id] = comp
}

func (self *Comida) Desenhar(tela *ebiten.Image) {
	if !self.ObterEnergia().Status {
		return
	}

	var raioMaior float64 = utils.COMIDA_TAMANHO_MUNDO
	var raioMedio float64 = utils.COMIDA_TAMANHO_MUNDO - 5.0
	var raioPequeno float64 = utils.COMIDA_TAMANHO_MUNDO - 8.0

	ebitenutil.DrawCircle(tela, (self.cenaJogo.GetCamera().GetX()+self.GetX())+(raioMaior/2), (self.cenaJogo.GetCamera().GetY()+self.GetY())+(raioMaior/2), raioMaior, cores.ROSA_CHOQUE)
	ebitenutil.DrawCircle(tela, (self.cenaJogo.GetCamera().GetX()+self.GetX())+(raioMaior/2), (self.cenaJogo.GetCamera().GetY()+self.GetY())+(raioMaior/2), raioMedio, cores.BRANCO)

	if self.ciclos > 100 {
		ebitenutil.DrawCircle(tela, (self.cenaJogo.GetCamera().GetX()+self.GetX())+(raioMaior/2), (self.cenaJogo.GetCamera().GetY()+self.GetY())+(raioMaior/2), raioPequeno, cores.ROSA_CHOQUE)

		ebitenutil.DrawCircle(tela, (self.cenaJogo.GetCamera().GetX()+self.GetX())+(raioMaior/2)-20, (self.cenaJogo.GetCamera().GetY()+self.GetY())+(raioMaior/2), raioPequeno, cores.ROSA_CHOQUE)
		ebitenutil.DrawCircle(tela, (self.cenaJogo.GetCamera().GetX()+self.GetX())+(raioMaior/2)+20, (self.cenaJogo.GetCamera().GetY()+self.GetY())+(raioMaior/2), raioPequeno, cores.ROSA_CHOQUE)

		ebitenutil.DrawCircle(tela, (self.cenaJogo.GetCamera().GetX()+self.GetX())+(raioMaior/2), (self.cenaJogo.GetCamera().GetY()+self.GetY())+(raioMaior/2)-20, raioPequeno, cores.ROSA_CHOQUE)
		ebitenutil.DrawCircle(tela, (self.cenaJogo.GetCamera().GetX()+self.GetX())+(raioMaior/2), (self.cenaJogo.GetCamera().GetY()+self.GetY())+(raioMaior/2)+20, raioPequeno, cores.ROSA_CHOQUE)
	}

}

func (self *Comida) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawCircle(tela, mapaX+(self.GetX()/config.PROPORCAO_MAPA), mapaY+(self.GetY()/config.PROPORCAO_MAPA), utils.COMIDA_TAMANHO_MAPA, cores.LARANJA)
}

func (self *Comida) ExisteComponente(id string) bool {
	_, existe := self.Componentes[id]
	return existe
}
