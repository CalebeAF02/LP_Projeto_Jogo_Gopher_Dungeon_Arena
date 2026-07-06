package cenas

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/assets/imagens"
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/interfaces"

	"github.com/hajimehoshi/ebiten/v2"
)

type CenaMenuIniciar struct {
	game           interfaces.IGame
	fontecache     assets.FonteCache
	aceitaComandos int
}

func (self *CenaMenuIniciar) SetFonteCache(cache assets.FonteCache) {
	self.fontecache = cache
}

func (self *CenaMenuIniciar) GetGame() interfaces.IGame {
	return self.game
}

func (self *CenaMenuIniciar) SetGame(game interfaces.IGame) {
	self.game = game
}

func (self *CenaMenuIniciar) Input() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) && self.aceitaComandos >= 300 {
		//fmt.Println("estou precionando o esc na cena menu !")
		self.game.Sair()
	}

	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		//fmt.Println("estou precionando o enter na cena menu !")

		self.game.IniciarJogo()
	}

	if ebiten.IsKeyPressed(ebiten.KeyP) {
		//fmt.Println("estou precionando o enter na cena menu !")

		self.game.MudarTelaProgresso()
	}
}

func (self *CenaMenuIniciar) Update() error {
	self.Input()
	if self.aceitaComandos <= 400 {
		self.aceitaComandos += 1
	}

	return nil
}

func (self *CenaMenuIniciar) Draw(tela *ebiten.Image) {

	imagens.DesenharImagemPreencherTela(
		tela,
		imagens.ImagemMenu,
		config.JANELA_LARGURA,
		config.JANELA_ALTURA,
	)

	//assets.EscreverTextoCentralizado(tela, cmi.fontecache.Titulo, 180, "GOPHER DUNGEON ARENA")
	assets.EscreverTextoCentralizado(tela, self.fontecache.Normal, 510, "[ ENTER ]            INICIAR")
	assets.EscreverTextoCentralizado(tela, self.fontecache.Normal, 570, "[ P ]            PROGRESSO")
	assets.EscreverTextoCentralizado(tela, self.fontecache.Normal, 640, "[ ESC ]            SAIR")

	//assets.EscreverTextoCentralizado(tela, cmi.fontecache.Rodape, 640, "Linguagens de Programação - LP")
	//assets.EscreverTextoCentralizado(tela, cmi.fontecache.Rodape, 680, "Universidade de Brasilia")
}

func (self *CenaMenuIniciar) GetNome() string {
	return "CENA_MENU_INICIAR"
}
