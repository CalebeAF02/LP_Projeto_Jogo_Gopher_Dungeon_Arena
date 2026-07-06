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

type MiniMapa struct {
	mundo          *geometria.Retangulo
	miniMapa       *geometria.Retangulo
	cameraMiniMapa *geometria.Retangulo
	posicao        *geometria.Ponto
	camera         *Camera
}

func NovoMiniMapa(mundo *geometria.Retangulo, posicao *geometria.Ponto, camera *Camera) *MiniMapa {
	miniMapa := geometria.NovoRetangulo(posicao.GetX(), posicao.GetY(), config.MAPA_LARGURA, config.MAPA_ALTURA)

	cameraMiniMapa := geometria.NovoRetangulo(0, 0, 0, 0)

	nMiniMapa := MiniMapa{mundo: mundo, miniMapa: miniMapa, cameraMiniMapa: cameraMiniMapa, posicao: posicao, camera: camera}
	return &nMiniMapa
}

func (self *MiniMapa) Dados() {
	fmt.Println("Mapa>>>>>>>>>>>> mundoPosX: ", self.mundo.GetX(), " | mundoPosY: ", self.mundo.GetY(), " | mundoLargura: ", self.mundo.GetLargura(), " | mundoAltura: ", self.mundo.GetAltura())
	fmt.Println("Mapa>>>>>>>>>>>> miniMapaPosX: ", self.miniMapa.GetX(), " | miniMapaPosY: ", self.miniMapa.GetY(), " | miniMapaLargura: ", self.miniMapa.GetLargura(), " | miniMapaAltura: ", self.miniMapa.GetAltura())
	fmt.Println("Mapa>>>>>>>>>>>> posicaoPosX: ", self.posicao.GetX(), " | posicaoPosY: ", self.posicao.GetY())
	fmt.Println("Mapa>>>>>>>>>>>> cameraPosX: ", self.camera.GetX(), " | cameraPosY: ", self.camera.GetY(), " | cameraLargura: ", self.camera.GetLargura(), " | cameraAltura: ", self.camera.GetAltura())

}

func (self *MiniMapa) GetX() float64 {
	return self.miniMapa.GetX()
}
func (self *MiniMapa) GetY() float64 {
	return self.miniMapa.GetY()
}

func (self *MiniMapa) SetPosicao(x float64, y float64) {
	self.miniMapa.SetPosicao(x, y)
}

func (self *MiniMapa) NovaCameraMiniMapa() {
	amareloPosX := self.miniMapa.GetX() - (self.camera.GetX() / config.PROPORCAO_MAPA)
	amareloPosY := self.miniMapa.GetY() - (self.camera.GetY() / config.PROPORCAO_MAPA)
	amareloLargura := self.miniMapa.GetLargura() / config.PROPORCAO_MUNDO
	amareloAltura := self.miniMapa.GetAltura() / config.PROPORCAO_MUNDO

	if self.mundo.EstaDentroDireto(amareloPosX, amareloPosY, amareloLargura, amareloAltura) {
		self.cameraMiniMapa.SetX(float64(amareloPosX))
		self.cameraMiniMapa.SetY(amareloPosY)
		self.cameraMiniMapa.SetLargura(amareloLargura)
		self.cameraMiniMapa.SetAltura(amareloAltura)
	}
}

func (self *MiniMapa) Mover() {
}

func (self *MiniMapa) Atualizar() {
	self.NovaCameraMiniMapa()
	//mm.Dados()
}

func (self *MiniMapa) Desenhar(tela *ebiten.Image) {
	//Tela Branca
	ebitenutil.DrawRect(tela, self.miniMapa.GetX(), self.miniMapa.GetY(), self.miniMapa.GetLargura(), self.miniMapa.GetAltura(), cores.BRANCO)

	//Tela Amarela
	//amareloPosX := mm.miniMapa.GetX() - (mm.camera.GetX() / config.PROPORCAO_MAPA)
	//amareloPosY := mm.miniMapa.GetY() - (mm.camera.GetY() / config.PROPORCAO_MAPA)
	//amareloLargura := mm.miniMapa.GetLargura() / config.PROPORCAO_MUNDO
	//amareloAltura := mm.miniMapa.GetAltura() / config.PROPORCAO_MUNDO
	self.NovaCameraMiniMapa()

	ebitenutil.DrawRect(tela, self.cameraMiniMapa.GetX(), self.cameraMiniMapa.GetY(), self.cameraMiniMapa.GetLargura(), self.cameraMiniMapa.GetAltura(), cores.AMARELO_CLARO)

	//Margem
	utils.MargemInterna(tela, self.miniMapa, utils.JOGADOR_TAMANHO_MAPA, cores.VERMELHO)
}

func (self *MiniMapa) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
}
