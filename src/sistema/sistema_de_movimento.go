package sistema

type SistemaMovimento struct{}

func (s *SistemaMovimento) Atualizar(g *Game) {

	for _, entidade := range g.entidades {

		entidade.Atualizar()
	}
}
