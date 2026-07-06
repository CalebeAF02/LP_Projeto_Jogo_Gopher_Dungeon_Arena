package cenas

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/interfaces"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	CONCLUIDO    = 3
	DESBLOQUEADO = 2
	BLOQUEADO    = 1
)

type CenaProgresso struct {
	game       interfaces.IGame
	fontecache assets.FonteCache
}

func (self *CenaProgresso) SetFonteCache(cache assets.FonteCache) {
	self.fontecache = cache
}

func (self *CenaProgresso) GetGame() interfaces.IGame {
	return self.game
}

func (self *CenaProgresso) SetGame(game interfaces.IGame) {
	self.game = game
}

func (self *CenaProgresso) Input() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		self.game.MudarTelaMenuIniciar()
	}
}

func (self *CenaProgresso) Update() error {
	self.Input()
	return nil

}

func (self *CenaProgresso) Draw(tela *ebiten.Image) {

	assets.EscreverTextoCentralizado(tela, self.fontecache.Titulo, 180, "GOPHER DUNGEON ARENA - NÍVEIS")

	var linha float64 = 180

	nivel := self.game.GetNivelCorrente()

	for i := 1; i <= 10; i++ {
		x := linha + float64((i-1)%5)*200.0
		y := 350.0 + float64((i-1)/5)*150.0

		if i < nivel {
			// níveis já concluídos → verde
			self.DrawNivel(tela, x, y, i, CONCLUIDO)
		} else if i == nivel {
			// nível corrente → destaque
			self.DrawNivel(tela, x, y, i, DESBLOQUEADO)
		} else {
			// níveis futuros → normal
			self.DrawNivel(tela, x, y, i, BLOQUEADO)
		}
	}

}

func (self *CenaProgresso) DrawNivel(tela *ebiten.Image, px float64, py float64, nivel int, modo int) {

	if modo == DESBLOQUEADO {
		ebitenutil.DrawRect(tela, px, py, 100, 100, cores.AMBAR)
	} else if modo == CONCLUIDO {
		ebitenutil.DrawRect(tela, px, py, 100, 100, cores.VERDE)
	} else {
		ebitenutil.DrawRect(tela, px, py, 100, 100, cores.VERMELHO)
	}

	ebitenutil.DrawRect(tela, px+10, py+10, 100-20, 100-20, cores.PRETO)

	rodape := &text.GoTextFace{
		Source: assets.Fonte,
		Size:   80,
	}

	opRodape := &text.DrawOptions{}

	if nivel < 10 {
		opRodape.GeoM.Translate(px+30, py)
	} else {
		opRodape.GeoM.Translate(px+10, py)
	}

	text.Draw(
		tela,
		strconv.Itoa(nivel),
		rodape,
		opRodape,
	)
}

func (self *CenaProgresso) GetNome() string {
	return "CENA_PROGRESSO"
}
