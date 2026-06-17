package assets

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"bytes"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var Fonte *text.GoTextFaceSource

type FonteCache struct {
	Titulo    *text.GoTextFace
	Subtitulo *text.GoTextFace
	Normal    *text.GoTextFace
	Rodape    *text.GoTextFace
}

func FonteCacheCriar() *FonteCache {

	titulo := &text.GoTextFace{
		Source: Fonte,
		Size:   72,
	}

	subtitulo := &text.GoTextFace{
		Source: Fonte,
		Size:   20,
	}

	normal := &text.GoTextFace{
		Source: Fonte,
		Size:   30,
	}

	rodape := &text.GoTextFace{
		Source: Fonte,
		Size:   16,
	}

	return &FonteCache{Titulo: titulo, Subtitulo: subtitulo, Normal: normal, Rodape: rodape}
}

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

func EscreverTextoCentralizado(tela *ebiten.Image, fonte *text.GoTextFace, py float64, texto string) {

	opTexto := &text.DrawOptions{}
	opTexto.PrimaryAlign = text.AlignCenter
	opTexto.SecondaryAlign = text.AlignCenter
	opTexto.GeoM.Translate(config.JANELA_LARGURA/2, py)

	text.Draw(
		tela,
		texto,
		fonte,
		opTexto,
	)

}

func EscreverTextoCentralizadoColorido(tela *ebiten.Image, tamanho float64, py float64, texto string, cor color.Color) {
	opTexto := &text.DrawOptions{}
	opTexto.PrimaryAlign = text.AlignCenter
	opTexto.SecondaryAlign = text.AlignCenter
	opTexto.GeoM.Translate(config.JANELA_LARGURA/2, py)

	fonte := &text.GoTextFace{
		Source: Fonte,
		Size:   tamanho,
	}

	// A mágica acontece aqui: define a cor do texto
	opTexto.ColorScale.ScaleWithColor(cor)

	text.Draw(
		tela,
		texto,
		fonte,
		opTexto,
	)
}

func EscreverTextoLocal(tela *ebiten.Image, fonte *text.GoTextFace, px float64, py float64, texto string) {

	opTexto := &text.DrawOptions{}
	opTexto.GeoM.Translate(px, py)

	text.Draw(
		tela,
		texto,
		fonte,
		opTexto,
	)

}

func EscreverNumeroLocal(tela *ebiten.Image, fonte *text.GoTextFace, px float64, py float64, numero int) {

	opTexto := &text.DrawOptions{}
	opTexto.GeoM.Translate(px, py)

	text.Draw(
		tela,
		strconv.Itoa(numero),
		fonte,
		opTexto,
	)

}
