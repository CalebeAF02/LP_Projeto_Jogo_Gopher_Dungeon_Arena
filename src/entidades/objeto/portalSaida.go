package objeto

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/componentes"
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"image/color"
	"math"

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
	offsetBarras  float64
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
	if b.ObterRecebendoTeletransporte().TemBot {

		// Avança o progresso da animação
		b.anguloRotacao += 0.002
		if b.anguloRotacao >= 1.0 {
			b.anguloRotacao -= 1.0
		}

		// --- NOVO: calcula deslocamento vertical das barras laterais ---
		// Usamos seno para dar movimento suave (vai e volta)
		// Valor oscila entre -1 e +1
		osc := math.Sin(b.anguloRotacao * 2 * math.Pi)

		// Multiplica por uma amplitude (pixels de deslocamento)
		amplitude := 10.0
		b.offsetBarras = osc * amplitude
		// Agora em Desenhar você usa posY + b.offsetBarras para desenhar as barras

		// --- Lógica de teletransporte continua igual ---
		if b.ObterRecebendoTeletransporte().Contagem > 0 {
			b.ObterRecebendoTeletransporte().Contagem -= 1
		}

		if b.ObterRecebendoTeletransporte().Contagem == 0 {
			bot_corpo_comp := b.ObterRecebendoTeletransporte().Bot.GetComponente(componentes.CORPO.String())
			bot_corpo := bot_corpo_comp.(*geometria.Retangulo)

			liberdade_comp := b.ObterRecebendoTeletransporte().Bot.GetComponente(componentes.ATIVIDADE.String())
			liberdade := liberdade_comp.(*componentes.Atividade)
			liberdade.Acao = componentes.AIVIDADE_MOVIMENTO

			bot_corpo.SetPosicao(b.ObterCorpo().GetX()+70, b.ObterCorpo().GetY()+70)

			b.ObterRecebendoTeletransporte().TemBot = false
		}
	}
}

func (b *PortalSaida) Desenhar(tela *ebiten.Image) {
	// 1. Pega as coordenadas X e Y da tela considerando a câmera
	posXX := b.cenaJogo.GetCamera().GetX() + b.GetX1()
	posY := b.cenaJogo.GetCamera().GetY() + b.GetY1()
	tamanho := float32(utils.PORTAL_ENTRADA_TAMANHO)

	// 2. Desenha o quadrado central (Verde)
	ebitenutil.DrawRect(tela, posXX, posY, float64(tamanho), float64(tamanho), cores.VERDE)

	// 3. Configurações dos retângulos laterais
	larguraRet := float64(tamanho / 5.0) // largura dos retângulos
	alturaRet := float64(tamanho * 1.2)  // altura maior que o quadrado
	offsetY := (alturaRet - float64(tamanho)) / 2.0

	// Esquerda (afasta/aproxima no eixo X)
	ebitenutil.DrawRect(tela,
		posXX-larguraRet-b.offsetBarras, posY-offsetY,
		larguraRet, alturaRet,
		cores.PRETO)

	ebitenutil.DrawRect(tela,
		(posXX-larguraRet-b.offsetBarras)+3, (posY-offsetY)+3,
		larguraRet-6, alturaRet-6,
		cores.VERDE)

	// Direita (afasta/aproxima no eixo X, em sentido contrário)
	ebitenutil.DrawRect(tela,
		posXX+float64(tamanho)+b.offsetBarras, posY-offsetY,
		larguraRet, alturaRet,
		cores.PRETO)

	// 4. Desenha o círculo centralizado no meio do portal
	raioCirculo := tamanho / 4.0
	centroX := float32(posXX) + (tamanho / 2.0)
	centroY := float32(posY) + (tamanho / 2.0)
	vector.DrawFilledCircle(tela, centroX, centroY, raioCirculo, cores.PRETO, true)

	recebendoTeletransporte_comp := b.GetComponente(componentes.RECEBENDO_TELETRANSPORTE.String())
	recebendoTeletransporte := recebendoTeletransporte_comp.(*componentes.RecebendoTeletransporte)

	if recebendoTeletransporte.TemBot {
		vector.DrawFilledCircle(tela, centroX, centroY, raioCirculo, cores.AMARELO, true)
		assets.EscreverNumero(tela, float64(centroX-10), float64(centroY-10), recebendoTeletransporte.Contagem, 15, cores.PRETO)
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
