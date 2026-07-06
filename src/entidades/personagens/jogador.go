package personagens

import (
	"Gopher_Dungeon_Arena/src/componentes"
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
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
	nJogador.AdicionarComponente(componentes.PONTUACAO.String(), &componentes.Pontuacao{Coletado: 0, Requisito: 1, EntreiNaSaida: false})

	// DEFINIR PONTUACAO DO JOGO AQUI == 3

	nJogador.entidade = &nJogador

	return &nJogador
}

func (self *Jogador) GetID() ecs.EntidadeID {
	return self.entidadeID
}

func (self *Jogador) ObterVida() *componentes.Vida {
	if sangue_comp := self.GetComponente(componentes.VIDA.String()); sangue_comp != nil {
		return sangue_comp.(*componentes.Vida)
	}
	return nil
}
func (self *Jogador) ObterCorpo() *geometria.Retangulo {
	if corpo_comp := self.GetComponente(componentes.CORPO.String()); corpo_comp != nil {
		return corpo_comp.(*geometria.Retangulo)
	}
	return nil
}
func (self *Jogador) ObterNivel() *componentes.Nivel {
	if nivel_comp := self.GetComponente(componentes.NIVEL.String()); nivel_comp != nil {
		return nivel_comp.(*componentes.Nivel)
	}
	return nil
}

func (self *Jogador) EstaVivo() bool {
	resp := self.ObterVida().EstaVivo()
	if !resp {
		self.SetCor(cores.CINZA_CLARO)
	}
	return resp
}

func (self *Jogador) CorrigeSangue() {
	self.ObterVida().CorrigeSangue(self.ObterNivel().Valor)
}

func (self *Jogador) Renasce() {
	self.ObterVida().Renasce(self.ObterNivel().Valor)
}

func (self *Jogador) TiraUmaVida() {
	self.ObterVida().TiraUmaVida()
}

func (self *Jogador) AcrescentaUmaVida() {
	if !self.ObterVida().AcrescentaUmaVida() {
		fmt.Println("A vida do jogador " + self.nome + " já está cheia!")
	}
}

func (self *Jogador) ResetaVida() {
	self.ObterVida().ResetaVida(3)
}

func (self *Jogador) ResetaSangue() {
	self.ObterVida().ResetaSangue(self.ObterNivel().Valor)
}

func (self *Jogador) PerdeSangue(rit int) {
	self.ObterVida().PerdeSangue(rit, self.ObterNivel().Valor)

	if self.ObterVida().Sangue <= 0 {
		self.Renasce()
	}
}

func (self *Jogador) GetEntidade() ecs.Entidade {
	return self.entidade
}

func (self *Jogador) GetNome() string {
	return self.nome
}
func (self *Jogador) GetPosicao() *geometria.Ponto {
	return geometria.NovoPonto(self.ObterCorpo().GetX(), self.ObterCorpo().GetY())
}
func (self *Jogador) GetX1() float64 {
	return self.GetPosicao().GetX()
}
func (self *Jogador) GetY1() float64 {
	return self.GetPosicao().GetY()
}
func (self *Jogador) GetX2() float64 {
	return self.GetPosicao().GetX() + utils.JOGADOR_TAMANHO_MUNDO
}
func (self *Jogador) GetY2() float64 {
	return self.GetPosicao().GetY() + utils.JOGADOR_TAMANHO_MUNDO
}
func (self *Jogador) GetLargura() float64 {
	return utils.JOGADOR_TAMANHO_MUNDO
}
func (self *Jogador) GetAltura() float64 {
	return utils.JOGADOR_TAMANHO_MUNDO
}
func (self *Jogador) GetCor() color.Color {
	return self.cor
}

func (self *Jogador) GetTipo() string {
	return entidades.JOGADOR.String()
}

func (self *Jogador) SetPosicao(x float64, y float64) {
	self.ObterCorpo().SetX(x)
	self.ObterCorpo().SetY(y)
}
func (self *Jogador) SetX(x float64) {
	self.GetPosicao().SetX(x)
}
func (self *Jogador) SetY(y float64) {
	self.GetPosicao().SetY(y)
}
func (self *Jogador) SetCor(cor color.Color) {
	self.cor = cor
}
func (self *Jogador) SetNivel(nivel int) {
	self.ObterNivel().Valor = nivel
	self.CorrigeSangue()
}

func (self *Jogador) Mover() {
	velocidade := float64(utils.JOGADOR_TAMANHO_MUNDO / config.PROPORCAO_MAPA)

	origemX := self.GetX1()
	origemY := self.GetY1()

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
			colisao := self.cenaJogo.GetSistemaColisao().VaiColidir(self.GetTipo(), self.GetEntidade(), corpoDeFiltro, testeCorpoX)
			// Verifica se o PRÓXIMO pixel está livre
			if self.cenaJogo.GetMundo().EstaNaMargemInterna(testeCorpoX, utils.JOGADOR_TAMANHO_MUNDO) &&
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
			colisao := self.cenaJogo.GetSistemaColisao().VaiColidir(self.GetTipo(), self.GetEntidade(), corpoDeFiltro, testeCorpoY)

			// Verifica se o PRÓXIMO pixel está livre
			if self.cenaJogo.GetMundo().EstaNaMargemInterna(testeCorpoY, utils.JOGADOR_TAMANHO_MUNDO) &&
				!colisao.Status {
				origemY = proximoY // Avança 1 pixel com segurança
			} else {
				break // Bateu seco! Para o loop imediatamente e cola no obstáculo
			}
		}
	}

	// --- APLICAÇÃO FINAL ---
	// Define a posição final onde o jogador conseguiu chegar sem interceptar nada
	self.SetPosicao(origemX, origemY)
}

func (self *Jogador) CarregarPontuacao() {
	pontuacaoComp := self.GetComponente(componentes.PONTUACAO.String())
	pontuacao := pontuacaoComp.(*componentes.Pontuacao)

	if pontuacao.Coletado >= pontuacao.Requisito {
		self.cenaJogo.ColetadoTudo(true)
		self.cenaJogo.SetFaltaPontuacao(0)
	} else {
		self.cenaJogo.SetFaltaPontuacao(pontuacao.Requisito - pontuacao.Coletado)
	}

	if pontuacao.EntreiNaSaida {
		self.cenaJogo.EntreiNaSaida()
	}

}

func (self *Jogador) Atualizar() {
	if self.EstaVivo() && !self.cenaJogo.Concluiu() && !self.cenaJogo.EntrouNaSaida() {
		self.Mover()

		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			self.Atira()
		}

		self.CarregarPontuacao()

	}
}

func (self *Jogador) Desenhar(tela *ebiten.Image) {

	if self.cenaJogo.Concluiu() && self.cenaJogo.EntrouNaSaida() {
		return
	}

	if self.entidade == nil {

	} else {
		ebitenutil.DrawRect(tela, self.cenaJogo.GetCamera().GetX()+self.GetX1(), self.cenaJogo.GetCamera().GetY()+self.GetY1()-10, float64(self.ObterVida().Sangue)/5, 5, cores.VERMELHO_ESCURO)

		ebitenutil.DrawRect(tela, self.cenaJogo.GetCamera().GetX()+self.GetX1(), self.cenaJogo.GetCamera().GetY()+self.GetY1(), utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO, self.GetCor())
		ebitenutil.DrawRect(tela, self.cenaJogo.GetCamera().GetX()+self.GetX1()+2, self.cenaJogo.GetCamera().GetY()+self.GetY1()+2, utils.JOGADOR_TAMANHO_MUNDO-4, utils.JOGADOR_TAMANHO_MUNDO-4, cores.BRANCO)
		ebitenutil.DrawRect(tela, self.cenaJogo.GetCamera().GetX()+self.GetX1()+4, self.cenaJogo.GetCamera().GetY()+self.GetY1()+4, utils.JOGADOR_TAMANHO_MUNDO-8, utils.JOGADOR_TAMANHO_MUNDO-8, self.GetCor())

		//ebitenutil.DrawRect(tela, j.cenaJogo.GetCamera().GetX()+j.GetX1(), j.cenaJogo.GetCamera().GetY()+j.GetY1(), utils.JOGADOR_TAMANHO_MUNDO, utils.JOGADOR_TAMANHO_MUNDO, j.GetCor())

		//ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX()+5, j.game.GetCamera().GetY()+j.GetY()+5, JOGADOR_TAMANHO_INTERNO, JOGADOR_TAMANHO_INTERNO, color.White)
		//ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX()+10, j.game.GetCamera().GetY()+j.GetY()+5, JOGADOR_TAMANHO_INTERNO, JOGADOR_TAMANHO_INTERNO, color.White)

		//ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX()+5, j.game.GetCamera().GetY()+j.GetY()+10, JOGADOR_TAMANHO_INTERNO, JOGADOR_TAMANHO_INTERNO, color.White)
		//ebitenutil.DrawRect(tela, j.game.GetCamera().GetX()+j.GetX()+10, j.game.GetCamera().GetY()+j.GetY()+10, JOGADOR_TAMANHO_INTERNO, JOGADOR_TAMANHO_INTERNO, color.White)

	}

}

func (self *Jogador) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawRect(tela, mapaX+(self.GetX1()/config.PROPORCAO_MAPA), mapaY+(self.GetY1()/config.PROPORCAO_MAPA), utils.JOGADOR_TAMANHO_MAPA, utils.JOGADOR_TAMANHO_MAPA, cores.AZUL)

	//ebitenutil.DrawRect(tela, (mapaX + (j.GetX() / utils.PROPORCAO_MAPA) + 1), (mapaY + (j.GetY() / utils.PROPORCAO_MAPA) + 1), JOGADOR_TAMANHO_INTERNO/2, JOGADOR_TAMANHO_INTERNO/2, cores.PRETO)
	//ebitenutil.DrawRect(tela, (mapaX + (j.GetX() / utils.PROPORCAO_MAPA) + 1), (mapaY + (j.GetY() / utils.PROPORCAO_MAPA) + 2), JOGADOR_TAMANHO_INTERNO/2, JOGADOR_TAMANHO_INTERNO/2, cores.PRETO)
	//ebitenutil.DrawRect(tela, (mapaX + (j.GetX() / utils.PROPORCAO_MAPA) + 2), (mapaY + (j.GetY() / utils.PROPORCAO_MAPA) + 1), JOGADOR_TAMANHO_INTERNO/2, JOGADOR_TAMANHO_INTERNO/2, cores.PRETO)
	//ebitenutil.DrawRect(tela, (mapaX + (j.GetX() / utils.PROPORCAO_MAPA) + 2), (mapaY + (j.GetY() / utils.PROPORCAO_MAPA) + 2), JOGADOR_TAMANHO_INTERNO/2, JOGADOR_TAMANHO_INTERNO/2, cores.PRETO)
}

func (self *Jogador) Atira() {
	//j.game.GerarBot(posX, posY)
}

func (self *Jogador) GetComponente(id string) interface{} {
	return self.Componentes[id]
}

func (self *Jogador) AdicionarComponente(id string, comp interface{}) {
	if self.Componentes == nil {
		self.Componentes = make(map[string]interface{})
	}
	self.Componentes[id] = comp
}
func (self *Jogador) ExisteComponente(id string) bool {
	_, existe := self.Componentes[id]
	return existe
}
