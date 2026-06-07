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
	cenaJogo    interfaces.ICenaJogo
	entidadeID  ecs.EntidadeID
	entidade    ecs.Entidade
	nome        string
	cor         color.Color
	Componentes map[string]interface{}
}

func NovoJogador(cj interfaces.ICenaJogo, n string) *Jogador {
	nEntidade := cj.CriarEntidade()
	corpo := geometria.NovoRetangulo(0, 0, utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO)
	nJogador := Jogador{cenaJogo: cj, entidadeID: nEntidade, nome: n, cor: color.White}
	cj.SetEntidade(nEntidade, &nJogador)

	nJogador.AdicionarComponente(componentes.CORPO.String(), corpo)
	nJogador.AdicionarComponente(componentes.VIDA.String(), &componentes.Vida{TipoOrganismo: entidades.JOGADOR.String(), Status: true, Quantidade: 3, Sangue: 100})
	nJogador.AdicionarComponente(componentes.NIVEL.String(), &componentes.Nivel{Valor: 1, Progressao: 0})

	nJogador.entidade = &nJogador

	return &nJogador
}

func (j *Jogador) GetID() ecs.EntidadeID {
	return j.entidadeID
}

func (j *Jogador) ObterVida() *componentes.Vida {
	if sangue_comp := j.GetComponente(componentes.VIDA.String()); sangue_comp != nil {
		return sangue_comp.(*componentes.Vida)
	}
	return nil
}
func (j *Jogador) ObterCorpo() *geometria.Retangulo {
	if corpo_comp := j.GetComponente(componentes.CORPO.String()); corpo_comp != nil {
		return corpo_comp.(*geometria.Retangulo)
	}
	return nil
}
func (j *Jogador) ObterNivel() *componentes.Nivel {
	if nivel_comp := j.GetComponente(componentes.NIVEL.String()); nivel_comp != nil {
		return nivel_comp.(*componentes.Nivel)
	}
	return nil
}

func (j *Jogador) EstaVivo() bool {
	resp := j.ObterVida().EstaVivo()
	if !resp {
		j.SetCor(cores.CINZA_CLARO)
	}
	return resp
}

func (j *Jogador) CorrigeSangue() {
	j.ObterVida().CorrigeSangue(j.ObterNivel().Valor)
}

func (j *Jogador) Renasce() {
	j.ObterVida().Renasce(j.ObterNivel().Valor)
}

func (j *Jogador) TiraUmaVida() {
	j.ObterVida().TiraUmaVida()
}

func (j *Jogador) AcrescentaUmaVida() {
	if !j.ObterVida().AcrescentaUmaVida() {
		fmt.Println("A vida do jogador " + j.nome + " já está cheia!")
	}
}

func (j *Jogador) ResetaVida() {
	j.ObterVida().ResetaVida(3)
}

func (j *Jogador) ResetaSangue() {
	j.ObterVida().ResetaSangue(j.ObterNivel().Valor)
}

func (j *Jogador) PerdeSangue(rit int) {
	j.ObterVida().PerdeSangue(rit, j.ObterNivel().Valor)

	if j.ObterVida().Sangue <= 0 {
		j.Renasce()
	}
}

func (j *Jogador) GetEntidade() ecs.Entidade {
	return j.entidade
}

func (j *Jogador) GetNome() string {
	return j.nome
}
func (j *Jogador) GetPosicao() *geometria.Ponto {
	return geometria.NovoPonto(j.ObterCorpo().GetX(), j.ObterCorpo().GetY())
}
func (j *Jogador) GetX1() float64 {
	return j.GetPosicao().GetX()
}
func (j *Jogador) GetY1() float64 {
	return j.GetPosicao().GetY()
}
func (j *Jogador) GetX2() float64 {
	return j.GetPosicao().GetX() + utils.JOGADOR_TAMANHO_MUNDO
}
func (j *Jogador) GetY2() float64 {
	return j.GetPosicao().GetY() + utils.JOGADOR_TAMANHO_MUNDO
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
	j.ObterCorpo().SetX(x)
	j.ObterCorpo().SetY(y)
}
func (j *Jogador) SetX(x float64) {
	j.GetPosicao().SetX(x)
}
func (j *Jogador) SetY(y float64) {
	j.GetPosicao().SetY(y)
}
func (j *Jogador) SetCor(cor color.Color) {
	j.cor = cor
}
func (j *Jogador) SetNivel(nivel int) {
	j.ObterNivel().Valor = nivel
	j.CorrigeSangue()
}

func (j *Jogador) Mover() {
	velocidade := float64(utils.JOGADOR_TAMANHO_MUNDO / config.PROPORCAO_MAPA)

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
		totalPassosX := int(velocidade)
		for i := 0; i < totalPassosX; i++ {
			proximoX := origemX + passoX
			testeCorpoX := geometria.NovoRetangulo(proximoX, origemY, utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO)
			colisao := j.cenaJogo.GetSistemaColisao().VaiColidir(j.GetTipo(), j.GetEntidade(), corpoDeFiltro, testeCorpoX)
			// Verifica se o PRÓXIMO pixel está livre
			if j.cenaJogo.GetMundo().EstaNaMargemInterna(testeCorpoX, utils.JOGADOR_TAMANHO_MUNDO) &&
				!colisao.Status {
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
		totalPassosY := int(velocidade)
		for i := 0; i < totalPassosY; i++ {
			proximoY := origemY + passoS
			// Importante: Usa o origemX já processado para validar quinas corretamente
			testeCorpoY := geometria.NovoRetangulo(origemX, proximoY, utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO)
			colisao := j.cenaJogo.GetSistemaColisao().VaiColidir(j.GetTipo(), j.GetEntidade(), corpoDeFiltro, testeCorpoY)

			// Verifica se o PRÓXIMO pixel está livre
			if j.cenaJogo.GetMundo().EstaNaMargemInterna(testeCorpoY, utils.JOGADOR_TAMANHO_MUNDO) &&
				!colisao.Status {
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
	ebitenutil.DrawRect(tela, j.cenaJogo.GetCamera().GetX()+j.GetX1(), j.cenaJogo.GetCamera().GetY()+j.GetY1()-10, float64(j.ObterVida().Sangue)/5, 5, cores.VERMELHO_ESCURO)

	ebitenutil.DrawRect(tela, j.cenaJogo.GetCamera().GetX()+j.GetX1(), j.cenaJogo.GetCamera().GetY()+j.GetY1(), utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO, j.GetCor())
	ebitenutil.DrawRect(tela, j.cenaJogo.GetCamera().GetX()+j.GetX1()+2, j.cenaJogo.GetCamera().GetY()+j.GetY1()+2, utils.JOGADOR_TAMANHO_MUNDO-4, utils.JOGADOR_TAMANHO_MUNDO-4, cores.BRANCO)
	ebitenutil.DrawRect(tela, j.cenaJogo.GetCamera().GetX()+j.GetX1()+4, j.cenaJogo.GetCamera().GetY()+j.GetY1()+4, utils.JOGADOR_TAMANHO_MUNDO-8, utils.JOGADOR_TAMANHO_MUNDO-8, j.GetCor())

	//ebitenutil.DrawRect(tela, j.cenaJogo.GetCamera().GetX()+j.GetX1(), j.cenaJogo.GetCamera().GetY()+j.GetY1(), utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO, j.GetCor())

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
func (e *Jogador) ExisteComponente(id string) bool {
	_, existe := e.Componentes[id]
	return existe
}
