package personagens

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/componentes"
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
	corpo   *geometria.Retangulo
	Componentes map[string]interface{}
}

func NovoJogador(game interfaces.IGame, n string) *Jogador {
	nEntidade := game.CriarEntidade()

	posicao := geometria.NovoPonto(0, 0)
	nJogador := Jogador{game: game, entidade: nEntidade, nome: n, vida: 2, sangue: 100, cor: color.White, Status: true, posicao: posicao, corpo: geometria.NovoRetangulo(posicao.GetX(), posicao.GetY(), utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO)}
	game.SetEntidade(nEntidade, &nJogador)

	nJogador.AdicionarComponente(componentes.CORPO.String(), nJogador.corpo)

	return &nJogador
}

func (j *Jogador) EstaVivo() bool {
	if j.vida > 0 && j.sangue > 0 {
		j.Status = true
		return j.Status
	}
	j.Status = false
	return j.Status
}

func (j *Jogador) Renasce() {
	if j.EstaVivo() {
		j.TiraUmaVida()

		if j.vida > 0 {
			j.ResetaSangue()
		} else {
			fmt.Println("O jogador " + j.nome + " morreu!")
		}

	} else {
		fmt.Println("O jogador " + j.nome + " já está morto!")
	}
}

func (j *Jogador) TiraUmaVida() {
	if j.vida > 0 {
		j.vida -= 1
	}
}

func (j *Jogador) AcrescentaUmaVida() {
	if j.vida < 3 {
		j.vida += 1
	} else {
		fmt.Println("A vida do jogador " + j.nome + " já está cheia!")
	}
}

func (j *Jogador) ResetaVida() {
	j.vida = 3
}

func (j *Jogador) ResetaSangue() {
	j.sangue = 100
}

func (j *Jogador) PerdeSangue(rit int) {
	j.sangue -= rit
}

func (j *Jogador) GetCorpo() *geometria.Retangulo {
	return j.corpo
}

func (j *Jogador) GetNome() string {
	return j.nome
}
func (j *Jogador) GetPosicao() *geometria.Ponto {
	return j.posicao
}
func (j *Jogador) GetX1() float64 {
	return j.posicao.GetX()
}
func (j *Jogador) GetY1() float64 {
	return j.posicao.GetY()
}
func (j *Jogador) GetX2() float64 {
	return j.posicao.GetX() + utils.JOGADOR_TAMANHO_MUNDO
}
func (j *Jogador) GetY2() float64 {
	return j.posicao.GetY() + utils.JOGADOR_TAMANHO_MUNDO
}
func (j *Jogador) GetLargura() float64 {
	return utils.JOGADOR_TAMANHO_MUNDO
}
func (j *Jogador) GetAltura() float64 {
	return utils.JOGADOR_TAMANHO_MUNDO
}
func (j *Jogador) GetCor() color.Color {
	return j.cor
}

func (j *Jogador) GetTipo() string {
	return entidades.JOGADOR.String()
}

func (j *Jogador) SetPosicao(x float64, y float64) {
	j.posicao.SetPosicao(x, y)
	j.corpo.SetX(x)
	j.corpo.SetY(y)
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

	origemX := j.GetX1()
	origemY := j.GetY1()

	// 1. Criamos a cópia estável para servir de filtro de auto-colisão no ECS
	corpoDeFiltro := geometria.NovoRetangulo(origemX, origemY, utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO)

	// --- PROCESSAMENTO DO EIXO X (Pixel por Pixel) ---
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		passoX := 1.0
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			passoX = -1.0
		}

		// Transforma o speed em inteiro para saber quantos passos de 1 pixel dar
		totalPassosX := int(speed)
		for i := 0; i < totalPassosX; i++ {
			proximoX := origemX + passoX
			testeCorpoX := geometria.NovoRetangulo(proximoX, origemY, utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO)

			// Verifica se o PRÓXIMO pixel está livre
			if j.game.GetMundo().EstaNaMargemInterna(testeCorpoX, utils.JOGADOR_TAMANHO_MUNDO) &&
				!j.game.VaiColidir(corpoDeFiltro, testeCorpoX) {
				origemX = proximoX // Avança 1 pixel com segurança
			} else {
				break // Bateu seco! Para o loop imediatamente e cola no obstáculo
			}
		}
	}

	// --- PROCESSAMENTO DO EIXO Y (Pixel por Pixel) ---
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		passoS := 1.0
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			passoS = -1.0
		}

		// Transforma o speed em inteiro para saber quantos passos de 1 pixel dar
		totalPassosY := int(speed)
		for i := 0; i < totalPassosY; i++ {
			proximoY := origemY + passoS
			// Importante: Usa o origemX já processado para validar quinas corretamente
			testeCorpoY := geometria.NovoRetangulo(origemX, proximoY, utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO)

			// Verifica se o PRÓXIMO pixel está livre
			if j.game.GetMundo().EstaNaMargemInterna(testeCorpoY, utils.JOGADOR_TAMANHO_MUNDO) &&
				!j.game.VaiColidir(corpoDeFiltro, testeCorpoY) {
				origemY = proximoY // Avança 1 pixel com segurança
			} else {
				break // Bateu seco! Para o loop imediatamente e cola no obstáculo
			}
		}
	}

	// --- APLICAÇÃO FINAL ---
	// Define a posição final onde o jogador conseguiu chegar sem interceptar nada
	j.SetPosicao(origemX, origemY)
}

func (j *Jogador) Atualizar() {
	if j.EstaVivo() {
		j.Mover()

		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			j.Atira()
		}
	}
}

func (j *Jogador) Desenhar(tela *ebiten.Image) {

	ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX1(), j.game.GetCamera().GetY()+j.GetY1(), utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO, j.GetCor())

	//ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX()+5, j.game.GetCamera().GetY()+j.GetY()+5, JOGADOR_TAMANHO_INTERNO, JOGADOR_TAMANHO_INTERNO, color.White)
	//ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX()+10, j.game.GetCamera().GetY()+j.GetY()+5, JOGADOR_TAMANHO_INTERNO, JOGADOR_TAMANHO_INTERNO, color.White)

	//ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX()+5, j.game.GetCamera().GetY()+j.GetY()+10, JOGADOR_TAMANHO_INTERNO, JOGADOR_TAMANHO_INTERNO, color.White)
	//ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX()+10, j.game.GetCamera().GetY()+j.GetY()+10, JOGADOR_TAMANHO_INTERNO, JOGADOR_TAMANHO_INTERNO, color.White)

}

func (j *Jogador) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawRect(tela, mapaX+(j.GetX1()/config.PROPORCAO_MAPA), mapaY+(j.GetY1()/config.PROPORCAO_MAPA), utils.JOGADOR_TAMANHO_MAPA, utils.JOGADOR_TAMANHO_MAPA, cores.AZUL)

	//ebitenutil.DrawRect(tela, (mapaX + (j.GetX() / utils.PROPORCAO_MAPA) + 1), (mapaY + (j.GetY() / utils.PROPORCAO_MAPA) + 1), JOGADOR_TAMANHO_INTERNO/2, JOGADOR_TAMANHO_INTERNO/2, cores.PRETO)
	//ebitenutil.DrawRect(tela, (mapaX + (j.GetX() / utils.PROPORCAO_MAPA) + 1), (mapaY + (j.GetY() / utils.PROPORCAO_MAPA) + 2), JOGADOR_TAMANHO_INTERNO/2, JOGADOR_TAMANHO_INTERNO/2, cores.PRETO)
	//ebitenutil.DrawRect(tela, (mapaX + (j.GetX() / utils.PROPORCAO_MAPA) + 2), (mapaY + (j.GetY() / utils.PROPORCAO_MAPA) + 1), JOGADOR_TAMANHO_INTERNO/2, JOGADOR_TAMANHO_INTERNO/2, cores.PRETO)
	//ebitenutil.DrawRect(tela, (mapaX + (j.GetX() / utils.PROPORCAO_MAPA) + 2), (mapaY + (j.GetY() / utils.PROPORCAO_MAPA) + 2), JOGADOR_TAMANHO_INTERNO/2, JOGADOR_TAMANHO_INTERNO/2, cores.PRETO)
}

func (e *Jogador) Atira() {
	//j.game.GerarBot(posX, posY)
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
