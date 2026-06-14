package cenas

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/interfaces"

	"github.com/hajimehoshi/ebiten/v2"
)

type CenaMenuIniciar struct {
	game       interfaces.IGame
	fontecache assets.FonteCache
}

func (cmi *CenaMenuIniciar) SetFonteCache(cache assets.FonteCache) {
	cmi.fontecache = cache
}

func (cmi *CenaMenuIniciar) GetGame() interfaces.IGame {
	return cmi.game
}

func (cmi *CenaMenuIniciar) SetGame(game interfaces.IGame) {
	cmi.game = game
}

func (cmi *CenaMenuIniciar) Input() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		//fmt.Println("estou precionando o esc na cena menu !")

		cmi.game.Sair()
	}

	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		//fmt.Println("estou precionando o enter na cena menu !")

		cmi.game.IniciarJogo()
	}

	if ebiten.IsKeyPressed(ebiten.KeyP) {
		//fmt.Println("estou precionando o enter na cena menu !")

		cmi.game.MudarTelaProgresso()
	}
}

func (cmi *CenaMenuIniciar) Update() error {
	cmi.Input()
	return nil

}

func (cmi *CenaMenuIniciar) Draw(tela *ebiten.Image) {

	assets.EscreverTextoCentralizado(tela, cmi.fontecache.Titulo, 180, "GOPHER DUNGEON ARENA")
	assets.EscreverTextoCentralizado(tela, cmi.fontecache.Subtitulo, 250, "Arena Paralela de Sobrevivencia")

	assets.EscreverTextoCentralizado(tela, cmi.fontecache.Normal, 390, "[ ENTER ]  INICIAR")
	assets.EscreverTextoCentralizado(tela, cmi.fontecache.Normal, 450, "[ P ]    PROGRESSO")
	assets.EscreverTextoCentralizado(tela, cmi.fontecache.Normal, 490, "[ ESC ]    SAIR")

	assets.EscreverTextoCentralizado(tela, cmi.fontecache.Rodape, 640, "Linguagens de Programação - LP")
	assets.EscreverTextoCentralizado(tela, cmi.fontecache.Rodape, 680, "Universidade de Brasilia")
}

func (cmi *CenaMenuIniciar) GetNome() string {
	return "CENA_MENU_INICIAR"
}
