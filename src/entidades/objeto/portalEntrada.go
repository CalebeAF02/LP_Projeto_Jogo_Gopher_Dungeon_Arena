package objeto

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/componentes"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"fmt"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type PortalEntrada struct {
	cenaJogo      interfaces.ICenaJogo
	entidadeID    ecs.EntidadeID
	entidade      ecs.Entidade
	Id            int64
	cor           color.Color
	posicao       *geometria.Ponto
	corpo         *geometria.Retangulo
	anguloRotacao float64 // Adicione este campo
	Componentes   map[string]interface{}
}

func NovoPortalEntrada(cj interfaces.ICenaJogo, id int64) *PortalEntrada {

	nEntidade := cj.CriarEntidade()
	posicao := geometria.NovoPonto(0, 0)
	nBot := PortalEntrada{cenaJogo: cj, entidadeID: nEntidade, Id: id, cor: cores.BRANCO, posicao: posicao, corpo: geometria.NovoRetangulo(posicao.GetX(), posicao.GetY(), utils.PORTAL_ENTRADA_TAMANHO, utils.PORTAL_ENTRADA_TAMANHO)}

	cj.SetEntidade(nEntidade, &nBot)
	nBot.AdicionarComponente(componentes.CORPO.String(), nBot.corpo)
	nBot.AdicionarComponente(componentes.ENVIANDO_TELETRANSPORTE.String(), &componentes.EnviandoTeletransporte{TemBot: false, Bot: nil})

	nBot.entidade = &nBot
	return &nBot
}

func (j *PortalEntrada) GetID() ecs.EntidadeID {
	return j.entidadeID
}

func (e *PortalEntrada) GetComponente(id string) interface{} {
	return e.Componentes[id]
}

func (e *PortalEntrada) AdicionarComponente(id string, comp interface{}) {
	if e.Componentes == nil {
		e.Componentes = make(map[string]interface{})
	}
	e.Componentes[id] = comp
}
func (e *PortalEntrada) AlterarComponente(id string, comp interface{}) {
	e.Componentes[id] = comp
}

func (e *PortalEntrada) ExisteComponente(id string) bool {
	_, existe := e.Componentes[id]
	return existe
}

func (b *PortalEntrada) Atualizar() {

	b.anguloRotacao += 0.002

	// Mantém o progresso sempre dentro do limite de 0.0 a 1.0 de forma eterna
	if b.anguloRotacao >= 1.0 {
		b.anguloRotacao -= 1.0
	}

	if b.ObterTeleTransporte().TemBot {
		if b.ObterTeleTransporte().Contagem > 0 {
			fmt.Printf("\t ++ Vou fazer teletransporte\n")
			b.ObterTeleTransporte().Contagem -= 1

			if b.ObterTeleTransporte().Contagem == 0 {
				fmt.Printf("\t ++ Teletransportar AGORAAAAAAAAAAAAAAA\n")

				//portal_saida_corpo_comp := b.ObterTeleTransporte().ConectadoSaida.GetComponente(componentes.CORPO.String())
				//portal_saida := portal_saida_corpo_comp.(*geometria.Retangulo)

				portal_saida_recebendo_teletransporte_comp := b.ObterTeleTransporte().ConectadoSaida.GetComponente(componentes.RECEBENDO_TELETRANSPORTE.String())
				portal_recebendo_teletransporte := portal_saida_recebendo_teletransporte_comp.(*componentes.RecebendoTeletransporte)

				if !portal_recebendo_teletransporte.TemBot {

					portal_recebendo_teletransporte.TemBot = true
					portal_recebendo_teletransporte.Contagem = 200

					botTele := b.ObterTeleTransporte().Bot

					portal_recebendo_teletransporte.Bot = botTele
					b.ObterTeleTransporte().TemBot = false

				} else {
					fmt.Printf("\t Nao tem como sair do teletransporte !!!")
					b.ObterTeleTransporte().Contagem = 200

				}

			}

		}
	}

}

func (b *PortalEntrada) ObterTeleTransporte() *componentes.EnviandoTeletransporte {
	if tele_comp := b.GetComponente(componentes.ENVIANDO_TELETRANSPORTE.String()); tele_comp != nil {
		return tele_comp.(*componentes.EnviandoTeletransporte)
	}
	return nil
}

func (b *PortalEntrada) ConectarSaida(portalSaida *PortalSaida) {
	b.ObterTeleTransporte().ConectadoSaida = portalSaida.entidade
}

func (b *PortalEntrada) Desenhar(tela *ebiten.Image) {
	// 1. Pega as coordenadas X e Y da tela considerando a câmera
	posXX := b.cenaJogo.GetCamera().GetX() + b.GetX1()
	posY := b.cenaJogo.GetCamera().GetY() + b.GetY1()
	tamanho := float32(utils.PORTAL_ENTRADA_TAMANHO)

	// 2. Desenha o quadrado maior original (Laranja)
	ebitenutil.DrawRect(tela, posXX, posY, float64(tamanho), float64(tamanho), cores.LARANJA)

	// 3. Configurações dos quadradinhos das quinas
	tamQuina := float64(tamanho / 5.0)

	// Margens externas exatas onde os quadradinhos devem deslizar
	minX := posXX - tamQuina
	maxX := posXX + float64(tamanho)
	minY := posY - tamQuina
	maxY := posY + float64(tamanho)

	// Desenha as 4 quinas
	for i := 0; i < 4; i++ {
		// Distribui as 4 quinas igualmente
		progresso := b.anguloRotacao + (float64(i) * 0.25)
		if progresso >= 1.0 {
			progresso -= 1.0
		}

		var quinaX, quinaY float64

		// Máquina de estados para garantir o trilho perfeito do quadrado
		if progresso < 0.25 { // 1. Lado Superior (Esquerda para a Direita)
			t := progresso / 0.25
			if t > 1.0 {
				t = 1.0
			}
			quinaX = minX + (t * (maxX - minX))
			quinaY = minY
		} else if progresso < 0.50 { // 2. Lado Direito (Cima para Baixo)
			t := (progresso - 0.25) / 0.25
			if t > 1.0 {
				t = 1.0
			}
			quinaX = maxX
			quinaY = minY + (t * (maxY - minY))
		} else if progresso < 0.75 { // 3. Lado Inferior (Direita para a Esquerda)
			t := (progresso - 0.50) / 0.25
			if t > 1.0 {
				t = 1.0
			}
			quinaX = maxX - (t * (maxX - minX))
			quinaY = maxY
		} else { // 4. Lado Esquerdo (Baixo para Cima)
			t := (progresso - 0.75) / 0.25
			if t > 1.0 {
				t = 1.0
			}
			quinaX = minX
			quinaY = maxY - (t * (maxY - minY))
		}

		// A) Desenha o quadradinho preto que se move
		ebitenutil.DrawRect(tela, quinaX, quinaY, tamQuina, tamQuina, cores.PRETO)

		// B) CALCULA O CENTRO E O RAIO DO CÍRCULO INTERNO DA QUINA
		// Raio do circulozinho (metade da quina dividida por 2 para deixar uma borda preta)
		raioQuina := float32(tamQuina / 4.0)

		// Centro real: posição inicial da quina + metade do tamanho dela
		centroQuinaX := float32(quinaX + (tamQuina / 2.0))
		centroQuinaY := float32(quinaY + (tamQuina / 2.0))

		// C) Desenha o círculo laranja centralizado dentro do quadradinho preto
		vector.DrawFilledCircle(tela, centroQuinaX, centroQuinaY, raioQuina, cores.LARANJA, true)
	}

	// 4. Desenha o círculo preto centralizado estático no meio do portal
	raioCirculo := tamanho / 4.0
	centroX := float32(posXX) + (tamanho / 2.0)
	centroY := float32(posY) + (tamanho / 2.0)
	vector.DrawFilledCircle(tela, centroX, centroY, raioCirculo, cores.PRETO, true)

	portalTransporte_com := b.GetComponente(componentes.ENVIANDO_TELETRANSPORTE.String())
	teletransporte := portalTransporte_com.(*componentes.EnviandoTeletransporte)

	if teletransporte.TemBot {
		vector.DrawFilledCircle(tela, centroX, centroY, raioCirculo, cores.AMARELO, true)

		titulo := &text.GoTextFace{
			Source: assets.Fonte,
			Size:   15,
		}

		opTitulo := &text.DrawOptions{}
		opTitulo.GeoM.Translate(float64(centroX)-10, float64(centroY)-10)
		opTitulo.ColorScale.ScaleWithColor(cores.PRETO)

		texto_valor := strconv.Itoa(teletransporte.Contagem)

		text.Draw(
			tela,
			texto_valor,
			titulo,
			opTitulo,
		)

	}

}

func (b *PortalEntrada) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawRect(tela, mapaX+(b.GetX1()/config.PROPORCAO_MAPA), mapaY+(b.GetY1()/config.PROPORCAO_MAPA), utils.BOT_TAMANHO_MAPA, utils.BOT_TAMANHO_MAPA, cores.VERMELHO)
}

func (b *PortalEntrada) GetX1() float64 {
	return b.posicao.GetX()
}

func (b *PortalEntrada) GetY1() float64 {
	return b.posicao.GetY()
}

func (b *PortalEntrada) GetX2() float64 {
	return b.posicao.GetX() + utils.BOT_TAMANHO_MUNDO
}

func (b *PortalEntrada) GetY2() float64 {
	return b.posicao.GetY() + utils.BOT_TAMANHO_MUNDO
}
func (b *PortalEntrada) GetLargura() float64 {
	return utils.BOT_TAMANHO_MUNDO
}

func (b *PortalEntrada) GetTipo() string {
	return entidades.PORTAL_ENTRADA.String()
}

func (b *PortalEntrada) SetPosicao(x float64, y float64) {
	b.posicao.SetPosicao(x, y)
	b.corpo.SetX(x)
	b.corpo.SetY(y)
}
