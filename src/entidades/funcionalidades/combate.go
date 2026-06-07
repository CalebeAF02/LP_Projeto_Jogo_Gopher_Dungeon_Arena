package funcionalidades

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/enum/componentes"
	"fmt"
)

func Simetria(s1 string, s2 string, c1 string, c2 string) bool {
	return (s1 == c1 && s2 == c2) || (s1 == c2 && s2 == c1)
}

func Combate_jogador_vs_bot(jogador ecs.Entidade, bot ecs.Entidade) {

	fmt.Println("-----------------------------")
	botVidaOrigemComp := bot.GetComponente(componentes.VIDA.String())
	botVidaOrigem := botVidaOrigemComp.(*componentes.Vida)
	fmt.Println("SANGUE_BOT-> Tinha : ", botVidaOrigem.Sangue)
	botVidaOrigem.PerdeSangue(30)
	fmt.Println("SANGUE_BOT-> Perdeu : ", 30)
	fmt.Println("SANGUE_BOT-> Possui : ", botVidaOrigem.Sangue)
	fmt.Println("SANGUE_BOT-> Vivo  : ", botVidaOrigem.Status)

	jogadorVidaOrigemComp := jogador.GetComponente(componentes.VIDA.String())
	jogadorVidaOrigem := jogadorVidaOrigemComp.(*componentes.Vida)
	fmt.Println("SANGUE_JOGADOR-> Tinha : ", jogadorVidaOrigem.Sangue)
	jogadorVidaOrigem.PerdeSangue(10)
	fmt.Println("SANGUE_JOGADOR-> Perdeu : ", 10)
	fmt.Println("SANGUE_JOGADOR-> Possui : ", jogadorVidaOrigem.Sangue)
	fmt.Println("SANGUE_JOGADOR-> Vivo  : ", jogadorVidaOrigem.Status)

	if !botVidaOrigem.EstaVivo("BOT") {
		botVidaOrigem.Status = false
	}

	if !jogadorVidaOrigem.EstaVivo("JOGADOR") {
		jogadorVidaOrigem.Status = false
	}

}
