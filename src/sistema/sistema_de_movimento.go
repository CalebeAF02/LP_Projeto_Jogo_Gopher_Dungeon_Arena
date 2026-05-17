package sistema

type SistemaMovimento struct{}

func (s *SistemaMovimento) Atualizar(g *Game) {

	for _, entidade := range g.GetEntidades() {

		entidade.Atualizar()
	}
}
