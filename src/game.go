package src

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/movimentacao"
	"Gopher_Dungeon_Arena/src/entidades/outros"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"image/color"
	"math/rand"

	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	proximo   int
	mundo     *geometria.Retangulo
	Entidades map[ecs.EntidadeID]ecs.Entidade
	aleatorio *rand.Rand

	mapa *geometria.Retangulo

	camera           geometria.Ponto
	cameraVelocidade geometria.Ponto

	movendoCamera int
}

func NovoGame() *Game {
	mundo := geometria.NovoRetangulo(0, 0, utils.MUNDO_LARGURA, utils.MUNDO_ALTURA)
	entidades := make(map[ecs.EntidadeID]ecs.Entidade)
	aleatorio := utils.GeradorAleatorio()

	g := Game{mundo: mundo, Entidades: entidades, aleatorio: aleatorio, camera: *geometria.NovoPonto(0, 0), movendoCamera: 0}

	g.SetCameraX(500)
	g.SetDirecaoX(10)
	g.SetDirecaoY(5)

	//Mapa
	mapa := geometria.NovoRetangulo(10, 10, utils.MUNDO_LARGURA/utils.PROPORCAO_MINI_MAPA, utils.MUNDO_ALTURA/utils.PROPORCAO_MINI_MAPA)
	g.SetMapa(mapa)

	// Jogadores
	j1 := personagens.NovoJogador(&g, "Jogador 1")
	j2 := personagens.NovoJogador(&g, "Jogador 2")
	j3 := personagens.NovoJogador(&g, "Jogador 3")

	j4 := personagens.NovoJogador(&g, "Jogador 4")
	j5 := personagens.NovoJogador(&g, "Jogador 5")
	j6 := personagens.NovoJogador(&g, "Jogador 6")

	j1.SetPosicao(100, 500)
	j2.SetPosicao(200, 300)
	j3.SetPosicao(300, 500)

	j4.SetPosicao(500, 100)
	j5.SetPosicao(300, 200)
	j6.SetPosicao(500, 300)

	// Times
	t1 := outros.NovoTime(&g, "Vermelhao - Time_Vermelho", cores.VERMELHO)
	t2 := outros.NovoTime(&g, "Azulzinhos - Time_Azul", cores.AZUL)

	// Gerenciando
	t1.Adicionnar(j1)
	t1.Adicionnar(j2)
	t1.Adicionnar(j3)
	t1.Posicoes()

	t2.Adicionnar(j4)
	t2.Adicionnar(j5)
	t2.Adicionnar(j6)
	t2.Posicoes()

	//Bot
	for id := 0; id < 1; id++ {
		g.CriarBot(&movimentacao.MovimentadorSimples{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorVertical{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorVerticalConstante{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorHorizontal{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorHorizontalConstante{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorDiagonal{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorLogicoLinha{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorLogicoDiagonal{}, g.OrganizaPosicaoAleatoriaBot())
		g.CriarBot(&movimentacao.MovimentadorLogicoDuplo{}, g.OrganizaPosicaoAleatoriaBot())
	}

	return &g
}

func (g *Game) Layout(l, a int) (int, int) {
	return utils.LARGURA, utils.ALTURA
}

func (g *Game) CriarEntidade() ecs.EntidadeID {
	entidade := ecs.EntidadeID(g.proximo)
	g.proximo++
	return entidade
}

func (g *Game) CriarBot(movendo interfaces.Movimentador, posicao *geometria.Ponto) {
	b := personagens.NovoBot(g, 0)
	b.SetPosicao(posicao.GetX(), posicao.GetY())
	b.SetMovimentacao(movendo)

	if movendo.GetTipo() == "LOGICO_LINHA" {
		b.SetCor(cores.AMARELO)
	} else if movendo.GetTipo() == "LOGICO_DIAGONAL" {
		b.SetCor(cores.VERDE)
	} else if movendo.GetTipo() == "LOGICO_DUPLO" {
		b.SetCor(cores.AZUL)
	} else if movendo.GetTipo() == "SIMPLES" {
		b.SetCor(cores.VERDE_LIMAO)
	} else if movendo.GetTipo() == "VERTICAL" {
		b.SetCor(cores.AMARELO_CLARO)
	} else if movendo.GetTipo() == "VERTICAL_CONSTANTE" {
		b.SetCor(cores.AMARELO_ESCURO)
	} else if movendo.GetTipo() == "HORIZONTAL" {
		b.SetCor(cores.MARROM)
	} else if movendo.GetTipo() == "HORIZONTAL_CONSTANTE" {
		b.SetCor(cores.MARROM_ESCURO)
	} else if movendo.GetTipo() == "DIAGONAL" {
		b.SetCor(cores.ROSA)
	}

	b.SetMovimentacao(movendo)
	//fmt.Printf("BOT <%s> | X: %f | Y: %f\n", b.GetMovendoTipo(), b.GetX(), b.GetY())
}

func (g *Game) CriarBotAleatorio() {

	for id := 0; id < 10; id++ {
		b := personagens.NovoBot(g, int64(id))
		b.SetPosicao(utils.XAleatorio(g.aleatorio), utils.YAleatorio(g.aleatorio))

		movimentacaoAleatoria := g.aleatorio.Intn(100)
		if movimentacaoAleatoria >= 0 && movimentacaoAleatoria < 15 {
			b.SetMovimentacao(&movimentacao.MovimentadorSimples{})
			b.SetCor(cores.BRANCO)
		} else if movimentacaoAleatoria >= 15 && movimentacaoAleatoria < 40 {
			b.SetMovimentacao(&movimentacao.MovimentadorVertical{})
			b.SetCor(cores.VERDE)
		} else if movimentacaoAleatoria >= 40 && movimentacaoAleatoria < 60 {
			b.SetMovimentacao(&movimentacao.MovimentadorHorizontalConstante{})
			b.SetCor(cores.LARANJA)
		} else if movimentacaoAleatoria >= 60 && movimentacaoAleatoria < 80 {
			b.SetMovimentacao(&movimentacao.MovimentadorVerticalConstante{})
			b.SetCor(cores.VERDE)
		} else {
			b.SetMovimentacao(&movimentacao.MovimentadorDiagonal{})
			b.SetCor(cores.CIANO)
		}

		if g.aleatorio.Intn(100) < 30 {
		} else {
			valor := g.aleatorio.Intn(100)
			if valor < 30 {

			} else if valor > 30 && valor < 50 {
				v2 := g.aleatorio.Intn(100)
				if v2 > 50 {
				}
			} else if valor > 50 && valor < 70 {

			} else {
				b.SetMovimentacao(&movimentacao.MovimentadorHorizontal{})
				b.SetCor(cores.LARANJA)
			}
		}
	}
}
func (g *Game) OrganizaPosicaoAleatoriaBot() *geometria.Ponto {
	x, y := float64(g.aleatorio.Intn(int(g.mundo.PosXmax(personagens.BOT_TAMANHO)))), float64(g.aleatorio.Intn(int(g.mundo.PosYmax(personagens.BOT_TAMANHO))))

	return geometria.NovoPonto(x, y)
}

func (g *Game) GetEntidades() map[ecs.EntidadeID]ecs.Entidade {
	return g.Entidades
}
func (g *Game) GetMundo() geometria.Retangulo {
	return *g.mundo
}
func (g *Game) GetAleatorio() *rand.Rand {
	return g.aleatorio
}
func (g *Game) GetLargura() float64 {
	return g.mundo.GetLargura()
}
func (g *Game) GetAltura() float64 {
	return g.mundo.GetAltura()
}
func (g *Game) GetCameraX() float64 {
	return g.camera.GetX()
}
func (g *Game) GetCameraY() float64 {
	return g.camera.GetY()
}

func (g *Game) SetEntidade(nEntidade ecs.EntidadeID, posicao ecs.Entidade) {
	g.Entidades[nEntidade] = posicao
}
func (g *Game) SetDirecaoY(dirX float64) {
	g.cameraVelocidade.SetX(dirX)
}
func (g *Game) SetDirecaoX(dirY float64) {
	g.cameraVelocidade.SetY(dirY)
}
func (g *Game) SetCameraX(cameraX float64) {
	g.camera.SetX(cameraX)
}
func (g *Game) SetCameraY(cameraY float64) {
	g.camera.SetY(cameraY)
}
func (g *Game) SetMapa(mapa *geometria.Retangulo) {
	g.mapa = mapa
}

func (g *Game) Update() error {

	g.movendoCamera += 1

	if g.movendoCamera >= 5 {

		g.camera.SetX(g.camera.GetX() + g.cameraVelocidade.GetX())
		limiteCameraX := utils.MUNDO_LARGURA + g.camera.GetX()
		if g.camera.GetX() >= 0 {
			fmt.Printf("\nSai da visao Minima do Eixo X\n")
			g.camera.SetX(0)
			g.cameraVelocidade.SetX(g.cameraVelocidade.GetX() * (-1))
		} else if limiteCameraX <= utils.LARGURA {
			fmt.Printf("\nSai da visao Maxima do Eixo X\n")
			g.camera.SetX((utils.MUNDO_LARGURA - utils.LARGURA) * (-1))
			g.cameraVelocidade.SetX(g.cameraVelocidade.GetX() * (-1))
		}

		g.camera.SetY(g.camera.GetY() + g.cameraVelocidade.GetY())

		limiteCameraY := utils.MUNDO_ALTURA + g.camera.GetY()
		if g.camera.GetY() >= 0 {
			fmt.Printf("\nSai da visao Minima do Eixo Y\n")
			g.camera.SetY(0)
			g.cameraVelocidade.SetY(g.cameraVelocidade.GetY() * (-1))
		} else if limiteCameraY <= utils.ALTURA {
			fmt.Printf("\nSai da visao Maxima do Eixo Y\n")
			g.camera.SetY((utils.MUNDO_ALTURA - utils.ALTURA) * (-1))
			g.cameraVelocidade.SetY(g.cameraVelocidade.GetY() * (-1))

		}

		g.movendoCamera = 0
		fmt.Printf("--------->> MUNDO MINI MAPA : [ %.0f | %.0f ] \n\t CAMERA [ X = %.0f | Y = %.0f ] \n\t LimiteCamera [ X = %.0f | Y = %.0f ]\n\n", utils.MUNDO_LARGURA, utils.MUNDO_ALTURA, g.camera.GetX(), g.camera.GetY(), limiteCameraX, limiteCameraY)
	}

	for _, entidade := range g.Entidades {
		entidade.Atualizar()
	}

	return nil
}

func (g *Game) Draw(tela *ebiten.Image) {
	tela.Fill(color.RGBA{20, 20, 20, 255})
	// Desenhar Time 1 (Red)
	for _, entidade := range g.Entidades {
		entidade.Desenhar(tela)
	}

	utils.Margem(tela, g.mapa, 5, cores.MARROM)

	//Tela Branca
	ebitenutil.DrawRect(tela, g.mapa.GetX(), g.mapa.GetY(), g.mapa.GetLargura(), g.mapa.GetAltura(), cores.BRANCO)

	//Tela Amarela
	ebitenutil.DrawRect(tela, g.mapa.GetX()-(g.camera.GetX()/utils.PROPORCAO_MINI_MAPA), g.mapa.GetY()-(g.camera.GetY()/utils.PROPORCAO_MINI_MAPA), g.mapa.GetLargura()/2, g.mapa.GetAltura()/2, cores.AMARELO)

	for _, entidade := range g.Entidades {
		entidade.DesenharMapa(tela, g.mapa.GetX(), g.mapa.GetY())
	}

}
