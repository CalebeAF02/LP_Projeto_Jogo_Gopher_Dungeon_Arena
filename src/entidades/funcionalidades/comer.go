package funcionalidades

import (
	"Gopher_Dungeon_Arena/src/componentes"
	"Gopher_Dungeon_Arena/src/ecs"
	"fmt"
)

func EncherBucho(jogador ecs.Entidade, comida ecs.Entidade) {

	fmt.Printf("++ Hora de comer !\n")

	energiaComp := comida.GetComponente(componentes.ENERGIA.String())
	energia := energiaComp.(*componentes.Energia)

	energia.Status = false
}
