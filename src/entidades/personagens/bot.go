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
	entidadeID  ecs.EntidadeID
	entidade    ecs.Entidade
	Id          int64
	cor         color.Color
	movendo     interfaces.Movimentador
	posicao     *geometria.Ponto
	corpo       *geometria.Retangulo
	Componentes map[string]interface{}
}

func NovoBot(cj interfaces.ICenaJogo, id int64) *Bot {

	nEntidade := cj.CriarEntidade()
	posicao := geometria.NovoPonto(0, 0)
	nBot := Bot{cenaJogo: cj, entidadeID: nEntidade, Id: id, cor: cores.BRANCO, posicao: posicao, corpo: geometria.NovoRetangulo(posicao.GetX(), posicao.GetY(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)}

	cj.SetEntidade(nEntidade, &nBot)
	nBot.AdicionarComponente(componentes.CORPO.String(), nBot.corpo)
	nBot.AdicionarComponente(componentes.SUB_TIPO.String(), &componentes.SubTipo{Valor: ""})
	nBot.AdicionarComponente(componentes.VIDA.String(), &componentes.Vida{TipoOrganismo: entidades.BOT.String(), Status: true, Quantidade: 1, Sangue: 100})
	nBot.AdicionarComponente(componentes.NIVEL.String(), &componentes.Nivel{Valor: 1, Progressao: 0})

	nBot.entidade = &nBot
	return &nBot
}

func (j *Bot) GetID() ecs.EntidadeID {
	return j.entidadeID
}

func (b *Bot) ObterVida() *componentes.Vida {
	if sangue_comp := b.GetComponente(componentes.VIDA.String()); sangue_comp != nil {
		return sangue_comp.(*componentes.Vida)
	}
	return nil
}

func (b *Bot) ObterNivel() *componentes.Nivel {
	if nivel_comp := b.GetComponente(componentes.NIVEL.String()); nivel_comp != nil {
		return nivel_comp.(*componentes.Nivel)
	}
	return nil
}

func (b *Bot) EstaVivo() bool {
	if !b.ObterVida().EstaVivo() {
		b.SetCor(cores.CINZA_CLARO)
		return false
	}
	return true
}

func (b *Bot) CorrigeSangue() {
	b.ObterVida().CorrigeSangue(b.ObterNivel().Valor)
}

func (b *Bot) PerdeSangue(rit int) {
	b.ObterVida().PerdeSangue(rit,b.ObterNivel().Valor)
}

func (b *Bot) ResetaSangue() {
	b.ObterVida().ResetaSangue(b.ObterNivel().Valor)
}

func (b *Bot) SetPosicao(x float64, y float64) {
	b.posicao.SetPosicao(x, y)
	b.corpo.SetX(x)
	b.corpo.SetY(y)
}

func (b *Bot) GetEntidade() ecs.Entidade {
	return b.entidade
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
	return b.ObterNivel().Valor
}

func (b *Bot) SetCor(c color.Color) {
	b.cor = c
}

func (b *Bot) SetNivel(nivel int) {
	b.ObterNivel().Valor = nivel
	b.CorrigeSangue()
}

func (b *Bot) SetNivelAleatorio() {
	nivel := b.cenaJogo.GetAleatorio().Intn(100)
	switch {
	case nivel >= 70:
		b.ObterNivel().Valor = 3
	case nivel >= 50:
		b.ObterNivel().Valor = 2
	default:
		b.ObterNivel().Valor = 1
	}
	b.CorrigeSangue()
}

func (b *Bot) Mover(r *rand.Rand) {
	posX := b.posicao.GetX()
	posY := b.posicao.GetY()

	if b.movendo != nil {
		b.movendo.Mover(b.entidade, b.cenaJogo.GetSistemaColisao(), b.cenaJogo.GetMundo(), b, r)
	}

	if b.cenaJogo.GetMundo().EstaNaMargemInterna(geometria.NovoRetangulo(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO), utils.BOT_TAMANHO_MUNDO) {
		//b.SetPosicao(posX, posY)
		//fmt.Println("Bot Passou nesta funcao !")
	}
}

func (b *Bot) SetMovimentacao(movendo interfaces.Movimentador) {
	b.movendo = movendo
	b.AlterarComponente(componentes.SUB_TIPO.String(), &componentes.SubTipo{Valor: movendo.GetTipo()})
}

func (b *Bot) Atualizar() {
	if b.ObterVida().EstaVivo() {
		b.Mover(b.cenaJogo.GetAleatorio())
	}
}

func (b *Bot) Desenhar(tela *ebiten.Image) {
	ebitenutil.DrawRect(tela, b.cenaJogo.GetCamera().GetX()+b.GetX1(), b.cenaJogo.GetCamera().GetY()+b.GetY1()-10, float64(b.ObterVida().Sangue)/5, 5, cores.VERMELHO_ESCURO)

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

func (e *Bot) ExisteComponente(id string) bool {
	_, existe := e.Componentes[id]
	return existe
}
