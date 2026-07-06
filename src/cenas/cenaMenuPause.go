package cenas

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/interfaces"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type CenaMenuPause struct {
	game           interfaces.IGame
	aceitaComandos int
}

func (self *CenaMenuPause) SetFonteCache(cache assets.FonteCache) {
}

func (self *CenaMenuPause) GetGame() interfaces.IGame {
	return self.game
}

func (self *CenaMenuPause) SetGame(game interfaces.IGame) {
	self.game = game
}

func (self *CenaMenuPause) Voltar() {
	self.game.Voltar()
}

func (self *CenaMenuPause) Sair() {
	self.game.MudarTelaMenuIniciar()
}

func (self *CenaMenuPause) Input() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) && self.aceitaComandos >= 100 {
		//fmt.Println("estou precionando o esc na cena pause !")
		self.Sair()
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && self.aceitaComandos >= 100 {
		//fmt.Println("estou precionando o space na cena pause !")
		self.Voltar()
	}
}

func (self *CenaMenuPause) Update() error {
	self.Input()
	if self.aceitaComandos <= 100 {
		self.aceitaComandos += 1
	}
	return nil
}

func (self *CenaMenuPause) Draw(tela *ebiten.Image) {

	titulo := &text.GoTextFace{
		Source: assets.Fonte,
		Size:   72,
	}

	menu := &text.GoTextFace{
		Source: assets.Fonte,
		Size:   30,
	}

	rodape := &text.GoTextFace{
		Source: assets.Fonte,
		Size:   16,
	}

	// Título
	opTitulo := &text.DrawOptions{}
	opTitulo.GeoM.Translate(470, 200)

	text.Draw(
		tela,
		"PAUSE",
		titulo,
		opTitulo,
	)

	// Continuar
	opContinuar := &text.DrawOptions{}
	opContinuar.GeoM.Translate(420, 380)

	text.Draw(
		tela,
		"[ SPACE ]  CONTINUAR",
		menu,
		opContinuar,
	)

	// Sair
	opSair := &text.DrawOptions{}
	opSair.GeoM.Translate(470, 450)

	text.Draw(
		tela,
		"[ ESC ]    SAIR",
		menu,
		opSair,
	)

	// Rodapé
	opRodape := &text.DrawOptions{}
	opRodape.GeoM.Translate(470, 680)

	text.Draw(
		tela,
		"Jogo Pausado",
		rodape,
		opRodape,
	)
}

func (self *CenaMenuPause) GetNome() string {
	return "CENA_MENU_PAUSE"
}
