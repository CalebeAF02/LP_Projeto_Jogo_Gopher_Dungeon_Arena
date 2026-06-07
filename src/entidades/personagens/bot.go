package personagens

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/enum/componentes"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"

	"Gopher_Dungeon_Arena/src/entidades/geometria"

	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bot struct {
	cenaJogo    interfaces.ICenaJogo
	entidade    ecs.EntidadeID
	Id          int64
	nivel       int
	sangue      int
	cor         color.Color
	status      bool
	movendo     interfaces.Movimentador
	posicao     *geometria.Ponto
	corpo       *geometria.Retangulo
	Componentes map[string]interface{}
}

type SubTipo struct {
	Valor string
}

func NovoBot(cj interfaces.ICenaJogo, id int64) *Bot {

	nEntidade := cj.CriarEntidade()
	posicao := geometria.NovoPonto(0, 0)
	nBot := Bot{cenaJogo: cj, entidade: nEntidade, Id: id, nivel: 1, sangue: 100, cor: cores.BRANCO, status: true, posicao: posicao, corpo: geometria.NovoRetangulo(posicao.GetX(), posicao.GetY(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)}

	cj.SetEntidade(nEntidade, &nBot)

	nBot.AdicionarComponente(componentes.CORPO.String(), nBot.corpo)
	nBot.AdicionarComponente(componentes.SUB_TIPO.String(), &SubTipo{Valor: ""})

	return &nBot
}

func (b *Bot) EstaVivo() bool {
	if b.sangue > 0 {
		b.status = true
		return b.status
	}
	b.status = false
	return b.status
}

func (b *Bot) ResetaSangue() {
	b.sangue = 100 * b.nivel
}

func (b *Bot) PerdeSangue(rit int) {
	b.sangue -= rit
}

func (b *Bot) SetPosicao(x float64, y float64) {
	b.posicao.SetPosicao(x, y)
	b.corpo.SetX(x)
	b.corpo.SetY(y)
}

func (b *Bot) GetCorpo() *geometria.Retangulo {
	corpo := geometria.NovoRetangulo(b.GetX1(), b.GetY1(), b.GetLargura(), b.GetAltura())
	return corpo
}

func (b *Bot) GetX1() float64 {
	return b.posicao.GetX()
}

func (b *Bot) GetY1() float64 {
	return b.posicao.GetY()
}

func (b *Bot) GetX2() float64 {
	return b.posicao.GetX() + utils.BOT_TAMANHO_MUNDO
}

func (b *Bot) GetY2() float64 {
	return b.posicao.GetY() + utils.BOT_TAMANHO_MUNDO
}
func (b *Bot) GetLargura() float64 {
	return utils.BOT_TAMANHO_MUNDO
}

func (b *Bot) GetAltura() float64 {
	return utils.BOT_TAMANHO_MUNDO
}

func (b *Bot) GetCor() color.Color {
	return b.cor
}

func (b *Bot) GetMovendoTipo() string {
	return b.movendo.GetTipo()
}

func (b *Bot) GetTipo() string {
	return entidades.BOT.String()
}

func (b *Bot) GetSubTipo() string {
	return b.movendo.GetTipo()
}

func (b *Bot) GetNivel() int {
	return b.nivel
}

func (b *Bot) SetCor(c color.Color) {
	b.cor = c
}

func (b *Bot) SetNivel(nivel int) {
	b.nivel = nivel
	b.CorrigeSangue()
}

func (b *Bot) SetNivelAleatorio() {
	nivel := b.cenaJogo.GetAleatorio().Intn(100)
	switch {
	case nivel >= 70:
		b.nivel = 3
	case nivel >= 50:
		b.nivel = 2
	default:
		b.nivel = 1
	}
	b.CorrigeSangue()
}

func (b *Bot) CorrigeSangue() {
	b.sangue = 100 * b.nivel
}

func (b *Bot) Mover(r *rand.Rand) {
	posX := b.posicao.GetX()
	posY := b.posicao.GetY()

	if b.movendo != nil {
		b.movendo.Mover(b.cenaJogo, b.cenaJogo.GetMundo(), b, r)
	}

	if b.cenaJogo.GetMundo().EstaNaMargemInterna(geometria.NovoRetangulo(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO), utils.BOT_TAMANHO_MUNDO) {
		//b.SetPosicao(posX, posY)
		//fmt.Println("Bot Passou nesta funcao !")
	}
}

func (b *Bot) SetMovimentacao(movendo interfaces.Movimentador) {
	b.movendo = movendo
	b.AlterarComponente(componentes.SUB_TIPO.String(), &SubTipo{Valor: movendo.GetTipo()})
}

func (b *Bot) Atualizar() {
	b.Mover(b.cenaJogo.GetAleatorio())
}

func (b *Bot) Desenhar(tela *ebiten.Image) {
	ebitenutil.DrawRect(tela, b.cenaJogo.GetCamera().GetX()+b.GetX1(), b.cenaJogo.GetCamera().GetY()+b.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO, b.GetCor())
}

func (b *Bot) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawRect(tela, mapaX+(b.GetX1()/config.PROPORCAO_MAPA), mapaY+(b.GetY1()/config.PROPORCAO_MAPA), utils.BOT_TAMANHO_MAPA, utils.BOT_TAMANHO_MAPA, cores.VERMELHO)
}

func (e *Bot) GetComponente(id string) interface{} {
	return e.Componentes[id]
}

func (e *Bot) AdicionarComponente(id string, comp interface{}) {
	if e.Componentes == nil {
		e.Componentes = make(map[string]interface{})
	}
	e.Componentes[id] = comp
}
func (e *Bot) AlterarComponente(id string, comp interface{}) {
	e.Componentes[id] = comp
}
