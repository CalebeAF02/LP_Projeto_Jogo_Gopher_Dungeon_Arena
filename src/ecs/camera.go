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

func (c *Camera) MoverCameraLivremente() {
	//c.movendoCamera += 1

	if c.movendoCamera >= 5 {
		c.posicao.SetX(c.posicao.GetX() + c.velocidade.GetX())
		c.posicao.SetY(c.posicao.GetY() + c.velocidade.GetY())

		c.movendoCamera = 0
	}
	//c.OrganizarCamera()
}

func (c *Camera) OrganizarCamera(posX float64, posY float64) {
	limiteCamerax2 := float64(config.MUNDO_LARGURA - config.JANELA_LARGURA)
	limiteCameraxY := float64(config.MUNDO_ALTURA - config.JANELA_ALTURA)

	if posX >= 0 {
		fmt.Println("Sai da visao Minima do Eixo X ", posX)
		c.SetX(0)
	} else if posX >= limiteCamerax2 {
		fmt.Println("Sai da visao Maxima do Eixo X ", posX)
		c.SetX(limiteCamerax2*(-1) - (config.JANELA_LARGURA / 2))
	} else {
		c.SetX(posX)
	}

	if posY >= 0 {
		fmt.Println("Sai da visao Minima do Eixo Y ", posY)
		c.SetY(0)
	} else if posY >= limiteCameraxY {
		fmt.Println("Sai da visao Maxima do Eixo Y ", posY)
		c.SetY((limiteCameraxY*(-1) - (config.JANELA_ALTURA / 2)))
	} else {
		c.SetY(posY)
	}
}

func (c *Camera) OrganizarCameraPeloJogador(j *geometria.Ponto) {
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

	c.SetX(novoX)
	c.SetY(novoY)
}

func (c *Camera) GetMundo() *geometria.Retangulo {
	return c.mundo
}
func (c *Camera) GetX() float64 {
	return c.posicao.GetX()
}
func (c *Camera) GetY() float64 {
	return c.posicao.GetY()
}

func (c *Camera) GetLargura() float64 {
	return c.mundo.GetLargura()
}
func (c *Camera) GetAltura() float64 {
	return c.mundo.GetAltura()
}

func (c *Camera) SetMundo(tela *geometria.Retangulo) {
	c.mundo = tela
}
func (c *Camera) SetX(cameraX float64) {
	c.posicao.SetX(cameraX)
}
func (c *Camera) SetY(cameraY float64) {
	c.posicao.SetY(cameraY)
}
func (c *Camera) SetVelocidadeX(dirY float64) {
	c.velocidade.SetY(dirY)
}
func (c *Camera) SetVelocidadeY(dirX float64) {
	c.velocidade.SetX(dirX)
}

func (c *Camera) Atualizar() {
	//c.OrganizarCamera()
	
}

func (c *Camera) Desenhar(tela *ebiten.Image) {
	//Tela Amarela
	//ebitenutil.DrawRect(tela, c.GetX(), c.GetY(), (c.GetLargura()/PROPORCAO_MINI_MAPA)/2, (c.GetAltura()/PROPORCAO_MINI_MAPA)/2, cores.AMARELO)
}
