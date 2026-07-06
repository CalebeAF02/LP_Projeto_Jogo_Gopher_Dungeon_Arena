package objeto

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/componentes"
	"Gopher_Dungeon_Arena/src/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Parede struct {
	cenaJogo    interfaces.ICenaJogo
	entidadeID  ecs.EntidadeID
	estrutura   *geometria.Retangulo
	Componentes map[string]interface{}
}

func NovaParede(cenaJogo interfaces.ICenaJogo, posicao *geometria.Ponto) *Parede {

	nEntidade := cenaJogo.CriarEntidade()
	nParede := Parede{cenaJogo: cenaJogo, entidadeID: nEntidade, estrutura: geometria.NovoRetangulo(posicao.GetX(), posicao.GetY(), utils.PAREDE_TAMANHO_MUNDO, utils.PAREDE_TAMANHO_MUNDO)}

	cenaJogo.SetEntidade(nEntidade, &nParede)

	nParede.AdicionarComponente(componentes.CORPO.String(), nParede.estrutura)

	return &nParede
}

func (self *Parede) GetID() ecs.EntidadeID {
	return self.entidadeID
}

func (self *Parede) GetX() float64 {
	return self.estrutura.GetX()
}

func (self *Parede) GetY() float64 {
	return self.estrutura.GetY()
}

func (self *Parede) GetLargura() float64 {
	return self.estrutura.GetLargura()
}

func (self *Parede) GetAltura() float64 {
	return self.estrutura.GetAltura()
}

func (self *Parede) GetTipo() string {
	return "PAREDE"
}

func (self *Parede) Atualizar() {

}

func (self *Parede) GetComponente(id string) interface{} {
	return self.Componentes[id]
}

func (self *Parede) AdicionarComponente(id string, comp interface{}) {
	if self.Componentes == nil {
		self.Componentes = make(map[string]interface{})
	}
	self.Componentes[id] = comp
}

func (self *Parede) Desenhar(tela *ebiten.Image) {
	ebitenutil.DrawRect(tela, self.cenaJogo.GetCamera().GetX()+self.GetX(), self.cenaJogo.GetCamera().GetY()+self.GetY(), utils.PAREDE_TAMANHO_MUNDO, utils.PAREDE_TAMANHO_MUNDO, cores.PRETO)
}

func (self *Parede) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawRect(tela, mapaX+(self.GetX()/config.PROPORCAO_MAPA), mapaY+(self.GetY()/config.PROPORCAO_MAPA), utils.PAREDE_TAMANHO_MAPA, utils.PAREDE_TAMANHO_MAPA, cores.PRETO)
}

func (self *Parede) ExisteComponente(id string) bool {
	_, existe := self.Componentes[id]
	return existe
}
