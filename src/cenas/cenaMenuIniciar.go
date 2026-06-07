package cenas

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/interfaces"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type CenaMenuIniciar struct {
	game interfaces.IGame
}

func NovaCenaMenuIniciar(game interfaces.IGame) *CenaMenuIniciar {
	return &CenaMenuIniciar{game: game}
}

func (cmi *CenaMenuIniciar) GetGame() interfaces.IGame {
	return cmi.game
}

func (cmi *CenaMenuIniciar) SetGame(game interfaces.IGame) {
	cmi.game = game
}

func (cmi *CenaMenuIniciar) Input() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		fmt.Println("estou precionando o esc na cena menu !")

		cmi.game.Sair()
	}

	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		fmt.Println("estou precionando o enter na cena menu !")

		cmi.game.IniciarJogo()
	}
}
func (cmi *CenaMenuIniciar) Update() error {
	cmi.Input()
	return nil

}

func (cmi *CenaMenuIniciar) Draw(tela *ebiten.Image) {

	titulo := &text.GoTextFace{
		Source: assets.Fonte,
		Size:   72,
	}

	subtitulo := &text.GoTextFace{
		Source: assets.Fonte,
		Size:   20,
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
	opTitulo.GeoM.Translate(180, 180)

	text.Draw(
		tela,
		"GOPHER DUNGEON ARENA",
		titulo,
		opTitulo,
	)

	// Subtítulo
	opSub := &text.DrawOptions{}
	opSub.GeoM.Translate(420, 250)

	text.Draw(
		tela,
		"Arena Paralela de Sobrevivencia",
		subtitulo,
		opSub,
	)

	// Menu
	opJogar := &text.DrawOptions{}
	opJogar.GeoM.Translate(470, 390)

	text.Draw(
		tela,
		"[ ENTER ]  INICIAR",
		menu,
		opJogar,
	)

	opSair := &text.DrawOptions{}
	opSair.GeoM.Translate(500, 450)

	text.Draw(
		tela,
		"[ ESC ]    SAIR",
		menu,
		opSair,
	)

	// Rodapé
	opRodape := &text.DrawOptions{}
	opRodape.GeoM.Translate(420, 680)

	text.Draw(
		tela,
		"Universidade de Brasilia",
		rodape,
		opRodape,
	)
}

func (cmi *CenaMenuIniciar) GetNome() string {
	return "CENA_MENU_INICIAR"
}
