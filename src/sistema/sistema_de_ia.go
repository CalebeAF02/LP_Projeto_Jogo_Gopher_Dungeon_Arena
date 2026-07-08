package sistema

import (
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/interfaces"
	"sync"
)

type SistemaIA struct{}

func (self *SistemaIA) Atualizar(cj interfaces.ICenaJogo) {
	entidades := cj.GetEntidades()

	var wg sync.WaitGroup
	decisions := make(chan BotDecision, len(entidades))

	for _, e := range entidades {
		if e.GetTipo() == "BOT" {
			if bot, ok := e.(*personagens.Bot); ok {
				wg.Add(1)
				go func(b *personagens.Bot) {
					defer wg.Done()
					// Exemplo de cálculo paralelo: checar estado e emitir decisão simples
					if b.EstaVivo() && b.PossoMeMover() {
						decisions <- BotDecision{EntidadeID: b.GetID(), TipoAcao: "INTEND_MOVE"}
					} else {
						decisions <- BotDecision{EntidadeID: b.GetID(), TipoAcao: "NONE"}
					}
				}(bot)
			}
		}
	}

	go func() {
		wg.Wait()
		close(decisions)
	}()

	// Por enquanto apenas iteramos sobre as decisões coletadas.
	for d := range decisions {
		_ = d // placeholder: futuras integrações processarão essas decisões
	}
}
