package personagens

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"image/color"

	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Jogador struct {
	game     interfaces.IGame
	entidade ecs.EntidadeID
	nome     string
	vida     int
	sangue   int
	cor      color.Color
	Status   bool
	vivendo  int
	posicao  *geometria.Ponto
}

func NovoJogador(game interfaces.IGame, n string) *Jogador {
	nEntidade := game.CriarEntidade()

	posicao := geometria.NovoPonto(0, 0)
	nJogador := Jogador{game: game, entidade: nEntidade, nome: n, vida: 2, sangue: 100, cor: color.White, Status: true, vivendo: 0, posicao: posicao}
	game.SetEntidade(nEntidade, &nJogador)

	return &nJogador
}

func (j *Jogador) estaVivo() bool {
	if j.vida == 0 {
		j.Status = false
	}
	return j.Status
}

func (j *Jogador) renasce() {
	if j.sangue == 0 {
		j.vida -= 1
		if j.estaVivo() {
			fmt.Println("O jogador " + j.nome + " morreu!")
		}
		j.resetaSangue()
	}
}

func (j *Jogador) resetaSangue() {
	if j.vida == 1 {
		j.sangue = 100
	}
}

func (j *Jogador) GetNome() string {
	return j.nome
}

func (j *Jogador) GetPosicao() *geometria.Ponto {
	return j.posicao
}

func (j *Jogador) SetPosicao(x float64, y float64) {
	j.posicao.SetPosicao(x, y)
}

func (j *Jogador) GetX() float64 {
	return j.posicao.GetX()
}

func (j *Jogador) GetY() float64 {
	return j.posicao.GetY()
}

func (j *Jogador) GetCor() color.Color {
	return j.cor
}

func (j *Jogador) SetCor(cor color.Color) {
	j.cor = cor
}

func (j *Jogador) GetVivendo() int {
	return j.vivendo
}

func (j *Jogador) GetTipo() string {
	return entidades.JOGADOR.String()
}

func (j *Jogador) Atualizar() {
	j.vivendo += 1
	if j.vivendo >= 30 {
		j.vivendo = 0
	}
}

func (jogador *Jogador) Desenhar(tela *ebiten.Image) {
	ebitenutil.DrawRect(tela, jogador.game.GetCameraX()+jogador.GetX(), jogador.game.GetCameraY()+jogador.GetY(), 20, 20, jogador.GetCor())

	ebitenutil.DrawRect(tela, jogador.game.GetCameraX()+jogador.GetX()+5, jogador.game.GetCameraY()+jogador.GetY()+5, 5, 5, color.White)
	ebitenutil.DrawRect(tela, jogador.game.GetCameraX()+jogador.GetX()+10, jogador.game.GetCameraY()+jogador.GetY()+5, 5, 5, color.White)

	ebitenutil.DrawRect(tela, jogador.game.GetCameraX()+jogador.GetX()+5, jogador.game.GetCameraY()+jogador.GetY()+10, 5, 5, color.White)
	ebitenutil.DrawRect(tela, jogador.game.GetCameraX()+jogador.GetX()+10, jogador.game.GetCameraY()+jogador.GetY()+10, 5, 5, color.White)

}

func (jogador *Jogador) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {

	ebitenutil.DrawRect(tela, mapaX+(jogador.GetX()/utils.PROPORCAO_MINI_MAPA), mapaX+(jogador.GetY()/utils.PROPORCAO_MINI_MAPA), 3, 3, cores.AZUL)

}
