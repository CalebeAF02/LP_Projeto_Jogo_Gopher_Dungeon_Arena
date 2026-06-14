package ecs

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/utils"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type BarJogador struct {
	mundo      *geometria.Retangulo
	barJogador *geometria.Retangulo
	posicao    *geometria.Ponto
}

func NovoBarJogador(mundo *geometria.Retangulo, posicao *geometria.Ponto, camera *Camera) *BarJogador {
	barJogador := geometria.NovoRetangulo(posicao.GetX(), posicao.GetY(), config.MAPA_LARGURA, config.MAPA_ALTURA)

	nBar := BarJogador{mundo: mundo, barJogador: barJogador, posicao: posicao}
	return &nBar
}

func (bj *BarJogador) Dados() {
	fmt.Println("Mapa>>>>>>>>>>>> mundoPosX: ", bj.mundo.GetX(), " | mundoPosY: ", bj.mundo.GetY(), " | mundoLargura: ", bj.mundo.GetLargura(), " | mundoAltura: ", bj.mundo.GetAltura())
	fmt.Println("Mapa>>>>>>>>>>>> miniMapaPosX: ", bj.barJogador.GetX(), " | miniMapaPosY: ", bj.barJogador.GetY(), " | miniMapaLargura: ", bj.barJogador.GetLargura(), " | miniMapaAltura: ", bj.barJogador.GetAltura())
	fmt.Println("Mapa>>>>>>>>>>>> posicaoPosX: ", bj.posicao.GetX(), " | posicaoPosY: ", bj.posicao.GetY())

}

func (bj *BarJogador) GetX() float64 {
	return bj.barJogador.GetX()
}
func (bj *BarJogador) GetY() float64 {
	return bj.barJogador.GetY()
}

func (bj *BarJogador) Mover() {
}

func (bj *BarJogador) Atualizar() {
	//bj.Dados()
}

func (bj *BarJogador) Desenhar(tela *ebiten.Image) {
	//Tela Branca
	ebitenutil.DrawRect(tela, bj.barJogador.GetX(), bj.barJogador.GetY(), bj.barJogador.GetLargura(), bj.barJogador.GetAltura(), cores.BRANCO)

	//Tela Amarela
	//amareloPosX := bj.barJogador.GetX() - (bj.camera.GetX() / config.PROPORCAO_MAPA)
	//amareloPosY := bj.barJogador.GetY() - (bj.camera.GetY() / config.PROPORCAO_MAPA)
	//amareloLargura := bj.barJogador.GetLargura() / config.PROPORCAO_MUNDO
	//amareloAltura := bj.barJogador.GetAltura() / config.PROPORCAO_MUNDO

	//Margem
	utils.MargemInterna(tela, bj.barJogador, utils.JOGADOR_TAMANHO_MAPA, cores.VERMELHO)
}

func (bj *BarJogador) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
}
