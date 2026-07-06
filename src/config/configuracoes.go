package config

import (
	"math/rand"
	"time"
)

const (
	JANELA_LARGURA  = 1280
	JANELA_ALTURA   = 720
	PROPORCAO_MUNDO = 2 // numeros pares

	MUNDO_LARGURA = JANELA_LARGURA * PROPORCAO_MUNDO
	MUNDO_ALTURA  = JANELA_ALTURA * PROPORCAO_MUNDO

	PROPORCAO_MAPA = PROPORCAO_MUNDO * 4 // numeros pares

	MAPA_LARGURA = MUNDO_LARGURA / PROPORCAO_MAPA
	MAPA_ALTURA  = MUNDO_ALTURA / PROPORCAO_MAPA

	MM1_POS_X_MAPA = 50
	MM1_POS_Y_MAPA = MAPA_ALTURA / 3

	MM2_POS_X_MAPA = JANELA_LARGURA - (MAPA_LARGURA + MAPA_ALTURA/4)
	MM2_POS_Y_MAPA = MAPA_ALTURA / 3

	MM3_POS_X_MAPA = JANELA_LARGURA - (MAPA_LARGURA + MAPA_ALTURA/4)
	MM3_POS_Y_MAPA = JANELA_ALTURA - (MAPA_ALTURA + (MAPA_ALTURA / 4))

	MM4_POS_X_MAPA = 50
	MM4_POS_Y_MAPA = JANELA_ALTURA - (MAPA_ALTURA + (MAPA_ALTURA / 4))

	NOME_JOGO = "Gopher_Dungeon_Arena"
)

func XAleatorio(self *rand.Rand) float64 {
	return float64(self.Intn(JANELA_LARGURA))
}

func YAleatorio(self *rand.Rand) float64 {
	return float64(self.Intn(JANELA_ALTURA))
}

func GeradorAleatorio() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
