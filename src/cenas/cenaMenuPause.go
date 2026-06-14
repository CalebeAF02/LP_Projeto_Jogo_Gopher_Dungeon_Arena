package cenas

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/interfaces"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type CenaMenuPause struct {
	game interfaces.IGame
}

func (cp *CenaMenuPause) SetFonteCache(cache assets.FonteCache) {
}

func (cmp *CenaMenuPause) GetGame() interfaces.IGame {
	return cmp.game
}

func (cmp *CenaMenuPause) SetGame(game interfaces.IGame) {
	cmp.game = game
}

func (cmp *CenaMenuPause) Voltar() {
	cmp.game.Voltar()
}

func (cmp *CenaMenuPause) Sair() {
	cmp.game.Sair()
}

func (cmp *CenaMenuPause) Input() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		//fmt.Println("estou precionando o esc na cena pause !")
		cmp.Sair()
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		//fmt.Println("estou precionando o space na cena pause !")
		cmp.Voltar()
	}
}

func (cmp *CenaMenuPause) Update() error {
	cmp.Input()
	return nil
}

func (cmp *CenaMenuPause) Draw(tela *ebiten.Image) {

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

func (cmp *CenaMenuPause) GetNome() string {
	return "CENA_MENU_PAUSE"
}
