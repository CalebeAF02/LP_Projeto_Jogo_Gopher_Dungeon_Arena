package main

import (
	"Gopher_Dungeon_Arena/src"
	"Gopher_Dungeon_Arena/src/config"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := src.NovoGame()
	ebiten.SetWindowSize(config.JANELA_LARGURA, config.JANELA_ALTURA)
	ebiten.SetWindowTitle(config.NOME_JOGO)
	//ebiten.SetTPS(ebiten.SyncWithFPS) // sincroniza lógica com FPS
	//ebiten.SetTPS(1)                  // força 60 ticks por segundo

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
