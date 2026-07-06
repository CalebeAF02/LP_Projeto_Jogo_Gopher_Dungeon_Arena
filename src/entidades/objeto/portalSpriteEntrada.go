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

type PortalSpriteEntrada struct {
	cenaJogo      interfaces.ICenaJogo
	entidadeID    ecs.EntidadeID
	entidade      ecs.Entidade
	Id            int64
	cor           color.Color
	posicao       *geometria.Ponto
	corpo         *geometria.Retangulo
	Componentes   map[string]interface{}
	spriteFatiado *ebiten.Image // Armazena o pedaço já recortado (Laranja) recebido do construtor
}

// NovoPortalSpriteEntrada agora recebe o pedaço pronto da imagem (spritePronto)
func NovoPortalSpriteEntrada(cj interfaces.ICenaJogo, id int64, spritePronto *ebiten.Image) *PortalSpriteEntrada {
	nEntidade := cj.CriarEntidade()
	posicao := geometria.NovoPonto(0, 0)

	nBot := PortalSpriteEntrada{
		cenaJogo:      cj,
		entidadeID:    nEntidade,
		Id:            id,
		cor:           cores.BRANCO,
		posicao:       posicao,
		corpo:         geometria.NovoRetangulo(posicao.GetX(), posicao.GetY(), utils.PORTAL_ENTRADA_TAMANHO, utils.PORTAL_ENTRADA_TAMANHO),
		spriteFatiado: spritePronto, // Guarda o pedaço da imagem diretamente aqui
	}

	// O método extrairSprite foi removido daqui porque o pacote 'portais' já faz isso fora da classe!

	cj.SetEntidade(nEntidade, &nBot)
	nBot.AdicionarComponente(componentes.CORPO.String(), nBot.corpo)

	nBot.entidade = &nBot
	return &nBot
}

func (self *PortalSpriteEntrada) GetID() ecs.EntidadeID {
	return self.entidadeID
}

func (self *PortalSpriteEntrada) GetComponente(id string) interface{} {
	return self.Componentes[id]
}

func (self *PortalSpriteEntrada) AdicionarComponente(id string, comp interface{}) {
	if self.Componentes == nil {
		self.Componentes = make(map[string]interface{})
	}
	self.Componentes[id] = comp
}

func (self *PortalSpriteEntrada) AlterarComponente(id string, comp interface{}) {
	self.Componentes[id] = comp
}

func (self *PortalSpriteEntrada) ExisteComponente(id string) bool {
	_, existe := self.Componentes[id]
	return existe
}

func (self *PortalSpriteEntrada) Atualizar() {
	// A lógica antiga do loop "for" de 4 quinas rodando sumiu.
	// O Update fica limpo e performático porque o portal agora usa o sprite fixo do PNG.
}

func (self *PortalSpriteEntrada) Desenhar(tela *ebiten.Image) {
	if self.spriteFatiado == nil {
		return
	}

	// 1. Pega as coordenadas X e Y da tela considerando a câmera
	posXX := self.cenaJogo.GetCamera().GetX() + self.GetX1()
	posY := self.cenaJogo.GetCamera().GetY() + self.GetY1()

	// 2. Configura as opções de desenho e redimensionamento do sprite
	op := &ebiten.DrawImageOptions{}

	// Trata a escala dinamicamente para casar com a constante de tamanho do seu mundo
	bounds := self.spriteFatiado.Bounds()
	escalaX := utils.PORTAL_ENTRADA_TAMANHO / float64(bounds.Dx())
	escalaY := utils.PORTAL_ENTRADA_TAMANHO / float64(bounds.Dy())
	op.GeoM.Scale(escalaX, escalaY)

	// Move para a posição correta na tela baseada no mundo e na câmera
	op.GeoM.Translate(posXX, posY)

	// 3. Renderiza o sprite na tela (Substitui todos os DrawRect e DrawFilledCircle de antes)
	tela.DrawImage(self.spriteFatiado, op)
}

func (self *PortalSpriteEntrada) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	ebitenutil.DrawRect(tela, mapaX+(self.GetX1()/config.PROPORCAO_MAPA), mapaY+(self.GetY1()/config.PROPORCAO_MAPA), utils.BOT_TAMANHO_MAPA, utils.BOT_TAMANHO_MAPA, cores.VERMELHO)
}

func (self *PortalSpriteEntrada) GetX1() float64 {
	return self.posicao.GetX()
}

func (self *PortalSpriteEntrada) GetY1() float64 {
	return self.posicao.GetY()
}

func (self *PortalSpriteEntrada) GetX2() float64 {
	return self.posicao.GetX() + utils.BOT_TAMANHO_MUNDO
}

func (self *PortalSpriteEntrada) GetY2() float64 {
	return self.posicao.GetY() + utils.BOT_TAMANHO_MUNDO
}

func (self *PortalSpriteEntrada) GetLargura() float64 {
	return utils.BOT_TAMANHO_MUNDO
}

func (self *PortalSpriteEntrada) GetTipo() string {
	return entidades.PORTAL_ENTRADA.String()
}

func (self *PortalSpriteEntrada) SetPosicao(x float64, y float64) {
	self.posicao.SetPosicao(x, y)
	self.corpo.SetX(x)
	self.corpo.SetY(y)
}
