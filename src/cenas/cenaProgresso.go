package cenas

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/interfaces"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CenaProgresso struct {
	game       interfaces.IGame
	fontecache assets.FonteCache
}

func (cp *CenaProgresso) SetFonteCache(cache assets.FonteCache) {
	cp.fontecache = cache
}

func (cp *CenaProgresso) GetGame() interfaces.IGame {
	return cp.game
}

func (cp *CenaProgresso) SetGame(game interfaces.IGame) {
	cp.game = game
}

func (cp *CenaProgresso) Input() {
	if ebiten.IsKeyPressed(ebiten.KeyV) {
		cp.game.MudarTelaMenuIniciar()
	}
}

func (cp *CenaProgresso) Update() error {
	cp.Input()
	return nil

}

func (cp *CenaProgresso) Draw(tela *ebiten.Image) {

	assets.EscreverTextoCentralizado(tela, cp.fontecache.Titulo, 180, "GOPHER DUNGEON ARENA - NÍVEIS")

	var linha float64 = 180

	cp.DrawNivel(tela, linha, 350, true)
	cp.DrawNivel(tela, linha+200, 350, false)
	cp.DrawNivel(tela, linha+400, 350, false)
	cp.DrawNivel(tela, linha+600, 350, false)
	cp.DrawNivel(tela, linha+800, 350, false)

	cp.DrawNivel(tela, linha, 500, false)
	cp.DrawNivel(tela, linha+200, 500, false)
	cp.DrawNivel(tela, linha+400, 500, false)
	cp.DrawNivel(tela, linha+600, 500, false)
	cp.DrawNivel(tela, linha+800, 500, false)

}

func (cp *CenaProgresso) DrawNivel(tela *ebiten.Image, px float64, py float64, desbloqueado bool) {

	if desbloqueado {
		ebitenutil.DrawRect(tela, px, py, 100, 100, cores.VERDE)
	} else {
		ebitenutil.DrawRect(tela, px, py, 100, 100, cores.VERMELHO)
	}

	ebitenutil.DrawRect(tela, px+10, py+10, 100-20, 100-20, cores.PRETO)

}

func (cp *CenaProgresso) GetNome() string {
	return "CENA_PROGRESSO"
}
