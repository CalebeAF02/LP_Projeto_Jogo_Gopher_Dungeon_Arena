package sistema

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/interfaces"
	"sync"
)

type SistemaMovimento struct{}

func (self *SistemaMovimento) Atualizar(cj interfaces.ICenaJogo) {

	//for _, entidade := range cj.GetEntidades() {

	//	entidade.Atualizar()
	//}

	entidades := cj.GetEntidades()

	// 1. Criamos um WaitGroup para gerenciar e esperar as Goroutines terminarem
	var wg sync.WaitGroup

	// 2. Disparar uma Goroutine para cada entidade rodar o seu método Atualizar() em paralelo
	for _, entidade := range entidades {

		wg.Add(1) // Avisa ao WaitGroup que uma nova rotina concorrente começou

		// Executa a Goroutine passando a entidade de forma isolada por parâmetro
		go func(e ecs.Entidade) {

			defer wg.Done() // Garante que avisa o término da execução ao finalizar a rotina

			// Processa toda a lógica interna de movimento/IA da entidade concorrentemente
			e.Atualizar()

		}(entidade)
	}

	// 3. Bloqueia a execução principal até que TODAS as Goroutines terminem o processamento
	wg.Wait()
}
