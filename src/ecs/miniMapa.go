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

func (mm *MiniMapa) Dados() {
	fmt.Println("Mapa>>>>>>>>>>>> mundoPosX: ", mm.mundo.GetX(), " | mundoPosY: ", mm.mundo.GetY(), " | mundoLargura: ", mm.mundo.GetLargura(), " | mundoAltura: ", mm.mundo.GetAltura())
	fmt.Println("Mapa>>>>>>>>>>>> miniMapaPosX: ", mm.miniMapa.GetX(), " | miniMapaPosY: ", mm.miniMapa.GetY(), " | miniMapaLargura: ", mm.miniMapa.GetLargura(), " | miniMapaAltura: ", mm.miniMapa.GetAltura())
	fmt.Println("Mapa>>>>>>>>>>>> posicaoPosX: ", mm.posicao.GetX(), " | posicaoPosY: ", mm.posicao.GetY())
	fmt.Println("Mapa>>>>>>>>>>>> cameraPosX: ", mm.camera.GetX(), " | cameraPosY: ", mm.camera.GetY(), " | cameraLargura: ", mm.camera.GetLargura(), " | cameraAltura: ", mm.camera.GetAltura())

}

func (mm *MiniMapa) GetX() float64 {
	return mm.miniMapa.GetX()
}
func (mm *MiniMapa) GetY() float64 {
	return mm.miniMapa.GetY()
}

func (mm *MiniMapa) SetPosicao(x float64, y float64) {
	mm.miniMapa.SetPosicao(x, y)
}

func (mm *MiniMapa) NovaCameraMiniMapa() {
	amareloPosX := mm.miniMapa.GetX() - (mm.camera.GetX() / config.PROPORCAO_MAPA)
	amareloPosY := mm.miniMapa.GetY() - (mm.camera.GetY() / config.PROPORCAO_MAPA)
	amareloLargura := mm.miniMapa.GetLargura() / config.PROPORCAO_MUNDO
	amareloAltura := mm.miniMapa.GetAltura() / config.PROPORCAO_MUNDO

	if mm.mundo.EstaDentroDireto(amareloPosX, amareloPosY, amareloLargura, amareloAltura) {
		mm.cameraMiniMapa.SetX(float64(amareloPosX))
		mm.cameraMiniMapa.SetY(amareloPosY)
		mm.cameraMiniMapa.SetLargura(amareloLargura)
		mm.cameraMiniMapa.SetAltura(amareloAltura)
	}
}

func (mm *MiniMapa) Mover() {
}

func (mm *MiniMapa) Atualizar() {
	mm.NovaCameraMiniMapa()
	//mm.Dados()
}

func (mm *MiniMapa) Desenhar(tela *ebiten.Image) {
	//Tela Branca
	ebitenutil.DrawRect(tela, mm.miniMapa.GetX(), mm.miniMapa.GetY(), mm.miniMapa.GetLargura(), mm.miniMapa.GetAltura(), cores.BRANCO)

	//Tela Amarela
	//amareloPosX := mm.miniMapa.GetX() - (mm.camera.GetX() / config.PROPORCAO_MAPA)
	//amareloPosY := mm.miniMapa.GetY() - (mm.camera.GetY() / config.PROPORCAO_MAPA)
	//amareloLargura := mm.miniMapa.GetLargura() / config.PROPORCAO_MUNDO
	//amareloAltura := mm.miniMapa.GetAltura() / config.PROPORCAO_MUNDO
	mm.NovaCameraMiniMapa()

	ebitenutil.DrawRect(tela, mm.cameraMiniMapa.GetX(), mm.cameraMiniMapa.GetY(), mm.cameraMiniMapa.GetLargura(), mm.cameraMiniMapa.GetAltura(), cores.AMARELO_CLARO)

	//Margem
	utils.MargemInterna(tela, mm.miniMapa, utils.JOGADOR_TAMANHO_MAPA, cores.VERMELHO)
}

func (mm *MiniMapa) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
}
