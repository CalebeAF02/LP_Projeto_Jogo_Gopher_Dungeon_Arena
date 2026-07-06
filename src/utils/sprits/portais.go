package portais

import (
	"image"
	_ "image/png" // Necessário para decodificar imagens PNG

	"github.com/hajimehoshi/ebiten/v2"
)

// Definição de constantes públicas para acessar os portais de qualquer lugar do projeto
const (
	PortalAzul = iota
	PortalRosa
	PortalVerdeClaro // Usado no PortalSpriteSaida
	PortalRoxo
	PortalAmarelo
	PortalVermelho
	PortalCiano
	PortalLaranja // Usado no PortalSpriteEntrada
	PortalAzulEscuro
)

// SpriteSheetPortal gerencia o fatiamento da imagem de portais
type SpriteSheetPortal struct {
	imagemOriginal *ebiten.Image
	subPortais     [9]*ebiten.Image
	larguraFrame   int
	alturaFrame    int
}

// NovoSpriteSheetPortal carrega a imagem de sprites e realiza o recorte matemático dos 9 portais
func NovoSpriteSheetPortal(imgImg *ebiten.Image) *SpriteSheetPortal {
	ss := &SpriteSheetPortal{
		imagemOriginal: imgImg,
	}

	// Pegamos o tamanho total da imagem e dividimos por 3 (já que é uma grade 3x3)
	bounds := imgImg.Bounds()
	ss.larguraFrame = bounds.Dx() / 3
	ss.alturaFrame = bounds.Dy() / 3

	// Fatiamento automático da grade 3x3
	indice := 0
	for linha := 0; linha < 3; linha++ {
		for coluna := 0; coluna < 3; coluna++ {
			// Calcula os limites (pixels) de onde começa e termina o frame atual
			x0 := coluna * ss.larguraFrame
			y0 := linha * ss.alturaFrame
			x1 := x0 + ss.larguraFrame
			y1 := y0 + ss.alturaFrame

			// Recorta e faz o cast correto para *ebiten.Image
			subImg := imgImg.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
			ss.subPortais[indice] = subImg
			indice++
		}
	}

	return ss
}

// ObterPortal retorna o sub-sprite correspondente ao índice (0 a 8) ou constante
func (self *SpriteSheetPortal) ObterPortal(tipo int) *ebiten.Image {
	if tipo < 0 || tipo >= len(self.subPortais) {
		return self.subPortais[PortalLaranja] // Fallback seguro caso mande um índice inválido
	}
	return self.subPortais[tipo]
}

// DesenharPortal desenha o portal específico diretamente na tela na posição desejada
func (self *SpriteSheetPortal) DesenharPortal(tela *ebiten.Image, tipo int, x, y float64) {
	portalImg := self.ObterPortal(tipo)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	tela.DrawImage(portalImg, op)
}
