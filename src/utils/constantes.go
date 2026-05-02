package utils

import (
	"math/rand"
	"time"
)

const (
	LARGURA             = 1240
	ALTURA              = 720
	PROPORCAO_MINI_MAPA = float64(10)

	MUNDO_LARGURA = float64(LARGURA * (PROPORCAO_MINI_MAPA / (PROPORCAO_MINI_MAPA / 2)))
	MUNDO_ALTURA  = float64(ALTURA * (PROPORCAO_MINI_MAPA / (PROPORCAO_MINI_MAPA / 2)))

	NOME_JOGO = "Gopher_Dungeon_Arena"
)

func XAleatorio(r *rand.Rand) float64 {
	return float64(r.Intn(LARGURA))
}

func YAleatorio(r *rand.Rand) float64 {
	return float64(r.Intn(ALTURA))
}

func GeradorAleatorio() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
