package informativos

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/interfaces"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func InformativoPerdeu(cj interfaces.ICenaJogo, tela *ebiten.Image) {

	cinzaTranslucido := color.RGBA{R: 50, G: 50, B: 50, A: 120}

	larguraTela := float64(tela.Bounds().Dx())
	alturaTela := float64(tela.Bounds().Dy())
	ebitenutil.DrawRect(tela, 0, 0, larguraTela, alturaTela, cinzaTranslucido)

	assets.EscreverTexto(
		tela,
		"Você Morreu !!!",
		(config.JANELA_LARGURA / 9),
		(config.JANELA_ALTURA/10)*2,
		150,
		cores.VERMELHO,
	)

	assets.EscreverTexto(
		tela,
		"Teve uma sequência de "+cj.GetContadorMortos()+" bots mortos",
		(config.JANELA_LARGURA/9)*2,
		(config.JANELA_ALTURA/10)*6,
		50,
		cores.VERMELHO,
	)

}
