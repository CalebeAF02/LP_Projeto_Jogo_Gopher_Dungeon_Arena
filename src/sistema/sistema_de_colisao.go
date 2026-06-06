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

	// Limitamos as tentativas para evitar loop infinito se o mapa estiver cheio
	for tentativas := 0; tentativas < 100; tentativas++ {
		x := float64(g.GetAleatorio().Intn(int(g.GetMundo().PosXmax(utils.BOT_TAMANHO_MUNDO))))
		y := float64(g.GetAleatorio().Intn(int(g.GetMundo().PosYmax(utils.BOT_TAMANHO_MUNDO))))

		// REUTILIZAÇÃO: Usamos diretamente o método do jogo para checar barreiras (paredes)
		corpoTemporario := geometria.NovoRetangulo(x, y, larguraBot, alturaBot)
		if !g.ColideComTipo(corpoTemporario, entidades.PAREDE.String()) {
			return geometria.NovoPonto(x, y)
		}
	}

	return nil
}

func (g *Game) VaiColidir(meuCorpoAtual *geometria.Retangulo, proximoCorpo *geometria.Retangulo) bool {
	for _, e := range g.GetEntidades() {
		tipo := e.GetTipo()
		if tipo == entidades.PAREDE.String() || tipo == entidades.JOGADOR.String() || tipo == entidades.BOT.String() {
			if corpoEntidade := e.GetComponente(componentes.CORPO.String()); corpoEntidade != nil {
				corpo := corpoEntidade.(*geometria.Retangulo)

				// EVITA AUTO-COLISÃO REAL:
				// Se a entidade da lista tiver exatamente a mesma posição X e Y do meu corpo atual,
				// significa que essa entidade SOU EU MESMO na tabela do ECS. Ignoramos!
				if corpo.GetX() == meuCorpoAtual.GetX() && corpo.GetY() == meuCorpoAtual.GetY() {
					continue
				}

				// Agora sim, testa se a minha PRÓXIMA posição vai bater em OUTRA entidade
				if proximoCorpo.Colide(corpo) {
					if meuCorpoAtual.Colide(corpo) {
						continue
					}
					return true
				}
			}
		}
	}
	return false
}

// ColideComTipo isola uma busca específica (útil para o Spawn ou lógicas de IA direcionadas)
func (g *Game) ColideComTipo(eu *geometria.Retangulo, tipoDesejado string) bool {
	for _, e := range g.GetEntidades() {
		if e.GetTipo() == tipoDesejado {
			if corpoEntidade := e.GetComponente(componentes.CORPO.String()); corpoEntidade != nil {
				if eu.Colide(corpoEntidade.(*geometria.Retangulo)) {
					return true
				}
			}
		}
	}
	return false
}

// Métodos auxiliares semanticamente limpos, reaproveitando a função genérica
func (g *Game) ColideComBarreiras(eu *geometria.Retangulo) bool {
	return g.ColideComTipo(eu, entidades.PAREDE.String())
}

func (g *Game) ColideComJogador(eu *geometria.Retangulo) bool {
	return g.ColideComTipo(eu, entidades.JOGADOR.String())
}

func (g *Game) ColideComBot(eu *geometria.Retangulo) bool {
	return g.ColideComTipo(eu, entidades.BOT.String())
}
