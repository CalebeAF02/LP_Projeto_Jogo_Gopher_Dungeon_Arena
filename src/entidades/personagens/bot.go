package personagens

import (
	"Gopher_Dungeon_Arena/src/componentes"
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
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
	Componentes map[string]interface{}
}

func NovoBot(cj interfaces.ICenaJogo, id int64) *Bot {

	nEntidade := cj.CriarEntidade()
	corpo := geometria.NovoRetangulo(0, 0, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO)
	nBot := Bot{cenaJogo: cj, entidadeID: nEntidade, Id: id}

	cj.SetEntidade(nEntidade, &nBot)
	nBot.AdicionarComponente(componentes.CORPO.String(), corpo)
	nBot.AdicionarComponente(componentes.SUB_TIPO.String(), &componentes.SubTipo{Valor: ""})
	nBot.AdicionarComponente(componentes.VIDA.String(), &componentes.Vida{TipoOrganismo: entidades.BOT.String(), Status: true, Quantidade: 1, Sangue: 100})
	nBot.AdicionarComponente(componentes.NIVEL.String(), &componentes.Nivel{Valor: 1, Progressao: 0})
	nBot.AdicionarComponente(componentes.ATIVIDADE.String(), &componentes.Atividade{Acao: componentes.AIVIDADE_MOVIMENTO})
	nBot.AdicionarComponente(componentes.MOVIMENTO.String(), &componentes.Movimento{Tipo: nil, Cor: cores.PRETO})

	nBot.entidade = &nBot
	return &nBot
}

func (self *Bot) GetID() ecs.EntidadeID {
	return self.entidadeID
}

func (self *Bot) ObterVida() *componentes.Vida {
	if sangue_comp := self.GetComponente(componentes.VIDA.String()); sangue_comp != nil {
		return sangue_comp.(*componentes.Vida)
	}
	return nil
}
func (self *Bot) ObterCorpo() *geometria.Retangulo {
	if corpo_comp := self.GetComponente(componentes.CORPO.String()); corpo_comp != nil {
		return corpo_comp.(*geometria.Retangulo)
	}
	return nil
}
func (self *Bot) ObterNivel() *componentes.Nivel {
	if nivel_comp := self.GetComponente(componentes.NIVEL.String()); nivel_comp != nil {
		return nivel_comp.(*componentes.Nivel)
	}
	return nil
}

func (self *Bot) ObterAtividade() *componentes.Atividade {
	if nivel_comp := self.GetComponente(componentes.ATIVIDADE.String()); nivel_comp != nil {
		return nivel_comp.(*componentes.Atividade)
	}
	return nil
}

func (self *Bot) ObterMovimento() *componentes.Movimento {
	if mov_comp := self.GetComponente(componentes.MOVIMENTO.String()); mov_comp != nil {
		return mov_comp.(*componentes.Movimento)
	}
	return nil
}

func (self *Bot) EstaVivo() bool {
	if !self.ObterVida().EstaVivo() {
		self.SetCor(cores.CINZA_CLARO)
		return false
	}
	return true
}

func (self *Bot) PossoMeMover() bool {
	return self.ObterAtividade().Acao == componentes.AIVIDADE_MOVIMENTO
}

func (self *Bot) CorrigeSangue() {
	self.ObterVida().CorrigeSangue(self.ObterNivel().Valor)
}

func (self *Bot) PerdeSangue(rit int) {
	self.ObterVida().PerdeSangue(rit, self.ObterNivel().Valor)
}

func (self *Bot) ResetaSangue() {
	self.ObterVida().ResetaSangue(self.ObterNivel().Valor)
}

func (self *Bot) SetPosicao(x float64, y float64) {
	self.ObterCorpo().SetX(x)
	self.ObterCorpo().SetY(y)
}

func (self *Bot) GetEntidade() ecs.Entidade {
	return self.entidade
}
func (self *Bot) GetPosicao() *geometria.Ponto {
	return geometria.NovoPonto(self.ObterCorpo().GetX(), self.ObterCorpo().GetY())
}
func (self *Bot) GetX1() float64 {
	return self.GetPosicao().GetX()
}

func (self *Bot) GetY1() float64 {
	return self.GetPosicao().GetY()
}

func (self *Bot) GetX2() float64 {
	return self.GetPosicao().GetX() + utils.BOT_TAMANHO_MUNDO
}

func (self *Bot) GetY2() float64 {
	return self.GetPosicao().GetY() + utils.BOT_TAMANHO_MUNDO
}
func (self *Bot) GetLargura() float64 {
	return utils.BOT_TAMANHO_MUNDO
}

func (self *Bot) GetAltura() float64 {
	return utils.BOT_TAMANHO_MUNDO
}

func (self *Bot) GetCor() color.Color {
	return self.ObterMovimento().Cor
}

func (self *Bot) GetMovendoTipo() string {
	return self.ObterMovimento().Tipo.GetTipo()
}

func (self *Bot) GetTipo() string {
	return entidades.BOT.String()
}

func (self *Bot) GetSubTipo() string {
	return self.ObterMovimento().Tipo.GetTipo()
}

func (self *Bot) GetNivel() int {
	return self.ObterNivel().Valor
}

func (self *Bot) SetCor(c color.Color) {
	self.ObterMovimento().Cor = c
}

func (self *Bot) SetNivel(nivel int) {
	self.ObterNivel().Valor = nivel
	self.CorrigeSangue()
}

func (self *Bot) SetNivelAleatorio() {
	nivel := self.cenaJogo.GetAleatorio().Intn(100)
	switch {
	case nivel >= 70:
		self.ObterNivel().Valor = 3
	case nivel >= 50:
		self.ObterNivel().Valor = 2
	default:
		self.ObterNivel().Valor = 1
	}
	self.CorrigeSangue()
}

func (self *Bot) Mover(r *rand.Rand) {
	posX := self.GetPosicao().GetX()
	posY := self.GetPosicao().GetY()

	if self.ObterMovimento().Tipo != nil {
		self.ObterMovimento().Tipo.Mover(self.entidade, self.cenaJogo.GetSistemaColisao(), self.cenaJogo.GetMundo(), self, r)
	}

	if self.cenaJogo.GetMundo().EstaNaMargemInterna(geometria.NovoRetangulo(posX, posY, utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO), utils.BOT_TAMANHO_MUNDO) {
		//b.SetPosicao(posX, posY)
		//fmt.Println("Bot Passou nesta funcao !")
	}
}

func (self *Bot) SetMovimentacao(movendo interfaces.Movimentador) {
	self.ObterMovimento().Tipo = movendo
	self.AlterarComponente(componentes.MOVIMENTO.String(), &componentes.Movimento{Tipo: movendo, Cor: movendo.GetCor()})
	self.AlterarComponente(componentes.SUB_TIPO.String(), &componentes.SubTipo{Valor: movendo.GetTipo()})
}

func (self *Bot) Atualizar() {
	if self.PossoMeMover() && self.ObterVida().EstaVivo() {
		self.Mover(self.cenaJogo.GetAleatorio())
	}
}

func (self *Bot) Desenhar(tela *ebiten.Image) {

	if self.PossoMeMover() {
		if self.ObterVida().Sangue < (100 * self.ObterNivel().Valor) {
			ebitenutil.DrawRect(tela, self.cenaJogo.GetCamera().GetX()+self.GetX1(), self.cenaJogo.GetCamera().GetY()+self.GetY1()-10, float64(self.ObterVida().Sangue)/5, 5, cores.VERMELHO_ESCURO)
		}
		ebitenutil.DrawRect(tela, self.cenaJogo.GetCamera().GetX()+self.GetX1(), self.cenaJogo.GetCamera().GetY()+self.GetY1(), utils.BOT_TAMANHO_MUNDO, utils.BOT_TAMANHO_MUNDO, self.GetCor())
	}
}

func (self *Bot) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	if self.PossoMeMover() {
		ebitenutil.DrawRect(tela, mapaX+(self.GetX1()/config.PROPORCAO_MAPA), mapaY+(self.GetY1()/config.PROPORCAO_MAPA), utils.BOT_TAMANHO_MAPA, utils.BOT_TAMANHO_MAPA, cores.VERMELHO)
	}
}

func (self *Bot) GetComponente(id string) interface{} {
	return self.Componentes[id]
}

func (self *Bot) AdicionarComponente(id string, comp interface{}) {
	if self.Componentes == nil {
		self.Componentes = make(map[string]interface{})
	}
	self.Componentes[id] = comp
}
func (self *Bot) AlterarComponente(id string, comp interface{}) {
	self.Componentes[id] = comp
}

func (self *Bot) ExisteComponente(id string) bool {
	_, existe := self.Componentes[id]
	return existe
}
