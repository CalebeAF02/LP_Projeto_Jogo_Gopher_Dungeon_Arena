package imagens

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed icone.png
var iconeBytes []byte

//go:embed menu.png
var menuBytes []byte

// --- VARIÁVEIS GLOBAIS DE IMAGEM ---

var IconeDoJogo *ebiten.Image
var ImagemMenu *ebiten.Image

func AplicarIconeJanela() {
	// 1. Carrega e configura o Ícone da Janela
	imgIcone, _, err := image.Decode(bytes.NewReader(iconeBytes))
	if err != nil {
		log.Printf("Erro ao carregar ícone da janela: %v", err)
	} else {
		IconeDoJogo = ebiten.NewImageFromImage(imgIcone)
		icones := []image.Image{imgIcone}
		ebiten.SetWindowIcon(icones)
	}

	// 2. Carrega a Imagem do Menu
	imgMenu, _, err := image.Decode(bytes.NewReader(menuBytes))
	if err != nil {
		log.Printf("Erro ao carregar imagem do menu: %v", err)
	} else {
		ImagemMenu = ebiten.NewImageFromImage(imgMenu)
	}
}

func DesenharImagemCentralizada(tela *ebiten.Image, img *ebiten.Image, posicaoY float64, fatorEscala float64, larguraJanela int) {
	if img == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}

	// 1. Aplica a escala primeiro
	op.GeoM.Scale(fatorEscala, fatorEscala)

	// 2. Calcula a largura já reduzida pela escala
	larguraReduzida := float64(img.Bounds().Dx()) * fatorEscala

	// 3. Calcula o X para centralizar perfeitamente baseado na largura da janela passada
	posicaoX := (float64(larguraJanela) / 2.0) - (larguraReduzida / 2.0)

	// 4. Move para a posição correta
	op.GeoM.Translate(posicaoX, posicaoY)

	// 5. Renderiza na tela
	tela.DrawImage(img, op)
}

func DesenharImagemPreencherTela(tela *ebiten.Image, img *ebiten.Image, larguraJanela, alturaJanela int) {
	if img == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}

	// 1. Obtém o tamanho original da imagem
	larguraOriginal := float64(img.Bounds().Dx())
	alturaOriginal := float64(img.Bounds().Dy())

	// 2. Calcula o fator de escala necessário para X e Y
	escalaX := float64(larguraJanela) / larguraOriginal
	escalaY := float64(alturaJanela) / alturaOriginal

	// 3. Aplica a escala para esticar/ajustar a imagem ao tamanho da tela
	op.GeoM.Scale(escalaX, escalaY)

	// 4. Desenha a imagem a partir do ponto (0,0) que é o canto superior esquerdo
	tela.DrawImage(img, op)
}
