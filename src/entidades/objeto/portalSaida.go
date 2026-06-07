package objeto

import (
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

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type PortalSaida struct {
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

func NovoPortalSaida(cj interfaces.ICenaJogo, id int64) *PortalSaida {

	nEntidade := cj.CriarEntidade()
	posicao := geometria.NovoPonto(0, 0)
	nBot := PortalSaida{cenaJogo: cj, entidadeID: nEntidade, Id: id, cor: cores.BRANCO, posicao: posicao, corpo: geometria.NovoRetangulo(posicao.GetX(), posicao.GetY(), utils.PORTAL_ENTRADA_TAMANHO, utils.PORTAL_ENTRADA_TAMANHO)}

	cj.SetEntidade(nEntidade, &nBot)
	nBot.AdicionarComponente(componentes.CORPO.String(), nBot.corpo)
	nBot.AdicionarComponente(componentes.RECEBENDO_TELETRANSPORTE.String(), &componentes.RecebendoTeletransporte{TemBot: false, Bot: nil, Contagem: 0})

	nBot.entidade = &nBot
	return &nBot
}

func (j *PortalSaida) GetID() ecs.EntidadeID {
	return j.entidadeID
}

func (e *PortalSaida) GetComponente(id string) interface{} {
	return e.Componentes[id]
}

func (e *PortalSaida) AdicionarComponente(id string, comp interface{}) {
	if e.Componentes == nil {
		e.Componentes = make(map[string]interface{})
	}
	e.Componentes[id] = comp
}
func (e *PortalSaida) AlterarComponente(id string, comp interface{}) {
	e.Componentes[id] = comp
}

func (e *PortalSaida) ExisteComponente(id string) bool {
	_, existe := e.Componentes[id]
	return existe
}

func (b *PortalSaida) ObterCorpo() *geometria.Retangulo {
	if corpo_comp := b.GetComponente(componentes.CORPO.String()); corpo_comp != nil {
		return corpo_comp.(*geometria.Retangulo)
	}
	return nil
}

func (b *PortalSaida) ObterRecebendoTeletransporte() *componentes.RecebendoTeletransporte {
	if tele_comp := b.GetComponente(componentes.RECEBENDO_TELETRANSPORTE.String()); tele_comp != nil {
		return tele_comp.(*componentes.RecebendoTeletransporte)
	}
	return nil
}

func (b *PortalSaida) Atualizar() {
	b.anguloRotacao += 0.002

	// Mantém o progresso sempre dentro do limite de 0.0 a 1.0 de forma eterna
	if b.anguloRotacao >= 1.0 {
		b.anguloRotacao -= 1.0
	}

	if b.ObterRecebendoTeletransporte().TemBot {

		if b.ObterRecebendoTeletransporte().Contagem > 0 {
			b.ObterRecebendoTeletransporte().Contagem -= 1
			fmt.Printf("Recebendo Teletransporte !!!\n")

		}

		if b.ObterRecebendoTeletransporte().Contagem == 0 {

			bot_corpo_comp := b.ObterRecebendoTeletransporte().Bot.GetComponente(componentes.CORPO.String())
			bot_corpo := bot_corpo_comp.(*geometria.Retangulo)

			liberdade_comp := b.ObterRecebendoTeletransporte().Bot.GetComponente(componentes.LIBERDADE.String())
			liberdade := liberdade_comp.(*componentes.Liberdade)
			liberdade.Status = true

			bot_corpo.SetPosicao(b.ObterCorpo().GetX()+70, b.ObterCorpo().GetY()+70)
			//b.ObterTeleTransporte().TemBot = false

			fmt.Printf("\tSai X \n", b.ObterCorpo().GetX()+70)
			fmt.Printf("\tSai Y \n", b.ObterCorpo().GetY()+70)

			b.ObterRecebendoTeletransporte().TemBot = false
			fmt.Printf("Saindo do teletransporte AGORAAAAAAAAAAAAAAAA !!!\n")

		}

	}

}

func (b *PortalSaida) Desenhar(tela *ebiten.Image) {
	// 1. Pega as coordenadas X e Y da tela considerando a câmera
	posXX := b.cenaJogo.GetCamera().GetX() + b.GetX1()
	posY := b.cenaJogo.GetCamera().GetY() + b.GetY1()
	tamanho := float32(utils.PORTAL_ENTRADA_TAMANHO)

	// 2. Desenha o quadrado maior original (Laranja)
	ebitenutil.DrawRect(tela, posXX, posY, float64(tamanho), float64(tamanho), cores.VERDE)

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
		vector.DrawFilledCircle(tela, centroQuinaX, centroQuinaY, raioQuina, cores.VERDE, true)
	}

	// 4. Desenha o círculo preto centralizado estático no meio do portal
	raioCirculo := tamanho / 4.0
	centroX := float32(posXX) + (tamanho / 2.0)
	centroY := float32(posY) + (tamanho / 2.0)
	vector.DrawFilledCircle(tela, centroX, centroY, raioCirculo, cores.PRETO, true)

	recebendoTeletransporte_comp := b.GetComponente(componentes.RECEBENDO_TELETRANSPORTE.String())
	recebendoTeletransporte := recebendoTeletransporte_comp.(*componentes.RecebendoTeletransporte)

	if recebendoTeletransporte.TemBot {
		vector.DrawFilledCircle(tela, centroX, centroY, raioCirculo, cores.AMARELO, true)
	}

}

func (b *PortalSaida) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawRect(tela, mapaX+(b.GetX1()/config.PROPORCAO_MAPA), mapaY+(b.GetY1()/config.PROPORCAO_MAPA), utils.BOT_TAMANHO_MAPA, utils.BOT_TAMANHO_MAPA, cores.VERMELHO)
}

func (b *PortalSaida) GetX1() float64 {
	return b.posicao.GetX()
}

func (b *PortalSaida) GetY1() float64 {
	return b.posicao.GetY()
}

func (b *PortalSaida) GetX2() float64 {
	return b.posicao.GetX() + utils.BOT_TAMANHO_MUNDO
}

func (b *PortalSaida) GetY2() float64 {
	return b.posicao.GetY() + utils.BOT_TAMANHO_MUNDO
}
func (b *PortalSaida) GetLargura() float64 {
	return utils.BOT_TAMANHO_MUNDO
}

func (b *PortalSaida) GetTipo() string {
	return entidades.PORTAL_SAIDA.String()
}

func (b *PortalSaida) SetPosicao(x float64, y float64) {
	b.posicao.SetPosicao(x, y)
	b.corpo.SetX(x)
	b.corpo.SetY(y)
}
