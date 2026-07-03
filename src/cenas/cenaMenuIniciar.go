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
	if ebiten.IsKeyPressed(ebiten.KeyEscape) && cmi.aceitaComandos >= 300 {
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
	if cmi.aceitaComandos <= 400 {
		cmi.aceitaComandos += 1
	}

	return nil
}

func (cmi *CenaMenuIniciar) Draw(tela *ebiten.Image) {

	imagens.DesenharImagemPreencherTela(
		tela,
		imagens.ImagemMenu,
		config.JANELA_LARGURA,
		config.JANELA_ALTURA,
	)

	//assets.EscreverTextoCentralizado(tela, cmi.fontecache.Titulo, 180, "GOPHER DUNGEON ARENA")
	assets.EscreverTextoCentralizado(tela, cmi.fontecache.Normal, 510, "[ ENTER ]            INICIAR")
	assets.EscreverTextoCentralizado(tela, cmi.fontecache.Normal, 570, "[ P ]            PROGRESSO")
	assets.EscreverTextoCentralizado(tela, cmi.fontecache.Normal, 640, "[ ESC ]            SAIR")

	//assets.EscreverTextoCentralizado(tela, cmi.fontecache.Rodape, 640, "Linguagens de Programação - LP")
	//assets.EscreverTextoCentralizado(tela, cmi.fontecache.Rodape, 680, "Universidade de Brasilia")
}

func (cmi *CenaMenuIniciar) GetNome() string {
	return "CENA_MENU_INICIAR"
}
