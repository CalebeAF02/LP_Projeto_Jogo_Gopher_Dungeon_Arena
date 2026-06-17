package funcionalidades

import (
	"Gopher_Dungeon_Arena/src/componentes"
	"Gopher_Dungeon_Arena/src/ecs"
	"fmt"
)

func ConcluirPartida(jogador ecs.Entidade) {

	pontuacaoComp := jogador.GetComponente(componentes.PONTUACAO.String())
	pontuacao := pontuacaoComp.(*componentes.Pontuacao)

	if pontuacao.Coletado >= pontuacao.Requisito {

		fmt.Printf("--->> Conclui o jogo !!!\n")

		pontuacao.EntreiNaSaida = true

	}

}
