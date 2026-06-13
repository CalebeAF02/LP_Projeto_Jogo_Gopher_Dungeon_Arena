package funcionalidades

import (
	"Gopher_Dungeon_Arena/src/componentes"
	"Gopher_Dungeon_Arena/src/ecs"
)



func TeleTransporta(portalEntrada ecs.Entidade, bot ecs.Entidade) {

	portalTransporte_com := portalEntrada.GetComponente(componentes.ENVIANDO_TELETRANSPORTE.String())
	teletransporte := portalTransporte_com.(*componentes.EnviandoTeletransporte)

	if teletransporte.TemBot {
		//fmt.Printf("Já tem bot !!!\n")
	} else {
		//fmt.Printf("Guardar bot ....\n")

		teletransporte.TemBot = true
		teletransporte.Bot = bot
		teletransporte.Contagem = 150

		liberdade_comp := bot.GetComponente(componentes.ATIVIDADE.String())
		liberdade := liberdade_comp.(*componentes.Atividade)
		liberdade.Acao = componentes.AIVIDADE_TELETRANSPORTE

	}
}
