package outros

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Time struct {
	game      interfaces.IGame
	entidade  ecs.EntidadeID
	nome      string
	jogadores []*personagens.Jogador
	cor       color.Color
}

func NovoTime(game interfaces.IGame, n string, cor color.Color) Time {
	nEntidade := game.CriarEntidade()

	nTime := Time{game: game, entidade: nEntidade, nome: n, cor: cor}

	game.SetEntidade(nEntidade, &nTime)

	return nTime
}

func (t *Time) Adicionnar(jogador *personagens.Jogador) {
	jogador.SetCor(t.cor)
	t.jogadores = append(t.jogadores, jogador)
}

func (t *Time) EstaVivo() bool {
	for _, jogador := range t.jogadores {
		if jogador.Status {
			return true
		}
	}
	return false
}

func (t *Time) GetNome() string {
	return t.nome
}

func (t *Time) GetJogadores() []*personagens.Jogador {
	return t.jogadores
}

func (t *Time) Posicoes() {
	fmt.Printf("Lista de Jogadores do %s\n", t.nome)

	for i, j := range t.jogadores {
		fmt.Printf("Jogador_%d: %s esta na posicao %s\n", i+1, j.GetNome(), j.GetPosicao().ToString())
	}
}

func (t *Time) GetTipo() string {
	return entidades.TIME.String()
}

func (t *Time) Atualizar() {

}

func (t *Time) Desenhar(tela *ebiten.Image) {
}

func (t *Time) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {

}
