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
	POS_X_MAPA   = JANELA_LARGURA - (MAPA_LARGURA + MAPA_ALTURA/4)
	POS_Y_MAPA   = MAPA_ALTURA / 3

	NOME_JOGO = "Gopher_Dungeon_Arena"
)

func XAleatorio(r *rand.Rand) float64 {
	return float64(r.Intn(JANELA_LARGURA))
}

func YAleatorio(r *rand.Rand) float64 {
	return float64(r.Intn(JANELA_ALTURA))
}

func GeradorAleatorio() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
