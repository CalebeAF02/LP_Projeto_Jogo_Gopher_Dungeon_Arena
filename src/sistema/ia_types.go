package sistema

import "Gopher_Dungeon_Arena/src/ecs"

// BotDecision representa uma decisão calculada pela IA para uma entidade
type BotDecision struct {
	EntidadeID ecs.EntidadeID
	TipoAcao   string
}
