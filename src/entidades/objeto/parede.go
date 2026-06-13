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

func (j *Parede) GetID() ecs.EntidadeID {
	return j.entidadeID
}

func (b *Parede) GetX() float64 {
	return b.estrutura.GetX()
}

func (b *Parede) GetY() float64 {
	return b.estrutura.GetY()
}

func (b *Parede) GetLargura() float64 {
	return b.estrutura.GetLargura()
}

func (b *Parede) GetAltura() float64 {
	return b.estrutura.GetAltura()
}

func (b *Parede) GetTipo() string {
	return "PAREDE"
}

func (b *Parede) Atualizar() {

}

func (e *Parede) GetComponente(id string) interface{} {
	return e.Componentes[id]
}

func (e *Parede) AdicionarComponente(id string, comp interface{}) {
	if e.Componentes == nil {
		e.Componentes = make(map[string]interface{})
	}
	e.Componentes[id] = comp
}

func (b *Parede) Desenhar(tela *ebiten.Image) {
	ebitenutil.DrawRect(tela, b.cenaJogo.GetCamera().GetX()+b.GetX(), b.cenaJogo.GetCamera().GetY()+b.GetY(), utils.PAREDE_TAMANHO_MUNDO, utils.PAREDE_TAMANHO_MUNDO, cores.PRETO)
}

func (b *Parede) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawRect(tela, mapaX+(b.GetX()/config.PROPORCAO_MAPA), mapaY+(b.GetY()/config.PROPORCAO_MAPA), utils.PAREDE_TAMANHO_MAPA, utils.PAREDE_TAMANHO_MAPA, cores.PRETO)
}

func (e *Parede) ExisteComponente(id string) bool {
	_, existe := e.Componentes[id]
	return existe
}
