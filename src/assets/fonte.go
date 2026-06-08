package assets

import (
	"Gopher_Dungeon_Arena/src/enum/cores"
	"bytes"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var Fonte *text.GoTextFaceSource

func init() {
	var err error

	Fonte, err = text.NewGoTextFaceSource(
		bytes.NewReader(IcebergTTF),
	)

	if err != nil {
		log.Fatal(err)
	}
}

func EscreverNumero(tela *ebiten.Image, px float64, py float64, numero int, tamanho float64, cor color.Color) {

	titulo := &text.GoTextFace{
		Source: Fonte,
		Size:   tamanho,
	}

	opTitulo := &text.DrawOptions{}
	opTitulo.GeoM.Translate(float64(px), float64(py))
	opTitulo.ColorScale.ScaleWithColor(cores.PRETO)

	texto_valor := strconv.Itoa(numero)

	text.Draw(
		tela,
		texto_valor,
		titulo,
		opTitulo,
	)

}
func EscreverTexto(
	tela *ebiten.Image,
	texto string,
	px float64,
	py float64,
	tamanho float64,
	cor color.Color,
) {

	fonte := &text.GoTextFace{
		Source: Fonte,
		Size:   tamanho,
	}

	op := &text.DrawOptions{}
	op.GeoM.Translate(px, py)
	op.ColorScale.ScaleWithColor(cor)

	text.Draw(
		tela,
		texto,
		fonte,
		op,
	)
}
