package ecs

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"fmt"

	//	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	mundo         *geometria.Retangulo
	posicao       *geometria.Ponto
	velocidade    *geometria.Ponto
	movendoCamera int
}

func NovaCamera(mundo *geometria.Retangulo) *Camera {
	nCamera := Camera{mundo: mundo, posicao: geometria.NovoPonto(0, 0), velocidade: geometria.NovoPonto(0, 0), movendoCamera: 0}
	return &nCamera
}

func (self *Camera) MoverCameraLivremente() {
	//c.movendoCamera += 1

	if self.movendoCamera >= 5 {
		self.posicao.SetX(self.posicao.GetX() + self.velocidade.GetX())
		self.posicao.SetY(self.posicao.GetY() + self.velocidade.GetY())

		self.movendoCamera = 0
	}
	//c.OrganizarCamera()
}

func (self *Camera) OrganizarCamera(posX float64, posY float64) {
	limiteCamerax2 := float64(config.MUNDO_LARGURA - config.JANELA_LARGURA)
	limiteCameraxY := float64(config.MUNDO_ALTURA - config.JANELA_ALTURA)

	if posX >= 0 {
		fmt.Println("Sai da visao Minima do Eixo X ", posX)
		self.SetX(0)
	} else if posX >= limiteCamerax2 {
		fmt.Println("Sai da visao Maxima do Eixo X ", posX)
		self.SetX(limiteCamerax2*(-1) - (config.JANELA_LARGURA / 2))
	} else {
		self.SetX(posX)
	}

	if posY >= 0 {
		fmt.Println("Sai da visao Minima do Eixo Y ", posY)
		self.SetY(0)
	} else if posY >= limiteCameraxY {
		fmt.Println("Sai da visao Maxima do Eixo Y ", posY)
		self.SetY((limiteCameraxY*(-1) - (config.JANELA_ALTURA / 2)))
	} else {
		self.SetY(posY)
	}
}

func (self *Camera) OrganizarCameraPeloJogador(j *geometria.Ponto) {
	// 1. Centraliza a câmera no jogador
	// A fórmula é: -(Posição_Jogador) + (Metade_da_Tela)
	novoX := -j.GetX() + (float64(config.JANELA_LARGURA) / 2)
	novoY := -j.GetY() + (float64(config.JANELA_ALTURA) / 2)

	// 2. Limites Máximos (o quanto a câmera pode "correr" para a esquerda/cima)
	// Lembre-se: em translação, valores são negativos ou zero.
	limiteMaxX := float64(config.MUNDO_LARGURA-config.JANELA_LARGURA) * -1
	limiteMaxY := float64(config.MUNDO_ALTURA-config.JANELA_ALTURA) * -1

	// 3. Trava no limite Mínimo (Esquerda/Topo)
	if novoX > 0 {
		novoX = 0
	}
	if novoY > 0 {
		novoY = 0
	}

	// 4. Trava no limite Máximo (Direita/Baixo)
	if novoX < limiteMaxX {
		novoX = limiteMaxX
	}
	if novoY < limiteMaxY {
		novoY = limiteMaxY
	}

	self.SetX(novoX)
	self.SetY(novoY)
}

func (self *Camera) GetMundo() *geometria.Retangulo {
	return self.mundo
}
func (self *Camera) GetX() float64 {
	return self.posicao.GetX()
}
func (self *Camera) GetY() float64 {
	return self.posicao.GetY()
}

func (self *Camera) GetLargura() float64 {
	return self.mundo.GetLargura()
}
func (self *Camera) GetAltura() float64 {
	return self.mundo.GetAltura()
}

func (self *Camera) SetMundo(tela *geometria.Retangulo) {
	self.mundo = tela
}
func (self *Camera) SetX(cameraX float64) {
	self.posicao.SetX(cameraX)
}
func (self *Camera) SetY(cameraY float64) {
	self.posicao.SetY(cameraY)
}
func (self *Camera) SetVelocidadeX(dirY float64) {
	self.velocidade.SetY(dirY)
}
func (self *Camera) SetVelocidadeY(dirX float64) {
	self.velocidade.SetX(dirX)
}

func (self *Camera) Atualizar() {
	//c.OrganizarCamera()
	
}

func (self *Camera) Desenhar(tela *ebiten.Image) {
	//Tela Amarela
	//ebitenutil.DrawRect(tela, c.GetX(), c.GetY(), (c.GetLargura()/PROPORCAO_MINI_MAPA)/2, (c.GetAltura()/PROPORCAO_MINI_MAPA)/2, cores.AMARELO)
}
