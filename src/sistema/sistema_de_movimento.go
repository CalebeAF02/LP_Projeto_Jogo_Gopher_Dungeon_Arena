package sistema

import "Gopher_Dungeon_Arena/src/interfaces"

type SistemaMovimento struct{}

func (s *SistemaMovimento) Atualizar(cj interfaces.ICenaJogo) {

	for _, entidade := range cj.GetEntidades() {

		entidade.Atualizar()
	}
}
