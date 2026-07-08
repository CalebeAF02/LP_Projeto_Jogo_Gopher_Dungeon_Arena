# Paralelismo em Go para Gopher Dungeon Arena

## 1. Objetivo

Este documento descreve como evoluir o motor do jogo de um loop sequencial para um pipeline paralelo usando recursos idiomáticos de Go. A ideia é aumentar desempenho e abertura para comportamentos de IA mais complexos, mantendo o jogo seguro contra condições de corrida.

### Escopo
- Paralelizar decisão de bots
- Garantir sincronização por frame
- Proteger estado compartilhado
- Introduzir comunicação via canais
- Preparar a base para multiplayer e IA avançada

---

## 2. Arquitetura atual do jogo

O projeto já está organizado em um motor ECS com as camadas:

- `main.go` → inicializa Ebiten e roda `Game`
- `src/game.go` → delega `Update()` e `Draw()` para a cena atual
- `src/cenas/cenaJogo.go` → `CenaJogo.Update()` chama os sistemas em sequência
- `src/sistema/` → implementa sistemas de atualização e desenho
- `src/entidades/personagens/bot.go` → implementa o bot e seus movimentos
- `src/sistema/sistema_de_colisao.go` → faz checagem de colisões consultando todas as entidades

Atualmente, o fluxo de atualização do jogo é:

1. `SistemaInput`
2. `SistemaIA`
3. `SistemaSpawn`
4. `SistemaMovimento`
5. `SistemaEntidades`
6. `SistemaDebug`

Isso significa que um frame só avança depois que todos os sistemas terminam em sequência.

---

## 3. Por que paralelizar

### Vantagens
- bots podem calcular decisões simultaneamente
- uso mais eficiente de múltiplos núcleos
- frame de jogo pode escalar melhor com mais inimigos
- separação clara entre decisão e aplicação

### Cuidados
- Ebiten exige que `Draw()` e manipulação de janelas rodem na thread principal
- leitura e escrita de estado compartilhado precisam ser sincronizadas
- colisões e atualizações de entidades podem gerar race conditions

---

## 4. Principais pontos de paralelização

### 4.1. `SistemaIA` — etapa 1

Local ideal para lançar goroutines.

O objetivo aqui é:
- iterar sobre as entidades do tipo `BOT`
- calcular a próxima ação de cada bot em paralelo
- armazenar a decisão em um comando/struct
- não aplicar movimento diretamente nesta etapa

Arquivo alvo:
- `src/sistema/sistema_de_ia.go`

Exemplo de fluxo:
- criar `decisionChan := make(chan BotDecision, N)`
- para cada bot, iniciar `go func(bot *personagens.Bot) { decisionChan <- bot.Decidir(); wg.Done() }(bot)`
- esperar `wg.Wait()` e fechar canal
- coletar decisões para usar na próxima fase

Esse é o primeiro passo do roteiro técnico.

### 4.2. `sync.WaitGroup` — etapa 2

Use `sync.WaitGroup` para garantir que todas as goroutines terminem antes de avançar o frame.

Onde aplicar:
- em `SistemaIA` para esperar a decisão de todos os bots
- em `SistemaMovimento` se cada entidade for atualizada em paralelo

Importante:
- nunca avance para a próxima etapa do frame antes do `wg.Wait()`
- `WaitGroup` só sincroniza término; ele não protege estado compartilhado

### 4.3. `sync.RWMutex` — etapa 3

O estado compartilhado principal é:
- `CenaJogo.entidades`
- contadores e flags como `contadorBotsMortos`, `coletadoTudo`, `entrouNaSaida`
- componentes de entidades que podem ser lidos/escritos por múltiplas goroutines
- `SistemaColisao` lendo todas as entidades

Recomendações:
- adicione em `src/cenas/cenaJogo.go`:
  - `entidadesLock sync.RWMutex`
- use `RLock()` ao ler entidades e `Lock()` ao modificar o mapa ou componentes mutuáveis
- exponha métodos seguros como:
  - `GetEntidadesSeguras()` ou `GetEntidadesSnapshot()`
  - `SetEntidade()` com lock
  - `RemoverEntidade()` com lock
  - `GetEntidade(id)` se preciso

Isso evita concorrência insegura ao acessar o map e ao iterar sobre ele.

### 4.4. Channels e CSP — etapa 4

O modelo idiomático de Go é:

> Don't communicate by sharing memory; share memory by communicating.

Nesse projeto, essa abordagem ajuda a reduzir mutexes e deixar o fluxo mais claro.

Possível arquitetura por canal:

- `decisionChan` recebe as decisões de IA
- um agregador central transforma decisão em `MoveCommand`
- `moveChan` entrega comandos à fase de movimentação
- um executor central aplica comandos sequencialmente ou em paralelo protegido

Benefícios:
- mantêm o estado compartilhado concentrado em um ponto
- tornam mais explícito o que é leitura e o que é escrita
- facilitam a extensão para rede / multiplayer

### 4.5. Multiplayer e IA avançada — etapa 5

Com a base paralela consolidada:
- cada jogador remoto também gera comandos por canal
- o servidor/cliente do jogo pode tratar comandos de rede igual a decisões de bots
- a IA do bot pode evoluir para pathfinding, busca de alvo e tomadas de decisão mais complexas

O primeiro objetivo é estabilizar as etapas anteriores antes de adicionar rede.

---

## 5. Mudanças propostas por arquivo

### `src/cenas/cenaJogo.go`

Adicionar:
- `entidadesLock sync.RWMutex`
- acesso seguro ao mapa de entidades
- métodos para ler e escrever o estado com proteção
- possivelmente `GetEntidadesSnapshot()` que retorna uma cópia se necessário

O objetivo é fazer com que qualquer sistema que consulte entidades use lock. Isso inclui `SistemaColisao`, `SistemaDesenho`, `SistemaDebug`, `SistemaIA` e `SistemaMovimento`.

### `src/sistema/sistema_de_ia.go`

Implementar:
- coleta de bots
- decisões em goroutines
- `WaitGroup`
- canal de decisão
- output de uma lista de intenções que será usada pela próxima fase

### `src/sistema/sistema_de_movimento.go`

Atualizar para:
- não depender diretamente de `e.Atualizar()` sem proteção
- processar comandos de movimento pré-calculados
- usar goroutines apenas se o acesso à colisão e à posição for protegido

Se você for manter a lógica atual, envolva `e.Atualizar()` com locks ou use um snapshot de estado.

### `src/sistema/sistema_de_colisao.go`

Melhorar segurança ao:
- acessar entidades através de getters seguros
- iterar sobre um snapshot de entidades ou segurar `RLock()` durante a verificação
- evitar condições em que uma entidade é removida durante a iteração

### `src/entidades/personagens/bot.go`

Refatorar `Bot.Atualizar()` em duas fases:
- `Decidir()` / `CalcularMovimento()`
- `ExecutarMovimento()`

Isso permite que a lógica de decisão rode em paralelo e que a aplicação do movimento seja sincronizada separadamente.

---

## 6. Exemplo de pipeline paralelo

### Fase A: Input e estado do jogador
- `SistemaInput` processa controle do jogador
- não paralelizar esta etapa

### Fase B: Decisão de IA
- `SistemaIA` lança goroutines por bot
- cada bot calcula sua próxima ação
- decisões são agregadas em `[]BotDecision`

### Fase C: Aplicação de movimento
- `SistemaMovimento` usa as decisões já calculadas
- cada bot executa sua movimentação
- o movimento pode ser paralelizado se a colisão/estado for protegido

### Fase D: Entidades e remoção
- `SistemaEntidades` aplica remoções e atualizações finais
- mantém o mapa consistente antes do desenho

### Fase E: Desenho
- `CenaJogo.Draw()` permanece no fluxo normal do Ebiten
- `SistemaDesenho` não deve ser paralelizado, pois ele lê o estado final do frame

Esse fluxo garante que o frame só finalize após todas as ações de IA e movimento concluírem.

---

## 7. Detalhes de implementação

### 7.1. Estruturas úteis

Crie structs de comando para separar intenção e execução:

```go
package sistema

import "Gopher_Dungeon_Arena/src/ecs"

type BotDecision struct {
    EntidadeID ecs.EntidadeID
    TipoAcao   string
    Direcao    string
    DestinoX   float64
    DestinoY   float64
}

type MoveCommand struct {
    EntidadeID ecs.EntidadeID
    DX         float64
    DY         float64
}
```

### 7.2. `CenaJogo` seguro

No `CenaJogo`:

```go
type CenaJogo struct {
    entidades     map[ecs.EntidadeID]ecs.Entidade
    entidadesLock sync.RWMutex
    // ... resto
}
```

E métodos:

```go
func (self *CenaJogo) GetEntidadesSnapshot() map[ecs.EntidadeID]ecs.Entidade {
    self.entidadesLock.RLock()
    defer self.entidadesLock.RUnlock()
    copia := make(map[ecs.EntidadeID]ecs.Entidade, len(self.entidades))
    for k, v := range self.entidades {
        copia[k] = v
    }
    return copia
}

func (self *CenaJogo) SetEntidade(nEntidade ecs.EntidadeID, e ecs.Entidade) {
    self.entidadesLock.Lock()
    defer self.entidadesLock.Unlock()
    self.entidades[nEntidade] = e
}

func (self *CenaJogo) RemoverEntidade(entidade ecs.EntidadeID) {
    self.entidadesLock.Lock()
    defer self.entidadesLock.Unlock()
    delete(self.entidades, entidade)
}
```

### 7.3. `SistemaIA` com goroutines

Implemente a lógica de IA em paralelo, por exemplo:

```go
func (self *SistemaIA) Atualizar(cj interfaces.ICenaJogo) {
    entidades := cj.GetEntidadesSnapshot()
    var wg sync.WaitGroup
    decisions := make(chan BotDecision, len(entidades))

    for _, e := range entidades {
        if bot, ok := e.(*personagens.Bot); ok {
            wg.Add(1)
            go func(bot *personagens.Bot) {
                defer wg.Done()
                decisions <- bot.CalcularDecisao()
            }(bot)
        }
    }

    go func() {
        wg.Wait()
        close(decisions)
    }()

    for decision := range decisions {
        // armazena em estrutura que será usada por SistemaMovimento
    }
}
```

### 7.4. `SistemaMovimento` com sincronização

No movimento, use um `WaitGroup` se quiser paralelizar cada entidade:

```go
func (self *SistemaMovimento) Atualizar(cj interfaces.ICenaJogo) {
    commands := cj.GetPendingMoveCommands()
    var wg sync.WaitGroup
    for _, cmd := range commands {
        wg.Add(1)
        go func(cmd MoveCommand) {
            defer wg.Done()
            aplicador.Mover(cmd)
        }(cmd)
    }
    wg.Wait()
}
```

Se usar canais, um consumidor central pode aplicar movimentos serialmente ou com bloqueios.

---

## 8. Testes e validação

### `go test -race ./...`

Esse comando é essencial para detectar erros de concorrência que não aparecem no comportamento normal do jogo.

### Teste manual

- rode o jogo com muitos bots
- observe se há travamentos ou comportamento errático
- verifique se colisões continuam consistentes
- compare o resultado com a versão sequencial

### Estratégia incremental

1. implemente getters seguros em `CenaJogo`
2. paralelize apenas `SistemaIA`
3. verifique `go test -race`
4. paralelize `SistemaMovimento`
5. valide colisões e remoções

---

## 9. Recomendações finais

- comece com `SistemaIA` antes de tocar em `SistemaMovimento`
- mantenha `Draw()` no fluxo normal do Ebiten
- use canais para comunicação de intenções
- use `RWMutex` apenas onde necessário
- evite modificar `entidades` enquanto outra goroutine itera sobre o map
- transforme a lógica de bot para duas fases: decisão e execução

Esse caminho abre a porta para ampliar o jogo com:
- IA mais avançada
- networking/multiplayer
- comandos remotos e replay
- agentes com pathfinding

---

## 10. Arquivo sugerido

Nome sugerido para este documento:
- `docs/paralelismo_go.md`

Mantê-lo no diretório `docs/` ajuda a separar o guia técnico da documentação geral do jogo.
