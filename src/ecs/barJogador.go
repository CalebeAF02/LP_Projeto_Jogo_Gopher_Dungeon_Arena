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

func (self *BarJogador) Dados() {
	fmt.Println("Mapa>>>>>>>>>>>> mundoPosX: ", self.mundo.GetX(), " | mundoPosY: ", self.mundo.GetY(), " | mundoLargura: ", self.mundo.GetLargura(), " | mundoAltura: ", self.mundo.GetAltura())
	fmt.Println("Mapa>>>>>>>>>>>> miniMapaPosX: ", self.barJogador.GetX(), " | miniMapaPosY: ", self.barJogador.GetY(), " | miniMapaLargura: ", self.barJogador.GetLargura(), " | miniMapaAltura: ", self.barJogador.GetAltura())
	fmt.Println("Mapa>>>>>>>>>>>> posicaoPosX: ", self.posicao.GetX(), " | posicaoPosY: ", self.posicao.GetY())

}

func (self *BarJogador) GetX() float64 {
	return self.barJogador.GetX()
}
func (self *BarJogador) GetY() float64 {
	return self.barJogador.GetY()
}

func (self *BarJogador) Mover() {
}

func (self *BarJogador) Atualizar() {
	//bj.Dados()
}

func (self *BarJogador) Desenhar(tela *ebiten.Image) {
	//Tela Branca
	ebitenutil.DrawRect(tela, self.barJogador.GetX(), self.barJogador.GetY(), self.barJogador.GetLargura(), self.barJogador.GetAltura(), cores.BRANCO)

	//Tela Amarela
	//amareloPosX := bj.barJogador.GetX() - (bj.camera.GetX() / config.PROPORCAO_MAPA)
	//amareloPosY := bj.barJogador.GetY() - (bj.camera.GetY() / config.PROPORCAO_MAPA)
	//amareloLargura := bj.barJogador.GetLargura() / config.PROPORCAO_MUNDO
	//amareloAltura := bj.barJogador.GetAltura() / config.PROPORCAO_MUNDO

	//Margem
	utils.MargemInterna(tela, self.barJogador, utils.JOGADOR_TAMANHO_MAPA, cores.VERMELHO)
}

func (self *BarJogador) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
}
