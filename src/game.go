package src

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/cenas"
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/nivel"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	CenaCorrente interfaces.ICena
	CenaJogo     interfaces.ICenaJogo
	Progresso    nivel.Progresso
}

func NovoGame() *Game {
	cenaCorrente := cenas.CenaMenuIniciar{}

	game := Game{CenaCorrente: &cenaCorrente, CenaJogo: nil}
	cenaJogo := cenas.NovoCenaJogo(&game)

	cenaCorrente.SetGame(&game)
	cenaCorrente.SetFonteCache(*assets.FonteCacheCriar())
	game.SetCenaJogo(cenaJogo)

	nivel.Converter()

	game.Progresso = nivel.CarregarProgresso()

	if game.Progresso.NivelCorrente == 0 {
		game.Progresso.NivelCorrente = 1
	}

	nivel.SalvarProgresso(game.Progresso)

	return &game
}

func (self *Game) GetCena() interfaces.ICena {
	return self.CenaCorrente
}
func (self *Game) GetCenaJogo() interfaces.ICenaJogo {
	return self.CenaJogo
}

func (self *Game) SetCenaJogo(cj interfaces.ICenaJogo) {
	self.CenaJogo = cj
}
func (self *Game) SetCena(cena interfaces.ICena) {
	self.CenaCorrente = cena
}

func (self *Game) Update() error {
	self.CenaCorrente.Update()

	//fmt.Println(g.CenaCorrente.GetNome())

	return nil
}

func (self *Game) Draw(tela *ebiten.Image) {
	self.CenaCorrente.Draw(tela)
}
func (self *Game) Layout(l, a int) (int, int) {
	return config.JANELA_LARGURA, config.JANELA_ALTURA
}

func (self *Game) IniciarJogo() {
	self.CenaJogo.ReIniciar()
	self.CenaCorrente = self.CenaJogo
}

func (self *Game) ReiniciarMudarTelaMenuIniciar() {

	self.SalvarProgresso()

	cenaCorrente := cenas.CenaMenuIniciar{}
	cenaCorrente.SetFonteCache(*assets.FonteCacheCriar())
	cenaCorrente.SetGame(self)
	self.CenaJogo = cenas.NovoCenaJogo(self)
	self.SetCena(&cenaCorrente)
}

func (self *Game) MudarTelaMenuIniciar() {
	cenaCorrente := cenas.CenaMenuIniciar{}
	cenaCorrente.SetFonteCache(*assets.FonteCacheCriar())
	cenaCorrente.SetGame(self)
	self.SetCena(&cenaCorrente)
}

func (self *Game) MudarTelaProgresso() {
	cenaCorrente := cenas.CenaProgresso{}
	cenaCorrente.SetFonteCache(*assets.FonteCacheCriar())
	cenaCorrente.SetGame(self)
	self.SetCena(&cenaCorrente)
}

func (self *Game) Pausar() {
	cenaPause := cenas.CenaMenuPause{}
	cenaPause.SetGame(self)
	self.SetCena(&cenaPause)
}

func (self *Game) Voltar() {
	self.SetCena(self.GetCenaJogo())
}

func (self *Game) Sair() {
	os.Exit(0)
}

func (self *Game) GetNome() string {
	return "GAME"
}

func (self *Game) SalvarProgresso() {
	nivel.SalvarProgresso(self.Progresso)
}

func (self *Game) GetNivelCorrente() int {
	return self.Progresso.NivelCorrente
}
