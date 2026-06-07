package funcionalidades

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/enum/componentes"
	"Gopher_Dungeon_Arena/src/utils"
)

func Simetria(s1 string, s2 string, c1 string, c2 string) bool {
	return (s1 == c1 && s2 == c2) || (s1 == c2 && s2 == c1)
}

func ReduzSangue(entidade ecs.Entidade, rit int) {

	botVidaOrigemComp := entidade.GetComponente(componentes.VIDA.String())
	botVidaOrigem := botVidaOrigemComp.(*componentes.Vida)

	botNivelOrigemComp := entidade.GetComponente(componentes.NIVEL.String())
	botNivelOrigem := botNivelOrigemComp.(*componentes.Nivel)

	//fmt.Println("SANGUE_BOT-> Tinha : ", botVidaOrigem.Sangue)
	botVidaOrigem.PerdeSangue(rit, botNivelOrigem.Valor)
	//fmt.Println("SANGUE_BOT-> Perdeu : ", rit)
	//fmt.Println("SANGUE_BOT-> Possui : ", botVidaOrigem.Sangue)
	//fmt.Println("SANGUE_BOT-> Vivo  : ", botVidaOrigem.Status)

	if !botVidaOrigem.EstaVivo() {
		botVidaOrigem.Status = false
	}
}

func CombateJogadorBot(jogador ecs.Entidade, bot ecs.Entidade) {
	ReduzSangue(jogador, utils.COMBATE_BOT_RIT)
	ReduzSangue(bot, utils.COMBATE_JOGADOR_RIT)
}
