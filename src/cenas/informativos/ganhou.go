package informativos

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/interfaces"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func InformativoGanhou(cj interfaces.ICenaJogo, tela *ebiten.Image) {

	cinzaTranslucido := color.RGBA{R: 50, G: 50, B: 50, A: 120}

	larguraTela := float64(tela.Bounds().Dx())
	alturaTela := float64(tela.Bounds().Dy())
	ebitenutil.DrawRect(tela, 0, 0, larguraTela, alturaTela, cinzaTranslucido)

	assets.EscreverTextoCentralizadoColorido(
		tela,
		150,
		300,
		"Você Ganhouuuuu !!!",
		cores.VERDE,
	)

	assets.EscreverTextoCentralizadoColorido(
		tela,
		50,
		450,
		"Aperte [ESC] para voltar ao menu iniciar !",
		cores.VERMELHO,
	)

}
