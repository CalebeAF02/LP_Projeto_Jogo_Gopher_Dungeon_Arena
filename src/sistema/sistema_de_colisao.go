package sistema

import (
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/componentes"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/utils"
)

func OrganizaPosicaoAleatoriaBot(g *Game) *geometria.Ponto {
	larguraBot := float64(utils.BOT_TAMANHO_MUNDO)
	alturaBot := float64(utils.BOT_TAMANHO_MUNDO)

	// Tentamos encontrar uma posição válida (limitamos as tentativas para evitar loop infinito)
	for tentativas := 0; tentativas < 100; tentativas++ {
		x := float64(g.aleatorio.Intn(int(g.mundo.PosXmax(utils.BOT_TAMANHO_MUNDO))))
		y := float64(g.aleatorio.Intn(int(g.mundo.PosYmax(utils.BOT_TAMANHO_MUNDO))))

		if PosicaoEstaLivre(g, x, y, larguraBot, alturaBot) {
			return geometria.NovoPonto(x, y)
		}
	}

	return nil
}

func PosicaoEstaLivre(g *Game, x, y float64, largura, altura float64) bool {
	// Cria um retângulo temporário para representar o corpo do bot na posição sorteada
	corpoBot := geometria.NovoRetangulo(x, y, largura, altura)

	for _, e := range g.GetEntidades() {
		if e.GetTipo() == entidades.PAREDE.String() {
			if corpoParede := e.GetComponente(componentes.CORPO.String()); corpoParede != nil {
				// Se colidir com qualquer parede, a posição não está livre
				if corpoBot.Colide(corpoParede.(*geometria.Retangulo)) {
					return false
				}
			}
		}
	}
	return true // Não colidiu com nenhuma parede
}
