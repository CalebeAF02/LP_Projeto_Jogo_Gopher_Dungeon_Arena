package personagens

import (
	"Gopher_Dungeon_Arena/src/config"
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
	game        interfaces.IGame
	entidade    ecs.EntidadeID
	nome        string
	vida        int
	sangue      int
	cor         color.Color
	Status      bool
	posicao     *geometria.Ponto
	Componentes map[string]interface{}
}

func NovoJogador(game interfaces.IGame, n string) *Jogador {
	nEntidade := game.CriarEntidade()

	posicao := geometria.NovoPonto(0, 0)
	nJogador := Jogador{game: game, entidade: nEntidade, nome: n, vida: 2, sangue: 100, cor: color.White, Status: true, posicao: posicao}
	game.SetEntidade(nEntidade, &nJogador)

	return &nJogador
}

func (j *Jogador) EstaVivo() bool {
	if j.vida == 0 {
		j.Status = false
	}
	return j.Status
}

func (j *Jogador) renasce() {
	if j.sangue == 0 {
		j.vida -= 1
		if j.EstaVivo() {
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
func (j *Jogador) GetX() float64 {
	return j.posicao.GetX()
}
func (j *Jogador) GetY() float64 {
	return j.posicao.GetY()
}
func (j *Jogador) GetCor() color.Color {
	return j.cor
}

func (j *Jogador) GetTipo() string {
	return entidades.JOGADOR.String()
}

func (j *Jogador) SetPosicao(x float64, y float64) {
	j.posicao.SetPosicao(x, y)
}
func (j *Jogador) SetX(x float64) {
	j.posicao.SetX(x)
}
func (j *Jogador) SetY(y float64) {
	j.posicao.SetY(y)
}
func (j *Jogador) SetCor(cor color.Color) {
	j.cor = cor
}

func (j *Jogador) Mover() {
	speed := float64(utils.JOGADOR_TAMANHO_MUNDO / config.PROPORCAO_MAPA)
	posX := j.posicao.GetX()
	posY := j.posicao.GetY()

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		posX = posX - speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		posX = posX + speed

	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		posY = posY - speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		posY = posY + speed
	}

	if j.game.GetMundo().EstaNaMargemInterna(geometria.NovoRetangulo(posX, posY, utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO), utils.JOGADOR_TAMANHO_MUNDO) {

		corpo := geometria.NovoRetangulo(posX, posY, utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO)

		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			//j.game.GerarBot(posX, posY)
		}

		if j.game.ColideComBarreiras(corpo) {
			return
		}

		j.SetPosicao(posX, posY)
	}
}

func (j *Jogador) Atualizar() {
	if j.EstaVivo() {
		j.Mover()
	}
}

func (j *Jogador) Desenhar(tela *ebiten.Image) {
	ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX(), j.game.GetCamera().GetY()+j.GetY(), utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO, j.GetCor())

	//ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX()+5, j.game.GetCamera().GetY()+j.GetY()+5, JOGADOR_TAMANHO_INTERNO, JOGADOR_TAMANHO_INTERNO, color.White)
	//ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX()+10, j.game.GetCamera().GetY()+j.GetY()+5, JOGADOR_TAMANHO_INTERNO, JOGADOR_TAMANHO_INTERNO, color.White)

	//ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX()+5, j.game.GetCamera().GetY()+j.GetY()+10, JOGADOR_TAMANHO_INTERNO, JOGADOR_TAMANHO_INTERNO, color.White)
	//ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX()+10, j.game.GetCamera().GetY()+j.GetY()+10, JOGADOR_TAMANHO_INTERNO, JOGADOR_TAMANHO_INTERNO, color.White)

}

func (j *Jogador) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawRect(tela, mapaX+(j.GetX()/config.PROPORCAO_MAPA), mapaY+(j.GetY()/config.PROPORCAO_MAPA), utils.JOGADOR_TAMANHO_MAPA, utils.JOGADOR_TAMANHO_MAPA, cores.AZUL)

	//ebitenutil.DrawRect(tela, (mapaX + (j.GetX() / utils.PROPORCAO_MAPA) + 1), (mapaY + (j.GetY() / utils.PROPORCAO_MAPA) + 1), JOGADOR_TAMANHO_INTERNO/2, JOGADOR_TAMANHO_INTERNO/2, cores.PRETO)
	//ebitenutil.DrawRect(tela, (mapaX + (j.GetX() / utils.PROPORCAO_MAPA) + 1), (mapaY + (j.GetY() / utils.PROPORCAO_MAPA) + 2), JOGADOR_TAMANHO_INTERNO/2, JOGADOR_TAMANHO_INTERNO/2, cores.PRETO)
	//ebitenutil.DrawRect(tela, (mapaX + (j.GetX() / utils.PROPORCAO_MAPA) + 2), (mapaY + (j.GetY() / utils.PROPORCAO_MAPA) + 1), JOGADOR_TAMANHO_INTERNO/2, JOGADOR_TAMANHO_INTERNO/2, cores.PRETO)
	//ebitenutil.DrawRect(tela, (mapaX + (j.GetX() / utils.PROPORCAO_MAPA) + 2), (mapaY + (j.GetY() / utils.PROPORCAO_MAPA) + 2), JOGADOR_TAMANHO_INTERNO/2, JOGADOR_TAMANHO_INTERNO/2, cores.PRETO)
}

func (e *Jogador) GetComponente(id string) interface{} {
	return e.Componentes[id]
}

func (e *Jogador) AdicionarComponente(id string, comp interface{}) {
	if e.Componentes == nil {
		e.Componentes = make(map[string]interface{})
	}
	e.Componentes[id] = comp
}
