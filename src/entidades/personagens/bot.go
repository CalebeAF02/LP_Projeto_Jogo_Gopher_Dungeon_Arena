package personagens

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/utils"

	"Gopher_Dungeon_Arena/src/entidades/geometria"

	"Gopher_Dungeon_Arena/src/interfaces"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bot struct {
	game     interfaces.IGame
	entidade ecs.EntidadeID
	Id       int64
	sangue   int
	cor      color.Color
	status   bool
	movendo  interfaces.Movimentador
	posicao  *geometria.Ponto
}

const (
	BOT_VELOCIDADE_RAPIDA = 7
	BOT_VELOCIDADE_NORMAL = 5
	BOT_VELOCIDADE_LENTA  = 3
	BOT_VELOCIDADE_MINIMA = 0
	BOT_VELOCIDADE_MAXIMA = 10
	BOT_CICLOS_REPETICAO  = 10
	BOT_TAMANHO           = 10
)

func NovoBot(game interfaces.IGame, id int64) *Bot {

	nEntidade := game.CriarEntidade()
	posicao := geometria.NovoPonto(0, 0)
	nBot := Bot{game: game, entidade: nEntidade, Id: id, sangue: 100, cor: cores.BRANCO, status: true, posicao: posicao}

	game.SetEntidade(nEntidade, &nBot)

	return &nBot
}

func (b *Bot) SetPosicao(x float64, y float64) {
	b.posicao.SetPosicao(x, y)
}

func (b *Bot) GetX() float64 {
	return b.posicao.GetX()
}

func (b *Bot) GetY() float64 {
	return b.posicao.GetY()
}

func (b *Bot) GetCor() color.Color {
	return b.cor
}

func (b *Bot) SetCor(c color.Color) {
	b.cor = c
}

func (b *Bot) Mover(r *rand.Rand) {
	if b.movendo != nil {
		b.movendo.Mover(b.game.GetMundo(), b, r)
	}
}

func (b *Bot) SetMovimentacao(movendo interfaces.Movimentador) {
	b.movendo = movendo
}

func (b *Bot) GetMovendoTipo() string {
	return b.movendo.GetTipo()
}

func (b *Bot) GetTipo() string {
	return entidades.BOT.String()
}

func (b *Bot) Atualizar() {
	b.Mover(b.game.GetAleatorio())
}

func (b *Bot) Desenhar(tela *ebiten.Image) {
	ebitenutil.DrawRect(tela, b.game.GetCameraX()+b.GetX(), b.game.GetCameraY()+b.GetY(), 10, 10, b.GetCor())
}

func (b *Bot) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {

	ebitenutil.DrawRect(tela, mapaX+(b.GetX()/utils.PROPORCAO_MINI_MAPA), mapaX+(b.GetY()/utils.PROPORCAO_MINI_MAPA), 2, 2, cores.VERMELHO)

}
