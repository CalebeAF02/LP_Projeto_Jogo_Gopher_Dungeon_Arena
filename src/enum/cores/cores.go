package cores

import (
	"image/color"
)

var (
	BRANCO = rgb(227, 242, 253)

	PRETO = rgb(33, 33, 33)

	AMBAR          = rgb(255, 193, 37)
	AMARELO        = rgb(255, 238, 88)
	AMARELO_CLARO  = rgb(255, 241, 118)
	AMARELO_ESCURO = rgb(251, 192, 45)

	VERMELHO        = rgb(239, 83, 80)
	VERMELHO_CLARO  = rgb(239, 154, 154)
	VERMELHO_ESCURO = rgb(183, 28, 28)

	AZUL        = rgb(3, 155, 229)
	AZUL_CLARO  = rgb(187, 222, 251)
	AZUL_ESCURO = rgb(13, 71, 161)

	CIANO        = rgb(0, 172, 193)
	CIANO_CLARO  = rgb(178, 235, 242)
	CIANO_ESCURO = rgb(0, 96, 100)

	LARANJA        = rgb(255, 112, 67)
	LARANJA_CLARO  = rgb(255, 171, 145)
	LARANJA_ESCURO = rgb(191, 54, 12)

	VERDE        = rgb(76, 175, 80)
	VERDE_CLARO  = rgb(129, 199, 132)
	VERDE_ESCURO = rgb(27, 94, 32)

	ROSA        = rgb(171, 71, 188)
	ROSA_CLARO  = rgb(244, 143, 177)
	ROSA_ESCURO = rgb(136, 14, 79)

	ROXO        = rgb(156, 39, 176)
	ROXO_CLARO  = rgb(179, 157, 219)
	ROXO_ESCURO = rgb(74, 20, 140)

	MARROM        = rgb(141, 110, 99)
	MARROM_CLARO  = rgb(161, 136, 127)
	MARROM_ESCURO = rgb(62, 39, 35)

	CINZA_CLARO  = rgb(189, 189, 189)
	CINZA_ESCURO = rgb(66, 66, 66)

	VERDE_LIMAO        = rgb(118, 255, 3)
	VERDE_LIMAO_CLARO  = rgb(204, 255, 144)
	VERDE_LIMAO_ESCURO = rgb(100, 221, 23)
)

func rgb(r uint8, g uint8, b uint8) color.Color {
	return color.RGBA{r, g, b, 255}
}
