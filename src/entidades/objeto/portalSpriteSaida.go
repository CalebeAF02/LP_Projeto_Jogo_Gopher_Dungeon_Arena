package objeto

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/componentes"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type PortalSpriteSaida struct {
	cenaJogo      interfaces.ICenaJogo
	entidadeID    ecs.EntidadeID
	entidade      ecs.Entidade
	Id            int64
	cor           color.Color
	posicao       *geometria.Ponto
	corpo         *geometria.Retangulo
	Componentes   map[string]interface{}
	spriteFatiado *ebiten.Image // Armazena apenas o pedaço recortado deste portal (Verde Claro)
}

func NovoPortalSpriteSaida(cj interfaces.ICenaJogo, id int64, spritePronto *ebiten.Image) *PortalSpriteSaida {
	nEntidade := cj.CriarEntidade()
	posicao := geometria.NovoPonto(0, 0)

	nBot := PortalSpriteSaida{
		cenaJogo:      cj,
		entidadeID:    nEntidade,
		Id:            id,
		cor:           cores.BRANCO,
		posicao:       posicao,
		corpo:         geometria.NovoRetangulo(posicao.GetX(), posicao.GetY(), utils.PORTAL_ENTRADA_TAMANHO, utils.PORTAL_ENTRADA_TAMANHO),
		spriteFatiado: spritePronto, // Guarda o pedaço da imagem recebido diretamente do gerenciador
	}

	// O método extrairSprite foi removido daqui porque o pacote 'portais' já entrega a imagem cortada!

	cj.SetEntidade(nEntidade, &nBot)
	nBot.AdicionarComponente(componentes.CORPO.String(), nBot.corpo)

	nBot.entidade = &nBot
	return &nBot
}

func (self *PortalSpriteSaida) GetID() ecs.EntidadeID {
	return self.entidadeID
}

func (self *PortalSpriteSaida) GetComponente(id string) interface{} {
	return self.Componentes[id]
}

func (self *PortalSpriteSaida) AdicionarComponente(id string, comp interface{}) {
	if self.Componentes == nil {
		self.Componentes = make(map[string]interface{})
	}
	self.Componentes[id] = comp
}

func (self *PortalSpriteSaida) AlterarComponente(id string, comp interface{}) {
	self.Componentes[id] = comp
}

func (self *PortalSpriteSaida) ExisteComponente(id string) bool {
	_, existe := self.Componentes[id]
	return existe
}

func (self *PortalSpriteSaida) Atualizar() {
	// Atualização lógica se necessário. Fica vazio e super leve!
}

func (self *PortalSpriteSaida) Desenhar(tela *ebiten.Image) {
	if self.spriteFatiado == nil {
		return
	}

	// 1. Pega as coordenadas X e Y da tela considerando a câmera
	posXX := self.cenaJogo.GetCamera().GetX() + self.GetX1()
	posY := self.cenaJogo.GetCamera().GetY() + self.GetY1()

	// 2. Configura as opções de desenho da imagem
	op := &ebiten.DrawImageOptions{}

	// Ajusta a escala do sprite para o tamanho padrão configurado no seu motor
	bounds := self.spriteFatiado.Bounds()
	escalaX := utils.PORTAL_ENTRADA_TAMANHO / float64(bounds.Dx())
	escalaY := utils.PORTAL_ENTRADA_TAMANHO / float64(bounds.Dy())
	op.GeoM.Scale(escalaX, escalaY)

	// Move para a posição corrigida do mundo
	op.GeoM.Translate(posXX, posY)

	// 3. Renderiza a textura na tela através da GPU de forma super rápida
	tela.DrawImage(self.spriteFatiado, op)
}

func (self *PortalSpriteSaida) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawRect(tela, mapaX+(self.GetX1()/config.PROPORCAO_MAPA), mapaY+(self.GetY1()/config.PROPORCAO_MAPA), utils.BOT_TAMANHO_MAPA, utils.BOT_TAMANHO_MAPA, cores.VERMELHO)
}

func (self *PortalSpriteSaida) GetX1() float64 {
	return self.posicao.GetX()
}

func (self *PortalSpriteSaida) GetY1() float64 {
	return self.posicao.GetY()
}

func (self *PortalSpriteSaida) GetX2() float64 {
	return self.posicao.GetX() + utils.BOT_TAMANHO_MUNDO
}

func (self *PortalSpriteSaida) GetY2() float64 {
	return self.posicao.GetY() + utils.BOT_TAMANHO_MUNDO
}

func (self *PortalSpriteSaida) GetLargura() float64 {
	return utils.BOT_TAMANHO_MUNDO
}

func (self *PortalSpriteSaida) GetTipo() string {
	return entidades.PORTAL_SAIDA.String()
}

func (self *PortalSpriteSaida) SetPosicao(x float64, y float64) {
	self.posicao.SetPosicao(x, y)
	self.corpo.SetX(x)
	self.corpo.SetY(y)
}
