package sistema

type SistemaCamera struct{}

func (s *SistemaCamera) Atualizar(g *Game) {
	// --- ATUALIZAÇÃO DA CÂMERA ---
	lTimes := g.GetTimes()

	if len(lTimes) > 0 && len(lTimes[0].GetJogadores()) > 0 {

		jogador := lTimes[0].GetJogador(0)

		g.camera.OrganizarCameraPeloJogador(jogador.GetPosicao())
	}
}
