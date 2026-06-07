package assets

import (
	"bytes"
	"log"

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
