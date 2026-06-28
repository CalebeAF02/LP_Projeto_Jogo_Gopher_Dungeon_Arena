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

func (c *Comida) ObterEnergia() *componentes.Energia {
	if energia_comp := c.GetComponente(componentes.ENERGIA.String()); energia_comp != nil {
		return energia_comp.(*componentes.Energia)
	}
	return nil
}

func (c *Comida) GetID() ecs.EntidadeID {
	return c.entidadeID
}

func (c *Comida) GetX() float64 {
	return c.estrutura.GetX()
}

func (c *Comida) GetY() float64 {
	return c.estrutura.GetY()
}

func (c *Comida) GetLargura() float64 {
	return c.estrutura.GetLargura()
}

func (c *Comida) GetAltura() float64 {
	return c.estrutura.GetAltura()
}

func (c *Comida) GetTipo() string {
	return "COMIDA"
}

func (c *Comida) Atualizar() {
	c.ciclos += 1

	if c.ciclos > 200 {
		c.ciclos = 0
	}
}

func (c *Comida) GetComponente(id string) interface{} {
	return c.Componentes[id]
}

func (c *Comida) AdicionarComponente(id string, comp interface{}) {
	if c.Componentes == nil {
		c.Componentes = make(map[string]interface{})
	}
	c.Componentes[id] = comp
}

func (c *Comida) Desenhar(tela *ebiten.Image) {
	if !c.ObterEnergia().Status {
		return
	}

	var raioMaior float64 = utils.COMIDA_TAMANHO_MUNDO
	var raioMedio float64 = utils.COMIDA_TAMANHO_MUNDO - 5.0
	var raioPequeno float64 = utils.COMIDA_TAMANHO_MUNDO - 8.0

	ebitenutil.DrawCircle(tela, (c.cenaJogo.GetCamera().GetX()+c.GetX())+(raioMaior/2), (c.cenaJogo.GetCamera().GetY()+c.GetY())+(raioMaior/2), raioMaior, cores.ROSA_CHOQUE)
	ebitenutil.DrawCircle(tela, (c.cenaJogo.GetCamera().GetX()+c.GetX())+(raioMaior/2), (c.cenaJogo.GetCamera().GetY()+c.GetY())+(raioMaior/2), raioMedio, cores.BRANCO)

	if c.ciclos > 100 {
		ebitenutil.DrawCircle(tela, (c.cenaJogo.GetCamera().GetX()+c.GetX())+(raioMaior/2), (c.cenaJogo.GetCamera().GetY()+c.GetY())+(raioMaior/2), raioPequeno, cores.ROSA_CHOQUE)

		ebitenutil.DrawCircle(tela, (c.cenaJogo.GetCamera().GetX()+c.GetX())+(raioMaior/2)-20, (c.cenaJogo.GetCamera().GetY()+c.GetY())+(raioMaior/2), raioPequeno, cores.ROSA_CHOQUE)
		ebitenutil.DrawCircle(tela, (c.cenaJogo.GetCamera().GetX()+c.GetX())+(raioMaior/2)+20, (c.cenaJogo.GetCamera().GetY()+c.GetY())+(raioMaior/2), raioPequeno, cores.ROSA_CHOQUE)

		ebitenutil.DrawCircle(tela, (c.cenaJogo.GetCamera().GetX()+c.GetX())+(raioMaior/2), (c.cenaJogo.GetCamera().GetY()+c.GetY())+(raioMaior/2)-20, raioPequeno, cores.ROSA_CHOQUE)
		ebitenutil.DrawCircle(tela, (c.cenaJogo.GetCamera().GetX()+c.GetX())+(raioMaior/2), (c.cenaJogo.GetCamera().GetY()+c.GetY())+(raioMaior/2)+20, raioPequeno, cores.ROSA_CHOQUE)
	}

}

func (c *Comida) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawCircle(tela, mapaX+(c.GetX()/config.PROPORCAO_MAPA), mapaY+(c.GetY()/config.PROPORCAO_MAPA), utils.COMIDA_TAMANHO_MAPA, cores.LARANJA)
}

func (c *Comida) ExisteComponente(id string) bool {
	_, existe := c.Componentes[id]
	return existe
}
